package api

type APIResponse[T any] struct {
	StatusCode int
	RawContent []byte
	Response   T
}

type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserResponse
	Token string `json:"token"`
}

func (a *API) Login(email, password string) (*APIResponse[LoginResponse], error) {
	u := LoginPath
	loginPaylod := LoginRequest{
		Email:    email,
		Password: password,
	}
	var apiResponse APIResponse[LoginResponse]
	resp, err := PostData(a.client, u, loginPaylod, apiResponse)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) GetCurrentUser() (*APIResponse[UserResponse], error) {
	u := CurrentUserPath
	return FetchResource[APIResponse[UserResponse]](a.client, u, nil)
}
