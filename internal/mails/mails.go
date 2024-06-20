package mails

func ExampleMail(recipient, subject string) error {
	html, err := renderHTMLMail("example", nil)
	if err != nil {
		return err
	}
	return sendMail(html, "This is an example mail", recipient, subject)
}
