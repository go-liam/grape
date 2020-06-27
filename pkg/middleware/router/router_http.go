package router

import "net/http"

func AccessControlAllowHTTP(res http.ResponseWriter, req *http.Request) bool {
	res.Header().Add("Access-Control-Allow-Headers", "access-control-allow-headers,access-control-allow-methods,access-control-allow-origin,cache-control,content-type,utoken,tokenjwt,x-ijt,Token")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,PATCH")
	if req.Method == "OPTIONS" {
		res.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}
