# Summary

Our results may seem redundant but these are necessary. We begin with scanning and monitoring of our source in a Git repository by enabling Snyk's GitHub integration. This provides immediate insights into open source dependency vulnerabilities in our application. However, as developers can feature branch new vulnerabilities may be introduced as these may take place as _out-of-band_ pipeline builds. However, since we are including the Snyk CircleCI Orb in our configuration file, we perform an additional scan and report our findings back. 

We apply the same workflow with our private container registry on Amazon ECR and perform a second scan of the container image in our CircleCI job again, invoking the Snyk Orb and reporting our findings.

Lastly, we are monitoring our Kubernetes cluster and gaining additional insights not only into the application and container image vulnerabilities but potential misconfigurations of security configuration of those workloads. Therefore, establishing a comprehensive process of checks and balances to secure our workloads at each phase of the development process.

