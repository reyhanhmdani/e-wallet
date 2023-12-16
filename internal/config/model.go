package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
	Email    Email
	Redis    Redis
	Queue    Redis
	Midtrans Midtrans
}

type Redis struct {
	Addr     string
	Password string
}

type Server struct {
	Host string
	Port string
}

type JWT struct {
	Key string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Email struct {
	Host      string
	Port      string
	User      string
	Password  string
	EmailFrom string
}

type Midtrans struct {
	Key    string
	IsProd bool
}
