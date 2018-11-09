package dao

import (
	"github.com/astaxie/beego/orm"
	"microTest/model"
)

type HelloDao struct {
}

func NewHelloDAO() *HelloDao {
	return &HelloDao{}
}

func (r *HelloDao) GetInfo() (*model.Hello, error) {
	c := &model.Hello{}
	o := orm.NewOrm()
	o.Using("default")

	//err := o.QueryTable("").One(c)
	//if err != nil {
	//	if err == orm.ErrNoRows {
	//		return nil, nil
	//	}
	//	return nil, err
	//}
	return c, nil
}
