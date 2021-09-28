---
description: >-
  This is an add-on module to the Infrastructure as Code 101 course in CircleCI
  Academy demonstrating how the Snyk Orb helps you easily scan for
  misconfigurations in your Terraform files.
---

# Scan Terraform with the Snyk Orb

## Lab Meta

> **Difficulty**: Intermediate
>
> **Time to Complete:** 15 minutes

## Introduction

Terraform makes creating and tearing down cloud infrastructure as easy as writing configuration files. In the Infrastructure as Code course in the CircleCI Academy, you created a workflow that uses Terraform to create a GKE cluster and deploy an application into it as part of a continuous delivery pipeline. 

According to the NSA, [misconfigurations are the top Cloud vulnerability](https://www.cloudhesive.com/blog-posts/misconfiguration-top-cloud-vulnerability/). In this add-on module, you'll add [Snyk Infrastructure as Code](https://snyk.io/product/infrastructure-as-code-security/) into the workflow to reinforce secure IaC development practices, ensuring your Terraform files aren't configured in ways that open up your cluster, and the applications running in them, to risks caused by cloud misconfiguration. Let's begin!

## Pre-Requisites:

#### Recommended Pre-Work

This lab assumes the following courses were completed in CircleCI Academy:

{% embed url="https://academy.circleci.com/infrastructure-as-code" caption="Infrastructure as Code 101 in the CircleCI Academy" %}

{% embed url="https://academy.circleci.com/orbs-course" caption="Orbs in the CircleCI Academy" %}

{% hint style="warning" %}
It's highly recommended you complete the courses before proceeding.
{% endhint %}

#### Sample Code

We'll use the same code used in the Infrastructure as Code course. It can be found on GitHub.

{% embed url="https://github.com/datapunkz/learn\_iac" %}

#### Snyk Account and Token

You'll need a Snyk Account to use the Snyk Orb. [Create a Snyk API Token](https://support.snyk.io/hc/en-us/articles/360004008278-Revoking-and-regenerating-Snyk-API-tokens), then [set an Environment Variable in CircleCI](https://circleci.com/docs/2.0/env-vars/#setting-an-environment-variable-in-a-project) called SNYK\_TOKEN with its value. 

When ready, continue to the next page. 

