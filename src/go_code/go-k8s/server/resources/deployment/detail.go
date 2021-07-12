package deploy

import (
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentDetail struct {
	Deploy
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

func GetDeploymentDetail(client kubernetes.Interface, ns string, deName string) (*DeploymentDetail, error) {
	log.Printf("begein get deployment %s detail.\n", deName)
	deployment, err := client.AppsV1().Deployments(ns).Get(deName, metav1.GetOptions{})
	if err != nil {
		log.Printf("get deployment detail failed, Error: %s", err.Error())
		return nil, err
	}
	deployDetail := toDeployDetail(deployment)
	log.Printf("get deployment %s detail completed.\n", deName)
	return &deployDetail, err
}

func toDeployDetail(deployment *v1.Deployment) DeploymentDetail {
	return DeploymentDetail{
		Deploy:    *toDeploy(*deployment),
		Namespace: deployment.ObjectMeta.Namespace,
		Labels:    deployment.ObjectMeta.Labels,
	}
}
