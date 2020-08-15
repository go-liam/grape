package site_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/test/mock/www/mock_site"
	"testing"
)

func TestSite_UpdateStatus(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_site.NewMockService(ctrl)
	ids := []int64{1}

	m.EXPECT().UpdateStatusByIDs(1, ids).Return(int64(1), nil).AnyTimes()
	e := new(SrvUpdateStatus)
	e.srv = m
	e.req = &reqStatus{Status: 1, IDs: []string{"1"}}
	e.UpdateStatus()
	assert.EqualValues(t, testdata.ConstWantOne, e.result)
}
