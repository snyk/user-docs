---
description: >-
  Welcome! This lab demonstrates a workflow using Docker Desktop, Docker Hub,
  GitHub, and Snyk.
---

# Build Secure Containers with Docker and Snyk

## Lab Meta

> **Difficulty**: Beginner 
>
> **Time**: Approximately 30 minutes

This lab walks you through the demo shown in [Docker Build: Build Secure Containers with Docker and Snyk](https://www.youtube.com/watch?v=RGCHHaYP9Lw). We'll secure a containerized application in three ways:

1. First we ensure our container uses a secure **base image**,
2. Next we address any vulnerable **application dependencies**
3. Finally, we check its **deployment manifests** for misconfiguration.

### Pre-Requisites <a id="pre-requisites"></a>

#### Snyk

You'll need a Snyk account for this lab. If you need one, [sign up free at snyk.io](https://snyk.co/SnykDockerAcademy)

{% hint style="info" %}
New to Snyk? Run `docker scan --login`after installing Docker Desktop and register with your Docker ID to unlock a special 200 scan free tier limit for Snyk Container, usually 100! 
{% endhint %}

{% hint style="info" %}
Snyk offers [unlimited tests for Open source projects](https://snyk.io/blog/snyk-projects-now-free/). We love open source! 
{% endhint %}

#### Docker

You'll need a Docker ID. If you need one, [register for Docker Hub](https://hub.docker.com/signup). Also download and install [Docker Desktop](https://www.docker.com/products/docker-desktop), and [enable the Kubernetes cluster](https://birthday.play-with-docker.com/kubernetes-docker-desktop/) it ships. 

#### GitHub 

You'll need a GitHub Account. If you need one, [sign up free at GitHub](https://github.com/join). We also use GitHub Actions in this Lab. If you're new to GitHub Actions, [Introduction to GitHub Actions](https://docs.github.com/en/actions/learn-github-actions/introduction-to-github-actions) is great 101 reading.

When you're ready, head on to the next page to get started!

