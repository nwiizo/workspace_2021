package main

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"
)

func Test_envtest(t *testing.T) {
	testEnv := &envtest.Environment{}

	restConfig, err := testEnv.Start()
	if err != nil {
		t.Error(err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		t.Error(err)
	}

	cm := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cm",
			Namespace: "default",
		},
		Data: map[string]string{
			"foo": "bar",
		},
	}

	fmt.Printf("cm: %s\n", pp.Sprint(cm))
	_, err = clientset.CoreV1().ConfigMaps("default").Create(context.TODO(), cm, metav1.CreateOptions{})
	if err != nil {
		t.Error(err)
	}
	got, err := clientset.CoreV1().ConfigMaps("default").Get(context.TODO(), "test-cm", metav1.GetOptions{})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("got: %s\n", pp.Sprint(got))
	if got.Name != cm.Name {
		t.Errorf("missmatch: got = %v, want = %v", got.Name, cm.Name)
	}

	err = testEnv.Stop()
	if err != nil {
		t.Error(err)
	}
}
