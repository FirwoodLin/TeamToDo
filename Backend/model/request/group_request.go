package request

type GroupCreateRequest struct {
	Description string `json:"description"`
	GroupName   string `json:"groupName"` // 群组名称
}
