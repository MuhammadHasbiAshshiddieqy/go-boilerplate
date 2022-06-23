package dto

type (
	UserResponse struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		RoleID string `json:"role_id"`
	}

	UserRequestCreate struct {
		Name     string `json:"name" validate:"required,min=3,max=64"`
		Password string `json:"password"`
		RoleID   string `json:"role_id" validate:"required"`
	}

	UserRequestUpdate struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		RoleID string `json:"role_id"`
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
