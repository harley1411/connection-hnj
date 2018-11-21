package connection

// Configuration ...
type configuration_type struct {
	postgresql Pgsql `json:"postgresql"`
}

// CONFIG
var configuration configuration_type
