package model

type OrderPageListReq struct {
	Username string `json:"user_name"`
	PageNo   uint   `json:"page_no"`
	Size     uint   `json:"size"`
}
