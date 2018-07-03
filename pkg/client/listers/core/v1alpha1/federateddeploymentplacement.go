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

// FederatedDeploymentPlacementLister helps list FederatedDeploymentPlacements.
type FederatedDeploymentPlacementLister interface {
	// List lists all FederatedDeploymentPlacements in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedDeploymentPlacement, err error)
	// FederatedDeploymentPlacements returns an object that can list and get FederatedDeploymentPlacements.
	FederatedDeploymentPlacements(namespace string) FederatedDeploymentPlacementNamespaceLister
	FederatedDeploymentPlacementListerExpansion
}

// federatedDeploymentPlacementLister implements the FederatedDeploymentPlacementLister interface.
type federatedDeploymentPlacementLister struct {
	indexer cache.Indexer
}

// NewFederatedDeploymentPlacementLister returns a new FederatedDeploymentPlacementLister.
func NewFederatedDeploymentPlacementLister(indexer cache.Indexer) FederatedDeploymentPlacementLister {
	return &federatedDeploymentPlacementLister{indexer: indexer}
}

// List lists all FederatedDeploymentPlacements in the indexer.
func (s *federatedDeploymentPlacementLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedDeploymentPlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedDeploymentPlacement))
	})
	return ret, err
}

// FederatedDeploymentPlacements returns an object that can list and get FederatedDeploymentPlacements.
func (s *federatedDeploymentPlacementLister) FederatedDeploymentPlacements(namespace string) FederatedDeploymentPlacementNamespaceLister {
	return federatedDeploymentPlacementNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedDeploymentPlacementNamespaceLister helps list and get FederatedDeploymentPlacements.
type FederatedDeploymentPlacementNamespaceLister interface {
	// List lists all FederatedDeploymentPlacements in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedDeploymentPlacement, err error)
	// Get retrieves the FederatedDeploymentPlacement from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedDeploymentPlacement, error)
	FederatedDeploymentPlacementNamespaceListerExpansion
}

// federatedDeploymentPlacementNamespaceLister implements the FederatedDeploymentPlacementNamespaceLister
// interface.
type federatedDeploymentPlacementNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedDeploymentPlacements in the indexer for a given namespace.
func (s federatedDeploymentPlacementNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedDeploymentPlacement, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedDeploymentPlacement))
	})
	return ret, err
}

// Get retrieves the FederatedDeploymentPlacement from the indexer for a given namespace and name.
func (s federatedDeploymentPlacementNamespaceLister) Get(name string) (*v1alpha1.FederatedDeploymentPlacement, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federateddeploymentplacement"), name)
	}
	return obj.(*v1alpha1.FederatedDeploymentPlacement), nil
}
