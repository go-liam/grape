package core

import "grape/internal/pkg/config"

type SiteItem struct {
	Path string
	Site string
}

var list []*config.SiteItem

//var httpSt = "http://"
//var serverIP = "192.168.31.235"

func DateInit() {
	list = config.GateWayDateInit()
}
