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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	projectjudgev1alpha1 "projectjudge/projects/wallclock-controller/pkg/apis/projectjudge/v1alpha1"
	versioned "projectjudge/projects/wallclock-controller/pkg/generated/clientset/versioned"
	internalinterfaces "projectjudge/projects/wallclock-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "projectjudge/projects/wallclock-controller/pkg/generated/listers/projectjudge/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WallClockInformer provides access to a shared informer and lister for
// WallClocks.
type WallClockInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WallClockLister
}

type wallClockInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewWallClockInformer constructs a new informer for WallClock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWallClockInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWallClockInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredWallClockInformer constructs a new informer for WallClock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWallClockInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectjudgeV1alpha1().WallClocks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectjudgeV1alpha1().WallClocks().Watch(context.TODO(), options)
			},
		},
		&projectjudgev1alpha1.WallClock{},
		resyncPeriod,
		indexers,
	)
}

func (f *wallClockInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWallClockInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *wallClockInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectjudgev1alpha1.WallClock{}, f.defaultInformer)
}

func (f *wallClockInformer) Lister() v1alpha1.WallClockLister {
	return v1alpha1.NewWallClockLister(f.Informer().GetIndexer())
}