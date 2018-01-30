package types

type MyConfig struct {
	PostgresScheme   string `json:"postgres_scheme"`
	PostgresBase     string `json:"postgres_base"`
	PostgresHost     string `json:"postgres_host"`
	PostgresPort     string `json:"postgres_port"`
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
}
