# AKS traffic manager 
AKS traffic manager is a proxy between Azure Kubernetes cluster's cloud provider and Azure resource manager. It helps with throttling related issues by delaying 429 responses from Azure resource manager.

# User Guide

All aks clusters already have aks-traffic-manager applied. But you need to manually apply aks traffic manager to aks-engine clusters. At below are the steps:

0. Log on the master node of your cluster 

1. Download aks-traffic-manager binary:

```curl -o aks-traffic-manager -L https://github.com/Azure/aks-traffic-manager/releases/download/v0.2/aks-traffic-manager```

```chmod a+rx aks-traffic-manager```

2. Download the azurestackcloud.json file:

```curl -o /etc/kubernetes/azurestackcloud.json -L https://raw.githubusercontent.com/Azure/aks-traffic-manager/master/azurestackcloud.json```

3. Reconfig azure.json to use azure stack:

```sed -i 's/AzurePublicCloud/AzureStackCloud/g' /etc/kubernetes/azure.json```

4. Edit ```/etc/kubernetes/manifests/kube-controller-manager.yaml``` to include the following environment variable definition:

```
      env:
        - name: AZURE_ENVIRONMENT_FILEPATH
          value: /etc/kubernetes/azurestackcloud.json
```

5. Run aks-traffic-manager ```./aks-traffic-manager```

6. You should now be able to see logs from aks-traffice-manager showing the ARM requests going through the proxy.

# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
