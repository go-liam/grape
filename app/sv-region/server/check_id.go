package server

import "grape/pkg/config"

func IpCheck(ip string) (int,string) {
	if ip == "127.0.0.1"{
		return config.RegionIDLocal,config.RegionHostLocal
	}
	if ip == "192.168.0.1"{
		return config.RegionIDTest,config.RegionHostTest
	}
	return config.RegionIDDefault,config.RegionHostDefault
}

func UserIDCheck(uid int64) (int,string) {
	if uid  < 100{
		return config.RegionIDLocal,config.RegionHostLocal
	}else if uid < 200 {
		return config.RegionIDTest,config.RegionHostTest
	}
	return config.RegionIDDefault,config.RegionHostDefault
}
