# Snyk runtime monitoring: uninstall the Snyk agent for Java

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

To remove the agent:

1. Remove these .jar and .properties files, which you added to your project during installation: `snyk-java-runtime-agent.jar snyk-agent.properties`
2. Remove the agent command-line argument from the Java command: `-javaagent:path/to/snyk-java-runtime-agent.jar`

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page/)
{% endhint %}

