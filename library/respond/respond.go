package respond

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"sync"
)

const RspBodyKey = "gallop_rsp_key"

var responderList []Responder
var once_resp_list sync.Once

func get_responder_list() []Responder {
	once_resp_list.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
			(XMLResponder)(nil),
			(FileResponder)(nil),
			(NativeResponder)(nil),
		}
	})
	return responderList
}
func Convert(handler interface{}) ghttp.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range get_responder_list() {
		r_ref := reflect.TypeOf(r)
		if h_ref.Type().ConvertibleTo(r_ref) {
			return h_ref.Convert(r_ref).Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type Responder interface {
	RespondTo() ghttp.HandlerFunc
}
type NativeResponder func(r *ghttp.Request)

func (s NativeResponder) RespondTo() ghttp.HandlerFunc {
	return func(request *ghttp.Request) {
		s(request)
	}
}

type StringResponder func(request *ghttp.Request) string

func (s StringResponder) RespondTo() ghttp.HandlerFunc {
	return func(request *ghttp.Request) {
		body := s(request)
		request.SetParam(RspBodyKey, body)
		request.Response.Writeln(body)
	}
}

type Json interface{}
type JsonResponder func(request *ghttp.Request) Json

func (j JsonResponder) RespondTo() ghttp.HandlerFunc {
	return func(req *ghttp.Request) {
		body := j(req)
		req.SetParam(RspBodyKey, body)
		req.Response.WriteJson(body)
	}
}

type XML interface{}

type XMLResponder func(request *ghttp.Request) XML

func (s XMLResponder) RespondTo() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		body := s(r).(string)
		r.SetParam(RspBodyKey, body)
		r.Response.WriteXml(body)
	}
}

type File struct {
	Data        []byte
	ContentType string
	FileName    string
}

type FileResponder func(request *ghttp.Request) File

func (f FileResponder) RespondTo() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		file := f(r)
		r.Response.SetBuffer(file.Data)
		r.Response.Header().Set("Content-Length", gconv.String(r.Response.BufferLength()))
		r.Response.Header().Set("Content-Type", file.ContentType)
		r.Response.Header().Set("Accept-Ranges", "bytes")
		r.Response.Header().Set("Content-Transfer-Encoding", "binary")
		r.Response.Header().Set("Content-Disposition", "attachment; filename="+file.FileName)
		r.Response.Buffer()
	}
}
