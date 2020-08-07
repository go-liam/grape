package errorcode

const (
	Success   = 0
	MsSuccess = "OK"

	//system
	System   = 10002
	MsSystem = "系统错误"

	// Parameter
	RequestParameter = 10100
	MsRequest        = "请求参数错误"

	UnActivation   = 10110
	MsUnActivation = "账号没激活"

	// data
	Database       = 11001
	MsDatabase     = "数据库操作不成功"
	DataInfoNull   = 11002
	MsDataInfoNull = "信息不存在"
	DatabaseAdd    = 11003
	MsDatabaseAdd  = "数据库添加操作不成功"
	DatabaseEdit   = 11004
	MsDatabaseEdit = "数据库更新操作不成功"

	// power
	TokenExpired   = 20101
	MsTokenExpired = "utoken失效"
	Refuse         = 20102
	MsRefuse       = "没操作权限"
	GetToken       = 20103
	MsUsGetToken   = "获取utoken错误"

	// http
	HTTPSourceData = 21001

	//http
	Int200 = 200
	Int404 = 404
)
