# YAML pipelines

{% hint style="info" %}
Comprehensive documentation on YAML schema for Azure Pipelines is available on Microsoft's [Azure DevOps documentation pages](https://docs.microsoft.com/en-us/azure/devops/pipelines/yaml-schema?view=azure-devops&tabs=schema%2Cparameter-schema).
{% endhint %}

An example azure-pipelines.yml is available below:

```yaml
trigger:
- master

resources:
- repo: self

variables:
  dockerRegistryServiceConnection: '$(Your-GUID)'
  imageRepository: '$(Your-Repo-Name)'
  containerRegistry: '$(Your-Registry-Name)'
  dockerfilePath: '$(Build.SourcesDirectory)/Dockerfile'
  manifestfilePath: '$(Build.SourcesDirectory)/manifests/**/*'
  namespaceApp: '$(Your-Name)'
  tag: '$(Build.BuildId)'
  vmImageName: 'ubuntu-latest'

stages:
- stage:
  jobs:
  - job:
    pool:
      vmImage: $(vmImageName)
    steps:
    - task: SnykSecurityScan@0
      inputs:
        serviceConnectionEndpoint: 'Snyk'
        testType: 'app'
        monitorOnBuild: true
        failOnIssues: true
    - task: Docker@2
      inputs:
        command: 'buildAndPush'
        repository: $(imageRepository)
        dockerfile: $(dockerfilePath)
        containerRegistry: $(dockerRegistryServiceConnection)
        tags: |
          $(tag)
    - task: KubernetesManifest@0
      inputs:
        action: 'deploy'
        kubernetesServiceConnection: '$(Your-Name)'
        namespace: '$(namespaceApp)'
        manifests: '$(manifestfilePath)'
        containers: '$(tag)'
        
```

The example above is only a reference. Your pipeline may differ. In this example, we are scanning our source code in Azure Repo, pushing a container image to ACR, and deploying to AKS. The Snyk Scan is defined in lines 24-29 above. 

