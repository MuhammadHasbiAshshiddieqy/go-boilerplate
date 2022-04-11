package dto

type (
	UserResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	UserRequestCreate struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	UserRequestUpdate struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	UserRequestPasswordUpdate struct {
		Password    string `json:"password"`
		NewPassword string `json:"new_password,omitempty"`
	}
	UserResponseToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	UserRequestLogin struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	UserRequestRefresh struct {
		RefreshToken string `json:"refresh_token"`
	}
)
