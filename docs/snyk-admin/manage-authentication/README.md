# Manage authentication

## Introduction

Snyk offers API tokens to enable integrations with third-party developer tools for authentication of your personal account, or alternatively in lieu of credentials to integrate a service account.

{% hint style="info" %}
For authentication purposes, the identity providers **do not require access to your repositories**, only your email address.
{% endhint %}

### Supported identity providers

With Snyk, you can use one of the following identity providers for authentication:

* GitHub
* Bitbucket
* Google
* Azure AD
* Docker ID
* Single Sign-On (SSO) - see [Setting up Single Sign-On (SSO) for authentication](../../enterprise-setup/set-up-single-sign-on-sso-for-authentication/)

{% hint style="info" %}
**Feature availability**\
Single sign-on is available with Enterprise plans. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

{% hint style="warning" %}
Logging in with a different provider from the one that you registered with when you first created your Snyk account will create a separate new Snyk account.
{% endhint %}
