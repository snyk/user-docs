# Public facing

Knowing that code is deployed tells you that there is a possibility that someone can exploit a flaw you are concerned about. That someone may be a well-trusted person within your Organization or a completely unknown external entity. Snyk can further narrow down the possibilities by determining if the package or Image is configured to be exposed to external traffic.

Cloud-native solutions such as AWS, Kubernetes, and GCP are used for many purposes, including running workloads and providing network connectivity to the workloads. This can be done in many ways, including using core networking primitives like services, ingresses, load balancers, or a combination of various networking solutions like Gloo, Istio, Envoy, and so on.

## Snyk analysis to determine the network connectivity paths

The Snyk platform analyzes various data sources to compute the network connectivity paths. This information is used by Snyk AppRisk Essentials to determine if packages and Images may be exposed to external traffic.&#x20;

Cloud-native solutions are deterministic in the way the network connectivity is configured. Snyk uses that knowledge to compute the answers based on the available information. For example,  to understand how the network connectivity can be configured, see Kubernetes [Services](https://kubernetes.io/docs/concepts/services-networking/service/) and [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/) documentation.&#x20;

{% hint style="info" %}
Risk-based prioritization currently supports the following configurations: Kubernetes services and ingress, and Gloo.
{% endhint %}

## Technical details for the Public facing risk factor

The Kubernetes Connector continuously monitors the Kubernetes events. These events are streamed to the Snyk platform continuously.&#x20;

Every hour, the data pipeline performs a reconciliation of the state of the cluster to create a snapshot. This snapshot is used to compute the network relationship between various resources, such as services to pods and ingresses to services. The same snapshot is used to extrapolate which images are running at that time.

The analysis is currently performed only based on the connectivity specification. The granularity of the analysis is performed on `port` and `http path` levels. Various constraints may be applied to configured paths, such as network policies, security groups, and firewalls. Snyk currently does not include the constraints in the computation.&#x20;

At the same interval, the data pipeline takes a snapshot of all Snyk Projects and data sources and extrapolates packages and images. This snapshot is used to determine which images and packages are known to Snyk for any given customer.&#x20;

Both snapshots are then compared, and evidence graphs are generated to determine the public-facing factors at that point in time.

