package metric

var Server *SvMetric

func init() {
	Server = new(SvMetric)
	Server.IsOpen = true
	Server.Init()
}
