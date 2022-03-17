package dto

type DashProStatusOutput struct {
	Id        int `json:"id" gorm:"primary_key"`
	OneStatus int `json:"one_status" form:"one_status"`
	TwoStatus int `json:"two_status" form:"two_status"`
	ThrStatus int `json:"thr_status" form:"thr_status"`
	FouStatus int `json:"fou_status" form:"fou_status"`
	FivStatus int `json:"fiv_status" form:"fiv_status"`
}
