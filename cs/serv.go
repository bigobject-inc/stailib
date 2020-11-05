package cs

// CSServ : configuration server API service
type CSServ interface {
	NewCS(setting APISetting) ConfigServerAPI
}
