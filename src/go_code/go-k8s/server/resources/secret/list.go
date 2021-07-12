package secret

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretList struct {
	Secrets []Secret
}

type Secret struct {
	Name       string            `json:"name"`
	Type       string            `josn:"type"`
	StringData map[string]string `json:"stringData"`
	Namespace  string            `json:"namespace"`
	CreateTime metav1.Time       `json:"createTime"`
}

func GetSecretList(client kubernetes.Interface, namespace string) (*SecretList, error) {
	log.Printf("getting secret list in namespace %s\n", namespace)
	secretList, err := client.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println("get secret failed")
		return nil, err
	}
	secretlist := toSecretList(secretList)
	return &SecretList{Secrets: secretlist}, nil
}

func toSecretList(secretList *v1.SecretList) []Secret {
	var secretlist []Secret
	for _, secret := range secretList.Items {
		sce := toSecret(secret)
		secretlist = append(secretlist, *sce)
	}
	return secretlist
}

func toSecret(secret v1.Secret) *Secret {
	return &Secret{
		Name:       secret.ObjectMeta.Name,
		Type:       string(secret.Type),
		StringData: secret.StringData,
		Namespace:  secret.ObjectMeta.Namespace,
		CreateTime: secret.ObjectMeta.CreationTimestamp,
	}
}
