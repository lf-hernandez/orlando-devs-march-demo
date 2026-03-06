package main

import "os"

type Config struct {
	Port string
	Env  string

	Version   string
	Commit    string
	BuildTime string

	FfA bool
	FfB bool
	FfC bool
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
		FfA:       os.Getenv("FF_A") == "true",
		FfB:       os.Getenv("FF_B") == "true",
		FfC:       os.Getenv("FF_C") == "true",
	}
}
