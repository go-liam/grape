package notice_api

import (
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/home_site/notice"
	"testing"
)

func TestGetModel(t *testing.T) {
	i := new(ReqModel)
	i.ID = "1"
	v := GetModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.ID)
}

func TestGetRespModel(t *testing.T) {
	i := new(notice.Model)
	i.Status = 1
	v := GetRespModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.Status)
}
