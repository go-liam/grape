package log_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/errorcode"
	"grape/configs/testdata"
	"grape/test/mock/sylog/mock_sylog"
	"testing"
)

func TestSite_UpdateStatus(t *testing.T) {
	e := new(SrvUpdateStatus)

	t.Run(testdata.ConstFail, func(t *testing.T) {
		//e.req = new(reqStatus)
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.RequestParameter, got.Code)
	})

	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	e.req = &reqStatus{BeforeTime: 1}

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		m := mock_sylog.NewMockService(ctrl)
		m.EXPECT().Delete(e.req.BeforeTime).Return(int64(1), nil).AnyTimes()
		e.srv = m
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.Success, got.Code)
	})

	t.Run(testdata.ConstFail, func(t *testing.T) {
		m := mock_sylog.NewMockService(ctrl)
		m.EXPECT().Delete(e.req.BeforeTime).Return(int64(0), nil).AnyTimes()
		e.srv = m
		got := e.UpdateStatus()
		assert.EqualValues(t, errorcode.Database, got.Code)
	})

}
