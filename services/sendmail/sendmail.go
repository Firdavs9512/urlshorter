package sendmail

import (
	"fmt"
	"os"

	"github.com/jenazads/gomail"
)

var email string = os.Getenv("APP_MAIL_EMAIL")
var pass string = os.Getenv("APP_MAIL_PASSWORD")

func send(to_email, service_name, message string) {
	gomailObject, _ := gomail.NewGoMail()

	// credentials
	gomailObject.Set("Username", email)
	gomailObject.Set("Password", pass)

	// servername
	gomailObject.Set("Servername", "smtp.gmail.com:465")

	// from
	gomailObject.Set("From", email)
	gomailObject.Set("From_name", service_name)

	// to
	gomailObject.Set("To", to_email)
	gomailObject.Set("To_name", to_email)

	// subject
	gomailObject.Set("Subject", "Email confirmation password")

	// body message
	msg := fmt.Sprintf("This is the fabulous body message\n\nGood Luck!!\n%s", message)
	gomailObject.Set("BodyMessage", msg)

	err := gomailObject.SendMessage()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func SendMail(mess, toemail string) {
	service := "Firdavs MCHJ"
	send(toemail, service, mess)
}
