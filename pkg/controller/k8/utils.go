package k8

import (
	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

const (
	AppKey = "flink-app"
)

func IsK8sObjectDoesNotExist(err error) bool {
	return k8serrors.IsNotFound(err) || k8serrors.IsGone(err) || k8serrors.IsResourceExpired(err)
}

func GetAppLabel(appName string) map[string]string {
	return map[string]string{
		AppKey: appName,
	}
}

func MergeLabels(first map[string]string, second map[string]string) map[string]string {
	result := make(map[string]string, len(first))
	for key, value := range first {
		result[key] = value
	}
	for key, value := range second {
		result[key] = value
	}
	return result
}

func MergePorts(first []coreV1.ContainerPort, second []coreV1.ContainerPort) []coreV1.ContainerPort {
	result := append([]coreV1.ContainerPort{}, first...)
	portMap := make(map[string]int, len(first)+len(second))
	for i, value := range first {
		portMap[value.Name] = i
	}
	for _, value := range second {
		if i, found := portMap[value.Name]; found {
			result[i] = value
		} else {
			result = append(result, value)
			portMap[value.Name] = len(result) - 1
		}
	}
	return result
}

func GetDeploymentWithName(deployments []v1.Deployment, name string) *v1.Deployment {
	if len(deployments) == 0 {
		return nil
	}
	for _, deployment := range deployments {
		if deployment.Name == name {
			return &deployment
		}
	}
	return nil
}
