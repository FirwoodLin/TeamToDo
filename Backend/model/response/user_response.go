package response

type UserResponse struct {
	// 比较通用的结构体，用户信息三要素
	// 场景：登陆/注册/更新/查询 用户信息后返回
	UserID     uint   `json:"userID"`
	UserName   string `json:"userName" `
	UserAvatar string `json:"userAvatar"`
}
type UserQueryResponse struct {
	// 查询自己的个人信息用
	UserID     uint   `json:"userID"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	UserAvatar string `json:"userAvatar"`
}
