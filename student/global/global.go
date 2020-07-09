package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"student/config"
)

var (
	GL_DB		*gorm.DB		//mysql
	GL_REDIS 	*redis.Client	//redis
	GL_VP		*viper.Viper	//配置文件读取
	GL_LOG		*logging.Logger //日志
	GL_CONFIG	*config.Server
)
