# 'toolPath' Error for Azure pipelines.

**Problem**:

Sometimes the user receives the following error for Snyk integration for Azure Pipelines:

```text
Parameter 'toolPath' cannot be null or empty
```

This can happen when NodeJS and NPM are not in the toolPath of the environment.

**Screenshot**:

![A screenshot of a cell phone

Description automatically generated](https://support.snyk.io/attachments/token/bVLPogyrKmwvYj4bmX10IHNSa/?name=image001.png)

**Resolution**:

> **Requirements**
>
> This extension requires that Node.js and npm be installed on the build agent. These are available by default on all Microsoft-hosted build agents. However, if you are using a self-hosted build agent, you may need to explicitly activate Node.js and npm and ensure they are in your [PATH](https://en.wikipedia.org/wiki/PATH_%28variable%29). This can be done using the [NodeTool task from Microsoft](https://docs.microsoft.com/en-us/azure/devops/pipelines/tasks/tool/node-js?view=azure-devops) prior to the SnykSecurityScan task in your pipeline.

  
You can find the same in our documentation [Azure Pipelines integration](https://support.snyk.io/hc/en-us/articles/360004127677-Azure-Pipelines-integration?source=search&auth_token=eyJhbGciOiJIUzI1NiJ9.eyJhY2NvdW50X2lkIjo5MTMyNjQ0LCJ1c2VyX2lkIjozNzk4ODE0MzAzNTgsInRpY2tldF9pZCI6MTEzMDgsImNoYW5uZWxfaWQiOjYzLCJ0eXBlIjoiU0VBUkNIIiwiZXhwIjoxNjI1MjM4MjE0fQ.9Tga_3tCqSFCIquWl8X2gXF-yuyGXzTk4JlClE7cXZ0). So we suggest you add the NodeTool task with NodeJS version 10 or above, before the Snyk task and try again. This should resolve the error with toolPath.

