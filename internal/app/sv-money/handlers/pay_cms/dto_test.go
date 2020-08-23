package pay_cms

import (
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/money/pay"
	"testing"
)

func TestGetModel(t *testing.T) {
	i := new(ReqModel)
	i.ID = "1"
	v := GetModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.ID)
}

func TestGetRespModel(t *testing.T) {
	i := new(pay.Model)
	i.Status = 1
	v := GetRespModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.Status)
}
