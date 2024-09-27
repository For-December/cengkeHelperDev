package models

import "cengkeHelperDev/src/dbmodels"

type RespData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewBadResp(msg string) RespData {
	return RespData{
		Code: 400,
		Data: nil,
		Msg:  msg,
	}
}

type DividePageWrapper[
	T dbmodels.PostRecord] struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	List     []T `json:"list"`
	Total    int `json:"total"`
}
