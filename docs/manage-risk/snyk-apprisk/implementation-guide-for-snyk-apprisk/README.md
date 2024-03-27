# Implementation guide for Snyk AppRisk

Implementing Snyk AppRisk brings many benefits to your products by enabling the Application Security Posture Management for all your used products. You can benefit from automated application asset discovery, customized security controls, and risk-based prioritization.

Snyk AppRisk helps application security teams implement, manage, and scale their Snyk-based developer security program. It allows you to pull in Assets from various sources into one central view to understand what security controls you have in place.

This guide helps you implement Snyk AppRisk at the Group level. You can find more details about adopting Snyk at the enterprise level by accessing the [Enterprise Implementation Guide](../../../implement-snyk/enterprise-implementation-guide/) documentation.

## Implementation strategy overview

The implementation guide is split into multiple phases, each of them allowing you to use a certain feature of Snyk AppRisk:

* [Level 1: Configure Snyk AppRisk and setup integrations](level-1-configure-snyk-apprisk-and-setup-integrations.md)
* [Level 2: Policies](level-2-policies.md)&#x20;
* [Level 3: Prioritization setup ](level-3-prioritization-setup.md)
* [Level 4: Associate the source code with the container images](level-4-associate-the-source-code-with-the-container-images.md)

## Implementation objectives

You have the flexibility to incorporate Snyk AppRisk at various levels of maturity. This allows you to choose whether to fully adopt and implement all functionalities or only a subset of the available functionalities.&#x20;

{% hint style="info" %}
The Visibility, Coverage, and Prioritization fields use the Limited, Good, and Great ranking scores.
{% endhint %}

### Level 1

**Required implementation**: Onboard your SCM

**Objectives**: Get visibility into applications (code assets that exist in SCM repositories) and understand Snyk coverage controls. Prioritize Container issues with the OS condition risk factor.

**Visibility**: Good

**Coverage**: Good

**Prioritization**: Limited

### Level 2

**Required implementation**: Implement policies

**Objectives**: Manage assets with business context and configure coverage gaps and alerts to meet business needs. Examples:

1. Asset management (tagging, classification)
2. Notification workflows
3. Coverage gaps

**Visibility**: Great

**Coverage**: Great

**Prioritization**: Limited

### Level 3

**Required implementation**: Prioritization setup - Kubernetes connector

**Objectives**: Onboard the Kubernetes connector to prioritize issues on whether a container image has the risk factor set as "deployed" and "public facing".

**Visibility**: Great

**Coverage**: Great

**Prioritization**: Limited

### Level 4

**Required implementation**: Use tags to associate your Snyk Code and Snyk Open Source Projects with images.

**Objectives**: Associate code assets with relevant container images to prioritize Snyk Open Source and Snyk Code issues based on "deployed" and "public facing" risk factors.

**Visibility**: Great

**Coverage**: Great

**Prioritization**: Great

## Resources

Use the following resources for more details about Snyk AppRisk:

* [Snyk AppRisk ](../getting-started-with-snyk-apprisk.md)documentation
* [Enterprise implementation guide](../../../implement-snyk/enterprise-implementation-guide/)
* [Snyk Learn](https://learn.snyk.io/)&#x20;



