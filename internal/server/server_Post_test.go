package server

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/wspectra/urlShort/internal/config"
	mock_store "github.com/wspectra/urlShort/internal/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_Post(t *testing.T) {
	//ARRANGE
	type mockBehavior func(r *mock_store.MockStore, url string)

	tests := []struct {
		name                 string
		inputBody            string
		inputUrl             string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"Url":"http://google.com"}`,
			inputUrl:  "http://google.com",
			mockBehavior: func(r *mock_store.MockStore, url string) {
				r.EXPECT().PostInfo(url).Return("rtr", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "rtr",
		},
		{
			name:      "invalid JSON",
			inputBody: `Url":"http://google.com"}`,
			inputUrl:  "",
			mockBehavior: func(r *mock_store.MockStore, url string) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: "invalid character 'U' looking for beginning of value\n",
		},
		{
			name:      "wrong JSON structure",
			inputBody: `{"myUrl":"http://google.com"}`,
			inputUrl:  "http://google.com",
			mockBehavior: func(r *mock_store.MockStore, url string) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: "Key: 'RequestStruct.Url' Error:Field validation for 'Url' failed on the 'required' tag\n",
		},
		{
			name:      "invalid URL",
			inputBody: `{"Url":"qewerfwet"}`,
			inputUrl:  "",
			mockBehavior: func(r *mock_store.MockStore, url string) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: "parse \"qewerfwet\": invalid URI for request\n",
		},
		{
			name:      "postInfo error",
			inputBody: `{"Url":"http://google.com"}`,
			inputUrl:  "http://google.com",
			mockBehavior: func(r *mock_store.MockStore, url string) {
				r.EXPECT().PostInfo(url).Return("", errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "some error\n",
		},
	}

	config.ConfPath = "../../configs/config.toml"
	server := NewServer()
	zerolog.SetGlobalLevel(zerolog.Disabled)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_store.NewMockStore(c)
			test.mockBehavior(repo, test.inputUrl)

			server.Store = repo

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/post",
				bytes.NewBufferString(test.inputBody))

			// ACT
			server.handlePost().ServeHTTP(w, req)

			// ASSERT
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}

}
