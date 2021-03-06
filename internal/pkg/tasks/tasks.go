package tasks

import (
	"log"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"github.com/stakater/Jamadar/internal/pkg/config"
	"github.com/stakater/Jamadar/internal/pkg/tasks/namespaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// Task represents the actual tasks and actions to be taken by Jamadar
type Task struct {
	clientset clientset.Interface
	actions   []actions.Action
	config    config.Config
}

// NewTask creates a new Task object
func NewTask(clientSet clientset.Interface, actions []actions.Action, conf config.Config) *Task {
	return &Task{
		clientset: clientSet,
		actions:   actions,
		config:    conf,
	}
}

// PerformTasks handles all the cleanup tasks
func (t *Task) PerformTasks() {
	functionMap := map[string]interface{}{
		"namespaces": t.performNamespaceDeletion,
		"default":    t.performDefault,
	}
	for _, resource := range t.config.Resources {
		functionMap[resource].(func())()
	}
}

// performNamespaceDeletion handles the deletion of namespaces
func (t *Task) performNamespaceDeletion() {
	log.Println("Starting to delete Namespaces")
	namespaceList, err := t.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting namespaces: %v", err)
		return
	}
	namespace := namespaces.NewNamespaceToDelete(t.clientset, t.actions, t.config)
	err = namespace.DeleteNamespaces(namespaceList)
	if err != nil {
		log.Printf("Error deleting namespaces: %v", err)
		return
	}
}

// performDefault is the Default implementation
func (t *Task) performDefault() {
	log.Println("Performing Default Tasks.")
}
