package secret

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretDetail struct {
	Secret
	Data map[string][]byte `json:"data"`
}

func GetSecretDetail(client kubernetes.Interface, ns string, secretName string) (*SecretDetail, error) {
	log.Printf("getting secret %s detail in namespace %s\n", secretName, ns)
	secret, err := client.CoreV1().Secrets(ns).Get(secretName, metav1.GetOptions{})
	if err != nil {
		log.Println("get secret detail failed")
		return nil, err
	}
	sec := toSecret(*secret)
	return &SecretDetail{*sec, secret.Data}, nil
}
