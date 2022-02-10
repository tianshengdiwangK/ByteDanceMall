package settings

import (
	"github.com/spf13/viper"
)

type ServerSettings struct {
	RunMode  string
	HttpPort string
}
type DBSettings struct {
	DbType   string
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

var (
	ServerSetting ServerSettings
	DBSetting     DBSettings
)

func Init_setting() error {
	vp := viper.New()
	vp.SetConfigName("configs")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config/")
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("Database", &DBSetting)
	if err != nil {
		return err
	}
	return nil
}
