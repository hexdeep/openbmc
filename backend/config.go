package main

type Config struct {
	Address         string
	Password        string
	DBFile          string
	CleanerInterval uint
	TokenDuration   uint
	DefaultSize     int
	SSL             SSLConfig
}

type SSLConfig struct {
	Cert string
	Key  string
}
