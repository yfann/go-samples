package vault

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/vault/api"
	"net/http"
	"strings"
	"time"
)

const (
	vaultAddr = "http://127.0.0.1:8200"
	rootToken="s.L5K1VjBfzknvz5WwcmjI386K"
)
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type Vault struct {
	VaultClient *api.Client
}

func ParseToken(out string)map[string]string{
	res := strings.Split(out, "\n")
	a :=make(map[string]string)
	for _,row :=range res{
		 rows:=strings.Split(row,":")
		 if len(rows)==2{
		 	if strings.Contains(rows[0],"Unseal Key") || strings.Contains(rows[0],"Root Token"){
		 		key:=strings.ReplaceAll(strings.TrimSpace(rows[0])," ","_")
				a[key]=strings.TrimSpace(rows[1])
			}
		 }
	}
	return a
}


func GetVaultClient() *api.Client{

	return getVaultClient(vaultAddr,rootToken)
}

func AddTransitKey(key string){
	c:=GetVaultClient()
	options := map[string]interface{}{
		"type": "aes256-gcm96",
	}
	result,err:=c.Logical().Write("/transit/keys/"+key, options)
	fmt.Println(result)
	fmt.Println(err)
}


func getVaultClient(url string,token string) *api.Client{
	vaultClientConfig := api.DefaultConfig()
	vaultClientConfig.HttpClient = cleanhttp.DefaultClient()
	vaultClient, err := api.NewClient(vaultClientConfig)
	if err != nil {
		glog.Fatalf("failed to create vault api client:  %s", err)
	}

	vaultClient.SetAddress(url)
	if token!=""{
		vaultClient.SetToken(token)
	}
	return vaultClient
}


func TestK8sAuth(){
	client:=getVaultClient(vaultAddr,"")
	options := map[string]interface{}{
		"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6InZENk5CbHpmb1RnM1RKTVJCRkVUMW5Qd1VVTHVTMEIwcmY0eFZ2YzZJNTgifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJ2YXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLXJxNWc0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJlNjAxMjE5Zi0xYjRlLTQ2N2UtYWNjMS1iYThjM2ZhZmI1NzYiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6dmF1bHQ6ZGVmYXVsdCJ9.TQc6a0m1FfxaUTw3quTOPzDRT8IjVywDCl5sujETkBM1XuK60DdUTwvTNZxi1_sdwWdzlotVZl7igALaT0DQacdn-N8sl6WVEbjYeS3FNn9dkG586jyu54wIcHtBSWm0_J7N-agItb3pxMF666wO_NmAUvf6yIYOtK4kwhEEclw52GxRmRAGUd9O5BKmzXob0eSZNrOigNvWRjtduubajGIhVqI1SF39eUkS4Tkd6tPzHCKXWsvT0OsPVv2SIOa59XBGA5dYh0NGAbpiJICgkVu4Am7s13LFg5umxd2pXWOQmwqcYcAyTfcn-EKAMXS_O78kCRDbLvf7Ex-mtIbYaQ",
		"role":"test",
	}
	//fake
	//options := map[string]interface{}{
	//	"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IlR5aUt6Qi04M3lRNnZ4M2ZxVTlSNGs4b1BJMEZiblBhUERlYmVxYV9DUFkifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJzZXJ2aWNlLWJyb2tlciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLXdwd3pkIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIyODhlYzJlNS02MzYwLTRlYWEtOWQ2Mi02MzFjZTJhZjhlNTciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6c2VydmljZS1icm9rZXI6ZGVmYXVsdCJ9.X5TxbJWLs6z0PSLuc28xlsdvA2_LUfUtatTH6PR-93gb2-Xzd9WFxEwzJy_aEl3YTg6HsArN-Ic9E6-xpAtTOznxLyMhLjzbq6cgJg1ch44G39z23RZXV3QZe1kaFu64KkYHPNYUXbiTNDl51jh-zmaencUM4TBwKx9uHLHb0_1rqEH2JmgUeJJOs5LGpvt7ui3Y5khVQ2E5yy5b0M759mgsSXX8_XvtgJalh3LUEyrD3R94QTpF-8FMhGXcwWmvmzb19ETc82SrU2AN6zrI1SyU7njNxxTr2CZJU1WT5_tbDizMgNi3ojYIen01t8WzORDEfMCj3TcE1HNlshb4pg",
	//	"role":"test",
	//}
	result,err:=client.Logical().Write("auth/kubernetes/login", options)
	fmt.Println(result.Auth.ClientToken)
	fmt.Println(err)
	client.SetToken(result.Auth.ClientToken)
	sec,err:=client.Logical().Read("/kv/test/")
	fmt.Println(sec.Data["sss"])
}
