package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// api returns 500 internal server error for some reason. idk why. so i will not going to test this one.
func (a *API) Login(email, password string) (*http.Response, error) {
	u := loginPath + "/"
	reqBody, err := json.Marshal(
		map[string]interface{}{
			"email":    email,
			"password": password,
		},
	)

	if err != nil {
		return nil, err
	}
	req, err := a.client.Post(u, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	return req, nil
}
