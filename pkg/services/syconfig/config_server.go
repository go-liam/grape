package syconfig

var sv InConfig

func init() {
	sv = new(SvMock)
}

func Info(key string, versionID int64) (*ConfigItem, error)  {
	return sv.Info(key,versionID)
}

func Add(item *ConfigItem) (int64, error) {
	return sv.Add(item)
}

// 最新的 50条
func List(key string) ([]*ConfigItem, error) {
	return sv.List(key)
}

// Publish :发布
func Publish(key string, versionID int64) (int, error) {
	return sv.Publish(key, versionID)
}

// Undo :取消发布
func Undo(key string, versionID int64) (int, error) {
	return sv.Undo(key, versionID)
}

func Delete(key string, versionID int64) (int, error) {
	return sv.Delete(key, versionID)
}
