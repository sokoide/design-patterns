package adapter

import (
	"fmt"
	"template-method-example/domain"
)

type Sms struct {
	domain.DefaultOtp
}

func (s *Sms) GenRandomOTP(len int) string {
	return "1234" // Override
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving %s\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "SMS OTP is " + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending message: %s\n", message)
	return nil
}

type Email struct {
	domain.DefaultOtp
}

func (s *Email) GenRandomOTP(len int) string {
	return "5678" // Override
}

func (s *Email) SaveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving %s\n", otp)
}

func (s *Email) GetMessage(otp string) string {
	return "EMAIL OTP is " + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending message: %s\n", message)
	return nil
}
