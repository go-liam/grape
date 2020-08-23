package region_cms

import (
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/region"
	"testing"
)

func TestGetModel(t *testing.T) {
	i := new(ReqModel)
	i.UserID = "1"
	v := GetModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.UserID)
}

func TestGetRespModel(t *testing.T) {
	i := new(region.Model)
	i.RegionID = 1
	v := GetRespModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.RegionID)
}
