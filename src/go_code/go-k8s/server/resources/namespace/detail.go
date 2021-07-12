package namespace

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceDetail struct {
	Namespace
}

func GetNamespaceDetail(client kubernetes.Interface, nsname string) (*NamespaceDetail, error) {
	log.Printf("begin get namespace %s detail.\n", nsname)
	ns, err := client.CoreV1().Namespaces().Get(nsname, metav1.GetOptions{})
	if err != nil {
		// panic(err)
		log.Printf("get namespace %s detail failed. Error: %s\n", nsname, err.Error())
		return nil, err
	}
	nsDetail := toNamespaceDetail(ns)
	return nsDetail, err
}

func toNamespaceDetail(ns *v1.Namespace) *NamespaceDetail {
	namespace := Namespace{
		Name:       ns.ObjectMeta.Name,
		Status:     ns.Status.Phase,
		CreateTime: ns.ObjectMeta.CreationTimestamp,
	}
	return &NamespaceDetail{Namespace: namespace}
}
