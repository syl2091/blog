package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：lege2091"
	c.Data["qq"] = "QQ：164378335"
	c.Data["tel"] = "Tel：***"
	c.TplName = "aboultme.html"
}
