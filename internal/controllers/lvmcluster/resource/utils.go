/*
Copyright © 2023 Red Hat, Inc.

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

package resource

import (
	"fmt"

	"github.com/openshift/lvm-operator/internal/controllers/constants"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func setDaemonsetNodeSelector(nodeSelector *corev1.NodeSelector, ds *appsv1.DaemonSet) {
	if nodeSelector != nil {
		ds.Spec.Template.Spec.Affinity = &corev1.Affinity{
			NodeAffinity: &corev1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: nodeSelector,
			},
		}
	} else {
		ds.Spec.Template.Spec.Affinity = nil
	}
}

func GetStorageClassName(deviceName string) string {
	return constants.StorageClassPrefix + deviceName
}

func GetVolumeSnapshotClassName(deviceName string) string {
	return constants.VolumeSnapshotClassPrefix + deviceName
}

func verifyDaemonSetReadiness(ds *appsv1.DaemonSet) error {
	// If the update strategy is not a rolling update, there will be nothing to wait for
	if ds.Spec.UpdateStrategy.Type != appsv1.RollingUpdateDaemonSetStrategyType {
		return nil
	}

	// Make sure all the updated pods have been scheduled
	if ds.Status.UpdatedNumberScheduled != ds.Status.DesiredNumberScheduled {
		return fmt.Errorf("the DaemonSet is not ready: %s/%s. %d out of %d expected pods have been scheduled", ds.Namespace, ds.Name, ds.Status.UpdatedNumberScheduled, ds.Status.DesiredNumberScheduled)
	}
	maxUnavailable, err := intstr.GetScaledValueFromIntOrPercent(ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable, int(ds.Status.DesiredNumberScheduled), true)
	if err != nil {
		// If for some reason the value is invalid, set max unavailable to the
		// number of desired replicas. This is the same behavior as the
		// `MaxUnavailable` function in deploymentutil
		maxUnavailable = int(ds.Status.DesiredNumberScheduled)
	}

	expectedReady := int(ds.Status.DesiredNumberScheduled) - maxUnavailable
	if !(int(ds.Status.NumberReady) >= expectedReady) {
		return fmt.Errorf("the DaemonSet is not ready: %s/%s. %d out of %d expected pods are ready", ds.Namespace, ds.Name, ds.Status.NumberReady, expectedReady)
	}
	return nil
}

func verifyDeploymentReadiness(dep *appsv1.Deployment) error {
	if len(dep.Status.Conditions) == 0 {
		return fmt.Errorf("the Deployment %s/%s is not ready as no condition was found", dep.Namespace, dep.Name)
	}
	for _, condition := range dep.Status.Conditions {
		if condition.Type == appsv1.DeploymentAvailable && condition.Status == corev1.ConditionFalse {
			return fmt.Errorf("the Deployment %s/%s has not reached minimum availability and is not ready: %v",
				dep.Namespace, dep.Name, condition)
		} else if condition.Type == appsv1.DeploymentProgressing && condition.Status == corev1.ConditionFalse {
			return fmt.Errorf("the Deployment %s/%s has not progressed and is not ready: %v",
				dep.Namespace, dep.Name, condition)
		} else if condition.Type == appsv1.DeploymentReplicaFailure && condition.Status == corev1.ConditionTrue {
			return fmt.Errorf("the Deployment %s/%s has a replica failure and is not ready: %v",
				dep.Namespace, dep.Name, condition)
		}
	}
	return nil
}
