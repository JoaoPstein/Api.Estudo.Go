package configs

var cfg *config

type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig stuct {
	Host string
	Port string
	User string
	Pass string
	Database string
}

func init(){
	
}