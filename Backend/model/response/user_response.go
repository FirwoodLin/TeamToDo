package response

type UserResponse struct {
	// 比较通用的结构体，用户信息三要素
	// 场景：登陆/注册/更新/查询 用户信息后返回
	UserID     uint
	UserName   string `json:"userName" `
	UserAvatar string `json:"userAvatar"`
}
type UserQueryResponse struct {
	UserID uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	//ViewCount int    `json:"view_count"`
	//Tel       string `json:"tel"`
	//IsAdmin   bool   `json:"is_admin"`
	//Nickname  string `json:"nickname"`
}
