# Deployment tips and reference architectures

{% hint style="info" %}
You can learn more about deployment and architecture, and more, with the [Universal Broker Snyk Learn ](https://learn.snyk.io/lesson/universal-broker/#3f799f7f-58a9-4225-53fb-9cc5b6913920)course.
{% endhint %}

The Universal Broker allows you to run multiple connections of any type using a single container. Snyk recommends running in High Availability Mode (on by default), to run multiple replicas.

On Kubernetes systems, the Helm chart creates a stateful set with numerous members, automatically creating multiple replicas.

In other orchestrators, you must create multiple replicas explicitly in your deployment configuration.

Usage of resources varies based on a number of factors, making it difficult to model the actual use of resources (CPU and memory) for the container. Each deployment is limited to a maximum of 25 connections to avoid exhaustion of resources.

If you find you need more resource segregation or allocation that does not fit well into a single container setup, you can create a new deployment with its set of connections to split the load across numerous Broker clients and containers.

Note that moving connections from deployment A to deployment B is not supported.

Contact your Snyk account team or [Snyk support](https://support.snyk.io/s/) if you need more assistance designing a Universal Broker architecture for your needs.

\
