package api

import (
	"testing"

	"github.com/oktayozkan0/akicli-go/client"
	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	tests := []struct {
		name    string
		opts    []client.ClientOption
		wantErr bool
	}{
		{
			name: "başarılı API oluşturma",
			opts: []client.ClientOption{
				client.WithBaseURL("https://test.example.com"),
				client.WithToken("test-token"),
			},
			wantErr: false,
		},
		{
			name:    "opsiyon olmadan API oluşturma",
			opts:    []client.ClientOption{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := NewAPI(tt.opts...)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, api)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, api)
				assert.NotNil(t, api.client)
			}
		})
	}
}

func TestFetchResource(t *testing.T) {
	type TestResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	tests := []struct {
		name    string
		path    string
		params  map[string]string
		wantErr bool
	}{
		{
			name:    "geçersiz path ile istek",
			path:    "/invalid-path",
			params:  map[string]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := client.NewClient(
				client.WithBaseURL("https://test.example.com"),
				client.WithToken("test-token"),
			)
			assert.NoError(t, err)

			result, err := FetchResource[TestResponse](c, tt.path, tt.params)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
