package main

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your kubeconfig file ")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("this is an error ", err)
	}
	fmt.Println(config)
}
