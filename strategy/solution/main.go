package main

import "fmt"

// Solid
// Satified: S, O

// pros:
// cons: complicated

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

func main() {
	s := NotificationService{
		notifier: EmailNotifier{},
	}

	s.SendNotification("Hello, world!")
}
