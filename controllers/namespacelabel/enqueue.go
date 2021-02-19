package controllers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

// EnqueueNamespaceFromObject is an EventHandler that enqueues requests with
// the name taken from the source namespace. Example:
// An event cause by an object called Foo in the namespace Bar will enqueue
// the following request:
// { Name: "Bar", Namespace: "" }
type EnqueueNamespaceFromObject struct{}

var _ handler.EventHandler = &EnqueueNamespaceFromObject{}

// Create implements EventHandler
func (e *EnqueueNamespaceFromObject) Create(evt event.CreateEvent, q workqueue.RateLimitingInterface) {
	if evt.Meta == nil {
		return
	}

	addNamespaceRequest(evt.Meta, q)
}

// Update implements EventHandler
func (e *EnqueueNamespaceFromObject) Update(evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
	if evt.MetaOld != nil {
		addNamespaceRequest(evt.MetaOld, q)
	}

	if evt.MetaNew != nil {
		addNamespaceRequest(evt.MetaNew, q)
	}
}

// Delete implements EventHandler
func (e *EnqueueNamespaceFromObject) Delete(evt event.DeleteEvent, q workqueue.RateLimitingInterface) {
	if evt.Meta == nil {
		return
	}

	addNamespaceRequest(evt.Meta, q)
}

// Generic implements EventHandler
func (e *EnqueueNamespaceFromObject) Generic(evt event.GenericEvent, q workqueue.RateLimitingInterface) {
	if evt.Meta == nil {
		return
	}

	addNamespaceRequest(evt.Meta, q)
}

func addNamespaceRequest(meta metav1.Object, q workqueue.RateLimitingInterface) {
	q.Add(ctrl.Request{NamespacedName: types.NamespacedName{
		Name: meta.GetNamespace(),
	}})
}
