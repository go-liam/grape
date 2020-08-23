package user_api

import (
	"context"
	"github.com/go-liam/util/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/account/user"
	"grape/test/mock/account/mock_user"
	"testing"
)

func TestGetList(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_user.NewMockService(ctrl)
	item := &user.Model{ID: 1, CreatedAt: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	pagination := &response.Pagination{PageSize: 1, Current: 1}
	ls := make([]*user.Model, 0)
	ls = append(ls, item)

	m.EXPECT().FindMulti(pagination, s).Return(ls, nil).AnyTimes()
	e := new(SrvList)
	e.srv = m
	e.current = pagination.Current
	e.pageSize = pagination.PageSize
	e.GetList()
	assert.EqualValues(t, testdata.ConstWantOne, len(e.list))
}
