package statefulset

import (
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type StatefulSetDetail struct {
	StatefulSet
	ServiceName string                `json:"serviceName"`
	Selector    *metav1.LabelSelector `json:"selector"`
}

func GetStatefulsetDetail(client kubernetes.Interface, ns string, name string) (*StatefulSetDetail, error) {
	log.Printf("begin get statefulset %s detail in namespace %s\n", name, ns)
	statefulSet, err := client.AppsV1().StatefulSets(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("begin get statefulset %s detail in namespace %s failed, Error: %s\n", name, ns, err.Error())
	}
	statefulsetDetail := toStatefulSetDetail(statefulSet)
	return statefulsetDetail, err
}

func toStatefulSetDetail(statefulSet *v1.StatefulSet) *StatefulSetDetail {
	statefulset := toStatefulSet(*statefulSet)
	return &StatefulSetDetail{
		StatefulSet: statefulset,
		ServiceName: statefulSet.Spec.ServiceName,
		Selector:    statefulSet.Spec.Selector,
	}
}
