package models

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id                int64
	Title             string
	CreateTime        time.Time `orm:"index"`
	Views             int64     `orm:"index"`
	TopicTime         time.Time `orm:"index"`
	TopicCount        int64
	TopicLasterUserId int64
}

// 评论
type Comment struct {
	Id         int64
	Tid        int64
	Name       string
	Content    string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Lables          string
	Category        string
	Content         string `orm:"size(5000)"`
	Attachment      string
	CreateTime      time.Time `orm:"index"`
	UpdateTime      time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"Index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Person struct {
	User_Id  int64
	Username string
	Password string
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic), new(Person), new(Comment))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_url"), 30)
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{
		Title:      name,
		CreateTime: time.Now(),
		TopicTime:  time.Now(),
	}

	// 查询数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	/*
		ParseInt func(s string, base int, bitSize int) (i int64, err error)
			返回字符串表示的整数值，接受正负号。

			base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；

			bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

			返回的err是*NumErr类型的，
			如果语法有误，err.Error = ErrSyntax；
			如果结果超出类型范围err.Error = ErrRange。
	*/
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	//第一个返回值为影响的行数
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)
	//ORM 以 QuerySeter 来组织查询，每个返回 QuerySeter 的方法都会获得一个新的 QuerySeter 对象。
	qs := o.QueryTable("category")
	//返回对应的结果集对象
	// All 的参数支持 *[]Type 和 *[]*Type 两种形式的 slice
	_, err := qs.All(&cates)
	return cates, err
}

//查询所有文章列表
func GetAllTopics(cate string, label string, isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()

	topics = make([]*Topic, 0)

	qs := o.QueryTable("topic")
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels_contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-create_time").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

//添加文章
func AddTopic(title, category, label, content, attachment string) error {
	//处理标签 strings.Join将slice拼接程字符串，第二个参数是其拼接符
	lable := "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	//先获取一个orm对象
	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Lables:     lable,
		Content:    content,
		Attachment: attachment,
		CreateTime: time.Now(),
		Category:   category,
		UpdateTime: time.Now(),
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}
	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		//如果err不存在，则更新分类文章数,只当分类存在时进行更新
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return err
}

//根据ID查询文章
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)
	topic.Lables = strings.Replace(strings.Replace(
		topic.Lables, "#", " ", -1), "$", "", -1)

	return topic, nil
}

//根据ID修改文章
func ModifyTopic(tid, title, category, label, content, attachment string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	//定义旧的分类名称
	var oldCate, oldAttachment string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttachment = topic.Attachment
		topic.Title = title
		topic.Lables = label
		topic.Content = content
		topic.Attachment = attachment
		topic.Category = category
		topic.UpdateTime = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}
	//更新分类统计
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		//根据旧的分类名称查询出分类然后此类文章数-1
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	// 删除旧的附件
	if len(oldAttachment) > 0 {
		os.Remove(path.Join("attachment", oldAttachment))
	}
	//新的分类名称文章数加1并更新
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return nil
}

//根据id删除文章
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	//一般可以创建一个全局的orm对象，注意资源共享
	o := orm.NewOrm()

	//根据id读取文章
	var oldCate string
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		//拿出文章原来分类
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}
	//原来文章如果存在分类，查出原有分类，文章分类数-1并更新
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", topic.Category).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return err
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:        tidNum,
		Name:       nickname,
		Content:    content,
		CreateTime: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	var tidNum int64
	reply := &Comment{Id: ridNum}
	//如果读取成功则删除评论,赋值文章Tid给tidNum
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	//根据文章Id查询所有评论并按照创建时间降序排列
	_, err = qs.Filter("tid", tidNum).OrderBy("-create_time").All(&replies)
	if err != nil {
		return err
	}

	//查询文章的最近回复时间和总的回复数
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].CreateTime
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}
	return err
}
