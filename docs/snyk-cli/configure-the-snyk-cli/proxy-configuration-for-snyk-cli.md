# Proxy configuration for Snyk CLI

{% hint style="info" %}
Proxy authentication for Snyk CLI is currently available on Windows only. Support for other operating systems will be available later.
{% endhint %}

## General proxy configuration

When you use the Snyk CLI behind a proxy, you must provide the proxy configuration by using the following environment variables:

```bash
HTTP_PROXY or http_proxy
HTTPS_PROXY or https_proxy
NO_PROXY or no_proxy
```

The `https` in the `HTTPS_PROXY` means that requests using __ `https` __ protocol use this proxy. The proxy itself does not need to use `https`.

For more information, see [Configure the Snyk CLI to connect to the Snyk API](configure-snyk-cli-to-connect-to-snyk-api.md) and [How can I use Snyk behind a proxy?](https://support.snyk.io/hc/en-us/articles/360000925358-How-can-I-use-Snyk-behind-a-proxy-)

## Proxy authentication

By default Snyk CLI tries to detect and apply proxy authentication.

If the proxy server requests proxy authentication (as indicated by a `PROXY-AUTHENTICATE` response header) and both server and CLI support the same authentication mechanism, the CLI authenticates as the user who is currently logged in to the operating system (SSO).

This is supported for the following authentication mechanism:

* Negotiate
  * Kerberos (on all OS)
  * NTLM (Windows NT LAN Manager)

## Configuration on Windows operating systems

On Windows operating systems (OS), Kerberos and NTLM authentication mechanisms are provided by the OS itself and available for all domain users.

Snyk CLI does not require any specific configuration.

## Disable proxy authentication

To disable authentication, specify the following command line parameter:

```jsx
--proxy-noauth
```

## Troubleshooting

If you have connection problems, enable debug output `-d` for helpful insights.
