package main

import "os"

type Config struct {
	Port string
	Env  string

	Version   string
	Commit    string
	BuildTime string

	FfPlaceHodler bool
}

func NewConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return Config{
		Port:      port,
		Env:       env,
		Version:   Version,
		Commit:    Commit,
		BuildTime: BuildTime,
	}
}
