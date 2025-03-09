package nginxcachepurger

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Sleep interval between requests, in seconds
const DefaultRequestSleepInterval uint = 6

type Config struct {
	PurgeBaseUrl         string
	RequestSleepInterval uint
}

func ConfigFromDotEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	purgeBaseUrl := os.Getenv("PURGE_BASE_URL")
	if purgeBaseUrl == "" {
		return nil, errors.New("missing PURGE_BASE_URL environement variable")
	}

	sleepInterval, _ := strconv.ParseUint(os.Getenv("REQUEST_SLEEP_INTERVAL"), 10, 32)
	sleepInterval32 := uint(sleepInterval)
	if sleepInterval32 == 0 {
		sleepInterval32 = DefaultRequestSleepInterval
	}

	c := &Config{
		PurgeBaseUrl:         purgeBaseUrl,
		RequestSleepInterval: sleepInterval32,
	}

	return c, nil
}
