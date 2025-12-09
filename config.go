package config

import (
    "log"
    "os"
    "strconv"
)

func GetEnv() string {
    return get("ENV")
}

func GetDataSourceURL() string {
    return get("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
    p := get("APPLICATION_PORT")
    port, err := strconv.Atoi(p)
    if err != nil {
        log.Fatalf("Invalid port %s", p)
    }
    return port
}

func get(k string) string {
    v := os.Getenv(k)
    if v == "" {
        log.Fatalf("%s environment variable missing", k)
    }
    return v
}
