package main

import (
	"github.com/BurntSushi/toml"
)

// Config represents default configuration
type Config struct {
	ListenAddr string

	// Mb
	BufferSize uint

	// Kb
	Bitrate     uint
	Codec       string
	Format      string
	ContentType string

	LogLevel int
}

var config *Config

func initConfig(filePath string) error {
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return err
	}

	return config.Validate()
}

// Validate validates config and fills empty fields
func (config *Config) Validate() error {
	if config.ListenAddr == "" {
		config.ListenAddr = defaultListenAddr
	}

	if config.BufferSize == 0 {
		config.BufferSize = defaultBufferSize
	}

	if config.Bitrate == 0 {
		config.Bitrate = defaultBitrate
	}

	if config.Codec == "" {
		config.Codec = defaultCodec
	}

	if config.Format == "" {
		config.Format = defaultFormat
	}

	if config.ContentType == "" {
		config.ContentType = defaultContentType
	}

	return nil
}
