/*
Copyright 2024 Intel Corporation

Licensed under the Apache License, Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	diskiov1alpha1 "github.com/intel/cloud-resource-scheduling-and-isolation/pkg/api/diskio/v1alpha1"
	versioned "github.com/intel/cloud-resource-scheduling-and-isolation/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/intel/cloud-resource-scheduling-and-isolation/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/intel/cloud-resource-scheduling-and-isolation/pkg/generated/listers/diskio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeDiskDeviceInformer provides access to a shared informer and lister for
// NodeDiskDevices.
type NodeDiskDeviceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeDiskDeviceLister
}

type nodeDiskDeviceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNodeDiskDeviceInformer constructs a new informer for NodeDiskDevice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeDiskDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeDiskDeviceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNodeDiskDeviceInformer constructs a new informer for NodeDiskDevice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeDiskDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiskioV1alpha1().NodeDiskDevices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiskioV1alpha1().NodeDiskDevices(namespace).Watch(context.TODO(), options)
			},
		},
		&diskiov1alpha1.NodeDiskDevice{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeDiskDeviceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeDiskDeviceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeDiskDeviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&diskiov1alpha1.NodeDiskDevice{}, f.defaultInformer)
}

func (f *nodeDiskDeviceInformer) Lister() v1alpha1.NodeDiskDeviceLister {
	return v1alpha1.NewNodeDiskDeviceLister(f.Informer().GetIndexer())
}
