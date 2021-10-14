package sample

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
)


func getAliLogs(){
	client := sls.CreateNormalInterface(p.Config.aliLogEndpoint, p.Config.aliLogAccessKeyID, p.Config.aliLogAccessKeySecret, "")
}