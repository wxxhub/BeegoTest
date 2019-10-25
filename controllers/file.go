package controllers

import (
	"github.com/astaxie/beego"
)

// FileController .
type FileController struct {
	beego.Controller
}

// Get .
func (f *FileController) Get() {
	f.TplName = "file.html"
}

// Post .
func (f *FileController) Post() {
	sub := f.GetString("sub")

	if sub == "上传" {
		// f.Ctx.WriteString("你点的上传按钮")
		file, head, err := f.GetFile("file")

		if err != nil {
			f.Ctx.WriteString("获取文件失败")
			return
		}

		defer file.Close()

		fileName := head.Filename

		err = f.SaveToFile("file", "upload/"+fileName)

		if err != nil {
			f.Ctx.WriteString("上传失败!")
		} else {
			f.Ctx.WriteString("上传成功!")
		}
	} else {
		// f.Ctx.WriteString("你点的下载按钮")
		f.Ctx.Output.Download("static/img/test.jpeg", "test.jpeg")
	}
	
}