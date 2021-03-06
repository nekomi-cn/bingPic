package client

import (
	"github.com/ZMuSiShui/bingImg/conf"

	log "github.com/sirupsen/logrus"
)

// 初始化日志
func InitLog() {
	if conf.Debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
	log.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
	})
}
