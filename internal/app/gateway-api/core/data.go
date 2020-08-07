package core

import (
	"grape/configs"
)

type SiteItem struct {
	Path string
	Site string
}

var list []*configs.SiteItem

//var httpSt = "http://"
//var serverIP = "192.168.31.235"

func DateInit() {
	list = configs.GateWayDateInit()
}
