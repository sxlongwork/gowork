package configmap

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapDetail struct {
	ConfigMap
	Data map[string]string `json:"data"`
}

func GetConfigMapDetail(client kubernetes.Interface, ns string, name string) (*ConfigMapDetail, error) {
	log.Printf("getting configmap %s in namespace %s\n", name, ns)
	configmap, err := client.CoreV1().ConfigMaps(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("get configmap %s failed, Error: %s\n", name, err.Error())
		return nil, err
	}
	return &ConfigMapDetail{*toConfigMap(*configmap), configmap.Data}, nil
}
