package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// BaseController BaseController
type BaseController struct {
	beego.Controller
}

// ResponseData ResponseData
type ResponseData struct {
	Status  int         `json:"status" description:"状态 1:成功,-1:失败"`
	Message string      `json:"message" description:"附加信息"`
	Data    interface{} `json:"data,omitempty" description:"数据"`
}

// GetSuccessfulResponseData GetSuccessfulResponseData
func GetSuccessfulResponseData(data interface{}) (rd *ResponseData) {
	rd = &ResponseData{1, "OK", data}
	return
}

// GetSuccessfulResponseDataX GetSuccessfulResponseData
func GetSuccessfulResponseDataX(message string, data interface{}) (rd *ResponseData) {
	rd = &ResponseData{1, message, data}
	return
}

// GetFailedResponseData GetFailedResponseData
func GetFailedResponseData(errorMessage string) (rd *ResponseData) {
	rd = &ResponseData{-1, errorMessage, nil}
	return
}

// GetFailedResponseDataForTokenError GetFailedResponseDataForTokenError
func GetFailedResponseDataForTokenError(errorMessage string) (rd *ResponseData) {
	rd = &ResponseData{-2, errorMessage, nil}
	return
}

// DoReturn DoReturn
func (c *BaseController) DoReturn(l interface{}, err error) {
	if err != nil {
		c.Data["json"] = GetFailedResponseData(err.Error())
	} else {
		c.Data["json"] = GetSuccessfulResponseData(l)
	}
	c.ServeJSON()
}

// DoReturn DoReturn
func (c *BaseController) DoReturnX(status int, message string, data any) {
	c.Data["json"] = &ResponseData{status, message, data}
	c.ServeJSON()
}

// DoReturn DoReturn
func (c *BaseController) DoReturnZ(result any) {
	c.Data["json"] = &result
	c.ServeJSON()
}
