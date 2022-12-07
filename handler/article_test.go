package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abdghn/alpha-indo-soft-be-test/models"
	"github.com/abdghn/alpha-indo-soft-be-test/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := models.Article{
		Author: "Aaditya",
		Title:  "Title",
		Body:   "Test Body",
	}

	testCases := []struct {
		name          string
		body          interface{}
		buildStubs    func(repository *mocks.MockRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			body: user,
			buildStubs: func(repository *mocks.MockRepository) {
				repository.EXPECT().
					Create(&user).
					Times(1).
					Return(&user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchArticle(t, recorder.Body, user)
			},
		},
		{
			name: "InternalServerError",
			body: user,
			buildStubs: func(repository *mocks.MockRepository) {
				repository.EXPECT().
					Create(gomock.Any()).
					Times(1).
					Return(nil, errors.New(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: user,
			buildStubs: func(repository *mocks.MockRepository) {
				repository.EXPECT().
					Create(user).
					Times(0).
					Return(nil, errors.New(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockRepository(ctrl)
			tc.buildStubs(repository)

			server := NewServer(repository)

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/articles")
			var request *http.Request
			if tc.name == "BadRequest" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetAll(t *testing.T) {
	articleList := []*models.Article{
		{
			Author: "Aaditya",
			Title:  "Title",
			Body:   "Test Body",
		},
	}

	queryParam := &models.GetArticle{
		Query:  "Title",
		Author: "Aaditya",
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			buildStubs: func(repository *mocks.MockRepository) {
				repository.EXPECT().
					GetArticles(queryParam.Query, queryParam.Author).
					Times(1).
					Return(articleList, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchArticles(t, recorder.Body, articleList)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(repository *mocks.MockRepository) {
				repository.EXPECT().
					GetArticles(queryParam.Query, queryParam.Author).
					Times(1).
					Return(nil, errors.New(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockRepository(ctrl)
			tc.buildStubs(repository)

			server := NewServer(repository)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/articles")
			var request *http.Request

			request = httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}
