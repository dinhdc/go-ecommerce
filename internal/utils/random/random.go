package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	otp := 100000 + r.Intn(900000)
	return otp
}
