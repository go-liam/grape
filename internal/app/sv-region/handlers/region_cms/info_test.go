package region_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/region"
	"grape/test/mock/region/mock_region"
	"testing"
)

func TestGetInfo(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_region.NewMockService(ctrl)
	back := &region.Model{UserID: 1, CreatedAt: 1}

	m.EXPECT().FindOne(back.UserID).Return(back, nil).AnyTimes()
	e := new(SrvInfo)
	e.srv = m
	e.id = back.UserID
	e.GetInfo()
	assert.EqualValues(t, testdata.ConstWantOne, e.info.UserID)
}
