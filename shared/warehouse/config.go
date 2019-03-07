package warehouse

import "time"

type Config struct {
	DefaultStorage string `default:"local" yaml:"defaultStorage"`

	Http struct {
		Timeout         time.Duration `default:"30s"`
		MaxConnsPerHost int           `default:"5" yaml:"maxConnsPerHost"`
	}

	LocalStorage struct {
		SavePath string `default:"local/warehouse"`
		Endpoint string `default:"http://localhost:80"`
	}

	S3 struct {
		Region     string `default:"ap-northeast-1" yaml:"region"`
		AccessKey  string `yaml:"accessKey"`
		SecretKey  string `yaml:"secretKey"`
		Token      string `default:"" yaml:"token"`
		Bucket     string `yaml:"bucket"`
		PathPrefix string `yaml:"prefix"`
	}

	Debug struct {
		DisableUserAuthValidation bool `default:"false" yaml:"disableUserAuthValidation"`
		DisableSchemaValidation   bool `default:"false" yaml:"disableSchemaValidation"`
	}
}
