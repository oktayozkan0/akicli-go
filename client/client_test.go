package client

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    []ClientOption
		wantErr bool
	}{
		{
			name: "başarılı client oluşturma",
			opts: []ClientOption{
				WithBaseURL("https://test.example.com"),
				WithToken("test-token"),
			},
			wantErr: false,
		},
		{
			name:    "baseURL olmadan client oluşturma",
			opts:    []ClientOption{WithToken("test-token")},
			wantErr: true,
		},
		{
			name:    "hiçbir opsiyon olmadan client oluşturma",
			opts:    []ClientOption{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.opts...)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
				assert.NotEmpty(t, client.baseURL)
				assert.NotEmpty(t, client.Token)
				assert.NotNil(t, client.httpClient)
			}
		})
	}
}

func TestWithTimeout(t *testing.T) {
	timeout := 5 * time.Second
	client, err := NewClient(
		WithBaseURL("https://test.example.com"),
		WithTimeout(timeout),
	)
	assert.NoError(t, err)
	assert.Equal(t, timeout, client.httpClient.Timeout)
}

func TestClient_Get(t *testing.T) {
	client, err := NewClient(
		WithBaseURL("https://test.example.com"),
		WithToken("test-token"),
	)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		path    string
		params  map[string]string
		wantErr bool
	}{
		{
			name:    "geçersiz URL ile istek",
			path:    "invalid-path",
			params:  map[string]string{},
			wantErr: true,
		},
		{
			name: "parametreli istek",
			path: "/test",
			params: map[string]string{
				"key": "value",
			},
			wantErr: true, // Gerçek bir sunucu olmadığı için hata alacağız
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Get(tt.path, tt.params)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestClient_Post(t *testing.T) {
	client, err := NewClient(
		WithBaseURL("https://test.example.com"),
		WithToken("test-token"),
	)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		path    string
		body    *bytes.Buffer
		wantErr bool
	}{
		{
			name:    "geçersiz URL ile istek",
			path:    "invalid-path",
			body:    bytes.NewBufferString(`{"test": "data"}`),
			wantErr: true,
		},
		{
			name:    "geçerli body ile istek",
			path:    "/test",
			body:    bytes.NewBufferString(`{"test": "data"}`),
			wantErr: true, // Gerçek bir sunucu olmadığı için hata alacağız
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Post(tt.path, tt.body)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
} 