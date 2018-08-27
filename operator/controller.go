package main

import (
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/record"

	clientset "github.com/lx/kubernete-controller/pkg/client/clientset/versioned"
	listers "github.com/lx/kubernete-controller/pkg/client/listers/samplecontroller/v1"

)

type Controller struct {
	kubeclientset kubernetes.Interface
	sampleclientset clientset.Interface

	deploymentsLister appslisters.DeploymentLister
	deploymentsSynced cache.InformerSynced
	fooLister listers.FooLister
	fooSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface

	recorder record.EventRecorder
}