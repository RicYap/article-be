package config

import (
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	config        *Config
	fileToWatch   []string
)

const (
	envDevelopment = "development"
	envStaging     = "staging"
	envProduction  = "production"
)

type (
	option struct {
		configFile string
	}
)

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(out, &config)
	if err != nil {
		return err
	}

	// ADD CONFIG FILE TO WATCH
	fileToWatch = append(fileToWatch, opt.configFile)

	return err
}

func PrepareWatchPath() {
	for _, path := range fileToWatch {
		viper.SetConfigFile(path)
		viper.SetConfigType("yaml")
		viper.WatchConfig()
	}
}

// Option ...
type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {

	configPath := "./files/etc/example/example.development.yaml"
	namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	env := string(namespace)
	if os.Getenv("GOPATH") == "" {
		configPath = "./example.development.yaml"
	}

	if env != "" {
		if strings.Contains(env, envStaging) {
			time.Sleep(30 * time.Second)
			configPath = "/vault/secrets/database.yaml"
		} else if strings.Contains(env, envProduction) {
			time.Sleep(30 * time.Second)
			configPath = "/vault/secrets/database.yaml"

		}
	}
	
	return configPath
}

// Get ...
func Get() *Config {
	return config
}
