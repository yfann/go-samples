package sample

import (
	"flag"
	"github.com/google/martian/log"
	"github.com/golang/glog"
)

func Testlog(){
	log.SetLevel(log.Info)
	log.Infof("log test:%s %s %s","ss","bb","cc")


	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()
	glog.Infof("glog test: %s %s","aa","bb")
}