package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	beerdbv1 "github.com/sttts/beerdb/apis/beerdb/v1"
	"github.com/sttts/beerdb/generated/clientset/versioned"
	"github.com/sttts/beerdb/generated/informers/externalversions"
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

	informers := externalversions.NewSharedInformerFactory(client, time.Minute*30)

	// react on events
	beerInformer := informers.Beerdb().V1().Beers().Informer()
	beerInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			beer := obj.(*beerdbv1.Beer)
			fmt.Printf("New beer %s/%s\n", beer.Namespace, beer.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			beer := newObj.(*beerdbv1.Beer)
			fmt.Printf("Updated beer %s/%s\n", beer.Namespace, beer.Name)
		},
		DeleteFunc: func(obj interface{}) {
			beer := obj.(*beerdbv1.Beer)
			fmt.Printf("Deleted beer %s/%s\n", beer.Namespace, beer.Name)
		},
	})

	// start informer
	stopCh := make(chan struct{})
	informers.Start(stopCh)
	informers.WaitForCacheSync(stopCh)

	<-stopCh
}
