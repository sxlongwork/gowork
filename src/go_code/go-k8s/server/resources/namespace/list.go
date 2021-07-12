package namespace

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceList struct {
	Namespaces []Namespace
}

type Namespace struct {
	Name       string            `json:"name"`
	Status     v1.NamespacePhase `json:"status"`
	CreateTime metav1.Time       `json:"createTime"`
}

func GetNamespacesList(client kubernetes.Interface) (*NamespaceList, error) {
	log.Printf("begin get namespace list.\n")
	nsList, err := client.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get namespace list failed, Error: %s", err.Error())
		return nil, err
	}
	myNsList := toNamespaceList(nsList)
	log.Printf("get namespace list completed")
	return &NamespaceList{Namespaces: myNsList}, err
}

func toNamespaceList(list *v1.NamespaceList) []Namespace {
	var nslist []Namespace
	for _, ns := range list.Items {
		namespace := toNamespace(ns)
		nslist = append(nslist, namespace)
	}
	return nslist
}

func toNamespace(ns v1.Namespace) Namespace {
	return Namespace{
		Name:       ns.ObjectMeta.Name,
		Status:     ns.Status.Phase,
		CreateTime: ns.ObjectMeta.CreationTimestamp,
	}
}
