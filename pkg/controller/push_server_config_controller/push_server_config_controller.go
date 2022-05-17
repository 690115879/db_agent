package controller

import (
	"db_agent/pkg/global"
	module "db_agent/pkg/module/push_server_config_module"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	pushServerConfigParty := global.Webapp.Party("/push_server_config")
	m := mvc.New(pushServerConfigParty)
	m.Handle(new(PushServerConfigController))
}

type PushServerConfigSetConfigReq struct {
	SendKey        string `json:"send_key,omitempty"`
	FtHost         string `json:"ft_host,omitempty"`
	DingdingUrl    string `json:"dingding_url,omitempty"`
	DingdingSecret string `json:"dingding_secret,omitempty"`
}

type PushServerConfigResponse struct {
	global.HttpResponse
	Data string `json:"data,omitempty"`
}

type PushServerConfigController struct {
}

// GET http://localhost:12580/push_server_config
func (c *PushServerConfigController) GetBy(key string) *PushServerConfigResponse {
	m := module.PushServerConfigModule{}
	return &PushServerConfigResponse{
		HttpResponse: global.HttpResponse{
			Code: global.OK,
			Msg:  global.OKMsg,
		},
		Data: m.GetValue(key),
	}
}

// POST http://localhost:12580/push_server_config
func (c *PushServerConfigController) Post(config *PushServerConfigSetConfigReq) *PushServerConfigResponse {
	m := module.PushServerConfigModule{}
	ret := m.SetConfig(&module.PushServerConfig{
		SendKey:        config.SendKey,
		FtHost:         config.FtHost,
		DingdingUrl:    config.DingdingUrl,
		DingdingSecret: config.DingdingSecret,
	})
	if ret != nil {
		return &PushServerConfigResponse{
			HttpResponse: global.HttpResponse{
				Code: global.DbError,
				Msg:  "SetConfig error",
			},
		}
	}

	return &PushServerConfigResponse{
		HttpResponse: global.HttpResponse{
			Code: global.OK,
			Msg:  global.OKMsg,
		},
	}

}
