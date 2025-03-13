package config

import "flag"

var (
	Auth     bool
	Tls      bool
	Jobs     bool
	Database string
)

func Flags() {
	flag.BoolVar(&Auth, "auth", false, "Specific Auth Requirement")
	flag.BoolVar(&Tls, "tls", false, "Specific TLS Requirement")
	flag.BoolVar(&Jobs, "jobs", false, "Specific using background jobs")
	flag.StringVar(&Database, "database", "dev", "Specific database to connect")
	flag.Parse()

}
