# Recap & Next Steps

## Congratulations! <a id="congratulations"></a>

You reached the end of this Lab, what a journey! We hope that as you progressed through the lab you saw how Snyk and GitHub Actions together can help facilitate a secure Continuous Integration and Continuous Delivery paradigm for your software delivery practice.‌

We hope you enjoyed this lab. Below we recap what you've accomplished, and provide additional resources to help you get more value out of Snyk,.

## Snyk Open Source <a id="snyk-open-source"></a>

In part 1, you used Snyk Open Source to find vulnerabilities in the Open Source components for the sample application. You configured the GitHub integration, created fix Pull Requests, and built a gate into the release process to ensure issues did not make their way into a Production branch.‌

Some things we didn't cover:

#### Automatic Dependency Upgrade Pull Requests <a id="automatic-dependency-upgrade-pull-requests"></a>

Why wait until vulnerabilities are published to upgrade your dependencies? Snyk can be configured to [automatically open pull requests on your behalf](https://support.snyk.io/hc/en-us/articles/360006581898-Upgrading-dependencies-with-automatic-PRs#:~:text=Enable%20Snyk%20to%20regularly%20check,available%20for%20the%20specific%20project.), to keep your dependencies up to date and healthy.‌

#### Open Snyk Pull Requests using a fixed GitHub Account <a id="open-snyk-pull-requests-using-a-fixed-github-account"></a>

Snyk allows you to [configure a specific GitHub account](https://support.snyk.io/hc/en-us/articles/360010843797-Opening-fix-and-upgrade-pull-requests-from-a-fixed-GitHub-account-) on whose behalf the fix and upgrade PRs will be opened. Our research shows that this increases the likelihood of a Fix PR getting merged, so check it out!‌

## Snyk Container <a id="snyk-container"></a>

In Part 2, you used Snyk Container to find issues introduced by your choice of base image. You imported your Dockerfile into Snyk, and selected a less vulnerable base image for your application. 

If you're working with containers, some resources worth checking out that we didn't cover:

#### Integrate with your Container Registry

Snyk can integrate with your Container Registry to easily import your container images for scanning with Snyk. We support [Docker Hub](https://support.snyk.io/hc/en-us/articles/360003916058-Configure-integration-for-Docker-Hub), [Azure Container Registry](https://support.snyk.io/hc/en-us/articles/360003915958-Configure-integration-for-ACR), [Google Container Registry](https://support.snyk.io/hc/en-us/articles/360003916118-Configure-integration-for-GCR), [Artifactory Container Registry](https://support.snyk.io/hc/en-us/articles/360003915998-Configuring-your-JFrog-Artifactory-container-registry-integration), and [AWS Elastic Container Registry](https://support.snyk.io/hc/en-us/articles/360003916078-Configure-integration-for-Amazon-Elastic-Container-Registry-ECR-). 

#### Use Snyk with Docker's Tools

Snyk and Docker partnered to bring Snyk scanning into Docker Desktop and Docker Hub. If you're using those tools in your development process, and want to learn more, check out our [Lab: Build Secure Containers with Docker and Snyk](../../docker/lab-build-secure-containers-with-docker-and-snyk/) to see it in action! 

## Snyk Infrastructure as Code

In Part 3, you used Snyk Infrastructure as Code to find and fix configuration issues in your Kubernetes deployment manifests, and set up a gate that can prevent configuration issues from accidentally making their way into a Production environment. 

#### Configure IaC Rule Severity

Infrastructure as Code rules are not meant to be one-size-fits-all. Different workloads have different security requirements and tolerances, that's why we allow you to change how Snyk IaC scores your application configurations. Learn how to [adjust the severity scoring for IaC](https://support.snyk.io/hc/en-us/articles/360006402818#UUID-c1919782-6bfa-b84b-a638-3913cee39fc5) rules in our Docs.

#### Kubernetes Integration

An extension of Snyk Container and Snyk IaC, we offer a Kubernetes Monitor that can identify configuration issues and container vulnerabilities in running workloads. To learn more, check out the [Snyk Kubernetes Integration Overview](https://support.snyk.io/hc/en-us/articles/360003916138-Kubernetes-integration-overview).

#### Terraform Support

We didn't cover it in this Lab, but Snyk can also scan Terraform files for configuration issues. To learn more about our Terraform support, check out how to [Scan and Fix issues in your Terraform files](https://support.snyk.io/hc/en-us/articles/360010916577-Scan-and-fix-security-issues-in-your-Terraform-files).



