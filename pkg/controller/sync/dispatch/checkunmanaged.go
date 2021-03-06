/*
Copyright 2019 The Kubernetes Authors.

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

package dispatch

import (
	"github.com/pkg/errors"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"

	"sigs.k8s.io/kubefed/pkg/controller/util"
)

type isNamespaceInHostClusterFunc func(clusterObj pkgruntime.Object) bool

type CheckUnmanagedDispatcher interface {
	OperationDispatcher

	CheckRemovedOrUnlabeled(clusterName string, isHostNamespace isNamespaceInHostClusterFunc)
}

type checkUnmanagedDispatcherImpl struct {
	dispatcher *operationDispatcherImpl

	targetName util.QualifiedName
	targetKind string
}

func NewCheckUnmanagedDispatcher(clientAccessor clientAccessorFunc, targetKind string, targetName util.QualifiedName) CheckUnmanagedDispatcher {
	dispatcher := newOperationDispatcher(clientAccessor, nil)
	return &checkUnmanagedDispatcherImpl{
		dispatcher: dispatcher,
		targetName: targetName,
		targetKind: targetKind,
	}
}

func (d *checkUnmanagedDispatcherImpl) Wait() (bool, error) {
	return d.dispatcher.Wait()
}

// CheckRemovedOrUnlabeled checks that a resource either does not
// exist in the given cluster, or if it does exist, that it does not
// have the managed label.
func (d *checkUnmanagedDispatcherImpl) CheckRemovedOrUnlabeled(clusterName string, isHostNamespace isNamespaceInHostClusterFunc) {
	d.dispatcher.incrementOperationsInitiated()
	const op = "check for deletion of resource or removal of managed label from"
	const opContinuous = "Checking for deletion of resource or removal of managed label from"
	go d.dispatcher.clusterOperation(clusterName, op, func(client util.ResourceClient) util.ReconciliationStatus {
		klog.V(2).Infof(eventTemplate, opContinuous, d.targetKind, d.targetName, clusterName)

		clusterObj, err := client.Resources(d.targetName.Namespace).Get(d.targetName.Name, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return util.StatusAllOK
		}
		if err != nil {
			wrappedErr := d.wrapOperationError(err, clusterName, op)
			runtime.HandleError(wrappedErr)
			return util.StatusError
		}
		if clusterObj.GetDeletionTimestamp() != nil {
			if isHostNamespace(clusterObj) {
				return util.StatusAllOK
			}
			err = errors.Errorf("resource is pending deletion")
			wrappedErr := d.wrapOperationError(err, clusterName, op)
			runtime.HandleError(wrappedErr)
			return util.StatusError
		}
		if !util.HasManagedLabel(clusterObj) {
			return util.StatusAllOK
		}
		err = errors.Errorf("resource still has the managed label")
		wrappedErr := d.wrapOperationError(err, clusterName, op)
		runtime.HandleError(wrappedErr)
		return util.StatusError
	})
}

func (d *checkUnmanagedDispatcherImpl) wrapOperationError(err error, clusterName, operation string) error {
	return wrapOperationError(err, operation, d.targetKind, d.targetName.String(), clusterName)
}
