package deploy

import (
	"k8s-manage/server/resources/common"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func FindOldReplicaSets(deployment *apps.Deployment, replicasetsList []*apps.ReplicaSet) ([]*apps.ReplicaSet, error) {
	var requiredRSs []*apps.ReplicaSet
	newRS := FindNewReplicaSet(deployment, replicasetsList)
	for _, rs := range replicasetsList {
		// Filter out new replica set
		if newRS != nil && rs.UID == newRS.UID {
			continue
		}
		if *(rs.Spec.Replicas) != 0 {
			requiredRSs = append(requiredRSs, rs)
		}
	}
	return requiredRSs, nil
}

func FindNewReplicaSet(deployment *apps.Deployment, replicasetsList []*apps.ReplicaSet) *apps.ReplicaSet {
	newRsTemplate := GetNewRsTemplate(deployment)
	for i := range replicasetsList {
		if common.EqualIgnoreHash(replicasetsList[i].Spec.Template, newRsTemplate) {
			// This is the new ReplicaSet.
			return replicasetsList[i]
		}
	}
	// new ReplicaSet does not exist.
	return nil
}

func GetNewRsTemplate(deployment *apps.Deployment) v1.PodTemplateSpec {
	return v1.PodTemplateSpec{
		ObjectMeta: deployment.Spec.Template.ObjectMeta,
		Spec:       deployment.Spec.Template.Spec,
	}
}
