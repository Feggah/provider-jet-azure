/*
Copyright 2021 The Crossplane Authors.

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

package network

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-azure/config/common"
)

const groupNetwork = "network"

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.Kind = "NetworkInterface"
		r.ShortGroup = groupNetwork
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("/Microsoft.Network/networkInterfaces")
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "loadBalancers", "name")
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.Kind = "VirtualNetwork"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"subnet"},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworks", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_virtual_network_gateway", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGateway"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"ip_configuration.subnet_id": config.Reference{
				Type: "Subnet",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworkGateways", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_virtual_network_peering", func(r *config.Resource) {
		r.Kind = "VirtualNetworkPeering"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			"virtual_network_name": config.Reference{
				Type:      "VirtualNetwork",
				Extractor: rconfig.APISPackagePath + "/rconfig.ExtractResourceName()",
			},
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"remote_virtual_network_id": config.Reference{
				Type: "VirtualNetwork",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/virtualNetworkPeerings/myvnet1peering
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworks", "virtual_network_name",
			"virtualNetworkPeerings", "name",
		)
	})

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_application", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	p.AddResourceConfigurator("azurerm_virtual_network_gateway_connection", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGatewayConnection"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
			// TODO(aru): as we no longer hold Azure ID of resources in external-name annotation
			// these references are currently not generated
			/*"virtual_network_gateway_id": config.Reference{
				Type: "VirtualNetworkGateway",
			},
			"peer_virtual_network_gateway_id": config.Reference{
				Type: "VirtualNetworkGateway",
			},*/
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"connections", "name",
		)
	})

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_workspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	p.AddResourceConfigurator("azurerm_virtual_wan", func(r *config.Resource) {
		r.Kind = "VirtualWAN"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualWans/testvwan
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualWans", "name",
		)
	})
}
