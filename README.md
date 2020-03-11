# AKS traffic manager 
AKS traffic manager is a proxy between Azure Kubernetes cluster's cloud provider and Azure resource manager. 

# Overview
 
# User Guide

0. Build the binary by running ``` go build ``` 

1. Build the docker image by running ``` docker build -t your-docker-id/aks-traffic-manager .```

2. To test the changes locally, run ```docker run -p 7788:7788 your-docker-id/aks-traffic-manager```

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
