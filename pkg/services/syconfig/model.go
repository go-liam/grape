package syconfig

const (
	StatusDefault = 0
	StatusPublish = 10
	StatusDelete = 44
)

type ConfigItem struct {
	Key string `json:"key"`
	VersionID int64 `json:"versionID"`
	Value string `json:"value"`
	Time int64 `json:"time"`
	Status int `json:"status"`  // 0 默认 , 10 发布 , 44 删除
}


