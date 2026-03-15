package main

import "os"

type Config struct {
	Port string
	Env  string

	Version   string
	Commit    string
	BuildTime string

	FeatureHello bool
}

func NewConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return Config{
		Port:         port,
		Env:          env,
		Version:      Version,
		Commit:       Commit,
		BuildTime:    BuildTime,
		FeatureHello: os.Getenv("FEATURE_HELLO") == "true",
	}
}
