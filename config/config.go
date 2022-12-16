package config

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/azer/logger"
	"github.com/bitspill/floutil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	appDir        string
	defaultAppDir = floutil.AppDataDir("oipd", false)
)

//go:embed defaults/config.example.yml
var defaultConfig []byte

func init() {
	logger.SetOutput(os.Stdout)

	loadDefaults()

	pflag.String("appdir", defaultAppDir, "Location of oip data directory and config file")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	appDir = viper.GetString("appdir")

	err = viper.ReadConfig(bytes.NewReader(defaultConfig))
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(appDir)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("error loading config file, utilizing defaults", logger.Attrs{"err": err})

		confFile := filepath.Join(appDir, "config.yml")
		_, err = os.Stat(confFile)
		if os.IsNotExist(err) {
			log.Info("config.yml not found, writing default config file", logger.Attrs{"confFile": confFile})
			err = os.MkdirAll(appDir, 0755)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(confFile, defaultConfig, 0600)
			if err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		}
	}
}

func loadDefaults() {
	// command line flag to change config directory
	viper.SetDefault("appdir", defaultAppDir)

	// Elastic defaults
	viper.SetDefault("elastic.host", "http://127.0.0.1:9200")
	viper.SetDefault("elastic.useCert", false)
	viper.SetDefault("elastic.certFile", "certs/oipd.pem")
	viper.SetDefault("elastic.certKey", "certs/oipd.key")
	viper.SetDefault("elastic.certRoot", "certs/root-ca.pem")

	// Flod defaults
	defaultFlodDir := floutil.AppDataDir("flod", false)
	defaultFlodCert := filepath.Join(defaultFlodDir, "rpc.cert")
	viper.SetDefault("flod.certFile", defaultFlodCert)
	viper.SetDefault("flod.tls", true)
	viper.SetDefault("flod.host", "127.0.0.1:8334")
	viper.SetDefault("flod.user", "user")
	viper.SetDefault("flod.pass", "pass")

	// Testnet defaults
	viper.SetDefault("oip.network", "mainnet")

	// HttpApi defaults
	viper.SetDefault("oip.api.listen", "127.0.0.1:1606")
	viper.SetDefault("oip.api.enabled", false)
}

func IsTestnet() bool {
	return viper.GetString("oip.network") != "mainnet"
}

func SetTestnet(testnet bool) {
	n := "mainnet"
	if testnet {
		n = "testnet"
	}
	viper.Set("oip.network", n)
}

func GetFilePath(key string) string {
	v := viper.GetString(key)
	if filepath.IsAbs(v) {
		return v
	}
	return filepath.Join(appDir, v)
}
