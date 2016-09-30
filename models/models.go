package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_NAME      = "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8"
	MYSQL_DRIVER = "mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Views           int64
	Updated         time.Time `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {

	orm.RegisterModel(new(Category), new(Topic))

	orm.RegisterDriver(MYSQL_DRIVER, orm.DRMySQL)
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", MYSQL_DRIVER, DB_NAME, maxIdle, maxConn)

}

func AddCategory(name string) error {

	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	//不存在
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	return err
}

func GetAllCategory() ([]*Category, error) {

	o := orm.NewOrm()

	cate := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cate)
	return cate, err
}

func DelCategory(id int64) error {

	o := orm.NewOrm()
	//_, err := o.QueryTable("category").Filter("id", id).Delete()

	cate := &Category{Id: id}
	_, err := o.Delete(cate)

	return err
}

func AddTopic(title, content string) error {

	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

func GetAllTopic(isOrder bool) ([]*Topic, error) {

	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isOrder {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err

}

func GetTopic(id int64) (*Topic, error) {

	o := orm.NewOrm()

	topic := &Topic{}
	qs := o.QueryTable("topic")
	err := qs.Filter("id", id).One(topic)
	if err != nil {
		return nil, err
	}

	_, err = qs.Update(orm.Params{"views": orm.ColValue(orm.ColAdd, 1)})
	return topic, err
}
