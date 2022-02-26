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

type ResourceSettings struct {
	ImagePath string
}

var (
	ServerSetting  ServerSettings
	DBSetting      DBSettings
	StaticResource ResourceSettings
)

func Init_setting() error {
	vp := viper.New()
	viper.AutomaticEnv()
	if viper.GetBool("IS_DEV") {
		vp.SetConfigName("configs.dev")
	} else {
		vp.SetConfigName("configs")
	}
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
	err = vp.UnmarshalKey("StaticResource", &StaticResource)
	if err != nil {
		return err
	}
	return nil
}
