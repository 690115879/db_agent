package module

import (
	"db_agent/pkg/global"
	"fmt"

	"gorm.io/gorm"
)

/*
CREATE TABLE IP (
 ID INTEGER PRIMARY KEY NOT NULL,
 IP CHAR(50) NOT NULL,
 Created_At TEXT,
 Updated_At TEXT,
 Deleted_At TEXT
 );
*/
type IpReporterModule struct {
}

type IP struct {
	gorm.Model
	IP string
}

func (c *IpReporterModule) GetIP() *IP {
	ip := &IP{}
	if err := global.DataDb.Table("IP").Last(ip).Error; err != nil {
		fmt.Printf("IpReporterModule GetIP Last error: %s", err)
		return nil
	}
	return ip
}

func (c *IpReporterModule) SetIP(ip string) *IP {
	mIP := &IP{}
	curIp := c.GetIP()
	if curIp == nil {
		if err := global.DataDb.Table("IP").
			FirstOrCreate(mIP, IP{IP: ip}).Error; err != nil {
			fmt.Printf("IpReporterModule SetIP DataDb.Table error %s", err)
			return nil
		}
		return mIP
	}

	if curIp.IP != ip {
		if err := global.DataDb.Table("IP").
			Where("ip = ?", curIp.IP).Update("IP", ip).Error; err != nil {
			fmt.Printf("IpReporterModule SetIP Update err: %s", err)
			return nil
		}
		return c.GetIP()
	}

	return curIp
}
