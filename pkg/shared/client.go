// Copyright 2018 BlueData Software, Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shared

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	Client *K8sClient
)

// init creates the REST API client that will be used for actions not
// supported through the operator SDK.
func init() {

	Client = newClientInCluster()
}

// newClientInCluster creates a k8s REST API client that will operate using
// the service account credentials of the KubeDirector pod.
func newClientInCluster() *K8sClient {

	config := getConfigFromServiceAccount()
	client := &K8sClient{
		Clientset:    getClientSet(config),
		ClientConfig: config,
	}
	return client
}

// getClientSet creates a k8s REST API client from the given config.
func getClientSet(
	config *rest.Config,
) kubernetes.Interface {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Fatal("getClientSet: ", err)
	}
	return clientset
}

// getConfigFromServiceAccount generates a client config using the local
// service account credentials.
func getConfigFromServiceAccount() *rest.Config {

	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Fatal("getConfigFromServiceAccount: ", err)
	}
	return config
}
