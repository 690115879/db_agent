package controller

import (
	"db_agent/pkg/global"
	module "db_agent/pkg/module/ip_reporter_module"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	ipReporterParty := global.Webapp.Party("/ip_reporter")
	m := mvc.New(ipReporterParty)
	m.Handle(new(IpReporterController))
}

type IpReporterRequest struct {
	Ip *string `json:"ip"`
}

type IpReporterResponse struct {
	global.HttpResponse
	Data string `json:"data,omitempty"`
}

type IpReporterController struct {
}

// Get http://localhost:12580/ip_reporter
func (c *IpReporterController) Get() *IpReporterResponse {
	module := module.IpReporterModule{}
	ip := module.GetIP()
	if ip == nil {
		return &IpReporterResponse{
			HttpResponse: global.HttpResponse{
				Code: global.DbError,
				Msg:  "GetIP error ip == nil",
			},
		}
	}
	return &IpReporterResponse{
		HttpResponse: global.HttpResponse{
			Code: global.OK,
			Msg:  global.OKMsg,
		},
		Data: ip.IP,
	}
}

// Post http://localhost:12580/ip_reporter
func (c *IpReporterController) Post(req *IpReporterRequest) *IpReporterResponse {
	if req == nil || req.Ip == nil {
		return &IpReporterResponse{
			HttpResponse: global.HttpResponse{
				Code: global.InputError,
				Msg:  global.InputErrorMsg,
			},
		}
	}
	module := module.IpReporterModule{}
	ip := module.SetIP(*req.Ip)
	if ip == nil {
		return &IpReporterResponse{
			HttpResponse: global.HttpResponse{
				Code: global.DbError,
				Msg:  "SetIp Error ip == nil",
			},
		}
	}

	return &IpReporterResponse{
		HttpResponse: global.HttpResponse{
			Code: global.OK,
			Msg:  global.OKMsg,
		},
		Data: ip.IP,
	}
}
