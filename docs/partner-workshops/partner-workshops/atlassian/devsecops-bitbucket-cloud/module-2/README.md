---
title: Module 2
chapter: true
weight: 40
---

# Module 2

## Container security\*

![](../../../../.gitbook/assets/snyk-bitbucket-flow-module-02.png)

By 2022, more than [75% of global organizations](https://snyk.io/blog/putting-container-security-in-the-hands-of-developers/) will be running containerized applications in production \(Gartner\). Alongside the widespread adoption, there has been a surge in container vulnerabilities, with a 4X increase in reported operating system vulnerabilities in 2018. And yet [80% of developers say they don’t test their container images](https://snyk.io/blog/shifting-docker-security-left/) during development – it’s either not their responsibility, or they are accustomed to someone down the road catching the issues – which makes scaling [container security](https://snyk.io/container-security/) a challenge for fast-growing businesses.

## Learning objective

In this module we will learn how to [secure your build workflow](https://snyk.io/blog/secure-your-build-workflow-on-bitbucket-pipes-with-snyk/) on Bitbucket Pipes with Snyk. Scanning and analyzing your Linux-based container project for known vulnerabilities is an important step in securing your environment by helping you identify and mitigate security vulnerabilities. The exercises in this module will help [secure your container](https://support.snyk.io/hc/en-us/articles/360003946897-Container-security-overview) by leveraging the Snyk Pipe for Bitbucket pipelines to scan the base image for its dependencies including:

* The operating system \(OS\) packages installed and managed by the package manager
* Key binaries—layers that were not installed through the package manager

Based on these results, Snyk will provide fix advice and guidance including:

* Origins of the vulnerabilities in your OS packages and key binaries
* Base image upgrade details or a recommendation to rebuild the image
* Dockerfile layer in which the affected package was introduced
* Fixed-in version of the operating system and key binary packages

Lastly, you will enable Snyk's integration for [Amazon Elastic Container Registry \(ECR\)](https://support.snyk.io/hc/en-us/articles/360003916078-Configure-integration-for-Amazon-Elastic-Container-Registry-ECR-) to continuously scan and monitor your container images.

## Homework: Learn more about Snyk  & Container security

* [Container security throughout the SDLC](https://snyk.io/blog/container-security-throughout-the-sdlc/)
* [Everything you wanted to know about addressing security vulnerabilities in Linux-based containers](https://snyk.io/blog/everything-you-wanted-to-know-about-addressing-security-vulnerabilities-in-linux-based-containers/)
* [Protect container images directly from your registries](https://snyk.io/blog/protect-docker-images-directly-from-your-container-registries/)

**Footnotes:**

1. **\*Snyk Blog -** [Putting container security in the hands of developers](https://snyk.io/blog/putting-container-security-in-the-hands-of-developers)

