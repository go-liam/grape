package configclient

import "grape/pkg/service/syconfig"

//type ConfigItem struct {
//	Key string `json:"key"`
//	VersionID int64 `json:"versionID"`
//	Value string `json:"value"`
//	Time int64 `json:"time"`
//}

// Info ("syai",0) 获取最新的配置
func Info(key string, versionID int64) (*syconfig.ConfigItem, error) {
	return syconfig.Info(key, versionID)
}

// List :
func List(key string) ([]*syconfig.ConfigItem, error) {
	return syconfig.List(key)
}
