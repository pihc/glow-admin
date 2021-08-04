package api

import (
	"glow-admin/library/respond"
	"io/ioutil"

	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct {
}

func (h *helloApi) Native(r *ghttp.Request) {
	r.Response.Writeln("helloApi world")
}

func (h *helloApi) String(r *ghttp.Request) string {
	return "helloApi world!"
}

func (h *helloApi) Json(r *ghttp.Request) respond.Json {
	return "this's json api"
}

func (h *helloApi) Xml(r *ghttp.Request) respond.XML {
	return "<a>123123</a>"
}
func (h *helloApi) File(r *ghttp.Request) respond.File {
	data, _ := ioutil.ReadFile("./go.mod")
	return respond.File{
		Data:        data,
		ContentType: "text/plain",
		FileName:    "go.mod",
	}
}
