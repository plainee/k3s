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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/rancher/k3s/pkg/apis/k3s.cattle.io/v1"
	clientset "github.com/rancher/k3s/pkg/generated/clientset/versioned/typed/k3s.cattle.io/v1"
	informers "github.com/rancher/k3s/pkg/generated/informers/externalversions/k3s.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
)

type Interface interface {
	Addon() AddonController
	ListenerConfig() ListenerConfigController
}

func New(controllerManager *generic.ControllerManager, client clientset.K3sV1Interface,
	informers informers.Interface) Interface {
	return &version{
		controllerManager: controllerManager,
		client:            client,
		informers:         informers,
	}
}

type version struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.K3sV1Interface
}

func (c *version) Addon() AddonController {
	return NewAddonController(v1.SchemeGroupVersion.WithKind("Addon"), c.controllerManager, c.client, c.informers.Addons())
}
func (c *version) ListenerConfig() ListenerConfigController {
	return NewListenerConfigController(v1.SchemeGroupVersion.WithKind("ListenerConfig"), c.controllerManager, c.client, c.informers.ListenerConfigs())
}
