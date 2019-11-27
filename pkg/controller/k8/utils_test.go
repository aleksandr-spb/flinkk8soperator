package k8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
)

func TestGetAppLabel(t *testing.T) {
	appName := "app_name"
	appLabel := GetAppLabel(appName)
	assert.Equal(t, map[string]string{
		"flink-app": appName,
	}, appLabel)
}

func TestMergeLabels(t *testing.T) {
	first := map[string]string{
		"zero": "0",
		"one":  "1",
	}
	second := map[string]string{
		"one":   "10",
		"two":   "20",
		"three": "30",
	}
	result := MergeLabels(first, second)
	assert.Equal(t, map[string]string{
		"zero":  "0",
		"one":   "10",
		"two":   "20",
		"three": "30",
	}, result)
}

func TestMergePorts(t *testing.T) {
	first := []coreV1.ContainerPort{
		{
			Name:          "rpc",
			ContainerPort: 6123,
		},
		{
			Name:          "ui",
			ContainerPort: 1001,
		},
	}
	second := []coreV1.ContainerPort{
		{
			Name:          "query",
			ContainerPort: 2000,
		},
		{
			Name:          "ui",
			ContainerPort: 8081,
		},
		{
			Name:          "metric",
			ContainerPort: 9249,
		},
		{
			Name:          "query",
			ContainerPort: 6124,
		},
	}
	result := MergePorts(first, second)
	assert.Equal(t, []coreV1.ContainerPort{
		{
			Name:          "rpc",
			ContainerPort: 6123,
		},
		{
			Name:          "ui",
			ContainerPort: 8081,
		},
		{
			Name:          "query",
			ContainerPort: 6124,
		},
		{
			Name:          "metric",
			ContainerPort: 9249,
		},
	}, result)
}

func TestMergePortsEmptyFirst(t *testing.T) {
	first := []coreV1.ContainerPort{}
	second := []coreV1.ContainerPort{
		{
			Name:          "ui",
			ContainerPort: 8081,
		},
		{
			Name:          "query",
			ContainerPort: 6124,
		},
		{
			Name:          "metric",
			ContainerPort: 9249,
		},
	}
	result := MergePorts(first, second)
	assert.Equal(t, []coreV1.ContainerPort{
		{
			Name:          "ui",
			ContainerPort: 8081,
		},
		{
			Name:          "query",
			ContainerPort: 6124,
		},
		{
			Name:          "metric",
			ContainerPort: 9249,
		},
	}, result)
}

func TestGetDeploymentWithName(t *testing.T) {
	name := "jm-name"
	dep := v1.Deployment{}
	dep.Name = name
	deployments := []v1.Deployment{
		dep,
	}
	actualDeployment := GetDeploymentWithName(deployments, name)
	assert.NotNil(t, actualDeployment)
	assert.Equal(t, dep, *actualDeployment)
}

func TestGetDeploymentNotExists(t *testing.T) {
	name := "jm-name"
	dep := v1.Deployment{}
	dep.Name = name
	deployments := []v1.Deployment{
		dep,
	}
	actualDeployment := GetDeploymentWithName(deployments, "random")
	assert.Nil(t, actualDeployment)
}
