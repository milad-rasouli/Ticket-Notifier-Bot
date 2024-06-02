package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	url         string
	contentType string
}

func (c *Config) Read() error {
	var (
		err error
	)
	err = godotenv.Load()
	if err != nil {
		return err
	}
	c.url = os.Getenv("BUS_TICKET_URL")
	c.contentType = os.Getenv("BUS_TICKET_CONTENT_TYPE")
	return nil
}
