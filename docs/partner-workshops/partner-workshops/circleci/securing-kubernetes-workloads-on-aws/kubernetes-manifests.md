# Kubernetes manifests

To simply the deployment of our application, we have provided two [configuration files](https://kubernetes.io/docs/concepts/configuration/overview/) which will create the necessary resources to run our application. These two files will consist of a Deployment and a Service and are available in the `./deployment` directory of the accompanying GitHub repository. Let's take a look at what's in each of these configuration files.

#### Deployment

A [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) will define the desired state and create a ReplicaSet. Here we are defining the container image we will use as well as `cpu` and `memory` limits among other things.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goof
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goof
      tier: frontend
  template:
    metadata:
      labels:
        app: goof
        tier: frontend
    spec:
      containers:
        - name: goof
          image: DOCKER_IMAGE_NAME
          resources:
            requests:
              cpu: 100m
              memory: 100Mi
          ports:
            - containerPort: 3001
            - containerPort: 9229
          env:
            - name: DOCKER
              value: "1"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goof-mongo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goof
      tier: backend
  template:
    metadata:
      labels:
        app: goof
        tier: backend
    spec:
      containers:
        - name: goof-mongo
          image: mongo
          ports:
            - containerPort: 27017
```

#### Service

A [service](https://kubernetes.io/docs/concepts/services-networking/service/) is an abstraction for exposing an application running on a set of Pods as a network service. Here, we are defining a two-tier application that consists of a `frontend` which is our _**goof**_ application s well as a `backend` which in our case will be a `mongodb` container image.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: goof
spec:
  type: LoadBalancer
  ports:
  - protocol: TCP
    port: 80
    targetPort: 3001
    name: "http"
  - protocol: TCP
    port: 9229
    targetPort: 9229
    name: "debug"
  selector:
    app: goof
    tier: frontend
---
apiVersion: v1
kind: Service
metadata:
  name: goof-mongo
spec:
  ports:
  - protocol: TCP
    port: 27017
    targetPort: 27017
    name: "mongo"
  selector:
    app: goof
    tier: backend
```

