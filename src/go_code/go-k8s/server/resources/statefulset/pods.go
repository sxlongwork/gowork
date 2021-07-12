package statefulset

import (
	"k8s-manage/server/resources/common"
	"k8s-manage/server/resources/pod"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetStatefulsetPods(client kubernetes.Interface, ns string, name string) (*pod.PodList, error) {
	log.Printf("Getting replication controller %s pods in namespace %s", name, ns)
	statefulSetDetail, err := client.AppsV1().StatefulSets(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("get stateful pods failed, Error: %s", err.Error())
		return nil, err
	}
	podList, err := client.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get pod list in namespace %s failed, Error: %s", ns, err.Error())
		return nil, err
	}
	pods := common.GetPodsByControllerRef(statefulSetDetail, podList)
	if err != nil {
		return nil, err
	}
	podsList := toPodsList(pods)
	return &pod.PodList{Pods: podsList}, err
}

func toPodsList(pods []v1.Pod) []pod.Pod {
	var podsList []pod.Pod
	for _, onePod := range pods {
		mypod := pod.ToPod(onePod)
		podsList = append(podsList, *mypod)
	}
	return podsList
}
