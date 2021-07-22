package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)
func main() {
	var outClusterScheme   = runtime.NewScheme()
	v1beta1.AddToScheme(outClusterScheme)
	kubeconfig := flag.String("kubeconfig", "C:\\Users\\310256235\\.kube\\config", "kubeconfig file")
	flag.Parse()
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := client.New(config,client.Options{
		Scheme: outClusterScheme,
	})
	si := &v1beta1.ServiceInstance{}
	serviceInstanceRef := client.ObjectKey{Name: "created-instance", Namespace: "operator"}
	ctx := context.TODO()
	err := clientset.Get(ctx, serviceInstanceRef, si)

	fmt.Println(err)
}
