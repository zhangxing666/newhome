package models

import ("github.com/astaxie/beego/orm"
_"github.com/go-sql-driver/mysql"

	"time"
)
type User struct {
	Id   int
	UserName string
	Password   string
}
type Article struct {
	Id       int
	ArticleTitle  string
	ArticleContent  string
	ArticleType  string
	AddTime     time.Time`orm:"type(datetime);auto_now_add"`
	ReadCount    int
	Img          string

}
type ArticleType struct {
	Id   int
	TypeName   string

}

func init(){
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/Publish?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}
