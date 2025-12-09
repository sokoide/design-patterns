package domain

import "fmt"

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string)
	SendNotification(string) error
}

// Template
type Otp struct {
	iOtp IOtp
}

func NewOtp(i IOtp) *Otp {
	return &Otp{iOtp: i}
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.iOtp.GenRandomOTP(otpLength)
	o.iOtp.SaveOTPCache(otp)
	message := o.iOtp.GetMessage(otp)
	err := o.iOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

// Default Implementations (Optional hooks)
type DefaultOtp struct{}

func (d *DefaultOtp) GenRandomOTP(len int) string {
	return fmt.Sprintf("%d%d%d%d", 1, 2, 3, 4) // Simplified
}

func (d *DefaultOtp) SaveOTPCache(otp string) {
	fmt.Printf("Saving OTP: %s to cache\n", otp)
}
