package main

type Config struct {
	Port            string
	Password        string
	DBFile          string
	CleanerInterval uint
	TokenDuration   uint
}

func (c *Config) Save() error {

	return nil
}
