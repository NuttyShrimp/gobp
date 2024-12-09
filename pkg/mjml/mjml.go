package mjml

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/studentkickoff/gobp/pkg/config"
	"go.uber.org/zap"
)

func Init() error {
	err := os.Mkdir("tmp", os.ModeDir)
	if err != nil && !os.IsExist(err) {
		return err
	}
	s, err := os.Stat("tmp/mjml-converter")
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if s != nil {
		return nil
	}

	zap.L().Info("MJML CLI not found, downloading it")
	ghAssetName := "mjml-converter-"

	switch runtime.GOARCH {
	case "amd64":
		ghAssetName += "x86_64-"
	case "arm64":
		ghAssetName += "aarch64-"
	default:
		zap.L().Panic("Unsupported arch for mjml cli")
	}

	switch runtime.GOOS {
	case "darwin":
		ghAssetName += "apple-darwin.tar.gz"
	case "linux":
		ghAssetName += "unknown-linux-"
		ghAssetName += config.GetDefaultString("mails.linker", "gnu")
		ghAssetName += ".tar.gz"
	default:
		zap.L().Panic("This CLI is not supported on your machine")
	}

	zap.L().Info("Downloading MJML CLI", zap.String("asset", ghAssetName), zap.String("arch", runtime.GOARCH), zap.String("os", runtime.GOOS))

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
