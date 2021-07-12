package replicaset

import (
	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ReplicaSetList struct {
	Replicaset []*ReplicaSet
}

type ReplicaSet struct {
	Name         string      `json:"name"`
	Desired      int32       `json:"desired"`
	CurrentCount int32       `json:"currentCount"`
	ReadyCount   int32       `json:"readyCount"`
	CreateTime   metav1.Time `json:"createTime"`
}

func GetReplicaSetList(client kubernetes.Interface, ns string, options metav1.ListOptions) ([]*apps.ReplicaSet, error) {
	var replicasetList []*apps.ReplicaSet
	replicasets, err := client.AppsV1().ReplicaSets(ns).List(options)
	if err != nil {
		return nil, err
	}
	for _, replicaset := range replicasets.Items {
		replicasetList = append(replicasetList, &replicaset)
	}
	return replicasetList, err
}

func ToReplicaSetList(replicasets []*apps.ReplicaSet) *ReplicaSetList {
	var replicaSetList []*ReplicaSet
	for _, replicaset := range replicasets {
		replicaSet := ToReplicaSet(replicaset)
		replicaSetList = append(replicaSetList, replicaSet)
	}
	return &ReplicaSetList{Replicaset: replicaSetList}
}

func ToReplicaSet(replicaset *apps.ReplicaSet) *ReplicaSet {
	return &ReplicaSet{
		Name:         replicaset.ObjectMeta.Name,
		Desired:      *replicaset.Spec.Replicas,
		CurrentCount: replicaset.Status.Replicas,
		ReadyCount:   replicaset.Status.ReadyReplicas,
		CreateTime:   replicaset.ObjectMeta.CreationTimestamp,
	}
}
