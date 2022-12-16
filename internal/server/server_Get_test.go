package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/wspectra/urlShort/internal/config"
	mock_store "github.com/wspectra/urlShort/internal/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_Get(t *testing.T) {
	//ARRANGE
	type mockBehavior func(r *mock_store.MockStore, url string)

	tests := []struct {
		name                 string
		shortUrl             string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "ok",
			shortUrl: "",
			mockBehavior: func(r *mock_store.MockStore, url string) {
				r.EXPECT().GetInfo(url).Return("https://www.google.com/", nil)
			},
			expectedStatusCode:   308,
			expectedResponseBody: "<a href=\"https://www.google.com/\">Permanent Redirect</a>.\n\n",
		},
		{
			name:     "Url not found",
			shortUrl: "",
			mockBehavior: func(r *mock_store.MockStore, url string) {
				r.EXPECT().GetInfo(url).Return("", errors.New("Url not found"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "Url not found\n",
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
			test.mockBehavior(repo, test.shortUrl)
			server.Store = repo

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)

			// ACT
			server.handleGet().ServeHTTP(w, req)

			// ASSERT
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}

}
