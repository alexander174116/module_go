package main

import (
	"fmt"
)

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

type SMSNotyfier struct {
	PhoneNumber string
}

func (e *EmailNotifier) Notify(message string) {
	fmt.Println("Email to", e.EmailAddress, ":", message)
}

func (s *SMSNotyfier) Notify(message string) {
	fmt.Println("SMS to", s.PhoneNumber, ":", message)
}

func SendNotifications(nots []Notifier, msg string) {
	for _, note := range nots {
		note.Notify(msg)
	}
}

func main() {
	mail_a := EmailNotifier{EmailAddress: "asdad@ya.ru"}
	mail_b := EmailNotifier{EmailAddress: "asLOLdad@ya.ru"}
	mail_c := EmailNotifier{EmailAddress: "asd123123ad@ya.ru"}
	sms_a := SMSNotyfier{PhoneNumber: "+79179058978"}
	sms_b := SMSNotyfier{PhoneNumber: "+79179798978"}
	sms_c := SMSNotyfier{PhoneNumber: "+73179008978"}
	nots := []Notifier{&mail_a, &mail_b, &mail_c, &sms_a, &sms_b, &sms_c}
	SendNotifications(nots, "yo")
}
