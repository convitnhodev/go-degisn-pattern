package main

import "fmt"

type NotificationService struct {
	notifierType string
}

// when have a lot of cases, we can use strategy pattern

func (s NotificationService) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Printf("Sendding message: %s (Sender: Email)\n", message)
	} else if s.notifierType == "sms" {
		fmt.Printf("Sendding message: %s (Sender: Sms)\n", message)
	}
}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNotification("Hello world!")
}
