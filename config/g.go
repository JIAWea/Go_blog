package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "go_mega"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")		// name of config file (without extension)

	viper.AddConfigPath(".")													// optionally look for config in the working directory
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))					// call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))		// path to look for the config file in

	err := viper.ReadInConfig()			// Find and read the config file
	if err != nil {						// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString func
func GetMysqlConnectingString() string {
    usr := viper.GetString("mysql.user")
    pwd := viper.GetString("mysql.password")
    host := viper.GetString("mysql.host")
    db := viper.GetString("mysql.db")
    charset := viper.GetString("mysql.charset")

    return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}

// 邮箱
func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	fmt.Printf("%s:%s,%s@%s \n", server, port, user, pwd)
	return
}

// 获取服务地址
func GetServerURL() (url string) {
	url = viper.GetString("server.url")
	return url
}