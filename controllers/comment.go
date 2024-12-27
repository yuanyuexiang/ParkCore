package controllers

import (
	"encoding/json"
	"errors"
	"park/models"
	"strconv"
	"strings"
)

type CommentController struct {
	BaseController
}

// GetOne GetOne
// @Title Get
// @Summary 获取期刊
// @Description 获取期刊根据ID
// @Param	id		path 	string	true		"期刊ID"
// @Success 200 {object} controllers.ResponseUser
// @router /:id [get]
func (c *CommentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCommentByID(id)
	c.DoReturn(v, err)
}

// GetAll GetAll
// @Title 查询期刊
// @Summary 查询期刊
// @Description 查询期刊 http://127.0.0.1:8080/park/v1/user
// @Param	query	query	string	false	"查询参数，用&分隔"
// @Param	sortby	query	string	false	"排序"
// @Param	order	query	string	false	"升序asc，降序desc"
// @Param	limit	query	string	false	"限制个数，默认10"
// @Param	offset	query	string	false	"偏移，默认0"
// @Success 200 {object} controllers.ResponseData
// @router / [get]
func (c *CommentController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	c.DoReturn(l, err)
}

// Post Post
// @Title Post
// @Summary 添加期刊
// @Description 添加期刊
// @Param	body		body 	models.User	true		"期刊结构体"
// @Success 200 {object} controllers.ResponseUser
// @router / [post]
func (c *CommentController) Post() {
	var v models.Comment
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	id, err := models.AddComment(&v)
	c.DoReturn(id, err)
}

// Put Put
// @Title 更新期刊
// @Summary 更新期刊
// @Description 更新期刊
// @Param	id		path 	string			true		"期刊ID"
// @Param	body	body 	models.User	true		"期刊结构体内容"
// @Success 200 {object} controllers.ResponseData
// @router /:id [put]
func (c *CommentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Comment{ID: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	err := models.UpdateCommentByID(&v)
	c.DoReturn(nil, err)
}

// DeleteList DeleteList
// @Title 批量删除期刊
// @Summary 批量删除期刊
// @Description create User
// @Param	body		body 	[]models.User	true		"只填id字段即可"
// @Success 201 {object} controllers.ResponseData
// @Failure 403 body is empty
// @router / [delete]
func (c *CommentController) DeleteList() {
	var v []models.Comment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	for _, comment := range v {
		err = models.DeleteComment(comment.ID)
	}
	c.DoReturn(nil, err)
}

// Delete Delete
// @Title 删除期刊
// @Summary 删除期刊
// @Description 删除期刊
// @Param	id		path 	string	true		"删除期刊的ID"
// @Success 200 {object} controllers.ResponseData
// @router /:id [delete]
func (c *CommentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	err := models.DeleteComment(id)
	c.DoReturn(nil, err)
}

// GetTest GetTest
// @Title Get
// @Summary 测试
// @Description 测试
// @Success 200 {object} controllers.ResponseData
// @router /test [get]
func (c *CommentController) GetTest() {
	models.Test()
	c.DoReturn(nil, nil)
}
