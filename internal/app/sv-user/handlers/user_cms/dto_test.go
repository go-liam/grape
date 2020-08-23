package user_cms

import (
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/account/user"
	"testing"
)

func TestGetModel(t *testing.T) {
	i := new(ReqModel)
	i.ID = "1"
	v := GetModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.ID)
}

func TestGetRespModel(t *testing.T) {
	i := new(user.Model)
	i.Status = 1
	v := GetRespModel(i)
	assert.EqualValues(t, testdata.ConstWantOne, v.Status)
}
