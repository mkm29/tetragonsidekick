package kubernetes

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	context *context.Context
	client  *rest.RESTClient
}

func NewKubernetesClient(ctx *context.Context) *KubernetesClient {
	return &KubernetesClient{
		context: ctx,
		client:  new(rest.RESTClient),
	}
}

func (k *KubernetesClient) Connect(kubeconfig string) error {
	// creates the in-cluster config
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

	// Register custom resources (operator)
	// v1alpha1.AddToScheme(scheme.Scheme)

	// crdConfig := *config
	// crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupVersion.Group, Version: v1alpha1.GroupVersion.Version}
	// crdConfig.APIPath = "/apis"
	// crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	// crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return err
	}
	k.client = client
	return nil
}

func (k *KubernetesClient) GetPods() (runtime.Object, error) {
	listOptions := metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=tetragon",
	}
	pods, err := k.client.Get().Namespace("kube-system").Resource("pods").VersionedParams(&listOptions, metav1.ParameterCodec).Do(*k.context).Get()
	// log.Printf("pods: %v", pods)
	if err != nil {
		return nil, err
	}
	return pods, nil
}
