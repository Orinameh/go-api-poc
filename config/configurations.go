package config

// Database structure
type Database struct {
	Name string
	User string
	Pass string
}

// Jwt structure
type Jwt struct {
	SecretKey []byte
}

// Config structure
type Config struct {
	Database Database
	Jwt      Jwt
}

// LoadConfig helps to load db configuration
func LoadConfig() Config {
	return Config{Database{"blockcoin", "username", "pwd"}, Jwt{[]byte("somerandomsecretkey")}}
}
