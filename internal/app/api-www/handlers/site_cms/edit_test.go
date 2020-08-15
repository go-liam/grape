package site_cms

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grape/configs/testdata"
	"grape/internal/pkg/data/home_site/site"
	"grape/test/mock/www/mock_site"
	"testing"
)

func TestSite_EditSite(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock_site.NewMockService(ctrl)
	item := &site.Model{
		ID:    1,
		Title: testdata.ConstWantString,
	}

	m.EXPECT().Update(item).Return(int64(1), nil).AnyTimes()
	e := new(SrvEdit)
	e.srv = m
	e.req = &ReqModel{ID: "1", Title: testdata.ConstWantString}
	e.EditSite()
	assert.EqualValues(t, testdata.ConstWantOne, e.result)
}
