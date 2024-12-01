/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"log"

	k8s "github.com/mkm29/tetragonsidekick/pkg/kubernetes"
)

var (
	kubeconfig string
	cluster    string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.StringVar(&cluster, "cluster", "", "cluster name")
	flag.Parse()
}

func main() {
	// create context
	ctx := context.Background()
	// Create Kubernetes client
	client := k8s.NewKubernetesClient(&ctx)
	// Connect to the cluster
	if err := client.Connect(kubeconfig); err != nil {
		panic(err)
	}
	// Get pods
	pods, err := client.GetPods()
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		// get Pod metadata
		log.Printf("pod: %v", pod.ObjectMeta)
	}
}
