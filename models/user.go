package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 学术期刊
// User User
type User struct {
	ID                  int64      `orm:"column(id);pk;auto" json:"id"`
	Number              int64      `orm:"column(number)" json:"number"`
	Name                string     `orm:"column(name)" json:"name"`
	Tag                 string     `orm:"column(tag)" json:"tag"`
	Content             string     `orm:"column(content)" json:"content"`
	Status              int8       `orm:"column(status)" json:"status" description:"上线下线 0:离线 1:在线"`
	Remarks             string     `orm:"column(remarks)" json:"remarks" description:"备注"`
	CreateTime          time.Time  `orm:"column(create_time)" json:"create_time"`
	UpdateTime          time.Time  `orm:"column(update_time)" json:"update_time"`
	Comments            []*Comment `orm:"reverse(many)" json:"comments,omitempty" description:"值域数组"`
}

func init() {
	orm.RegisterModel(new(User))
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	if userList, err := GetAllUserList(nil, nil, nil, nil, 0, -1); err == nil {
	// 		for _, v := range userList {
	// 			v.Status = 0
	// 			UpdateUserByID(&v)
	// 		}
	// 	}
	// }()
}

// AddUser insert a new User into database and returns
// last inserted ID on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	v := &User{Name: m.Name}
	if err = o.Read(v, "Name"); err == nil {
		return 0, errors.New("已经存在这个期刊")
	}
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	id, err = o.Insert(m)

	// 创建模板 测试使用
	return
}

// GetUserByNumber retrieves User by number. Returns error if
// number doesn't exist
func GetUserByName(name string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Name: name}
	if err = o.Read(v, "Name"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserByID retrieves User by ID. Returns error if
// ID doesn't exist
func GetUserByID(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{ID: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "Comments")
		if err != nil {
			fmt.Println(err)
		}
		for _, m := range v.Comments {
			m.User = nil
		}
		return v, nil
	}
	return nil, err
}

// GetAllUserX retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (m map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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
					// 	m.User = nil
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

// GetAllUserList retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUserList(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (list []User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

// UpdateUserByID updates User by ID and returns error if
// the record to be updated doesn't exist
func UpdateUserByID(m *User) (err error) {
	o := orm.NewOrm()
	v := User{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.UpdateTime = time.Now()
		var num int64
		if num, err = o.Update(m); err != nil {
			fmt.Println("Number of records updated in database:", num, err)
		}

		PushData(&Information{Type: 0, Data: m, Message: "推送期刊状态信息"})
	}
	return
}

// UpdateUserStatusByID updates User by ID and returns error if
// the record to be updated doesn't exist
func UpdateUserStatusByNumber(m *User) (v User, err error) {
	o := orm.NewOrm()
	v = User{Name: m.Name}
	// ascertain id exists in the database
	if err = o.Read(&v, "Number"); err == nil {
		v.UpdateTime = time.Now()
		v.Status = m.Status
		var num int64
		if num, err = o.Update(&v, "Status", "UpdateTime"); err != nil {
			fmt.Println("Number of records updated in database:", num, err)
		}
	}
	PushData(&Information{Type: 0, Data: v, Message: "推送期刊状态信息"})
	return
}

// DeleteUser deletes User by ID and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteAllUser deletes all User
func DeleteAllUser() (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	return qs.Filter("id__gt", 0).Delete()
}

func Test() {
	if user, err := GetUserByID(1); err == nil {
		user.Status = 1
		// 连接状态
		PushData(&Information{Type: 0, Data: user, Message: "===========连接状态============="})
		// 通知特征数据已更新
		PushData(&Information{Type: 2, Data: nil, Message: "=============通知特征数据已更新==========="})
	}
}
