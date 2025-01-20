package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 车位
// ParkingSpot ParkingSpot
type ParkingSpot struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`
	Address    string    `orm:"column(address)" json:"address" description:"钱包地址"`
	Name       string    `orm:"column(name)" json:"name" description:"名称昵称"`
	Status     int8      `orm:"column(status)" json:"status" description:"上线下线 0:离线 1:在线"`
	Longitude  int64     `orm:"column(longitude)" json:"longitude" description:"经度"`
	Latitude   int64     `orm:"column(latitude)" json:"latitude" description:"纬度"`
	Renter     string    `orm:"column(renter)" json:"renter" description:"租户地址"`
	RentPrice  int64     `orm:"column(rent_price)" json:"rent_price" description:"租金"`
	Content    string    `orm:"column(content)" json:"content" description:"介绍"`
	Remarks    string    `orm:"column(remarks)" json:"remarks" description:"备注"`
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`
	UpdateTime time.Time `orm:"column(update_time)" json:"update_time"`
}

func init() {
	orm.RegisterModel(new(ParkingSpot))
}

// AddParkingSpot insert a new ParkingSpot into database and returns
// last inserted ID on success.
func AddParkingSpot(m *ParkingSpot) (id int64, err error) {
	o := orm.NewOrm()
	v := &ParkingSpot{Name: m.Name}
	if err = o.Read(v, "Name"); err == nil {
		return 0, errors.New("已经存在这个用户")
	}
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	id, err = o.Insert(m)

	// 创建模板 测试使用
	return
}

// GetParkingSpotByNumber retrieves ParkingSpot by number. Returns error if
// number doesn't exist
func GetParkingSpotByName(name string) (v *ParkingSpot, err error) {
	o := orm.NewOrm()
	v = &ParkingSpot{Name: name}
	if err = o.Read(v, "Name"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetParkingSpotByID retrieves ParkingSpot by ID. Returns error if
// ID doesn't exist
func GetParkingSpotByID(id int64) (v *ParkingSpot, err error) {
	o := orm.NewOrm()
	v = &ParkingSpot{ID: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "Comments")
		if err != nil {
			fmt.Println(err)
		}
		return v, nil
	}
	return nil, err
}

// GetAllParkingSpotX retrieves all ParkingSpot matches certain condition. Returns empty list if
// no records exist
func GetAllParkingSpot(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (m map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkingSpot))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []ParkingSpot
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if count, err := qs.Count(); err == nil {
			ml := make([]interface{}, 0)
			if len(fields) == 0 {
				for _, v := range l {
					// _, err = o.LoadRelated(&v, "Comments")
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// for _, m := range v.Comments {
					// 	m.ParkingSpot = nil
					// }
					ml = append(ml, v)
				}
			} else {
				// trim unused fields
				for _, v := range l {
					m := make(map[string]interface{})
					val := reflect.ValueOf(v)
					for _, fname := range fields {
						m[fname] = val.FieldByName(fname).Interface()
					}
					ml = append(ml, m)
				}
			}

			m = make(map[string]interface{})

			m["list"] = ml
			m["total"] = count

			return m, nil
		}
	}
	return nil, err
}

// GetAllParkingSpotList retrieves all ParkingSpot matches certain condition. Returns empty list if
// no records exist
func GetAllParkingSpotList(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (list []ParkingSpot, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkingSpot))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []ParkingSpot
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			// for _, v := range l {
			// 	list = append(list, v)
			// }
			list = append(list, l...)
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
			}
		}
		return list, nil
	}
	return nil, err
}

// UpdateParkingSpotByID updates ParkingSpot by ID and returns error if
// the record to be updated doesn't exist
func UpdateParkingSpotByID(m *ParkingSpot) (err error) {
	o := orm.NewOrm()
	v := ParkingSpot{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.UpdateTime = time.Now()
		var num int64
		if num, err = o.Update(m); err != nil {
			fmt.Println("Number of records updated in database:", num, err)
		}

		PushData(&Information{Type: 0, Data: m, Message: "推送用户状态信息"})
	}
	return
}

// UpdateParkingSpotStatusByID updates ParkingSpot by ID and returns error if
// the record to be updated doesn't exist
func UpdateParkingSpotStatusByNumber(m *ParkingSpot) (v ParkingSpot, err error) {
	o := orm.NewOrm()
	v = ParkingSpot{Name: m.Name}
	// ascertain id exists in the database
	if err = o.Read(&v, "Number"); err == nil {
		v.UpdateTime = time.Now()
		v.Status = m.Status
		var num int64
		if num, err = o.Update(&v, "Status", "UpdateTime"); err != nil {
			fmt.Println("Number of records updated in database:", num, err)
		}
	}
	PushData(&Information{Type: 0, Data: v, Message: "推送用户状态信息"})
	return
}

// DeleteParkingSpot deletes ParkingSpot by ID and returns error if
// the record to be deleted doesn't exist
func DeleteParkingSpot(id int64) (err error) {
	o := orm.NewOrm()
	v := ParkingSpot{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ParkingSpot{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteAllParkingSpot deletes all ParkingSpot
func DeleteAllParkingSpot() (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkingSpot))
	return qs.Filter("id__gt", 0).Delete()
}
