package user_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/account/user"
	"grape/test/mock/account/mock_user"
	"testing"
)

func TestGetInfo(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_user.NewMockService(ctrl)
	back := &user.Model{ID: 1, CreatedAt: 1}

	m.EXPECT().FindOne(back.ID).Return(back, nil).AnyTimes()
	e := new(SrvInfo)
	e.srv = m
	e.id = back.ID
	e.GetInfo()
	assert.EqualValues(t, testdata.ConstWantOne, e.info.ID)
}
