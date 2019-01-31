package controllers

import ("github.com/astaxie/beego"
"github.com/astaxie/beego/orm"
	"Publish/models"
)
type UserController struct {
	beego.Controller
}
func (this *UserController)ShowRegister(){
	this.TplName="register.html"
}
func (this *UserController)HandleRegister(){
	userName:=this.GetString("userName")
	password:=this.GetString("password")
	beego.Info(userName)
	o:=orm.NewOrm()
	if userName==""||password==""{
		beego.Error("获取用户注册信息失败")
		this.Redirect("/register",302)
		return
	}
	var user models.User
	user.UserName=userName
	user.Password=password
	_,err:=o.Insert(&user)
	if err !=nil {
		beego.Error("用户注册失败")
		this.TplName="register.html"
		return
	}
	this.Redirect("/login",302)
}
func (this *UserController)ShowLogin(){
	this.TplName="login.html"
}
func (this *UserController)HandleLogin(){
	//从客户端获取用户信息    用户名和密码
	userName:=this.GetString("userName")
	password:=this.GetString("password")
	//对获取的信息进行校验
	if userName==""||password==""{
		beego.Error("获取用户数据失败")
		this.TplName="login.html"
		return
	}
	//处理获取到的数据
	o:=orm.NewOrm()
	var user models.User
	user.UserName=userName
	err:=o.Read(&user,"UserName")
	if err!=nil {
		beego.Error("用户名不存在")
		this.TplName="login.html"
	}
	if user.Password!=password{
		beego.Error("用户密码错误")
		this.TplName="login.html"
		return
	}
	this.Redirect("/index",302)
}
//展示页面
