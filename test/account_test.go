package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"likes_handler/internal/controllers"
	"likes_handler/internal/mocks/mock_controllers"
)

func TestIncrementLikeCounter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mock_controllers.NewMockIDatabase(ctrl)
	dbMock.EXPECT().
		AddLikeToAccount(gomock.Eq("Target"), gomock.Eq("Me")).
		Return(nil)
	dbMock.EXPECT().
		SelectAmountLikesForAccount(gomock.Eq("Target")).
		Return(1, nil)
	initTestDatabase(dbMock, ctrl)

	a := controllers.NewAccount("Me")
	likesCount, err := a.IncrementLikeCounter("Target")

	assert.Equal(t, 1, likesCount)
	assert.Equal(t, nil, err)
}

func initTestDatabase(dbMock controllers.IDatabase, ctrl *gomock.Controller) {
	factoryMock := mock_controllers.NewMockIFactory(ctrl)
	factoryMock.EXPECT().
		NewIDatabase(gomock.Eq("host:port"), gomock.Eq("pass")).
		Return(dbMock)
	controllers.InitControllersFactory(factoryMock)
	controllers.InitDatabase("host", "port", "pass")
}
