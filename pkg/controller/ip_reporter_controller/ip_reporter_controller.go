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
	Code global.HttpResponseCode `json:"code"`
	Msg  string                  `json:"msg"`
	Data string                  `json:"data,omitempty"`
}

type IpReporterController struct {
}

// Get http://localhost:12580/ip_reporter
func (c *IpReporterController) Get() *IpReporterResponse {
	module := module.IpReporterModule{}
	ip := module.GetIP()
	if ip == nil {
		return &IpReporterResponse{
			Code: global.DbError,
			Msg:  "GetIP error ip == nil",
		}
	}
	return &IpReporterResponse{
		Code: global.OK,
		Msg:  string(global.OKMsg),
		Data: ip.IP,
	}
}

// Put http://localhost:12580/ip_reporter
func (c *IpReporterController) Put(req *IpReporterRequest) *IpReporterResponse {
	if req == nil || req.Ip == nil {
		return &IpReporterResponse{
			Code: global.InputError,
			Msg:  string(global.InputErrorMsg),
		}
	}
	module := module.IpReporterModule{}
	ip := module.SetIP(*req.Ip)
	if ip == nil {
		return &IpReporterResponse{
			Code: global.DbError,
			Msg:  "SetIp Error ip == nil",
		}
	}

	return &IpReporterResponse{
		Code: global.OK,
		Msg:  string(global.OKMsg),
		Data: ip.IP,
	}
}
