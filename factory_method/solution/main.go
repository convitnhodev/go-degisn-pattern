package main

import "fmt"

// obligate user user static function

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: #{message} {Sender: Email}")
}

type SmsNotifier struct{}

func (SmsNotifier) Send(message string) {
	fmt.Printf("Sending message: #{message} {Sender: Sms}")
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	if t == "" {
		return SmsNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("email"),
	}

	s.SendNotification("Hello, world!")
}
