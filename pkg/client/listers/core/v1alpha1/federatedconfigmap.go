/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedConfigMapLister helps list FederatedConfigMaps.
type FederatedConfigMapLister interface {
	// List lists all FederatedConfigMaps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedConfigMap, err error)
	// FederatedConfigMaps returns an object that can list and get FederatedConfigMaps.
	FederatedConfigMaps(namespace string) FederatedConfigMapNamespaceLister
	FederatedConfigMapListerExpansion
}

// federatedConfigMapLister implements the FederatedConfigMapLister interface.
type federatedConfigMapLister struct {
	indexer cache.Indexer
}

// NewFederatedConfigMapLister returns a new FederatedConfigMapLister.
func NewFederatedConfigMapLister(indexer cache.Indexer) FederatedConfigMapLister {
	return &federatedConfigMapLister{indexer: indexer}
}

// List lists all FederatedConfigMaps in the indexer.
func (s *federatedConfigMapLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedConfigMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedConfigMap))
	})
	return ret, err
}

// FederatedConfigMaps returns an object that can list and get FederatedConfigMaps.
func (s *federatedConfigMapLister) FederatedConfigMaps(namespace string) FederatedConfigMapNamespaceLister {
	return federatedConfigMapNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedConfigMapNamespaceLister helps list and get FederatedConfigMaps.
type FederatedConfigMapNamespaceLister interface {
	// List lists all FederatedConfigMaps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedConfigMap, err error)
	// Get retrieves the FederatedConfigMap from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedConfigMap, error)
	FederatedConfigMapNamespaceListerExpansion
}

// federatedConfigMapNamespaceLister implements the FederatedConfigMapNamespaceLister
// interface.
type federatedConfigMapNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedConfigMaps in the indexer for a given namespace.
func (s federatedConfigMapNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedConfigMap, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedConfigMap))
	})
	return ret, err
}

// Get retrieves the FederatedConfigMap from the indexer for a given namespace and name.
func (s federatedConfigMapNamespaceLister) Get(name string) (*v1alpha1.FederatedConfigMap, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedconfigmap"), name)
	}
	return obj.(*v1alpha1.FederatedConfigMap), nil
}
