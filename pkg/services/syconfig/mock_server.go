package syconfig

var ServerMock *SvMock

func init() {
	ServerMock = new(SvMock)
	ServerMock.dataInit()
}

type SvMock struct {
	list    []*ConfigItem
	hadData bool
}

func (sv *SvMock) Info(key string, versionID int64) (*ConfigItem, error) {
	sv.dataInit()
	for _, v := range sv.list {
		if (v.Key == key && v.VersionID == versionID) || (v.Key == key && 0 == versionID) {
			return v, nil
		}
	}
	return new(ConfigItem), nil
}

func (sv *SvMock) Add(item *ConfigItem) (int64, error) {
	sv.dataInit()
	sv.list = append(sv.list, item)
	return 1, nil
}

// List 最新的 50条
func (sv *SvMock) List(key string) ([]*ConfigItem, error) {
	sv.dataInit()
	return sv.list, nil
}

// Publish :发布
func (sv *SvMock) Publish(key string, versionID int64) (int, error) {
	sv.dataInit()
	for _, v := range sv.list {
		if v.Key == key && v.VersionID == versionID {
			v.Status = StatusPublish
		}
	}
	return 1, nil
}

// Undo :取消发布
func (sv *SvMock) Undo(key string, versionID int64) (int, error) {
	sv.dataInit()
	for _, v := range sv.list {
		if v.Key == key && v.VersionID == versionID {
			v.Status = StatusDefault
		}
	}
	return 1, nil
}

func (sv *SvMock) Delete(key string, versionID int64) (int, error) {
	sv.dataInit()
	ls := make([]*ConfigItem, 0)
	for _, v := range sv.list {
		if v.Key == key && v.VersionID == versionID {
			continue
		}
		ls = append(ls, v)
	}
	sv.list = ls
	return 1, nil
}

func (sv *SvMock) dataInit() {
	if !sv.hadData {
		sv.hadData = true
		sv.list = make([]*ConfigItem, 0)
		sv.list = append(sv.list, &ConfigItem{Key: "sv-ai", Value: "box-add1", VersionID: 2020051201, Time: 1589338667})
		sv.list = append(sv.list, &ConfigItem{Key: "sv-ai", Value: "box-add2", VersionID: 2020051202, Time: 1589338668})
		sv.list = append(sv.list, &ConfigItem{Key: "sv-ai", Value: "box-add3", VersionID: 2020051203, Time: 1589338669})
		sv.list = append(sv.list, &ConfigItem{Key: "sv-oss", Value: "box-add4", VersionID: 2020051301, Time: 1589338661})
	}
}
