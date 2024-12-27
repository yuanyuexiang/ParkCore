package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 评论
// Comment
type Comment struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`
	Name       string    `orm:"column(name)" json:"name"`   // 用户名
	Score      int64     `orm:"column(score)" json:"score"` // 评分
	Content    string    `orm:"column(content)" json:"content" description:"评论内容"`
	Image      string    `orm:"column(image)" json:"image" description:"image"`
	CreateTime time.Time `orm:"column(create_time)" json:"create_time"`
	UpdateTime time.Time `orm:"column(update_time)" json:"update_time"`
	User   *User `orm:"rel(fk);on_delete(cascade)" json:"user,omitempty"` // RelForeignKey relation
}

// type Comment struct {
//     User *User`orm:"rel(fk)"` // RelForeignKey relation
// }

func init() {
	orm.RegisterModel(new(Comment))
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
func AddComment(m *Comment) (id int64, err error) {
	o := orm.NewOrm()
	// v := &Comment{Name: m.Name}
	// if err = o.Read(v, "Name"); err == nil {
	// 	return 0, errors.New("已经存在这个期刊")
	// }
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	id, err = o.Insert(m)

	// 创建模板 测试使用
	return
}

// GetUserByNumber retrieves User by number. Returns error if
// number doesn't exist
func GetCommentByName(name string) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{Name: name}
	if err = o.Read(v, "Name"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserByID retrieves User by ID. Returns error if
// ID doesn't exist
func GetCommentByID(id int64) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{ID: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "User")
		v.User.Comments = nil
		if err != nil {
			fmt.Println(err)
		}
		return v, nil
	}
	return nil, err
}

// GetAllUserX retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (m map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comment))
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

	var l []Comment
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if count, err := qs.Count(); err == nil {
			ml := make([]interface{}, 0)
			if len(fields) == 0 {
				for _, v := range l {
					if _, err = o.LoadRelated(&v, "User"); err == nil {
						v.User.Comments = nil
					}
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
func GetAllCommentList(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (list []Comment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comment))
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

	var l []Comment
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
func UpdateCommentByID(m *Comment) (err error) {
	o := orm.NewOrm()
	v := Comment{ID: m.ID}
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

// DeleteUser deletes User by ID and returns error if
// the record to be deleted doesn't exist
func DeleteComment(id int64) (err error) {
	o := orm.NewOrm()
	v := Comment{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Comment{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteAllUser deletes all User
func DeleteAllComment() (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comment))
	return qs.Filter("id__gt", 0).Delete()
}
