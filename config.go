package config

import "os"

func GetPort() string {
	p := os.Getenv("APPLICATION_PORT")
	if p == "" { return ":3001" }
	return ":" + p
}

func GetDataSourceURL() string {
	return os.Getenv("DATA_SOURCE_URL")
}
