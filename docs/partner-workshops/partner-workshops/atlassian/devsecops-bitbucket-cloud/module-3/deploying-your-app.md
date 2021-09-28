# Deploying your app

The final step in your sample [bitbucket-pipelines.yml](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/08298e9b6d3108d33bd54a4839e92884c79c8597/bitbucket-pipelines.yml#lines-57:79) file will be to take the container image that you have scanned with Snyk and securely stored in Amazon ECR and deploy this to your Amazon EKS cluster.

```yaml
deploy-app: &deploy-app
  - step:
      name: "Deploy application"
      deployment: staging
      script:
        - pipe: atlassian/aws-eks-kubectl-run:1.2.4
          variables:
            AWS_ACCESS_KEY_ID: '$AWS_ACCESS_KEY_ID'
            AWS_SECRET_ACCESS_KEY: '$AWS_SECRET_ACCESS_KEY'
            AWS_DEFAULT_REGION: '$AWS_DEFAULT_REGION'
            CLUSTER_NAME: '$AWS_EKS_CLUSTER'
            KUBECTL_COMMAND: 'apply'
            RESOURCE_PATH: "./deployment/goof-service.yaml"
        - envsubst < ./deployment/goof-deployment-template.yaml > ./deployment/goof-deployment.yaml
        - cat ./deployment/goof-deployment.yaml
        - pipe: atlassian/aws-eks-kubectl-run:1.2.4
          variables:
            AWS_ACCESS_KEY_ID: '$AWS_ACCESS_KEY_ID'
            AWS_SECRET_ACCESS_KEY: '$AWS_SECRET_ACCESS_KEY'
            AWS_DEFAULT_REGION: '$AWS_DEFAULT_REGION'
            CLUSTER_NAME: '$AWS_EKS_CLUSTER'
            KUBECTL_COMMAND: 'apply'
            RESOURCE_PATH: './deployment/goof-deployment.yaml'
```

In this example, we are leveraging the [aws-eks-kubectl-run](https://bitbucket.org/atlassian/aws-eks-kubectl-run) pipe to apply our [service](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/master/deployment/goof-service.yaml) and [deployment](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/master/deployment/goof-deployment-template.yaml) manifests against our running cluster. Here, we are referencing some of our previously defined _repository variables_ but we are also invoking the [`envsubst`](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html#index-envsubst) linux command to substitute the value of one of our variables.

The [goof-deployment-template.yaml](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/08298e9b6d3108d33bd54a4839e92884c79c8597/deployment/goof-deployment-template.yaml#lines-19) file in the `./deployments` directory of our repository contains two variables `${AWS_ECR_URI}` and ${BITBUCKET\_COMMIT} which we are substituting with the value of our docker tag, allowing us to pull the correct image from Amazon ECR.

```yaml
    spec:
      containers:
        - name: goof
          image: ${AWS_ECR_URI}:${BITBUCKET_COMMIT}
```

The [goof-service.yaml](https://bitbucket.org/snyk/patterns-library-atlassian-aws/src/master/deployment/goof-service.yaml) file is creating our `service` and deploying our _frontend_ app as `type: LoadBalancer`, exposing this on the standard `http` port `80`.

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

Let's proceed to the next section.

