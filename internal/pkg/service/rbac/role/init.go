package role

var Server InRole

func init() {
	Server = new(SvRoleMock) // ServerMock
}
