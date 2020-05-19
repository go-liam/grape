package syconfig



type InConfig interface {
	Add(item *ConfigItem) (int64,error)
	List(key string) ([]*ConfigItem,error)
	Publish(key string, versionID int64) (int,error)
	Undo(key string, versionID int64) (int,error)
	Delete(key string, versionID int64) (int,error)
	Info(key string, versionID int64) (*ConfigItem, error)
}
