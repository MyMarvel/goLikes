package tests

import (
	"bytes"
	"encoding/json"
	"likes_handler/mocks/mock_controllers"
	"likes_handler/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestAddLikeRoute(t *testing.T) {
	r := routes.GenerateRoutes()
	requestBody := map[string]string{
		"targetAccount":  "Target",
		"currentAccount": "Me",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	meAccountMock := mock_controllers.NewMockIAccount(ctrl)
	meAccountMock.EXPECT().
		IsAccountLikes(gomock.Eq("Target")).
		Return(false, nil)
	meAccountMock.EXPECT().
		IncrementLikeCounter(gomock.Eq("Target")).
		Return(1, nil)

	factoryMock := mock_controllers.NewMockIFactory(ctrl)
	factoryMock.EXPECT().
		NewIAccount(gomock.Eq("Me")).
		Return(meAccountMock)
	routes.InitControllersFactory(factoryMock)

	w := performRequest(r, http.MethodPost, "/api/v1.0/likes/doLike", requestBody)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddLikeAlreadyLiked(t *testing.T) {
	r := routes.GenerateRoutes()
	requestBody := map[string]string{
		"targetAccount":  "Target",
		"currentAccount": "Me",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	meAccountMock := mock_controllers.NewMockIAccount(ctrl)
	meAccountMock.EXPECT().
		IsAccountLikes(gomock.Eq("Target")).
		Return(true, nil)

	factoryMock := mock_controllers.NewMockIFactory(ctrl)
	factoryMock.EXPECT().
		NewIAccount(gomock.Eq("Me")).
		Return(meAccountMock)
	routes.InitControllersFactory(factoryMock)

	w := performRequest(r, http.MethodPost, "/api/v1.0/likes/doLike", requestBody)
	assert.Equal(t, http.StatusAlreadyReported, w.Code)
}

func TestGetLikeCountRoute(t *testing.T) {
	r := routes.GenerateRoutes()
	requestBody := map[string]string{
		"targetAccount": "Target",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountMock := mock_controllers.NewMockIAccount(ctrl)
	accountMock.EXPECT().
		GetLikesCount(gomock.Eq("Target")).
		Return(999, nil)

	factoryMock := mock_controllers.NewMockIFactory(ctrl)
	factoryMock.EXPECT().
		NewIAccount(gomock.Eq("Target")).
		Return(accountMock)
	routes.InitControllersFactory(factoryMock)

	w := performRequest(r, http.MethodPost, "/api/v1.0/likes/getLikeCount", requestBody)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, nil, err)
	value, exists := response["likes"]
	assert.Equal(t, true, exists)
	assert.Equal(t, "999", value)
}

func performRequest(r http.Handler, method, path string, body map[string]string) *httptest.ResponseRecorder {
	params, _ := json.Marshal(body)
	requestBodyBytes := bytes.NewBuffer(params)
	req := httptest.NewRequest(method, path, requestBodyBytes)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
