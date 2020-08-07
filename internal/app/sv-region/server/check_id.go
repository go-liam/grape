package server

import (
	"grape/configs"
)

func IpCheck(ip string) (int, string) {
	if ip == "127.0.0.1" {
		return configs.RegionIDLocal, configs.RegionHostLocal
	}
	if ip == "192.168.0.1" {
		return configs.RegionIDTest, configs.RegionHostTest
	}
	return configs.RegionIDDefault, configs.RegionHostDefault
}

func UserIDCheck(uid int64) (int, string) {
	if uid < 100 {
		return configs.RegionIDLocal, configs.RegionHostLocal
	} else if uid < 200 {
		return configs.RegionIDTest, configs.RegionHostTest
	}
	return configs.RegionIDDefault, configs.RegionHostDefault
}
