package cs

// APISetting specify Web address and login info
type APISetting struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
