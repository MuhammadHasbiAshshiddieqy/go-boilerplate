package dto

type (
	AccessDetails struct {
		AccessUuid string
		UserId     string
	}
	TokenDetails struct {
		AccessToken  string
		RefreshToken string
		AccessUuid   string
		RefreshUuid  string
		AtExpires    int64
		RtExpires    int64
	}
)
