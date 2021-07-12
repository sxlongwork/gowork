package service

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceDetail struct {
	Service
	Namespace string   `json:"namespace"`
	Selector  []string `json:"selector"`
}

func GetServiceDetail(client kubernetes.Interface, ns string, name string) (*ServiceDetail, error) {
	log.Printf("begin get service %s detail\n", name)
	service, err := client.CoreV1().Services(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("get service %s detail failed, Error: %s", name, err.Error())
	}
	svcDetail := toServiceDetail(service)
	return svcDetail, err
}

func toServiceDetail(service *v1.Service) *ServiceDetail {
	svc := toService(*service)
	var selectors []string
	for k, v := range service.Spec.Selector {
		str := k + "=" + v
		selectors = append(selectors, str)
	}
	return &ServiceDetail{
		Service:   *svc,
		Namespace: service.ObjectMeta.Namespace,
		Selector:  selectors,
	}
}
