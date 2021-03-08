package cs

// APISetting specify Web address and login info
type APISetting struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Token    TokenAPI
}

type TokenAPI struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
