package pod

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodDetail struct {
	Pod
	RestartCount int32       `json:"restartCount"`
	Namespace    string      `json:"namespace"`
	Containers   []Container `json:"containers"`
}

type Container struct {
	Name         string            `json:"name"`
	Ready        bool              `json:"ready"`
	State        v1.ContainerState `json:"status"`
	RestartCount int32             `json:"restartCount"`
	Image        string            `json:"image"`
}

func GetPodDetail(client kubernetes.Interface, namespace string, podname string) (*PodDetail, error) {
	log.Printf("begin pod %s detail in namespace %s.\n", podname, namespace)
	pod, err := client.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
	if err != nil {
		log.Printf("get pod %s detail failed in namespace %s. Error: %s\n", podname, namespace, err.Error())
		return nil, err
	}
	podDetail := toPodDetail(pod)
	log.Printf("get pod %s detail in namespace %s completed.\n", podname, namespace)
	return podDetail, err
}

func toPodDetail(pod *v1.Pod) *PodDetail {
	return &PodDetail{
		Pod:          *ToPod(*pod),
		RestartCount: getRestartCount(*pod),
		Namespace:    pod.ObjectMeta.Namespace,
	}
}

func getRestartCount(pod v1.Pod) int32 {
	var restartCount int32 = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartCount += containerStatus.RestartCount
	}
	return restartCount
}

func GetPodContainers(client kubernetes.Interface, namespace string, podname string) (*PodDetail, error) {
	log.Printf("begin get containers in pod %s in namespace %s.\n", podname, namespace)
	pod, err := client.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
	if err != nil {
		log.Printf("get containers in pod %s failed, Error: %s\n", podname, err.Error())
		return nil, err
	}
	podDetail := toPodContainers(pod)
	log.Printf("get containers in pod %s in namespace %s completed.\n", podname, namespace)
	return podDetail, err
}

func toPodContainers(pod *v1.Pod) *PodDetail {
	return &PodDetail{
		Containers: getContains(*pod),
	}
}

func getContains(pod v1.Pod) []Container {
	var containers []Container
	for _, containerStatus := range pod.Status.ContainerStatuses {
		container := Container{
			Name:         containerStatus.Name,
			Ready:        containerStatus.Ready,
			State:        containerStatus.State,
			RestartCount: containerStatus.RestartCount,
			Image:        containerStatus.Image,
		}
		containers = append(containers, container)
	}
	return containers

}
