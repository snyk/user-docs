# Triaging issues

### Introduction

Every vulnerability has conditions that must be met, for that vulnerability to pose a threat (to be exploitable).

Vulnerable condition examples include:

* Requiring a user interaction
* Having a specific port number open
* Requiring a certain CPU architecture
* Enabling a specific setting

#### It's all about the context

Some vulnerabilities have multiple different conditions that must all be met, for that vulnerability to be exploitable. A vulnerable open source package can be exploited in some applications, but not in others.

Exploitability depends on the context: for example, the environment, the settings, and the way the developer uses this package.

Vulnerabilities that are not exploitable are unlikely to pose a security threat to your application and can therefore be de-prioritized accordingly.

### Triage Assistant

{% hint style="info" %}
Currently, this feature is only available to **Java** (Gradle and Maven) ecosystem, when using **GitHub** as the source, and when **Snyk Code** is enabled.
{% endhint %}

In the context of your application, the Triage Assistant evaluates the vulnerable conditions, which helps you determine the exploitability of your application.

‌Snyk Code ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) engine is used to read your first-party code and to check the conditions for the vulnerabilities found by Snyk Open Source (SCA).

{% hint style="info" %}
To provide this feature, Snyk takes a temporary copy of your Git repository contents.

For more information see [how-snyk-handles-your-data.md](../../more-info/how-snyk-handles-your-data.md "mention")
{% endhint %}

#### Vulnerable Conditions

Jackson Vulnerable Conditions:

* **Vulnerable version**: The [Jackson package](https://snyk.io/vuln/maven:com.fasterxml.jackson.core%3Ajackson-databind) (**com.fasterxml.jackson.core:jackson-databind vulnerabilities**) should be in a specific version, that we know is vulnerable.
* **Specific setting**: A specific setting, or functionality need to be enabled, in our case it’s the [**Polymorphic Type Handling**](https://github.com/FasterXML/jackson-docs/wiki/JacksonPolymorphicDeserialization) feature.
  * You can check if this setting is enabled in your code by looking for one of the following:
    * `@JsonSubTypes` annotation was used.
    * `@JsonTypeInfo` annotation was used on a Class.
    * `enableDefaultTyping()` is used to enable Polymorphic Typing .
    * `enableDefaultTypingAsProperty()` is used to enable Polymorphic Typing.
* **User interactivity**: The application needs to accept JSON input from the user.
* **Specific gadget**: A “gadget”, which is a class or function, needs to be available within the executing scope of the application.

All the conditions must be met for the vulnerability to be exploitable.

{% hint style="info" %}
This feature is currently in **preview**, and might be changed.
{% endhint %}

#### Vulnerability with Exploit Maturity but not exploitable?

A vulnerability may have exploits available in the wild or detailed explanations of how to exploit it, but as long as not all the conditions are not met, the vulnerability will remain unexploitable.
