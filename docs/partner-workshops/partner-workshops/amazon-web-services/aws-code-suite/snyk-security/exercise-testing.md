# Exercise - Testing

In the `Containerize Application` lab you saw how to build your application. In this exercise you will try to run your build, which will fail due to security vulnerabilities being found. While normally done during the code development phase, we will take you through the process of fixing the vulnerability, and then re-running the exercise to see the build succeed.

_Save changes_:

```text
git commit -am "snyk"
```

_Push_:

```text
git push -f codecommit master
```

Now in `CodeBuild`, look at your **build history**. Note it may take a minute or two for the new scan to run.

![](../../../../.gitbook/assets/snyk_4_build.png)

Let’s look at why this failed. We see security vulnerabilities were found and we’re told **how** to fix it!

```text
Testing /usr/src/app...
✗ Medium severity vulnerability found in org.primefaces:primefaces
Description: Cross-site Scripting (XSS)
Info: https://snyk.io/vuln/SNYK-JAVA-ORGPRIMEFACES-31642
Introduced through: org.primefaces:primefaces@6.1
From: org.primefaces:primefaces@6.1
Remediation:
Upgrade direct dependency org.primefaces:primefaces@6.1 to org.primefaces:primefaces@6.2 (triggers upgrades to org.primefaces:primefaces@6.2)
✗ Medium severity vulnerability found in org.primefaces:primefaces
Description: Cross-site Scripting (XSS)
Info: https://snyk.io/vuln/SNYK-JAVA-ORGPRIMEFACES-31643
Introduced through: org.primefaces:primefaces@6.1
From: org.primefaces:primefaces@6.1
Remediation:
Upgrade direct dependency org.primefaces:primefaces@6.1 to org.primefaces:primefaces@6.2 (triggers upgrades to org.primefaces:primefaces@6.2)
Organisation: sample-integrations
Package manager: maven
Target file: pom.xml
Open source: no
Project path: /usr/src/app
Tested 37 dependencies for known vulnerabilities, found 2 vulnerabilities, 2 vulnerable paths.
The command '/bin/sh -c ./snyk test' returned a non-zero code: 1
[Container] 2018/11/09 03:46:22 Command did not exit successfully docker build --build-arg snyk_auth_token=$SNYK_AUTH_TOKEN -t $REPOSITORY_URI:latest . exit status 1
[Container] 2018/11/09 03:46:22 Phase complete: BUILD Success: false
[Container] 2018/11/09 03:46:22 Phase context status code: COMMAND_EXECUTION_ERROR Message: Error while executing command: docker build --build-arg snyk_auth_token=$SNYK_AUTH_TOKEN -t $REPOSITORY_URI:latest .. Reason: exit status 1
```

