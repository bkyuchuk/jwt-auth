package common

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/jwo_auth/internal/common/logger"
)

type Config struct {
	ServerEndpoint                       string
	ServerReadTimeout                    time.Duration
	ServerWriteTimeout                   time.Duration
	ServerHeaderAccessControlAllowOrigin string
	JwtSigningKey                        string
	JwtValidityInHours                   time.Duration
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	cwd, err := os.Getwd()

	if err != nil {
		logger.Error("Failed geting cwd.")
		return nil, err
	}

	file, err := os.OpenFile(filepath.Join(filepath.Dir(cwd), "config.json"), os.O_RDONLY, 0666)

	if err != nil {
		logger.Error("Failed opening config file.")
		return nil, err
	}

	defer file.Close()
	content, err := io.ReadAll(file)

	if err != nil {
		logger.Error("Failed reading from config file.")
		return nil, err
	}

	json.Unmarshal(content, config)
	return config, nil
}
