package main

import (
	"fmt"
	"sort"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/batch"
	apiscertificates "k8s.io/kubernetes/pkg/apis/certificates"
	"k8s.io/kubernetes/pkg/apis/core"

	"k8s.io/kubernetes/pkg/registry/apps/replicaset"
	"k8s.io/kubernetes/pkg/registry/batch/job"
	"k8s.io/kubernetes/pkg/registry/certificates/certificates"
	"k8s.io/kubernetes/pkg/registry/core/configmap"
	"k8s.io/kubernetes/pkg/registry/core/event"
	"k8s.io/kubernetes/pkg/registry/core/namespace"
	"k8s.io/kubernetes/pkg/registry/core/node"
	"k8s.io/kubernetes/pkg/registry/core/persistentvolume"
	"k8s.io/kubernetes/pkg/registry/core/persistentvolumeclaim"
	"k8s.io/kubernetes/pkg/registry/core/pod"
	"k8s.io/kubernetes/pkg/registry/core/replicationcontroller"
	"k8s.io/kubernetes/pkg/registry/core/secret"
)

func main() {
	strategies := []struct {
		name    string
		getAttr func(obj runtime.Object) (labels.Set, fields.Set, error)
		obj     runtime.Object
	}{
		{
			name:    "configmap",
			getAttr: configmap.GetAttrs,
			obj:     &core.ConfigMap{},
		},
		{
			name:    "event",
			getAttr: event.GetAttrs,
			obj:     &core.Event{},
		},
		{
			name:    "namespace",
			getAttr: namespace.GetAttrs,
			obj:     &core.Namespace{},
		},
		{
			name:    "node",
			getAttr: node.GetAttrs,
			obj:     &core.Node{},
		},
		{
			name:    "persistentvolume",
			getAttr: persistentvolume.GetAttrs,
			obj:     &core.PersistentVolume{},
		},
		{
			name:    "persistentvolumeclaim",
			getAttr: persistentvolumeclaim.GetAttrs,
			obj:     &core.PersistentVolumeClaim{},
		},
		{
			name:    "pod",
			getAttr: pod.GetAttrs,
			obj:     &core.Pod{},
		},
		{
			name:    "replicationcontroller",
			getAttr: replicationcontroller.GetAttrs,
			obj:     &core.ReplicationController{},
		},
		{
			name:    "secret",
			getAttr: secret.GetAttrs,
			obj:     &core.Secret{},
		},
		{
			name:    "replicaset",
			getAttr: replicaset.GetAttrs,
			obj:     &apps.ReplicaSet{},
		},
		{
			name:    "job",
			getAttr: job.GetAttrs,
			obj:     &batch.Job{},
		},
		{
			name:    "certificatesigningrequest",
			getAttr: certificates.GetAttrs,
			obj:     &apiscertificates.CertificateSigningRequest{},
		},
	}
	for _, strategy := range strategies {
		_, fields, err := strategy.getAttr(strategy.obj)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n%q filterable fields:\n", strategy.name)
		names := make([]string, 0, len(fields))
		for field := range fields {
			names = append(names, field)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("- %s\n", name)
		}
	}
}
