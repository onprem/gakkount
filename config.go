package main

import "os"

type config struct {
	dsn           string
	hydraAdminURL string
}

func getConfig() config {
	d, _ := os.LookupEnv("DSN")
	if d == "" {
		d = "postgresql://edu:secret@127.0.0.1:5432/edu?sslmode=disable"
	}

	h, _ := os.LookupEnv("HYDRA_ADMIN_URL")
	if h == "" {
		h = "http://localhost:4445"
	}
	return config{
		dsn:           d,
		hydraAdminURL: h,
	}
}
