package node

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeDetail struct {
	Node
	PodCIDRs      string     `json:"podCIDRS"`
	Unschedulable bool       `json:"unschdulable"`
	Taints        []v1.Taint `json:"taints"`
	Addresses     []v1.NodeAddress
}

func GetNodeDetail(client kubernetes.Interface, nodeName string) (*NodeDetail, error) {
	log.Printf("begin node %s detail.\n", nodeName)
	nodeDetail, err := client.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Printf("get node %s detail failed, Error: %s\n", nodeName, err.Error())
		return nil, err
	}
	nodedetail := toNodeDetail(nodeDetail)
	return nodedetail, err
}

func toNodeDetail(nodedetail *v1.Node) *NodeDetail {
	return &NodeDetail{
		Node:          toNode(*nodedetail),
		PodCIDRs:      nodedetail.Spec.PodCIDR,
		Unschedulable: nodedetail.Spec.Unschedulable,
		Taints:        nodedetail.Spec.Taints,
		Addresses:     nodedetail.Status.Addresses,
	}
}
