package utils

import (
	"math/rand"
	"time"
)

// GenerateInviteCode 生成邀请码（6位）
func GenerateInviteCode() string {
	const codeLength = 6
	const codeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//rand.Seed(time.Now().UnixNano())
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		code[i] = codeChars[rng.Intn(len(codeChars))]
	}
	return string(code)
}
