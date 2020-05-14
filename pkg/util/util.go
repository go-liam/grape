package util

import models "grape/pkg/model"

/*Nil :
 */
var (
	DataListNil = make([]int, 0)           // []
	DataItemNil = map[string]interface{}{} // {}
)

// APIResponseData :
func APIResponseData(code int, message string, data interface{}) *models.APIResponse {
	res := new(models.APIResponse)
	res.Code = code
	if message == "" {
		res.Message = "OK"
	} else {
		res.Message = message
	}
	if data != nil {
		res.Data = data
	} else {
		res.Data = map[string]interface{}{}
	}
	return res
}
