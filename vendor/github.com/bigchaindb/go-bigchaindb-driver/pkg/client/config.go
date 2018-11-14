package client

import (
	"flag"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ClientConfig struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Options struct {
	pathFlag *flag.Flag
	paths    []string

	cfgType string
	readers []io.Reader

	// TODO review viper package
	viper *viper.Viper
}

type OptionFunc func(*Options) error

func Load(cfg interface{}, optFuncs ...OptionFunc) error {
	var err error
	opts := &Options{}

	for _, optFunc := range optFuncs {
		err = optFunc(opts)

		if err != nil {
			return err
		}
	}

	if opts.viper == nil {
		opts.viper = viper.New()
	}

	if opts.cfgType != "" && len(opts.readers) != 0 {
		err = MergeInReaders(opts.viper, opts.cfgType, opts.readers)

		if err != nil {
			return err
		}
	}

	err = opts.viper.Unmarshal(cfg)

	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config files")
	}

	return nil
}

func FromReaders(cfgType string, readers ...io.Reader) OptionFunc {
	return func(o *Options) error {
		o.cfgType = cfgType
		o.readers = readers

		return nil
	}
}

// MergeInConfigs merges the viper configs found in several readers into a single one
func MergeInReaders(v *viper.Viper, cfgType string, readers []io.Reader) error {
	v.SetConfigType(cfgType)

	for _, reader := range readers {
		err := v.MergeConfig(reader)

		if err != nil {
			return errors.Wrap(err, "Could not merge config readers")
		}
	}

	return nil
}
