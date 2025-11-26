package entities

type AuthData struct {
	Host     string `json:"host" validate:"required,hostname"`
	Port     int    `json:"port" validate:"gt=0"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Database string `json:"database" validate:"required"`
	SSLMode  string `json:"sslmode"`
}
