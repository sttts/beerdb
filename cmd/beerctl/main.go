package main

import (
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/sttts/beerdb/generated/clientset/versioned"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
	if err != nil {
		panic(err)
	}

	client, err := versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	beers, err := client.BeerdbV1().Beers("").List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, b := range beers.Items {
		fmt.Printf("%s/%s\n", b.Namespace, b.Name)
	}
}
