# Triaging issues



### Introduction

Every vulnerability has conditions that must be met, for that vulnerability to pose a threat (to be exploitable).

Vulnerable condition examples include:

* Requiring a user interaction.
* Having a specific port number open.
* Requiring a certain CPU architecture.
* Enabling a specific setting.

#### It's all about the context

Some vulnerabilities have multiple different conditions that must all be met,  for that vulnerability to be exploitable. A vulnerable open source package can be exploited in some applications, but not in others. 

Exploitability depends on the context: for example, the environment, the settings, and the way the developer uses this package.

Vulnerabilities that are not exploitable are unlikely to pose a security threat to your application and can therefore be de-prioritized accordingly.

### Triage Assistant

{% hint style="info" %}
Currently, this feature is only available to **Java** (Gradle and Maven) ecosystem, when using GitHub as the SCM, and when **Snyk Code** is enabled.

To enable this feature, within Snyk, go to **Settings → Snyk Preview **and enable this feature.
{% endhint %}

In the context of your application, the Triage Assistant evaluates the vulnerable conditions, which helps you determine the exploitability of your application.

‌Snyk Code (SAST) engine is used to read your first-party code and to check the conditions for the vulnerabilities found by Snyk Open Source (SCA).

#### Vulnerable Conditions

Jackson Vulnerable Conditions:

* **Vulnerable version**: The jackson package (**com.fasterxml.jackson.core:jackson-databind vulnerabilities**) should be in a specific version, that we know is vulnerable.
* **Specific setting**: A specific setting, or functionality need to be enabled, in our case it’s the [**Polymorphic Type Handling**](https://github.com/FasterXML/jackson-docs/wiki/JacksonPolymorphicDeserialization) feature. 
* **User interactivity**: The application needs to accept JSON input from the user. 
* **Specific gadget**: A “gadget”, which is a class or function, needs to be available within the executing scope of the application.

All the conditions must be met for the vulnerability to be exploitable.

{% hint style="info" %}
This feature is currently in preview, and might be changed.
{% endhint %}

#### Vulnerability with Exploit Maturity but not exploitable?

A vulnerability may have exploits available in the wild or detailed explanations of how to exploit it, but as long as not all the conditions are not met, the vulnerability will remain unexploitable.
