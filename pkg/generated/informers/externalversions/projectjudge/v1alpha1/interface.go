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
	internalinterfaces "projectjudge/projects/wallclock-controller/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Timezones returns a TimezoneInformer.
	Timezones() TimezoneInformer
	// WallClocks returns a WallClockInformer.
	WallClocks() WallClockInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Timezones returns a TimezoneInformer.
func (v *version) Timezones() TimezoneInformer {
	return &timezoneInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// WallClocks returns a WallClockInformer.
func (v *version) WallClocks() WallClockInformer {
	return &wallClockInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
