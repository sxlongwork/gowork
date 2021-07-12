package deploy

import (
	"k8s-manage/server/resources/replicaset"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentList struct {
	Deploys []Deploy
}

type Deploy struct {
	Name              string                `json:"name"`
	ReadyReplicas     int32                 `json:"readyReplicas"`     // Total number of ready pods targeted by this deployment.
	Replicas          int32                 `json:"replicas"`          // Number of desired pods
	AvailableReplicas int32                 `json:"availableReplicas"` // Total number of available pods targeted by this deployment.
	Selector          *metav1.LabelSelector `json:"selector"`          // Label selector for pods
	CreateTime        metav1.Time           `json:"createTime"`
}

func GetDeploymentList(client kubernetes.Interface, ns string) (*DeploymentList, error) {
	log.Printf("begin get deployment list.\n")
	deploymentList, err := client.AppsV1().Deployments(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("get deployment list failed, Error: %s\n", err.Error())
	}
	deployList := toDeployList(deploymentList)
	log.Printf("get deploy list completed.\n")
	return &DeploymentList{Deploys: deployList}, err
}

func toDeployList(deploymentList *v1.DeploymentList) []Deploy {
	var deployList []Deploy
	for _, deployment := range deploymentList.Items {
		deploy := *toDeploy(deployment)
		deployList = append(deployList, deploy)
	}
	return deployList
}

func toDeploy(deployment v1.Deployment) *Deploy {
	return &Deploy{
		Name:              deployment.ObjectMeta.Name,
		ReadyReplicas:     deployment.Status.ReadyReplicas,
		Replicas:          *deployment.Spec.Replicas,
		AvailableReplicas: deployment.Status.AvailableReplicas,
		Selector:          deployment.Spec.Selector,
		CreateTime:        deployment.ObjectMeta.CreationTimestamp,
	}
}

func GetDeploymentOldReplicasetList(client kubernetes.Interface, ns string, deploymentName string) (*replicaset.ReplicaSetList, error) {
	log.Printf("begin get deployment %s oldreplicaset.\n", deploymentName)
	deployment, err := client.AppsV1().Deployments(ns).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Printf("get deploy %s failed, Error: %s\n", deploymentName, err.Error())
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		log.Printf("get deploy %s label selector failed, Error: %s\n", deploymentName, err.Error())
		return nil, err
	}
	options := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	relicasetsList, err := replicaset.GetReplicaSetList(client, ns, options)
	if err != nil {
		log.Printf("get replicasets failed, Error: %s\n", err.Error())
		return nil, err
	}
	oldreplicasets, err := FindOldReplicaSets(deployment, relicasetsList)
	if err != nil {
		log.Printf("get oldreplicasets failed, Error: %s\n", err.Error())
		return nil, err
	}
	oldReplicaSetList := replicaset.ToReplicaSetList(oldreplicasets)
	return oldReplicaSetList, nil
}
func GetDeploymentNewReplicaset(client kubernetes.Interface, ns string, deploymentName string) (*replicaset.ReplicaSet, error) {
	log.Printf("begin get deployment %s newreplicaset.\n", deploymentName)
	deployment, err := client.AppsV1().Deployments(ns).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Printf("get deploy %s failed, Error: %s\n", deploymentName, err.Error())
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		log.Printf("get deploy %s label selector failed, Error: %s\n", deploymentName, err.Error())
		return nil, err
	}
	options := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	relicasetsList, err := replicaset.GetReplicaSetList(client, ns, options)
	if err != nil {
		log.Printf("get replicaset list failed, Error: %s\n", err.Error())
		return nil, err
	}
	oldreplicaset := FindNewReplicaSet(deployment, relicasetsList)
	oldReplicaSet := replicaset.ToReplicaSet(oldreplicaset)

	return oldReplicaSet, nil
}
