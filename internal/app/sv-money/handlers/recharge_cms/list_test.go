package recharge_cms

import (
	"context"
	"github.com/go-liam/util/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/money/recharge"
	"grape/test/mock/money/mock_recharge"
	"testing"
)

func TestGetList(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_recharge.NewMockService(ctrl)
	item := &recharge.Model{ID: 1, CreatedAt: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	pagination := &response.Pagination{PageSize: 1, Current: 1}
	ls := make([]*recharge.Model, 0)
	ls = append(ls, item)

	m.EXPECT().FindMulti(pagination, s).Return(ls, nil).AnyTimes()
	e := new(SrvList)
	e.srv = m
	e.current = pagination.Current
	e.pageSize = pagination.PageSize
	e.GetList()
	assert.EqualValues(t, testdata.ConstWantOne, len(e.list))
}
