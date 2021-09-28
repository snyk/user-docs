# Clone and run the sample application

## Step 1: Get the sample app and clone the Repo

A template GitHub Repo is provided for this workshop using [Snyk's Goof application](https://github.com/snyk/goof#goof---snyks-vulnerable-demo-app).

{% embed url="https://github.com/snyk-partners/docker-academy" %}

Click "Use this Template" to copy the Repo into your GitHub account.

{% hint style="danger" %}
To properly copy-paste the commands in these instructions name your repo `docker-academy`. Otherwise, ensure the command uses the correct repo name.
{% endhint %}

If you want to copy-paste commands, set an environment variable for your GitHub ID.

```bash
# Set an environment variable for your GitHub ID
GithubId=<<your_github_id>>
```

Now, clone the Repo to your local environment, then `cd` into it.

```bash
# Clone the Repo and cd 
git clone https://www.github.com/$GithubId/docker-academy && cd docker-academy
```

## Step 2: Test the Application locally

### Install Dependencies

To run our application locally, we need to install its dependencies first.

```bash
# Install npm dependencies and generate lockfile
npm install --package-lock
```

### Start the Mongo database

The app requires MongoDB to work. Instead of installing Mongo, we'll run it in a Docker container.

```bash
# Start detached mongo with container port 27017 mapped to host port 27017
docker run -d -p 27017:27017 mongo
```

You can verify that mongo is running by running `docker ps`.

### Start the Application

Now you can run the application. 

```bash
node app.js
```

Once it starts, it will be available at [http://localhost:3001](http://localhost:3001). Verify it works by adding a few items into the todo list. 

![](../../../.gitbook/assets/todo.png)

Success! Now that we know it works, let's package it in a container for distribution.

## Build the Docker Image

If you want to copy-paste commands, set an environment variable for your Docker ID.

```bash
# Set an environment variable for your GitHub ID
DockerId=<<your_docker_id>>
```

The `Dockerfile` in the repo tells `docker build` how to build the container. To learn more about Dockerfile commands, visit [Docker's documentation](https://docs.docker.com/engine/reference/builder/). A summary is provided below:

{% code title="Dockerfile" %}
```bash
FROM node:10.4 ## Use Node 10.4 as our base image

RUN mkdir /usr/src/goof ## Create the directory for our application
RUN mkdir /tmp/extracted_files
COPY . /usr/src/goof ## Copy the current directory files into the image
WORKDIR /usr/src/goof ## Set the working directory

RUN npm update ## Ensure npm is up to date
RUN npm install ## Install dependencies
EXPOSE 3001 ## Expose the port
EXPOSE 9229
ENTRYPOINT ["npm", "start"] ## The command 
```
{% endcode %}

When ready, build and tag the container image to get it ready to push into Docker Hub.

```bash
# Build the container and tag it dev
docker build -t $DockerId/goof:dev .
```

When the build finishes, push the container to Docker Hub.

```bash
# Push to Docker Hub
docker push $DockerId/goof:dev
```

## Deploy the application to Kubernetes

To ensure the container deploys correctly into Kubernetes without incurring cloud costs, we use the Kubernetes cluster shipped with Docker Desktop. Before deploying the app, you'll need to make a change to the app's deployment manifest. Open the file `goof-deployment.yaml` in a text editor.

Find these lines, and insert your Docker ID where indicated. 

```yaml
spec:
  containers:
  - image: <<DOCKERID>>/goof:dev #Edit with your Docker Hub ID
    name: goof
    imagePullPolicy: Always
```

Save your changes. Now you're ready to deploy the application. Run the following commands:

```bash
# Create a namespace
kubectl create ns snyk-docker

# Set the current context to use the new namespace
kubectl config set-context --current --namespace snyk-docker

# Spin up the goof deployment and service
kubectl create -f goof-deployment.yaml,goof-mongo-deployment.yaml
```

To check the status of the pods as the application comes up, use the following command:

```bash
kubectl get pods
```

Once both are running, the application should now be accessible in [http://localhost:3001](http://localhost:3001). Success!

Now that we know our application works when deployed into Kubernetes, we'll set up a continuous integration and continuous delivery pipeline to re-build our container as code changes are made.

