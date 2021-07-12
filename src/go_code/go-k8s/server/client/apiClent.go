package client

import (
	"k8s-manage/server/args"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", args.Holder.GetKubeconfig())
	if err != nil {
		// panic(err.Error())
		log.Printf("build config from flags failed,  kubeconfig=%s\n", args.Holder.GetKubeconfig())
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("get config failed.\n")
		panic(err.Error())
	}
	return client
}
