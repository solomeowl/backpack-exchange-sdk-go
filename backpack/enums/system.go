package enums

// SystemStatus represents the system status.
type SystemStatus string

const (
	SystemStatusOk          SystemStatus = "Ok"
	SystemStatusMaintenance SystemStatus = "Maintenance"
)

// Status matches the OpenAPI enum name for system status.
type Status = SystemStatus
