package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config LilHelperConfig
var VTool *viper.Viper
var HelperJwt JwtConfig

type LilHelperConfig struct {
	Domain   string
	Database DatabaseConfig
	Jwt      LilHelperJwtConfig
}

type LilHelperJwtConfig struct {
	Helper JwtViperConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Config   string
}

type JwtViperConfig struct {
	Realm          string
	IdentityKey    string `mapstructure:"identity_key"`
	Secret         string
	Timeout        int
	MaxRefresh     int    `mapstructure:"max_refresh"`
	CookieName     string `mapstructure:"cookie_name"`
	CookiePath     string `mapstructure:"cookie_path"`
	CookieDomain   string `mapstructure:"cookie_domain"`
	SecureCookie   bool   `mapstructure:"secure_cookie"`
	CookieHTTPOnly bool   `mapstructure:"cookie_http_only"`
}

func (c *JwtViperConfig) JwtConfig() (jwtConfig JwtConfig) {
	jwtConfig.Realm = c.Realm
	jwtConfig.IdentityKey = c.IdentityKey
	jwtConfig.Secret = c.Secret
	jwtConfig.Timeout = time.Duration(c.Timeout) * time.Minute
	jwtConfig.MaxRefresh = time.Duration(c.MaxRefresh) * time.Minute
	if jwtConfig.CookieName = c.CookieName; jwtConfig.CookieName == "" {
		jwtConfig.CookieName = "x-token"
	}
	if jwtConfig.CookiePath = c.CookiePath; jwtConfig.CookiePath == "" {
		jwtConfig.CookiePath = "/"
	}
	jwtConfig.CookieDomain = c.CookieDomain
	jwtConfig.SecureCookie = c.SecureCookie
	jwtConfig.CookieHTTPOnly = c.CookieHTTPOnly

	return
}

type JwtConfig struct {
	Realm          string
	IdentityKey    string
	Secret         string
	Timeout        time.Duration
	MaxRefresh     time.Duration
	CookieName     string
	CookiePath     string
	CookieDomain   string
	SecureCookie   bool
	CookieHTTPOnly bool
}

func init() {
	v := viper.New()
	v.SetConfigName("lilhelper")
	v.AddConfigPath("static/config/")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file failed: %w \n", err))
	} else if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unmarshal config file failed: %w \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := v.Unmarshal(&Config)
		if err != nil {
			log.Println(err)
		}
	})
	HelperJwt = Config.Jwt.Helper.JwtConfig()
	if Config.Domain == "" {
		Config.Domain = "127.0.0.1:8080"
	}
	VTool = v
}
