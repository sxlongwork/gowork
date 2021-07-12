package service

import (
	"k8s-manage/server/resources/pod"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func GetServicePod(client kubernetes.Interface, namespace string, svcName string) (*pod.PodList, error) {
	log.Printf("Getting service %s pods in namespace %s\n", svcName, namespace)
	service, err := client.CoreV1().Services(namespace).Get(svcName, metav1.GetOptions{})
	if err != nil {
		log.Println("get service pods failed")
		return nil, err
	}
	if service.Spec.Selector == nil {
		log.Println("service.Spec.Selector is nil")
		return nil, err
	}
	labelSelector := labels.SelectorFromSet(service.Spec.Selector)
	podList, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector.String(),
		FieldSelector: fields.Everything().String(),
	})
	if err != nil {
		log.Println("getting service pod list failed")
		return nil, err
	}
	podlist := pod.ToPodList(podList)
	return &pod.PodList{Pods: podlist}, nil
}
