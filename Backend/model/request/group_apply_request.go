package request

import "time"

type InviteType int

const (
	InviteCode InviteType = iota
	InviteLink
)

// GroupCodeRequest 创建邀请码请求
type GroupCodeRequest struct {
	GroupID   uint       `json:"-"`
	InviterID uint       `json:"-"`
	Code      string     `json:"-"`
	Type      InviteType `json:"-"`
	ExpireAt  time.Time  `json:"expireAt"`
}
