package main

import (
	"fmt"
	"template-method-example/adapter"
	"template-method-example/domain"
)

func main() {
	fmt.Println("=== Template Method Pattern ===")

	smsOTP := &adapter.Sms{}
	o := domain.NewOtp(smsOTP)
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &adapter.Email{}
	o2 := domain.NewOtp(emailOTP)
	o2.GenAndSendOTP(4)
}
