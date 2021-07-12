package service

import (
	"fmt"
	"k8s-manage/server/resources/common"
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceList struct {
	Services []*Service `json:"services"`
}

type Service struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	ClusterIP  string   `json:"clusterIP"`
	ExternalIP []string `json:"externalIP"`
	Ports      string   `json:"ports"`
}

func GetServiceList(client kubernetes.Interface, ns string) (*ServiceList, error) {
	log.Println("begin get service list.")
	servicelist, err := client.CoreV1().Services(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get service list failed, ERROR: %s", err.Error())
	}
	slist := toServiceList(servicelist)
	return slist, err
}

func toServiceList(servicelist *v1.ServiceList) *ServiceList {
	var sList []*Service
	for _, svc := range servicelist.Items {
		s := toService(svc)
		sList = append(sList, s)
	}
	return &ServiceList{Services: sList}
}

func toService(svc v1.Service) *Service {
	return &Service{
		Name:       svc.ObjectMeta.Name,
		Type:       string(svc.Spec.Type),
		ClusterIP:  svc.Spec.ClusterIP,
		ExternalIP: svc.Spec.ExternalIPs,
		Ports:      toStringPorts(svc),
	}
}

func toStringPorts(svc v1.Service) string {
	sports := ""
	for _, port := range svc.Spec.Ports {
		sport := ""
		if svc.Spec.Type == v1.ServiceType(common.SERVICE_TYPE_NODE) {
			sport = fmt.Sprintf("%d:%d/%s", port.Port, port.NodePort, port.Protocol)
		} else {
			sport = fmt.Sprintf("%d/%s", port.Port, port.Protocol)
		}

		sports = sports + "," + sport
	}
	sports = strings.TrimLeft(sports, ",")
	return sports
}


