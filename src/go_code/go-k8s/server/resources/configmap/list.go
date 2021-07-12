package configmap

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapList struct {
	ConfigMaps []ConfigMap
}

type ConfigMap struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	CreateTime metav1.Time
}

func GetConfigMapList(client kubernetes.Interface, namespace string) (*ConfigMapList, error) {
	log.Printf("getting configmap list in namespace %s", namespace)
	configMapList, err := client.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println("get configmap list failed")
	}
	configmaplist := toConfigMapList(configMapList)
	return &ConfigMapList{configmaplist}, nil
}

func toConfigMapList(configMapList *v1.ConfigMapList) []ConfigMap {
	var configmapList []ConfigMap
	for _, config := range configMapList.Items {
		configmap := toConfigMap(config)
		configmapList = append(configmapList, *configmap)
	}
	return configmapList
}

func toConfigMap(configMap v1.ConfigMap) *ConfigMap {
	return &ConfigMap{
		Name:       configMap.ObjectMeta.Name,
		Namespace:  configMap.ObjectMeta.Namespace,
		CreateTime: configMap.ObjectMeta.CreationTimestamp,
	}
}
