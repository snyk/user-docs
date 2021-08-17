# Snyk runtime monitoring: install the Snyk agent for Java

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

Add the [Snyk agent](https://static.snyk.io/resources/runtime/agent/java/snyk-java-runtime-agent.zip) to your project in order to start using Snyk Runtime Monitoring for your Java applications.

1. Download the runtime agent and unzip the archive.
2. From the extracted files, copy:snyk-java-runtime-agent.jaralongside your application files.
3. Create a snyk-agent.properties file at the same location as the agent jar file.

   The properties file should contain the Snyk project ID.

   For example: `projectId=0462e42b-c92f-4b48-bac8-81eb3ff7f43e`

4. Add the agent as a command-line argument to the Java command that you use to start your application.

   For example: `-javaagent:path/to/snyk-java-runtime-agent.jar -jar my-app.jar`

5. Add the jar file as follows:

   If you are using Apache Maven, add the following file to your Maven options environment variable: `-javaagent:path/to/snyk-java-runtime-agent.jar`

   If you are using JavaEE containers such as GlassFish, add the following file to the JVM Options: `-javaagent:path/to/snyk-java-runtime-agent.jar`.

6. Once you have successfully added the agent, logging begins to appear in the stderr \(standard error\) file shortly after the JVM \(Java virtual machine\) starts up, as in the following example:

   ```text
   //Example of successful Java agent installation snyk-agent initialisation: startup snyk-agent initialisation: trying: /opt/app-1/agent/snyk-agent.properties snyk-agent initialisation: switching logging to: /opt/app-1/agent/snyk-logs/agent-log-2001-02-03T04:05:06.log
   ```

