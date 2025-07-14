package totp

import (
	"fmt"
	"time"

	hotp "github.com/binary141/hotp-go"
)

type Totp struct {
	t0       int64
	timeStep int64
	hotp     hotp.Hotp
}

func (totp Totp) GenerateOtpAuth() string {
	params := totp.hotp.GenerateOtpAuthParams()
	return fmt.Sprintf("otpauth://totp/%s", params)
}

func CreateTotp(secret string, digits int, t0 int64, timeStep int64, label string) Totp {
	// func CreateHotp(secret string, counter uint64, digits int) Hotp {
	// we get this when calling Calculate as the counter is based on the time in which you need to validate
	var counter uint64 = 0
	otp := hotp.CreateHotp(secret, counter, digits, label)

	return Totp{
		t0:       t0,
		timeStep: timeStep,
		hotp:     otp,
	}

}

func (totp *Totp) SetHasher(hashFunc hotp.HashFunc) error {
	return totp.hotp.SetHashFunc(hashFunc)

}

func (totp Totp) Calculate() (string, error) {
	totp.hotp.SetCounter(totp.getCounterFromCurrentTime())

	return totp.hotp.Calculate()
}

func (totp Totp) getCounterFromCurrentTime() uint64 {
	return getCounterFromTime(time.Now(), totp.t0, totp.timeStep)
}

func getCounterFromTime(t time.Time, t0 int64, timeStep int64) uint64 {
	t1 := t.UTC()

	startOfMinute := t1.Truncate(time.Minute)

	timeStepOffset := startOfMinute.Add(time.Duration(timeStep) * time.Second)

	timeToUse := startOfMinute.Unix()

	// if the current time is past that of the time step,
	// we need to "slide" the time over and use that for the unix time
	if t1.Unix() >= (timeStepOffset.Unix()) {
		timeToUse = timeStepOffset.Unix()
	}

	counter := uint64((timeToUse - t0) / timeStep)

	return counter
}
