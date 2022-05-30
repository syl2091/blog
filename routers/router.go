package routers

import (
	"blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	//退出
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	//标签
	beego.Router("/tags", &controllers.TagsController{})
	//相册
	beego.Router("/album", &controllers.AlbumController{})
	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
	//关于我
	beego.Router("/aboutme", &controllers.AboutMeController{})
}
