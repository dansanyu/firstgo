package config

import "os"

type Config struct {
	DbHost     string // 数据库地址
	DbUser     string // 数据库
	DbName     string // 数据库用户名
	DbPassword string // 数据库密码
	DbPort     string // 数据库端口号
}

func LoadConfig() Config {
	return Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}
