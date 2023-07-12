# Insights assets

Insights works by analyzing the assets described on this page.

## Images

Images are assets that represent a Docker image. Snyk Container performs security scans on the Docker images. Images can be mapped one to many with the Snyk Project created by the scans performed by Snyk Container. Docker images have natural IDs, which are represented by SHAs. Insights uses this natural ID to correlate the same images even if they are mapped to different Snyk Projects.

## Kubernetes resources

Kubernetes resources are assets that represent a Kubernetes object. The Kubernetes Connector collects resource information from the Kubernetes clusters.&#x20;

Kubernetes resources do not map to the Snyk Projects. These are internal entities used to compute certain risk factors, further detailed below, which can be related to the packages and images.

## Packages

Packages are assets that represent a software package. Snyk Open Source and Snyk Code products perform security scans on files. These files represent the package manager declaration and the source code of a software package, respectively. A package is a representation of that software package.

Packages can be mapped one to one with the Snyk Projects created by the scans performed by the Snyk Open Source and Snyk Code. All the issues identified by these products and attributed to these Projects will be mapped to the package entity.&#x20;

Package is a very coarse abstraction. It does not have versions. It is a representation of the current state of the software package at a point in time. The point in time is determined by the time when the Insights processing pipeline completes and the state of Snyk Projects at that time.&#x20;

Snyk Open Source uses the word package to refer to the third-party dependencies declared in the package dependency manifest. Insights does not currently expose the granularity of the third-party dependencies. However, from the Insights data model point of view, there is no distinction between third-party and first-party packages. These would be represented as a package object at the point in time.
