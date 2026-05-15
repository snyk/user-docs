---
description: This page explains how to transition away from affected images.
---

# Snyk Images migration

You can [create your own custom images](../user-defined-custom-images-for-cli.md) to use. This option is available for any user affected by the removal of the Snyk Images listed on this page.

Creating a custom image is the preferred solution, as it should guarantee compatibility with your system. However, if you cannot create a custom image, you can use an available alternative image.

{% hint style="info" %}
Snyk recommends using a pinned version of a software package unless stated otherwise.
{% endhint %}

## snyk/snyk:docker-\* <a href="#snyk-snyk-docker" id="snyk-snyk-docker"></a>

See the following table for the `snyk/snyk:docker-*` images to be removed

| Image                  | Based on     |
| ---------------------- | ------------ |
| snyk/snyk:docker-18.09 | docker:18.09 |
| snyk/snyk:docker-19.03 | docker:19.03 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:docker`
* `snyk/snyk:docker-latest`

{% hint style="info" %}
`snyk/snyk:docker` is based on Docker stable and is preferred.
{% endhint %}

## snyk/snyk:golang-\* <a href="#snyk-snyk-golang" id="snyk-snyk-golang"></a>

See the following table for the `snyk/snyk:golang-*` images to be removed.

| Image                 | Based on    |
| --------------------- | ----------- |
| snyk/snyk:golang-1.12 | golang:1.12 |
| snyk/snyk:golang-1.13 | golang:1.13 |
| snyk/snyk:golang-1.14 | golang:1.14 |
| snyk/snyk:golang-1.15 | golang:1.15 |
| snyk/snyk:golang-1.16 | golang:1.16 |
| snyk/snyk:golang-1.17 | golang:1.17 |
| snyk/snyk:golang-1.18 | golang:1.18 |
| snyk/snyk:golang-1.19 | golang:1.19 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:golang`
* `snyk/snyk:golang-1.20`
* `snyk/snyk:golang-1.21`
* `snyk/snyk:golang-1.22`

{% hint style="info" %}
It is recommended to use a pinned Golang version such as `snyk/snyk:golang-1.22` instead of `snyk/snyk:golang,` which is based on Golang latest.
{% endhint %}

## snyk/snyk:gradle-\* <a href="#snyk-snyk-gradle" id="snyk-snyk-gradle"></a>

See the following table for the `snyk/snyk:gradle-*` images to be removed.

| Image                      | Based on         |
| -------------------------- | ---------------- |
| snyk/snyk:gradle-6.4       | gradle:6.4       |
| snyk/snyk:gradle-6.4-jdk11 | gradle:6.4-jdk11 |
| snyk/snyk:gradle-6.4-jdk14 | gradle:6.4-jdk14 |
| snyk/snyk:gradle-6.4-jdk8  | gradle:6.4-jdk8  |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:gradle`
* `snyk/snyk:gradle-jdk11`
* `snyk/snyk:gradle-jdk12`
* `snyk/snyk:gradle-jdk13`
* `snyk/snyk:gradle-jdk14`
* `snyk/snyk:gradle-jdk16`
* `snyk/snyk:gradle-jdk17`
* `snyk/snyk:gradle-jdk18`
* `snyk/snyk:gradle-jdk19`
* `snyk/snyk:gradle-jdk20`
* `snyk/snyk:gradle-jdk21`
* `snyk/snyk:gradle-jdk8`

{% hint style="info" %}
- `snyk/snyk:gradle` uses the latest stable version of Gradle and Java,  8.8 and JDK 21 at the time of publication.
- The Gradle images use Gradle 8.8 unless otherwise specified, along with the specified JDK.
{% endhint %}

## snyk/snyk:maven-3-jdk-19 <a href="#snyk-snyk-maven-3-jdk-19" id="snyk-snyk-maven-3-jdk-19"></a>

`snyk/snyk:maven-3-jdk-19` will be removed from Docker. Users affected by this removal can update to the following images:

* `snyk/snyk:maven`
* `snyk/snyk:maven-3-jdk-11`
* `snyk/snyk:maven-3-jdk-17`
* `snyk/snyk:maven-3-jdk-20`
* `snyk/snyk:maven-3-jdk-21`
* `snyk/snyk:maven-3-jdk-22`
* `snyk/snyk:maven-3-jdk-8`

{% hint style="info" %}
`snyk/snyk/maven` uses the latest stable version of Maven, 3.9 at the time of publication.
{% endhint %}

## snyk/snyk:dotnet-\* <a href="#snyk-snyk-dotnet" id="snyk-snyk-dotnet"></a>

See the following table for the `snyk/snyk:dotnet-*` images to be removed.

| Image                | Based on                              |
| -------------------- | ------------------------------------- |
| snyk/snyk:dotnet-2.1 | mcr.microsoft.com/dotnet/core/sdk:2.1 |
| snyk/snyk:dotnet-2.2 | mcr.microsoft.com/dotnet/core/sdk:2.2 |
| snyk/snyk:dotnet-3.0 | mcr.microsoft.com/dotnet/core/sdk:3.0 |
| snyk/snyk:dotnet-3.1 | mcr.microsoft.com/dotnet/core/sdk:3.1 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:dotnet`
* `snyk/snyk:dotnet-8.0`

{% hint style="info" %}
It is recommended to use the pinned version `snyk/snyk:dotnet-8.0,` which is based on the active LTS version of .NET, 8.0 at the time of publication.
{% endhint %}

## snyk/snyk:node-\* <a href="#snyk-snyk-node" id="snyk-snyk-node"></a>

See the following table for the `snyk/snyk:node-*` images to be removed.

| Image             | Based on |
| ----------------- | -------- |
| snyk/snyk:node-8  | node:8   |
| snyk/snyk:node-10 | node:10  |
| snyk/snyk:node-12 | node:12  |
| snyk/snyk:node-13 | node:13  |
| snyk/snyk:node-14 | node:14  |
| snyk/snyk:node-15 | node:15  |
| snyk/snyk:node-16 | node:16  |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:node`
* `snyk/snyk:node-18`
* `snyk/snyk:node-20`
* `snyk/snyk:node-22`

{% hint style="info" %}
It is recommend to use a pinned version such as `snyk/snyk:node-22`, which uses LTS version 22.
{% endhint %}

## snyk/snyk:python-\* <a href="#snyk-snyk-python" id="snyk-snyk-python"></a>

See the following table for the `snyk/snyk:python-*` images to be removed.

| Image                | Based on   |
| -------------------- | ---------- |
| snyk/snyk:python-2.7 | python:2.7 |
| snyk/snyk:python-3.6 | python:3.6 |
| snyk/snyk:python-3.7 | python:3.7 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:python`
* `snyk/snyk:python-alpine`
* `snyk/snyk:python-3.8`
* `snyk/snyk:python-3.9`
* `snyk/snyk:python-3.10`
* `snyk/snyk:python-3.11`

{% hint style="info" %}
It is recommend to use a pinned version such as `snyk/snyk:python-3.11`, which uses LTS version 3.11.
{% endhint %}

## snyk/snyk:ruby-\* <a href="#snyk-snyk-ruby" id="snyk-snyk-ruby"></a>

See the following table for the `snyk/snyk:ruby-*` images to be removed.

| Image              | Based on |
| ------------------ | -------- |
| snyk/snyk:ruby-2.4 | ruby:2.4 |
| snyk/snyk:ruby-2.5 | ruby:2.5 |
| snyk/snyk:ruby-2.6 | ruby:2.6 |
| snyk/snyk:ruby-2.7 | ruby:2.7 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:ruby`
* `snyk/snyk:ruby-alpine`
* `snyk/snyk:ruby-3.3`

{% hint style="info" %}
It is recommended to use the pinned version `snyk/snyk:ruby-3.3`, which is based on Ruby 3.3
{% endhint %}

## snyk/snyk:sbt & snyk/snyk:scala <a href="#snyk-snyk-sbt-and-snyk-snyk-scala" id="snyk-snyk-sbt-and-snyk-snyk-scala"></a>

See the following table for the `snyk/snyk:sbt` and `snyk/snyk:scala` images to be removed.

| Image           | Based on                                  |
| --------------- | ----------------------------------------- |
| snyk/snyk:sbt   | hseeberger/scala-sbt:8u212\_1.2.8\_2.13.0 |
| snyk/snyk:scala | hseeberger/scala-sbt:8u212\_1.2.8\_2.13.0 |

Users affected by the removal of these images can update to the following images:

* `snyk/snyk:ruby`
* `snyk/snyk:ruby-alpine`
* `snyk/snyk:ruby-3.3`

{% hint style="info" %}
It is recommended to use the pinned version snyk/snyk:sbt1.10.0-scala3.4.2, which is based on Scala:3.4.2-sbt:1.10.0
{% endhint %}
