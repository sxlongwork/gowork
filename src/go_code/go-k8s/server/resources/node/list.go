package node

import (
	"k8s-manage/server/resources/pod"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
)

type NodeList struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Name       string            `json:"name"`
	Status     v1.NodePhase      `json:"status"`
	Labels     map[string]string `json:"labels"`
	CreateTime metav1.Time       `json:"createTime"`
}

func GetNodesList(client kubernetes.Interface) (*NodeList, error) {
	log.Printf("begin get nodes list in cluster.\n")
	nodeList, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get nodes list failed, Error: %s\n", err)
		return nil, err
	}
	nodeslist := toNodeList(nodeList)
	return &NodeList{Nodes: nodeslist}, err
}

func toNodeList(nodelist *v1.NodeList) []Node {
	var nodes []Node
	for _, node := range nodelist.Items {
		myNode := toNode(node)
		nodes = append(nodes, myNode)
	}
	return nodes
}

func toNode(node v1.Node) Node {
	return Node{
		Name:       node.ObjectMeta.Name,
		Status:     node.Status.Phase,
		Labels:     node.ObjectMeta.Labels,
		CreateTime: node.ObjectMeta.CreationTimestamp,
	}
}

func GetNodePods(client kubernetes.Interface, nodeName string) (*pod.PodList, error) {
	log.Printf("begin get all pods in node %s\n", nodeName)
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))
	if err != nil {
		log.Printf("parseSelector failed, Error: %s", err.Error())
		return nil, err
	}
	// 获取指定的pods
	pods, err := client.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{FieldSelector: fieldSelector.String()})
	if err != nil {
		log.Printf("get all pods in node %s failed. Error: %s\n", nodeName, err.Error())
		return nil, err
	}
	podList := pod.ToPodList(pods)
	log.Printf("get all pods completed.\n")
	return &pod.PodList{Pods: podList}, err
}
