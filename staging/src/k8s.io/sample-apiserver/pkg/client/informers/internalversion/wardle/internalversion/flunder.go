/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package internalversion

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	wardle "k8s.io/sample-apiserver/pkg/apis/wardle"
	clientset_internalversion "k8s.io/sample-apiserver/pkg/client/clientset/internalversion"
	internalinterfaces "k8s.io/sample-apiserver/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/sample-apiserver/pkg/client/listers/wardle/internalversion"
)

// FlunderInformer provides access to a shared informer and lister for
// Flunders.
type FlunderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FlunderLister
}

type flunderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFlunderInformer constructs a new informer for Flunder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlunderInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFlunderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFlunderInformer constructs a new informer for Flunder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlunderInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Wardle().Flunders(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Wardle().Flunders(namespace).Watch(options)
			},
		},
		&wardle.Flunder{},
		resyncPeriod,
		indexers,
	)
}

func (f *flunderInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFlunderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flunderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wardle.Flunder{}, f.defaultInformer)
}

func (f *flunderInformer) Lister() internalversion.FlunderLister {
	return internalversion.NewFlunderLister(f.Informer().GetIndexer())
}
