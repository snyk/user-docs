# Vulnerable conditions

Vulnerabilities that are not exploitable are unlikely to pose a security threat to your application and can therefore be de-prioritized for fixing.

Every vulnerability has conditions that must be met for the vulnerability to pose a threat, that is, to be exploitable.

Examples of vulnerable conditions include conditions that meet any of the following criteria:

* Require a user interaction
* Have a specific port number open
* Require a certain CPU architecture
* Enable a specific setting

Some vulnerabilities have multiple different conditions that must all be met for that vulnerability to be exploitable. A vulnerable open-source package can be exploited in some applications but not in others.

Exploitability depends on the context. Examples of context include the environment, the settings, and how the developer uses a package.

## Snyk Triage Assistant

{% hint style="info" %}
**Feature availability**\
This feature is available only for the Java (Gradle and Maven) ecosystem when you use GitHub as the source and Snyk Code is enabled.
{% endhint %}

In the context of your application, the Triage Assistant evaluates the vulnerable conditions, which helps you determine the exploitability of your application.

‌Snyk Code ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) engine is used to read your first-party code and to check the conditions for the vulnerabilities found by Snyk Open Source (SCA).

{% hint style="info" %}
To provide this feature, Snyk takes a temporary copy of your Git repository contents. For more information, see [How Snyk handles your data](../../snyk-data-and-governance/how-snyk-handles-your-data.md).
{% endhint %}

## Assessment of Jackson Vulnerable Conditions

Jackson Vulnerable Conditions must meet all of the following conditions for the vulnerability to be exploitable.

* Vulnerable version: The [Jackson package](https://snyk.io/vuln/maven:com.fasterxml.jackson.core%3Ajackson-databind) (com.fasterxml.jackson.core:jackson-databind vulnerabilities) should be a specific version that Snyk knows is vulnerable.
* Specific setting: A specific setting or functionality must be enabled, in this case, the [Polymorphic Type Handling](https://github.com/FasterXML/jackson-docs/wiki/JacksonPolymorphicDeserialization) feature.
  * You can check whether this setting is enabled in your code by looking for one of the following:
    * `@JsonSubTypes` annotation was used.
    * `@JsonTypeInfo` annotation was used on a Class.
    * `enableDefaultTyping()` is used to enable Polymorphic Typing.
    * `enableDefaultTypingAsProperty()` is used to enable Polymorphic Typing.
* User interactivity: The application needs to accept JSON input from the user.
* Specific gadget: A “gadget,” which is a class or function, must be available within the executing scope of the application.

## When is a vulnerability with Exploit Maturity not exploitable?

A vulnerability may have exploits available in the wild or detailed explanations of how to exploit it, but as long as not all the conditions are not met, the vulnerability will remain unexploitable.
