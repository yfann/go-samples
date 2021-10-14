package sample

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/google/martian/log"
	"time"
)


func GetAliLogs(){
	client := sls.CreateNormalInterface("http://cn-shanghai-intranet.log.aliyuncs.com", "xx", "xx", "")

	from,to,_:=parseFromTo("2021-06-01","2021-06-30")

	rep,err:=client.GetLogs("hscp","pod-metering-data","",from,to,"*",100,0,false)

	log.Infof("",rep)
	log.Errorf("",err)
}


func parseFromTo(f,t string)(int64,int64,error){
	from,err:=time.Parse("2006-01-02 15:04:05",f)
	if err!=nil{
		log.Errorf("time parse:",err)
		return 0,0,err
	}
	to,err:=time.Parse("2006-01-02 15:04:05",t)
	if err!=nil{
		log.Errorf("time parse:",err)
		return 0,0,err
	}
	return from.Unix(),to.Unix(),nil
}