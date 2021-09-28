---
title: DevSecOps with Snyk
chapter: true
weight: 1
---

# AWS Code Suite

{% embed url="https://youtu.be/cgsMsHO4Nbc" %}

In this workshop, you will learn how to add security testing to a CI/CD pipeline of a container application using AWS CodeCommit, AWS CodeBuild, and AWS CodePipeline. The modules contained in this workshop will provide you with step-by-step instructions for committing, building, testing, and deploying software in an automation fashion. You will also learn about some basic security tests and where to instrument them in the software development lifecycle.

## Objectives

* Gain familiarity with the workflow of a modern application
* Learn where to add security testing to a CI/CD pipeline
* Learn about AWS services used to orchestrate testing 

## What we will cover in this workshop

* Setup of a Cloud9 environment
* Usage of AWS CloudFormation to automate the deployment of infrastructure
* Deployment of Amazon Elastic Container Service
* Deploy and use a modernized pipeline using AWS CodePipeline, CodeCommit, and CodeBuild 
* Instrument a couple of security testing/scanning tools

## Sample reference architecture

At the conclusion of this workshop, you will end up with various AWS services provisioned in your AWS account. The following diagram illustrates some of these services and is intended as a sample reference architecture.

![](../../../.gitbook/assets/aws-pipeline.png)

## Workshop flow

Each section or module contained in this workshop is designed to guide you through each step of the process to build the architecture referenced above. This is accomplished by using AWS Cloud 9 as our starting point along with a \`git clone\` of the content from our repository. Everything you need is provided to you including sample code, AWS CloudFormation templates, and detailed instructions. We will be using the AWS CLI from our Cloud9 instance to deploy the CloudFormation templates and build out our environment.

{% hint style="info" %}
The examples and sample code provided in this workshop are intended to be consumed as instructional content.
{% endhint %}

