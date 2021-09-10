package congo

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	data map[string]interface{}
}

func New() *Config {
	return &Config{data: map[string]interface{}{}}
}

func (c *Config) Get(key string) interface{} {
	return c.data[key]
}

func (c *Config) GetString(key string) string {
	value := c.Get(key)
	if value == nil {
		return ""
	}

	return value.(string)
}

func (c *Config) load(r io.Reader) error {
	return json.NewDecoder(r).Decode(&c.data)
}

func (c *Config) LoadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.load(file)
}
