package utils

import (
	"TeamToDo/model/request"
	"math/rand"
	"time"
)

// GenerateInviteCode 生成邀请码（6位）
// 长一点的邀请码接在link后面就是邀请链接
func GenerateInviteCode(inviteType request.InviteType) string {
	const codeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var codeLength int
	if inviteType == request.InviteCode {
		codeLength = 6
	} else {
		codeLength = 30
	}
	//rand.Seed(time.Now().UnixNano())
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		code[i] = codeChars[rng.Intn(len(codeChars))]
	}
	return string(code)
}
