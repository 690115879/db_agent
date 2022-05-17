package module

import (
	"db_agent/pkg/global"
	"fmt"

	"gorm.io/gorm"
)

/*
CREATE TABLE PUSH_SERVER_CONFIG (
 ID INTEGER PRIMARY KEY NOT NULL,
 Send_Key CHAR(50) NOT NULL,
 Ft_Host CHAR(100) NOT NULL,
 Dingding_Url CHAR(150) NOT NULL,
 Dingding_Secret CHAR(100) NOT NULL,
 Created_At TEXT,
 Updated_At TEXT,
 Deleted_At TEXT
 );
*/
type PushServerConfigModule struct {
}

type PushServerConfig struct {
	gorm.Model
	SendKey        string
	FtHost         string
	DingdingUrl    string
	DingdingSecret string
}

func (c *PushServerConfigModule) GetValue(key string) string {
	var value string
	err := global.DataDb.Table("Push_Server_Config").Model(&PushServerConfig{}).Select(key).Last(&value).Error
	if err != nil {
		fmt.Printf("PushServerConfigModule GetValue key=%s, err: %s", key, err)
	}
	return value
}

func (c *PushServerConfigModule) SetConfig(config *PushServerConfig) *PushServerConfig {
	if config == nil {
		return nil
	}

	lastConf := &PushServerConfig{}
	err := global.DataDb.Table("Push_Server_Config").Last(&lastConf).Error
	if err == nil {
		if config.SendKey != "" {
			lastConf.SendKey = config.SendKey
		}

		if config.FtHost != "" {
			lastConf.FtHost = config.FtHost
		}

		if config.DingdingUrl != "" {
			lastConf.DingdingUrl = config.DingdingUrl
		}

		if config.DingdingSecret != "" {
			lastConf.DingdingSecret = config.DingdingSecret
		}

		lastConf.ID += 1
		err := global.DataDb.Table("Push_Server_Config").FirstOrCreate(lastConf).Error
		if err != nil {
			fmt.Printf("PushServerConfigModule SetConfig FirstOrCreate error: %s", err)
			return nil
		}
		return lastConf
	} else {
		if config.SendKey == "" || config.FtHost == "" || config.DingdingUrl == "" ||
			config.DingdingSecret == "" {
			return nil
		}
		err = global.DataDb.Table("Push_Server_Config").Create(config).Error
		if err != nil {
			fmt.Printf("PushServerConfigModule SetConfig Create error: %s", err)
			return nil
		}
		return config
	}

}
