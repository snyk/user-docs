# DevSecOps with Bitbucket Cloud

![](../../../.gitbook/assets/finding-open-source-vulnerabilities-within-the-bitbucket-workflow-.png)

Join [Snyk](https://snyk.io/) and [Atlassian](https://www.atlassian.com/) for this _**hands-on virtual workshop**_ where we will guide you on implementing security best-practices early on in your workflow to build an automated and **secure** [Continuous Integration \(CI\)](https://www.atlassian.com/continuous-delivery/continuous-integration) & [Continuous Delivery \(CD\)](https://www.atlassian.com/continuous-delivery) pipeline.

## Welcome!

![](../../../.gitbook/assets/mm.png)

You will begin this workshop as the newest member of a Mythical 500 company: **Mythical Mysfits**. It's your first day in the office and your predecessor\(s\) hastily \(i.e. manually\) deployed an "enterprise-ready" piece of software for group collaboration: [goof](https://github.com/snyk/goof). According to your colleagues, _goof_ became wildly popular as _"The BESTest todo app evar"_. You, however, are skeptical. Fortunately, your company also recently purchased Snyk and Atlassian Bitbucket Cloud!

In this workshop, you will learn patterns for shift-left security leveraging [Atlassian Bitbucket](https://www.atlassian.com/software/bitbucket), [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines), and [Snyk](https://snyk.io). These techniques will enable you to implement scanning of your container-based workloads running on [Amazon Elastic Kubernetes Service \(Amazon EKS\)](https://aws.amazon.com/eks/) and [Amazon Elastic Container Registry \(ECR\)](https://aws.amazon.com/ecr/) and how to use these patterns to release features and functionality at a faster pace that includes security at each step.

## Learning Objectives

![](../../../.gitbook/assets/snyk-bitbucket-flow.png)

* Understand [Software Composition Analysis \(SCA\)](https://snyk.io/blog/what-is-software-composition-analysis-sca-and-does-my-company-need-it/) and its importance in your developer workflows.
* Discover vulnerabilities in your application open source dependencies.
* Implement Snyk's [container image](https://snyk.io/blog/detecting-vulnerabilities-in-container-images/) scanning in your [Continuous integration](https://aws.amazon.com/devops/continuous-integration/) and 

  [Continuous delivery](https://aws.amazon.com/devops/continuous-delivery/) \(CI/CD\) pipeline.

* [Fix insecure Kubernetes configuration at the source](https://snyk.io/blog/fix-insecure-kubernetes-configuration/).
* Leverage the [Snyk Pipe](https://bitbucket.org/product/features/pipelines/integrations?p=snyk/snyk-scan) for Bitbucket Pipelines to secure your application.

## Intended Audience

* Developers
* Security/Application Teams
* DevOps/DevSecOps Engineers
* Cloud/Solutions Architects

## Content Structure

We have structured the various subjects covered in this workshop into specific modules. Each module will provide context on the theory behind the techniques presented as well as hands-on examples.

### Module 1 - Scanning & monitoring application source code

Enable the Snyk Open Source [integration](https://solutions.snyk.io/snyk-academy/open-source/create-source-control-integration) to Bitbucket and [import your SCM project](https://solutions.snyk.io/snyk-academy/open-source/import-scm-project). Understand transitive dependencies and how Snyk can generate automatic pull requests to streamline your process.

### Module 2 - Scanning & monitoring container images

Enable Snyk Container [integration](https://support.snyk.io/hc/en-us/articles/360003916078-Configure-integration-for-Amazon-Elastic-Container-Registry-ECR-) to Amazon Elastic Container Registry \(ECR\) and [import](https://solutions.snyk.io/snyk-academy/container/container-registry-and-image-import) your container images. Learn how Snyk provides base image ugprade recommendations.

### Module 3 - Scanning & monitoring for insecure Kubernetes configurations

[Install the Snyk controller](https://support.snyk.io/hc/en-us/articles/360011128137-Install-the-Snyk-controller-on-Amazon-Elastic-Kubernetes-Service-Amazon-EKS-) on Amazon Elastic Kubernetes Service \(Amazon EKS\) and [add workloads for scanning](https://support.snyk.io/hc/en-us/articles/360003947117-Adding-Kubernetes-workloads-for-security-scanning). Understand test results, how to interpret Snyk's [Priority Score](https://support.snyk.io/hc/en-us/articles/360010906897-Snyk-Priority-Score-and-Kubernetes), and how to fix configuration issues.

### Module 4 - Fixing known issues & monitoring

In this module, you will go through guided exercises that demonstrate how to fix for vulnerabilities and insecure configurations. You will apply what you learned in the previous modules and apply fixes to your application, container image, and Kubernetes configuration to secure your application.

{% hint style="info" %}
To make the most effective use of this content, you should be able to run basic Unix commands. You should also possess familiarity with AWS services, basic cloud concepts and general understanding of software development methodologies.
{% endhint %}

