package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	URL         string
	ContentType string
}

func (c *Config) Read() error {
	var (
		err error
	)
	err = godotenv.Load()
	if err != nil {
		return err
	}
	c.URL = os.Getenv("BUS_TICKET_URL")
	c.ContentType = os.Getenv("BUS_TICKET_CONTENT_TYPE")
	return nil
}
