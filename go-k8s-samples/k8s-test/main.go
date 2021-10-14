package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)
func main() {
	var outClusterScheme   = runtime.NewScheme()
	v1beta1.AddToScheme(outClusterScheme)
	v1.AddToScheme(outClusterScheme)
	kubeconfig := flag.String("kubeconfig", "C:\\Users\\310256235\\.kube\\config", "kubeconfig file")
	flag.Parse()
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	c, _ := client.New(config,client.Options{
		Scheme: outClusterScheme,
	})
	listPods(c)
}


func listPods(clientset client.Client){
	podList :=&v1.PodList{}
	opts :=[]client.ListOption{
		//client.InNamespace("kube-system"),
	}
	ctx := context.TODO()
	err := clientset.List(ctx, podList, opts...)
	fmt.Println(podList)
	fmt.Println(err)
}

func getServiceInstance(clientset client.Client){
	si := &v1beta1.ServiceInstance{}
	serviceInstanceRef := client.ObjectKey{Name: "created-instance", Namespace: "operator"}
	ctx := context.TODO()
	err := clientset.Get(ctx, serviceInstanceRef, si)
	fmt.Println(err)
}