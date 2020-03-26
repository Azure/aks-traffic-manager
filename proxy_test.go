package main

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCCPAutoThrust(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CCP Auto thrust Test suite")
}

var _ = Describe("testing enable TCP Reset for load balancer PUT", func() {
	It("should pass for stanard LB with correct api version", func() {

		inputLoadBalancerBody := `
		{
			"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
			"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"location":"westus2",
			"name":"rg1",
			"properties":{
			"backendAddressPools":[],
			"frontendIPConfigurations":[],
			"inboundNatPools":[],
			"inboundNatRules":[
				{
					"etag":"W/\"3de09ce5-fac8-4139-90e5-94e18da2da87\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/SSH-k8s-master-27630031-0",
					"name":"SSH-k8s-master-27630031-0",
					"properties":{
						"backendIPConfiguration":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/k8s-master-27630031-nic-0/ipConfigurations/ipconfig1",
						"resourceGroup":"rg1"
						},
						"backendPort":22,
						"enableDestinationServiceEndpoint":false,
						"enableFloatingIP":false,
						"enableTcpReset":false,
						"frontendIPConfiguration":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/k8s-master-lbFrontEnd-27630031",
						"resourceGroup":"rg1"
						},
						"frontendPort":22,
						"idleTimeoutInMinutes":4,
						"protocol":"Tcp",
						"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/inboundNatRules"
				}
			],
			"loadBalancingRules":[
				{
					"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/af18b09251eab4d97b180fc40bb145fe-TCP-443",
					"name":"af18b09251eab4d97b180fc40bb145fe-TCP-443",
					"properties":{
						"allowBackendPortConflict":false,
						"backendAddressPool":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/backendAddressPool1",
						"resourceGroup":"rg1"
						},
						"backendPort":443,
						"disableOutboundSnat":false,
						"enableDestinationServiceEndpoint":false,
						"enableFloatingIP":true,
						"enableTcpReset":false,
						"frontendIPConfiguration":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/af18b09251eab4d97b180fc40bb145fe",
						"resourceGroup":"rg1"
						},
						"frontendPort":443,
						"idleTimeoutInMinutes":4,
						"loadDistribution":"Default",
						"probe":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/probes/af18b09251eab4d97b180fc40bb145fe-TCP-443",
						"resourceGroup":"rg1"
						},
						"protocol":"Tcp",
						"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/loadBalancingRules"
				},
				{
					"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/lbrule2",
					"name":"lbrule2",
					"properties":{
					"allowBackendPortConflict":false,
					"backendAddressPool":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/backendAddressPool1",
						"resourceGroup":"rg1"
					},
					"backendPort":443,
					"disableOutboundSnat":false,
					"enableDestinationServiceEndpoint":false,
					"enableFloatingIP":true,
					"enableTcpReset":false,
					"frontendIPConfiguration":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/af18b09251eab4d97b180fc40bb145fe",
						"resourceGroup":"rg1"
					},
					"frontendPort":443,
					"idleTimeoutInMinutes":4,
					"loadDistribution":"Default",
					"probe":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/probes/af18b09251eab4d97b180fc40bb145fe-TCP-443",
						"resourceGroup":"rg1"
					},
					"protocol":"Tcp",
					"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/loadBalancingRules"
				}
			],
			"outboundRules":[
				{
					"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/LBOutboundRule",
					"name":"LBOutboundRule",
					"properties":{
						"allocatedOutboundPorts":1024,
						"backendAddressPool":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/rg1/backendAddressPools/rg1",
						"resourceGroup":"rg1"
						},
						"enableTcpReset":false,
						"frontendIPConfigurations":[
						{
							"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/rg1/frontendIPConfigurations/k8s-agent-outbound",
							"resourceGroup":"rg1"
						}
						],
						"idleTimeoutInMinutes":30,
						"protocol":"All",
						"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/outboundRules"
				}
			],
			"probes":[
		
			],
			"provisioningState":"Succeeded",
			"resourceGuid":"17fa8211-62b2-4234-840b-e65cc5cf7332"
			},
			"resourceGroup":"rg1",
			"sku":{
				"name":"Standard"
			},
			"type":"Microsoft.Network/loadBalancers"
		}
		`
		output, err := enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())

		var jsonBody map[string]interface{}
		json.Unmarshal(output, &jsonBody)

		numOfLBRules := 0
		properties := jsonBody["properties"].(map[string]interface{})
		loadBalancingRules := properties["loadBalancingRules"].([]interface{})
		for _, lbrule := range loadBalancingRules {
			rule := lbrule.(map[string]interface{})
			ruleProperties := rule["properties"].(map[string]interface{})
			Expect(ruleProperties["enableTcpReset"]).To(BeTrue())
			numOfLBRules++
		}

		Expect(numOfLBRules).To(Equal(2), "There should be 2 LB rules.")

		numOfOBRules := 0
		outboundRules := properties["outboundRules"].([]interface{})
		for _, obrule := range outboundRules {
			rule := obrule.(map[string]interface{})
			ruleProperties := rule["properties"].(map[string]interface{})
			Expect(ruleProperties["enableTcpReset"]).To(BeTrue())
			numOfOBRules++
		}
		Expect(numOfOBRules).To(Equal(1), "There is 1 OutBound rule.")

		numOfNatRules := 0
		inboundNatRules := properties["inboundNatRules"].([]interface{})
		for _, natrule := range inboundNatRules {
			rule := natrule.(map[string]interface{})
			ruleProperties := rule["properties"].(map[string]interface{})
			Expect(ruleProperties["enableTcpReset"]).To(BeTrue())
			numOfNatRules++
		}
		Expect(numOfNatRules).To(Equal(1), "There is 1 Nat rule.")
	})

	It("should skip for old api versions or wrong http method or wrong URL", func() {

		inputLoadBalancerBody := `
		{
			"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
			"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"location":"westus2",
			"name":"rg1",
			"properties":{
			"backendAddressPools":[],
			"frontendIPConfigurations":[],
			"inboundNatPools":[],
			"inboundNatRules":[],
			"loadBalancingRules":[
				{
					"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/af18b09251eab4d97b180fc40bb145fe-TCP-443",
					"name":"af18b09251eab4d97b180fc40bb145fe-TCP-443",
					"properties":{
						"allowBackendPortConflict":false,
						"backendAddressPool":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/backendAddressPool1",
						"resourceGroup":"rg1"
						},
						"backendPort":443,
						"disableOutboundSnat":false,
						"enableDestinationServiceEndpoint":false,
						"enableFloatingIP":true,
						"frontendIPConfiguration":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/af18b09251eab4d97b180fc40bb145fe",
						"resourceGroup":"rg1"
						},
						"frontendPort":443,
						"idleTimeoutInMinutes":4,
						"loadDistribution":"Default",
						"probe":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/probes/af18b09251eab4d97b180fc40bb145fe-TCP-443",
						"resourceGroup":"rg1"
						},
						"protocol":"Tcp",
						"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/loadBalancingRules"
				}
			],
			"outboundRules":[
				{
					"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
					"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/LBOutboundRule",
					"name":"LBOutboundRule",
					"properties":{
						"allocatedOutboundPorts":1024,
						"backendAddressPool":{
						"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/rg1/backendAddressPools/rg1",
						"resourceGroup":"rg1"
						},
						"frontendIPConfigurations":[
						{
							"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/rg1/frontendIPConfigurations/k8s-agent-outbound",
							"resourceGroup":"rg1"
						}
						],
						"idleTimeoutInMinutes":30,
						"protocol":"All",
						"provisioningState":"Succeeded"
					},
					"resourceGroup":"rg1",
					"type":"Microsoft.Network/loadBalancers/outboundRules"
				}
			],
			"probes":[],
			"provisioningState":"Succeeded",
			"resourceGuid":"17fa8211-62b2-4234-840b-e65cc5cf7332"
			},
			"resourceGroup":"rg1",
			"sku":{
			"name":"Standard"
			},
			"type":"Microsoft.Network/loadBalancers"
		}
		`

		output, err := enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-06-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")

		output, err = enableTCPReset(
			"GET",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")

		output, err = enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")
	})

	It("should skip for basic LB", func() {

		inputLoadBalancerBody := `
		{
			"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
			"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"location":"westus2",
			"name":"rg1",
			"properties":{
			"backendAddressPools":[],
			"frontendIPConfigurations":[],
			"inboundNatPools":[],
			"inboundNatRules":[],
			"loadBalancingRules":[],
			"outboundRules":[],
			"probes":[],
			"provisioningState":"Succeeded",
			"resourceGuid":"17fa8211-62b2-4234-840b-e65cc5cf7332"
			},
			"resourceGroup":"rg1",
			"type":"Microsoft.Network/loadBalancers"
		}
		`

		output, err := enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")

		inputLoadBalancerBody = `
		{
			"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
			"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"location":"westus2",
			"name":"rg1",
			"properties":{
			"backendAddressPools":[],
			"frontendIPConfigurations":[],
			"inboundNatPools":[],
			"inboundNatRules":[],
			"loadBalancingRules":[],
			"outboundRules":[],
			"probes":[],
			"provisioningState":"Succeeded",
			"resourceGuid":"17fa8211-62b2-4234-840b-e65cc5cf7332"
			},
			"sku":{
				"name":"Basic"
			},
			"resourceGroup":"rg1",
			"type":"Microsoft.Network/loadBalancers"
		}
		`

		output, err = enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")
	})

	It("should skip if resource type is not load balancer", func() {

		inputLoadBalancerBody := `
		{
			"etag":"W/\"2e1cb000-775f-4fa7-988a-d5e4c8db4fe1\"",
			"id":"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"location":"westus2",
			"name":"rg1",
			"properties":{
			"backendAddressPools":[],
			"frontendIPConfigurations":[],
			"inboundNatPools":[],
			"inboundNatRules":[],
			"loadBalancingRules":[],
			"outboundRules":[],
			"probes":[],
			"provisioningState":"Succeeded",
			"resourceGuid":"17fa8211-62b2-4234-840b-e65cc5cf7332"
			},
			"resourceGroup":"rg1",
			"sku":{
				"name":"Standard"
			},
			"type":"Microsoft.Network/virtualNetworks"
		}
		`

		output, err := enableTCPReset(
			"PUT",
			"/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1",
			"2018-07-01",
			[]byte(inputLoadBalancerBody))

		Expect(err).To(BeNil())
		Expect(string(output)).To(Equal(inputLoadBalancerBody), "Should make no changes to the request body")
	})
})
