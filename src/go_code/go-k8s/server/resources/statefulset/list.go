package statefulset

import (
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type StatefulSetList struct {
	StatefulsetList []StatefulSet `json:"statefulsetList"`
}

type StatefulSet struct {
	Name          string      `json:"name"`
	Replicas      int32       `json:"replicas"`      // replicas is the desired number of replicas of the given Template.
	ReadyReplicas int32       `json:"readyReplicas"` // readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	Namespace     string      `json:"namespace"`
	CreateTime    metav1.Time `json:"createTime"`
}

func GetStatefulSetList(client kubernetes.Interface, ns string) (*StatefulSetList, error) {
	log.Println("begin get statefulset list")
	statefulSetList, err := client.AppsV1().StatefulSets(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Println("get stateful list failed")
		return nil, err
	}
	statefulsetList := toStatefulSetList(statefulSetList)
	log.Println("get statefulset list completed")
	return &StatefulSetList{StatefulsetList: statefulsetList}, err
}

func toStatefulSetList(statefulSetList *v1.StatefulSetList) []StatefulSet {
	var statefulsetlist []StatefulSet
	for _, statefulSet := range statefulSetList.Items {
		statefulset := toStatefulSet(statefulSet)
		statefulsetlist = append(statefulsetlist, statefulset)
	}
	return statefulsetlist
}

func toStatefulSet(statefulSet v1.StatefulSet) StatefulSet {
	return StatefulSet{
		Name:          statefulSet.ObjectMeta.Name,
		Replicas:      *statefulSet.Spec.Replicas,
		ReadyReplicas: statefulSet.Status.ReadyReplicas,
		Namespace:     statefulSet.ObjectMeta.Namespace,
		CreateTime:    statefulSet.ObjectMeta.CreationTimestamp,
	}
}
