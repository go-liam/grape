package config

type SiteItem struct {
	Path string
	Site string
}

//var ServerList []*SiteItem
//var httpSt = "http://"
//var serverIP = "192.168.31.235"

func GateWayDateInit() []*SiteItem {
	var httpSt = "http://"
	var ServerList []*SiteItem
	ServerList = make([]*SiteItem, 0)
	ServerList = append(ServerList, &SiteItem{Path: ServerAPIApp.Name, Site: httpSt + ServerAPIApp.Host + ServerAPIApp.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerAPICMS.Name, Site: httpSt + ServerAPICMS.Host + ServerAPICMS.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerAPIWWW.Name, Site: httpSt + ServerAPIWWW.Host + ServerAPIWWW.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerAPIWX.Name, Site: httpSt + ServerAPIWX.Host + ServerAPIWX.Port})

	ServerList = append(ServerList, &SiteItem{Path: ServerAuthJWT.Name, Site: httpSt + ServerAuthJWT.Host + ServerAuthJWT.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerAuthRBAC.Name, Site: httpSt + ServerAuthRBAC.Host + ServerAuthRBAC.Port})

	ServerList = append(ServerList, &SiteItem{Path: ServerSvAI.Name, Site: httpSt + ServerSvAI.Host + ServerSvAI.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvConfig.Name, Site: httpSt + ServerSvConfig.Host + ServerSvConfig.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvLog.Name, Site: httpSt + ServerSvLog.Host + ServerSvLog.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvMQ.Name, Site: httpSt + ServerSvMQ.Host + ServerSvMQ.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvNotice.Name, Site: httpSt + ServerSvNotice.Host + ServerSvNotice.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvOSS.Name, Site: httpSt + ServerSvOSS.Host + ServerSvOSS.Port})

	ServerList = append(ServerList, &SiteItem{Path: ServerSvUser.Name, Site: httpSt + ServerSvUser.Host + ServerSvUser.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvMoney.Name, Site: httpSt + ServerSvMoney.Host + ServerSvMoney.Port})
	ServerList = append(ServerList, &SiteItem{Path: ServerSvRegion.Name, Site: httpSt + ServerSvRegion.Host + ServerSvRegion.Port})
	return ServerList
}
