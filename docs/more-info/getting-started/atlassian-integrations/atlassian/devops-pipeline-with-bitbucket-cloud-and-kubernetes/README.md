---
description: Build and test code and infrastructure using multiple integrations
---

# DevOps Pipeline with Bitbucket Cloud and Kubernetes

In this workshop, we'll work through a pipeline to checkout, build, test, and deploy your code as a container to running environments.  This use case is typical of what most users encounter, and the workshop is designed to show you how Snyk can plumb into your everyday activities.

This workshop starts with a java version of our popular vulnerable application named Goof.  The source in Github is [https://github.com/snyk-labs/java-goof](https://github.com/snyk-labs/java-goof) and we'll work through a pipeline example highlighting multiple integrations along the way.

## Learning Objectives

We'll work through examples to highlight the following concepts with hands-on examples:

* We'll start with your code base in a Bitbucket repository, and enable Snyk to provide timely information about Open-Source Vulnerabilities.
* We'll integrate Snyk into a working Atlasssian Bitbucket Pipeline to reveal vulnerabilities in your container images as part of your CI/CD pipeline
* We'll deploy your container to a running environment and reveal additional runtime vulnerabilities in your Kubernetes manifests.
* We'll show the correlation between the vulnerabilities and exploits.  Then we'll apply fixes to  your code to show remediation.

## Target audience

This workshop is designed for these members of your development team:

* Developers writing code and able to fix vulnerabilities with timely feedback.
* Security Teams interested in using results from security scans to help development teams prioritize security fixes.
* DevOps/DevSecOps engineers charged with the responsibility to ensure a well-running and compliant pipeline delivering code on time and with security.
* Any other member with the shared interest of delivering code with better security.

## Content Structure

This workshop is developed in several modules to address the interests of different stages of your team's workflow.  As a representative use case, you may find the use cases align very well with your team's structure.   If your pipeline operations are different, you should still see how the desired outcomes for each use case still map to your processes.



### Module 1 - Scanning and monitoring source code at a developer workstation

In this module, you enable Snyk to automatically scan your repository and quickly provide results to your entire team.  This example shows you how you and your team can leverage Snyk to easily create pull requests in your Bitbucket repository.&#x20;

### Module 2 - Scanning and monitoring container images

Next, you'll enable Snyk to monitor your container images as you prepare them for deliver to AWS ECR.  We'll use the feedback to make a change to a base image to address a vulnerability to improve the posture of your container image.

### Module 3 - Exploiting a container and Kubernetes cluster

In this module, we'll demonstrate a container exploit and expand the example into a cluster exploit.   We'll remediate the vulnerability.  This module requires features available to Snyk users in a paid Tier.







