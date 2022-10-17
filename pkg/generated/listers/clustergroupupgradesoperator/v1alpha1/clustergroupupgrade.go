/*
Copyright 2021.

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
	v1alpha1 "github.com/openshift-kni/cluster-group-upgrades-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterGroupUpgradeLister helps list ClusterGroupUpgrades.
// All objects returned here must be treated as read-only.
type ClusterGroupUpgradeLister interface {
	// List lists all ClusterGroupUpgrades in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterGroupUpgrade, err error)
	// ClusterGroupUpgrades returns an object that can list and get ClusterGroupUpgrades.
	ClusterGroupUpgrades(namespace string) ClusterGroupUpgradeNamespaceLister
	ClusterGroupUpgradeListerExpansion
}

// clusterGroupUpgradeLister implements the ClusterGroupUpgradeLister interface.
type clusterGroupUpgradeLister struct {
	indexer cache.Indexer
}

// NewClusterGroupUpgradeLister returns a new ClusterGroupUpgradeLister.
func NewClusterGroupUpgradeLister(indexer cache.Indexer) ClusterGroupUpgradeLister {
	return &clusterGroupUpgradeLister{indexer: indexer}
}

// List lists all ClusterGroupUpgrades in the indexer.
func (s *clusterGroupUpgradeLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterGroupUpgrade, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterGroupUpgrade))
	})
	return ret, err
}

// ClusterGroupUpgrades returns an object that can list and get ClusterGroupUpgrades.
func (s *clusterGroupUpgradeLister) ClusterGroupUpgrades(namespace string) ClusterGroupUpgradeNamespaceLister {
	return clusterGroupUpgradeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterGroupUpgradeNamespaceLister helps list and get ClusterGroupUpgrades.
// All objects returned here must be treated as read-only.
type ClusterGroupUpgradeNamespaceLister interface {
	// List lists all ClusterGroupUpgrades in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterGroupUpgrade, err error)
	// Get retrieves the ClusterGroupUpgrade from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterGroupUpgrade, error)
	ClusterGroupUpgradeNamespaceListerExpansion
}

// clusterGroupUpgradeNamespaceLister implements the ClusterGroupUpgradeNamespaceLister
// interface.
type clusterGroupUpgradeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterGroupUpgrades in the indexer for a given namespace.
func (s clusterGroupUpgradeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterGroupUpgrade, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterGroupUpgrade))
	})
	return ret, err
}

// Get retrieves the ClusterGroupUpgrade from the indexer for a given namespace and name.
func (s clusterGroupUpgradeNamespaceLister) Get(name string) (*v1alpha1.ClusterGroupUpgrade, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustergroupupgrade"), name)
	}
	return obj.(*v1alpha1.ClusterGroupUpgrade), nil
}
