package response

type UserBasicInfoResp struct {
	ID       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}
