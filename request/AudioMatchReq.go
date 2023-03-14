package request

type AudioMatchReq struct {
	Uid         int    `json:"uid"`
	ChannelName string `json:"channel_name"`
}
