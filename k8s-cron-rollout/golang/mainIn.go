package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	namespace      string
	deploymentName string
	changeCause    string
)

func init() {
	flag.StringVar(&namespace, "namespace", "", "K8s namespace.")
	flag.StringVar(&deploymentName, "deployment-name", "", "Deployment name to rollout.")
	flag.StringVar(&changeCause, "change-cause", "cronjob execution", "Change cause reason for annotation kubernetes.io/change-cause")
	flag.Parse()
}

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// set the namespace
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	fmt.Printf("Listing deployments in namespace %q:\n", namespace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	// list deployment information
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
	// set the annotation and restart
	data := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%s","kubernetes.io/change-cause":"%s"}}}}}`, time.Now().String(), changeCause)
	_, err = deploymentsClient.Patch(context.Background(), deploymentName, types.StrategicMergePatchType, []byte(data), metav1.PatchOptions{FieldManager: "kubectl-rollout"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Rollout execution success. Check running: kubectl rollout history deployment/<deployment-name>")
}
