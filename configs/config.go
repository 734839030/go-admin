package configs

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	Web *Web
	Log *Log
	Db  *DB
}

type Web struct {
	Addr string
	// seconds
	ShutdownTime time.Duration
}

type Log struct {
	Filename string
	// M
	MaxSize int
	// Day
	MaxAge          int
	Level           string
	ReportCaller    bool
	OutputToConsole bool
}
type DB struct {
	Dsn          string
	MaxIdleConns int
	// S
	ConnMaxLifetime time.Duration
	// S
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
}

var (
	C    Config
	once sync.Once
)

const (
	configPath = "configs"
	configName = "seezoon"
	configType = "yml"
)

func init() {
	once.Do(func() {
		workDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		viper := viper.New()
		viper.AddConfigPath(workDir + string(os.PathSeparator) + configPath)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&C); err != nil {
			panic(err)
		}

		if marshal, err := json.Marshal(&C); err != nil {
			panic(err)
		} else {
			log.Printf("config is:%s\n", marshal)
		}
	})
}
