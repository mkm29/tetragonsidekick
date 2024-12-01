package kubernetes

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	context   *context.Context
	clientset *kubernetes.Clientset
}

func NewKubernetesClient(ctx *context.Context) *KubernetesClient {
	return &KubernetesClient{
		context:   ctx,
		clientset: nil,
	}
}

// func (k *KubernetesClient) Connect(kubeconfig string) error {
// 	// creates the in-cluster config
// 	var config *rest.Config
// 	var err error
// 	if kubeconfig == "" {
// 		log.Printf("using in-cluster configuration")
// 		config, err = rest.InClusterConfig()
// 	} else {
// 		log.Printf("using configuration from '%s'", kubeconfig)
// 		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	// Register custom resources (operator)
// 	// v1alpha1.AddToScheme(scheme.Scheme)

// 	crdConfig := *config
// 	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "", Version: "v1"}
// 	crdConfig.APIPath = "/apis"
// 	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
// 	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

// 	client, err := rest.UnversionedRESTClientFor(&crdConfig)
// 	if err != nil {
// 		return err
// 	}
// 	k.client = client
// 	return nil
// }

func (k *KubernetesClient) Connect(kubeconfig string) error {
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		return err
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	k.clientset = cs
	return nil
}

func (k *KubernetesClient) GetPods() (*corev1.PodList, error) {
	listOptions := metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=tetragon",
	}
	pods, err := k.clientset.CoreV1().Pods("kube-system").List(*k.context, listOptions)
	if err != nil {
		return nil, err
	}
	return pods, nil
}
