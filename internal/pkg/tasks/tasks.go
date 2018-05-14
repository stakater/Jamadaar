package tasks

import (
	"log"

	"github.com/stakater/Jamadaar/internal/pkg/actions"
	"github.com/stakater/Jamadaar/internal/pkg/tasks/namespaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// PerformTasks handles all the cleanup tasks
func PerformTasks(clientset clientset.Interface, actions []actions.Action, age string) {
	performNamespaceDeletion(clientset, actions, age)
}

// performNamespaceDeletion handles the deletion of namespaces
func performNamespaceDeletion(clientset clientset.Interface, actions []actions.Action, age string) {
	log.Println("Starting to delete Namespaces")
	namespaceList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting namespaces: %v", err)
		return
	}

	err = namespaces.DeleteNamespaces(clientset, namespaceList, actions, age)
	if err != nil {
		log.Printf("Error deleting namespaces: %v", err)
		return
	}
}