# Snyk test using Maven

## Testing SPC using the Snyk Maven plugin

Snyk CLI enables developers to test and monitor their application using the CLI. If developers want Snyk to be part of the daily build, you can use the Snyk Maven plugin. The steps below will show you how to configure Maven to run the Snyk plugin. We will configure it to run in our local deployment. 

From the root directory of the SPC application

```text
vi pom.xml
```

In the plugin section, find and change the &lt;org&gt; name to match the Snyk organization you created earlier.  We will pass the SNYK\_TOKEN via the maven command line in the next step.

{% hint style="info" %}
See the next step to find your organization name in Snyk.
{% endhint %}

```text
<plugin>
  <groupId>io.snyk</groupId>
  <artifactId>snyk-maven-plugin</artifactId>
  <version>1.2.6</version>
  <executions>
    <execution>
      <id>snyk-test</id>
      <phase>test</phase>
      <goals>
        <goal>test</goal>
      </goals>
    </execution>
    <execution>
      <id>snyk-monitor</id>
      <phase>install</phase>
      <goals>
        <goal>monitor</goal>
      </goals>
    </execution>
  </executions>
  <configuration>
    <apiToken>${SNYK_TOKEN}</apiToken>
    <failOnSeverity>false</failOnSeverity>
    <!-- Replace with your Org name -->
    <org>springone-workshop</org>
  </configuration>
</plugin>
```

Your organization name can be retrieved under settings.

![](../../../.gitbook/assets/getting_org_name.png)

## Execute Maven to view Snyk results

After configuring the pom.xml file, we are ready to execute a Snyk test and view the results. Issue the command below with your personal API token.

```text
mvn snyk:test -DSNYK_TOKEN=add_your_personal_token
```

{% hint style="info" %}
We retrieved your personal API token during authentication.
{% endhint %}

The Synk test results will display color-coded vulnerability information similar to the output below. The Synk plugin will provide details on each vulnerability and its associated severity level. Developers can choose to fail the build based on severity level or allow the build to pass.

```text
Downloading from central: https://repo.maven.apache.org/maven2/org/apache/logging/log4j/log4j-api-java9/2.13.3/log4j-api-java9-2.13.3.pom
[WARNING] The POM for org.apache.logging.log4j:log4j-api-java9:zip:2.13.3 is missing, no dependency information available
[WARNING] ✗ medium severity vulnerability found on com.google.guava:guava@18.0
[WARNING] - desc: Deserialization of Untrusted Data
[WARNING] - info: https://snyk.io/vuln/SNYK-JAVA-COMGOOGLEGUAVA-32236
[WARNING] - from: org.springframework.samples:spring-petclinic@2.3.1.BUILD-SNAPSHOT > org.springframework.boot:spring-boot-starter-data-jpa@2.3.1.RELEASE > org.springframework.data:spring-data-jpa@2.3.1.RELEASE > com.querydsl:querydsl-apt@4.3.1 > com.querydsl:querydsl-codegen@4.3.1 > com.querydsl:querydsl-core@4.3.1 > com.google.guava:guava@18.0
[WARNING]
[WARNING] ✗ high severity vulnerability found on log4j:log4j@1.2.16
[WARNING] - desc: Deserialization of Untrusted Data
[WARNING] - info: https://snyk.io/vuln/SNYK-JAVA-LOG4J-572732
[WARNING] - from: org.springframework.samples:spring-petclinic@2.3.1.BUILD-SNAPSHOT > org.springframework.boot:spring-boot-starter-data-jpa@2.3.1.RELEASE > org.hibernate:hibernate-core@5.4.17.Final > org.jboss.logging:jboss-logging@3.4.1.Final > log4j:log4j@1.2.16
[WARNING]
[WARNING] ✗ medium severity vulnerability found on org.apache.tomcat.embed:tomcat-embed-core@9.0.36
[WARNING] - desc: Denial of Service (DoS)
[WARNING] - info: https://snyk.io/vuln/SNYK-JAVA-ORGAPACHETOMCATEMBED-584427
[WARNING] - from: org.springframework.samples:spring-petclinic@2.3.1.BUILD-SNAPSHOT > org.springframework.boot:spring-boot-starter-web@2.3.1.RELEASE > org.springframework.boot:spring-boot-starter-tomcat@2.3.1.RELEASE > org.apache.tomcat.embed:tomcat-embed-core@9.0.36
[WARNING]
[WARNING] ✗ high severity vulnerability found on org.hibernate:hibernate-core@5.4.17.Final
[WARNING] - desc: SQL Injection
[WARNING] - info: https://snyk.io/vuln/SNYK-JAVA-ORGHIBERNATE-584563
[WARNING] - from: org.springframework.samples:spring-petclinic@2.3.1.BUILD-SNAPSHOT > org.springframework.boot:spring-boot-starter-data-jpa@2.3.1.RELEASE > org.hibernate:hibernate-core@5.4.17.Final
```

{% hint style="info" %}
If you experience an issue with Snyk URL the org or API token is incorrect.
{% endhint %}

{% hint style="info" %}
This build will take about 45 seconds to complete.
{% endhint %}



