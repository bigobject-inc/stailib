package cs

/*
// STAILibrary STAI library
type STAILibrary interface {
	GetServConfigServer() cs.ConfigServerAPI
}*/

// ConfigServerAPI configuration server API service
type ConfigServerAPI interface {
	API(method string, postfix string) (body []byte, err error)
	ClearCache() (err error)
	Defer() (err error)
}
