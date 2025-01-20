package controllers

import (
	"encoding/json"
	"errors"
	"park/models"
	"strconv"
	"strings"
)

type ParkingSpotController struct {
	BaseController
}

// GetOne GetOne
// @Title Get
// @Summary 获取车位
// @Description 获取车位根据ID
// @Param	id		path 	string	true		"车位ID"
// @Success 200 {object} controllers.ResponseData
// @router /:id [get]
func (c *ParkingSpotController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetParkingSpotByID(id)
	c.DoReturn(v, err)
}

// GetAll GetAll
// @Title 查询车位
// @Summary 查询车位
// @Description 查询车位 http://127.0.0.1:8080/park/v1/parking-spot
// @Param	query	query	string	false	"查询参数，用&分隔"
// @Param	sortby	query	string	false	"排序"
// @Param	order	query	string	false	"升序asc，降序desc"
// @Param	limit	query	string	false	"限制个数，默认10"
// @Param	offset	query	string	false	"偏移，默认0"
// @Success 200 {object} controllers.ResponseData
// @router / [get]
func (c *ParkingSpotController) GetAll() {
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
	l, err := models.GetAllParkingSpot(query, fields, sortby, order, offset, limit)
	c.DoReturn(l, err)
}

// Post Post
// @Title Post
// @Summary 绑定车位
// @Description 绑定车位
// @Param	body		body 	models.ParkingSpot	true		"车位结构体"
// @Success 200 {object} controllers.ResponseData
// @router / [post]
func (c *ParkingSpotController) Post() {
	var v models.ParkingSpot
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	id, err := models.AddParkingSpot(&v)
	c.DoReturn(id, err)
}

// Put Put
// @Title 更新车位
// @Summary 更新车位
// @Description 更新车位
// @Param	id		path 	string			true		"车位ID"
// @Param	body	body 	models.ParkingSpot	true		"车位结构体内容"
// @Success 200 {object} controllers.ResponseData
// @router /:id [put]
func (c *ParkingSpotController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.ParkingSpot{ID: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	err := models.UpdateParkingSpotByID(&v)
	c.DoReturn(nil, err)
}

// DeleteList DeleteList
// @Title 批量删除车位
// @Summary 批量删除车位
// @Description 批量删除车位
// @Param	body		body 	[]models.ParkingSpot	true		"只填id字段即可"
// @Success 201 {object} controllers.ResponseData
// @Failure 403 body is empty
// @router / [delete]
func (c *ParkingSpotController) DeleteList() {
	var v []models.ParkingSpot
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	for _, parking_spot := range v {
		err = models.DeleteParkingSpot(parking_spot.ID)
	}
	c.DoReturn(nil, err)
}

// Delete Delete
// @Title 删除车位
// @Summary 删除车位
// @Description 删除车位
// @Param	id		path 	string	true		"删除车位的ID"
// @Success 200 {object} controllers.ResponseData
// @router /:id [delete]
func (c *ParkingSpotController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	err := models.DeleteParkingSpot(id)
	c.DoReturn(nil, err)
}
