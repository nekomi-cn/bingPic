package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/ZMuSiShui/bingImg/client"
	"github.com/ZMuSiShui/bingImg/conf"
)

func init() {
	flag.BoolVar(&conf.Help, "-help", false, "help message")
	flag.BoolVar(&conf.Help, "h", false, "help message")
	flag.BoolVar(&conf.Debug, "-debug", false, "start with debug mode")
	flag.BoolVar(&conf.Debug, "d", false, "start with debug mode")
	flag.BoolVar(&conf.Version, "-version", false, "print version info")
	flag.BoolVar(&conf.Version, "v", false, "print version info")
	flag.BoolVar(&conf.Update, "update", false, "update software")
	flag.StringVar(&conf.WriteToFile, "w", "image/", "download Bing Image")
	flag.Parse()
}

func Init() bool {
	client.InitLog()
	return true
}

func main() {
	if conf.Help {
		fmt.Printf("%v\n", conf.Usage)
		return
	}
	if conf.Version {
		fmt.Printf("Version: %s\n", conf.VERSION)
		return
	}
	if !Init() {
		os.Exit(1)
	}
	if conf.Debug {
		log.Info("Set Debug Mode")
	}
	log.Info("Starting download Bing Image")
	client.DownImage()
}
