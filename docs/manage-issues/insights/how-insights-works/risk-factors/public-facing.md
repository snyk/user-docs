# Public facing

Knowing if something is deployed gives you information that there is a potential that someone can abuse the flaw you worry about. That someone however may be a well-trusted person within your Organization or a completely unknown external entity.  Insights can further narrow down that list by determining if the package or Image is configured to be exposed to external traffic.

Cloud-native solutions such AWS, Kubernetes, and GCP are not only used to run workloads, but to provide network connectivity to the workloads, among many other things. This can be done in many ways by using core networking primitives like services, ingresses, load balancers or a combination of various networking solutions like Gloo, Istio, Envoy and so on.

The Snyk platform performs analysis on various data sources to compute the network connectivity paths. This information is used by Insights to determine if packages and Images may be exposed to external traffic.&#x20;

Cloud-native solutions are deterministic in the way the network connectivity is configured. Snyk uses that knowledge to compute the answers based on the available information. For example,  read Kubernetes [Services](https://kubernetes.io/docs/concepts/services-networking/service/) and [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/) documentation to understand how the network connectivity can be configured.&#x20;

{% hint style="info" %}
Insights currently supports the following configurations: Kubernetes services and ingress, and Gloo.
{% endhint %}

Note the following **technical details**:

The Kubernetes Connector continuously monitors the Kubernetes events. These events are streamed to the Snyk platform continuously.&#x20;

On a regular schedule, every hour, the data pipeline performs a reconciliation of the state of the cluster to create a snapshot. This snapshot is used to compute the network relationship between various resources like services to pods and ingresses to services. The same snapshot is used to extrapolate which images are running at that time.

Analysis is currently only performed based on the connectivity specification. The granularity of the analysis is performed on port and http path levels. There are various constraints that may be applied to configured paths such as network policies, security groups, and firewalls. Insights currently does not include the constraints in the computation.&#x20;

At the same interval, the data pipeline takes a snapshot of all Snyk Projects and data sources and extrapolates packages and images. This snapshot is used to determine which images and packages are known to Snyk for any given customer.&#x20;

Both snapshots are then compared and evidence graphs are generated to determine the public-facing facts at that point in time.

