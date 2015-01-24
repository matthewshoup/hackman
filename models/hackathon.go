package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Hackathon struct {
	Id           int
	Name         string `orm:"size(128)"`
	Description  string `orm:"size(128)"`
	Organization string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Hackathon))
}

// AddHackathon insert a new Hackathon into database and returns
// last inserted Id on success.
func AddHackathon(m *Hackathon) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// CreateHackathon checks for an existing hackathon and create one
// if it's not already there by same name.
func CreateHackathon(m *Hackathon) {
  o := orm.NewOrm()
  hackathon := Hackathon{Name: m.Name}

  err := o.Read(&hackathon, "Name")
  if err == orm.ErrNoRows {
    beego.Info("no result found")
    o.Insert(m)
  } else if err == orm.ErrMissPK {
    beego.Info("no primary key found")
  } else {
    beego.Info(hackathon.Name)
  }
}

// GetHackathonById retrieves Hackathon by Id. Returns error if
// Id doesn't exist
func GetHackathonById(id int) (v *Hackathon, err error) {
	o := orm.NewOrm()
	v = &Hackathon{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHackathon retrieves all Hackathon matches certain condition. Returns empty list if
// no records exist
func GetAllHackathon(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Hackathon))
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Hackathon
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
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
		return ml, nil
	}
	return nil, err
}

// UpdateHackathon updates Hackathon by Id and returns error if
// the record to be updated doesn't exist
func UpdateHackathonById(m *Hackathon) (err error) {
	o := orm.NewOrm()
	v := Hackathon{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHackathon deletes Hackathon by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHackathon(id int) (err error) {
	o := orm.NewOrm()
	v := Hackathon{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Hackathon{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
