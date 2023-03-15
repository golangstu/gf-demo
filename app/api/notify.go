package api

import (
	"gf-demo/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type notifyApi struct {
}

var Notify = notifyApi{}

func (*notifyApi) Index(request *ghttp.Request) {
	response.JsonExit(request, 0, "success", nil)
}
