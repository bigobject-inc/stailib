package cs

// APISetting specify Web address and login info
type APISetting struct {
	Address           string `json:"address"`
	Port              string `json:"port"`
	ClientCredentials bool   `json:"clientCredentials"` // if true, then User become client_id,  Password become client_secret
	User              string `json:"user"`
	Password          string `json:"password"`
	CacheTimeout      int    `json:"cacheTimeout"`
}
