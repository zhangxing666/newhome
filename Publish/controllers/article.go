package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Publish/models"
	"time"
	"path"
)

type ArticleController struct {
	beego.Controller
}
func (this *ArticleController)ShowIndex(){
	o:=orm.NewOrm()
	var articles []models.Article
	qs:=o.QueryTable("Article")
	qs.All(&articles)
    this.Data["articles"]=articles
	this.TplName="index.html"
}
//展示添加文章页面
func (this *ArticleController)ShowAdd(){
	this.TplName="add.html"
}
//添加文章操作
func (this *ArticleController)HandleAdd(){
     //获取页面添加的内容
	articleName:=this.GetString("articleName")
	selectType:=this.GetString("select")
	content:=this.GetString("content")
	file,head,err:=this.GetFile("uploadName")
	defer file.Close()
	if articleName==""||selectType==""||content==""||err!=nil {
		beego.Error("插入数据不能为空")
		this.TplName="add.html"
		return
	}
	if head.Size>5000000000000{
		beego.Error("文件太大，清重新上传")
		this.TplName="add.html"
		return

	}
	ext:=path.Ext(head.Filename)
	if ext!=".jpg"&&ext!="png"{
		beego.Error("图片上传格式有误")
		this.TplName="add.html"
		return
	}
	beego.Info(head.Filename,head.Header)
	o:=orm.NewOrm()
	var article models.Article
	article.ArticleTitle=articleName
	article.ArticleType=selectType
	article.ArticleContent=content
	fileName:=time.Now().Format("20160504150201")
	this.SaveToFile("uploadName","./static/Img"+fileName+ext)
	article.Img="/static/Img"+fileName+ext
	_,err=o.Insert(&article)
	if err!=nil{
		beego.Error(err)
	}
     this.Redirect("/index",302)
}
func (this *ArticleController)ShowContent(){
	id,err:=this.GetInt("id")
	if err!=nil{
		beego.Error("获取文章失败")
		this.TplName="index.html"
		return
	}

	o:=orm.NewOrm()

	var article models.Article
	o.QueryTable("Article").Filter("Id",id).All(&article)
	article.ReadCount++
	o.Update(&article)
	this.Data["article"]=article
	this.TplName="content.html"
}
func (this *ArticleController)ShowUpdate(){
	id,err:=this.GetInt("id")
	if err!=nil{
		beego.Error("获取更新文章失败")
		this.TplName="index.html"
		return
	}

	o:=orm.NewOrm()

	var article models.Article
	o.QueryTable("Article").Filter("Id",id).All(&article)
	this.Data["article"]=article
	this.TplName="update.html"
}
func (this *ArticleController)HandleUpdate(){
	id,_:=this.GetInt("id")
	articleName:=this.GetString("articleName")
	content:=this.GetString("content")
	file,head,err:=this.GetFile("uploadName")

	if articleName==""||content==""||err!=nil {
		beego.Error("插入数据不能为空")
		this.TplName="add.html"
		return
	}
	defer file.Close()
	if head.Size>5000000000000{
		beego.Error("文件太大，清重新上传")
		this.TplName="add.html"
		return

	}
	ext:=path.Ext(head.Filename)
	if ext!=".jpg"&&ext!="png"{
		beego.Error("图片上传格式有误")
		this.TplName="add.html"
		return
	}
	beego.Info(head.Filename,head.Header)
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	err = o.Read(&article)
	//更新
	if err != nil{
		beego.Error("更新数据不存在")
		this.TplName = "update.html"
		return
	}
	article.ArticleTitle=articleName
	article.ArticleContent=content
	fileName:=time.Now().Format("20160504150201")
	this.SaveToFile("uploadName","./static/Img"+fileName+ext)
	article.Img="/static/Img"+fileName+ext
	_,err=o.Update(&article)
	if err!=nil{
		beego.Error("文章更新失败")
	}
	this.Redirect("/index",302)
}
func (this *ArticleController)HandleDelete(){
	id,err:=this.GetInt("id")
	if  err!=nil {
		beego.Error("id获取失败")
		this.TplName="index.html"
		return
	}
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	err=o.Read(&article)
	if err!=nil{
		beego.Error("需要删除的文章不存在")
		this.TplName="index.html"
		return
	}
	o.Delete(&article)
	this.Redirect("/index",302)

}
func (this *ArticleController)ShowAddType(){
	this.TplName="addType.html"
}