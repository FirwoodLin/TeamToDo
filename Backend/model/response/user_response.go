package response

type UserSignInResponse struct {
	UserID     uint
	UserName   string `json:"userName" `
	UserAvatar string `json:"userAvatar"`
}
type UserQueryResponse struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	IsAdmin   bool   `json:"is_admin"`
	UserID    uint   `json:"id"`
	ViewCount int    `json:"view_count"`
}
