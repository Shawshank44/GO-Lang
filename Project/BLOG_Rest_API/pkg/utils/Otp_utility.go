package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
)

func GenerateOTP(length int) (string, error) {
	otp := ""

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", ErrorHandler(err, "Unable to generate OTP")
		}
		otp += fmt.Sprintf("%d", n)
	}
	return otp, nil
}

func SendOTPEmail(toEmail, otp, Subject string) error {
	from := "no-reply@blogadmins.dev"
	smtpHost := "localhost"
	smtpPort := "1025"

	message := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"Your OTP is: %s\nThis OTP expires in 5 minutes.\r\n",
		toEmail, Subject, otp,
	))

	addr := smtpHost + ":" + smtpPort
	return smtp.SendMail(addr, nil, from, []string{toEmail}, message)

}
