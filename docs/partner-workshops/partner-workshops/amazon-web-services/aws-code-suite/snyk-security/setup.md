# Setup

### Save Password to Session Manager

Run the following command, replacing `abc123`, with **your unique token**. This places the token in the session parameter manager.

```text
aws ssm put-parameter --name "snykAuthToken" --value "abc123" --type SecureString
```

### Setup the Application Scanning

We want to insert testing with **Snyk** after `maven` has built the application. The simplest method is to insert commands to download, authorize and run the **Snyk** commands after Mvn has built the application/dependency tree.

In `modules/snyk/Dockerfile`, we have inserted the following commands to perform these actions

Set an environment variable from a value passed to the `docker build` command, this will contain the token for **Snyk**. By using an environment variable, **Snyk** will automatically detect the token when used.

```text
#~~~~~~~SNYK Variable~~~~~~~~~~~~
# Declare Snyktoken as a build-arg
ARG snyk_auth_token
# Set the SNYK_TOKEN environment variable
ENV SNYK_TOKEN=${snyk_auth_token}
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```

Download **Snyk**, run a test, looking for medium to high severity issues, and if the build succeeds, post the results to **Snyk** for monitoring and reporting. If a new vulnerability is found, you will be notified.

```text
# package the application
RUN mvn package -Dmaven.test.skip=true

#~~~~~~~SNYK test~~~~~~~~~~~~
# download, configure and run snyk. Break build if vulns present, post results to `https://snyk.io/`
RUN curl -Lo ./snyk "https://github.com/snyk/snyk/releases/download/v1.210.0/snyk-linux"
RUN chmod -R +x ./snyk
#Auth set through environment variable
RUN ./snyk test --severity-threshold=medium
RUN ./snyk monitor
```

### Setting Up Docker Scanning

Later in the build process, a docker image is created. We want to analyze it for vulnerabilities. We will do this in `buildspec.yml`. First, pull the **Snyk** token `snykAuthToken` from the `parameter store`:

```text
env:
  parameter-store:
    SNYK_AUTH_TOKEN: "snykAuthToken"
```

In the `prebuild` phase, we will install Snyk

```text
phases:
  pre_build:
    commands:

      - PWDUTILS=$(pwd)
      - curl -Lo ./snyk "https://github.com/snyk/snyk/releases/download/v1.210.0/snyk-linux"
      - chmod -R +x ./snyk
```

In the `build` phase we will pass the token to the **docker compose** command where it will be retrieved in the Dockerfile code we previously setup to test the application:

```text
build:
    commands:
      - docker build --build-arg snyk_auth_token=$SNYK_AUTH_TOKEN -t $REPOSITORY_URI:latest .
```

Next we will **authorize** the **Snyk** instance for testing the Docker image that’s produced. If it passes we will pass the results to **Snyk** for monitoring and reporting.

```text
build:
    commands:
      - $PWDUTILS/snyk auth $SNYK_AUTH_TOKEN
      - $PWDUTILS/snyk test --docker $REPOSITORY_URI:latest
      - $PWDUTILS/snyk monitor --docker $REPOSITORY_URI:latest
      - docker tag $REPOSITORY_URI:latest $REPOSITORY_URI:$IMAGE_TAG
```

In terminal, **navigate** to this folder:

```text
cd ~/environment/aws-modernization-workshop/
```

To try this module, let us copy the **Snyk** versions over to our build:

```text
cp modules/snyk/Dockerfile modules/containerize-application/Dockerfile
cp modules/snyk/buildspec.yml buildspec.yml
```

### 

