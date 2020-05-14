package user

var Server InUser

func init()  {
	Server = new(SvUserMock) // ServerMock
}
