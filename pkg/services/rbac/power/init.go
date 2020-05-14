package power

var Server InPower

func init()  {
	Server = new(SvPowerMock)
}
