package mjml

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func Init() error {
	err := os.Mkdir("tmp", os.ModeDir)
	if !os.IsExist(err) {
		return err
	}
	_, err = os.Stat("tmp/mjm-converter")
	if !os.IsNotExist(err) {
		return err
	}

	fmt.Println("mjml cli not installed!")
	ghAssetName := "mjml-converter-"

	switch runtime.GOARCH {
	case "amd64":
		ghAssetName += "x86_64-"
	case "arm64":
		ghAssetName += "aarch64-"
	default:
		panic("Unsupported arch for mjml cli")
	}

	switch runtime.GOOS {
	case "darwin":
		ghAssetName += "apple-darwin.tar.gz"
	case "linux":
		// No proper way to check if gnu or musl. So going for gnu
		ghAssetName += "unknown-linux-gnu.tar.gz"
	default:
		panic("This CLI is not supported on your machine")
	}

	fmt.Printf("%s - %s - %s", runtime.GOARCH, runtime.GOOS, ghAssetName)

	return fetchRelease(ghAssetName)
}

func fetchRelease(release string) error {
	out, err := os.Create("tmp/mrml-cli.tar.gz")
	if err != nil {
		return err
	}
	// nolint:errcheck // we do not care if it errors
	defer out.Close()

	resp, err := http.Get(fmt.Sprintf("https://github.com/StudentKickOff/mrml-cli/releases/latest/download/%s", release))
	if err != nil {
		return err
	}
	// nolint:errcheck // we do not care if it errors
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	cmd := exec.Command("tar", "-xvf", "tmp/mrml-cli.tar.gz", "-C", "tmp")

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("mjml cli has been installed to /tmp folder!")
	return nil
}

func Convert(template string) (string, error) {
	output, err := exec.Command("tmp/mjml-converter", template).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
