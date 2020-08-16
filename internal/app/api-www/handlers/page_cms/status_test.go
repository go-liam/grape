package page_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/errorcode"
	"grape/configs/testdata"
	"grape/test/mock/www/mock_page"
	"testing"
)

func TestSite_UpdateStatus(t *testing.T) {
	e := new(SrvUpdateStatus)

	t.Run(testdata.ConstFail, func(t *testing.T) {
		e.req = new(reqStatus)
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.RequestParameter, got.Code)
	})

	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	ids := []int64{1}
	e.req = &reqStatus{Status: 1, IDs: []string{"1"}}

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		m := mock_page.NewMockService(ctrl)
		m.EXPECT().UpdateStatusByIDs(1, ids).Return(int64(1), nil).AnyTimes()
		e.srv = m
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.Success, got.Code)
	})

	t.Run(testdata.ConstFail, func(t *testing.T) {
		m := mock_page.NewMockService(ctrl)
		m.EXPECT().UpdateStatusByIDs(1, ids).Return(int64(0), nil).AnyTimes()
		e.srv = m
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.Database, got.Code)
	})

}
