package sendto

import "fmt"

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Test email from Go!\n"
	body := "<html><body><h1>Hello World!</h1></body></html>"
	msg := []byte(subject + mime + body)
return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{Address: from, Name: "test"},
		To: to,
		Subject: "OTP Verification",
		Body: fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}
	messageMail :=
}
