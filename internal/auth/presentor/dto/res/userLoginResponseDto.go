package res

type UserLoginResponseDto struct {
	StatusCode int
	Message    string
	Tokens     TokensMap
}

type TokensMap struct {
	AccessToken  string
	RefreshToken string
}
