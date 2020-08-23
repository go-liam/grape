package recharge_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/errorcode"
	"grape/configs/testdata"
	"grape/internal/pkg/data/money/recharge"
	"grape/test/mock/money/mock_recharge"
	"testing"
)

func TestAdd(t *testing.T) {
	e := new(SrvAdd)

	t.Run(testdata.ConstFail, func(t *testing.T) {
		e.req = new(ReqModel)
		got := e.Add()
		assert.EqualValues(t, errorcode.RequestParameter, got.Code)
	})

	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	t.Run(testdata.ConstSuccess, func(t *testing.T) {
		m := mock_recharge.NewMockService(ctrl)
		item := &recharge.Model{
			ID:     1,
			UserID: 1,
		}
		m.EXPECT().Create(item).Return(int64(1), nil).AnyTimes()

		e.srv = m
		e.req = &ReqModel{ID: "1", UserID: 1}
		e.Add()
		assert.EqualValues(t, testdata.ConstWantOne, e.result)
	})

	t.Run(testdata.ConstFail, func(t *testing.T) {
		m := mock_recharge.NewMockService(ctrl)
		item := &recharge.Model{
			ID:     1,
			UserID: 1,
		}
		m.EXPECT().Create(item).Return(int64(0), nil).AnyTimes()

		e.srv = m
		e.req = &ReqModel{ID: "1", UserID: 1}
		e.Add()
		assert.EqualValues(t, 0, e.result)
	})
}
