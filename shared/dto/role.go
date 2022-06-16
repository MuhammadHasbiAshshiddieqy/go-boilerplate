package dto

type (
	RoleResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	RoleRequestCreate struct {
		Name string `json:"name"`
	}

	RoleRequestUpdate struct {
		ID   string `param:"id"`
		Name string `json:"name"`
	}
)
