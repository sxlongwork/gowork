package pod

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodList struct {
	Pods []Pod `json:"pods"`
}

type Pod struct {
	PodName    string      `json:"podName"`
	Status     v1.PodPhase `json:"status"`
	PodIP      string      `json:"podIP"`
	NodeName   string      `json:"nodeName"`
	CreateTime metav1.Time `json:"createTime"`
}

func GetPodList(client kubernetes.Interface, namespace string) (*PodList, error) {
	log.Printf("begin get pods in namespace %s.\n", namespace)
	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get pod list in namespace%s failed, Error: %s\n", namespace, err.Error())
		return nil, err
	}
	podlist := ToPodList(pods)
	log.Printf("get pods in namespace %s completed.\n", namespace)
	return &PodList{Pods: podlist}, err
}

func ToPodList(pods *v1.PodList) []Pod {
	var mpods []Pod
	for _, pod := range pods.Items {
		mpod := *ToPod(pod)
		mpods = append(mpods, mpod)
	}
	return mpods
}

func ToPod(pod v1.Pod) *Pod {
	return &Pod{
		PodName:    pod.ObjectMeta.Name,
		Status:     pod.Status.Phase,
		PodIP:      pod.Status.PodIP,
		NodeName:   pod.Spec.NodeName,
		CreateTime: pod.ObjectMeta.CreationTimestamp,
	}
}
