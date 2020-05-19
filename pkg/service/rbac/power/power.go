package power

import "strings"

// CheckPowerByTag :
func CheckPowerByTag(tag string, list []*ModelPower) bool {
	for _, v := range list {
		if strings.Compare(v.PowerTag, tag) == 0 {
			return true
		}
	}
	return false
}

// GetPowerIDByTag :取 powerID
func GetPowerIDByTag(tag string, list []*ModelPower) int {
	for _, v := range list {
		if strings.Compare(v.PowerTag, tag) == 0 {
			return v.PowerID
		}
	}
	return 0
}
