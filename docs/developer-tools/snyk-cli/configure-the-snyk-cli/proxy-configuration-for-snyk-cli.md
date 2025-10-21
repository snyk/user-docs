# Proxy configuration for Snyk CLI

## General proxy configuration

When you use the Snyk CLI behind a proxy, you must provide the proxy configuration by using the following environment variables:

`HTTP_PROXY` or `http_proxy`

`HTTPS_PROXY` or `https_proxy`

`NO_PROXY` or `no_proxy`

The `https` in the `HTTPS_PROXY` means that requests using `https` protocol use this proxy. The proxy itself does not need to use `https`.

For more information, see [Configure the Snyk CLI to connect to the Snyk API](configure-snyk-cli-to-connect-to-snyk-api.md) and [How can I use Snyk behind a proxy?](https://support.snyk.io/s/article/How-can-I-use-Snyk-behind-a-proxy)

## Proxy authentication

By default Snyk CLI tries to detect and apply proxy authentication.

If the proxy server requests proxy authentication (as indicated by a `PROXY-AUTHENTICATE` response header), and both server and CLI support the same authentication mechanism, the CLI authenticates as the user who is currently logged in to the operating system (SSO).

This is supported for the following authentication mechanism:

* Negotiate
  * Kerberos (on all OS)
  * NTLM (Windows NT LAN Manager)

## Configuration on Windows operating systems

On Windows operating systems (OS), the Kerberos and NTLM authentication mechanisms are provided by the OS itself and available for all domain users.

Snyk CLI does not require any specific configuration.

## Configuration on non-Windows operating systems (Linux, macOS)

On non-Windows operating systems, Snyk CLI also supports SSO but, in addition, must be configured with the following environment variables.

```bash
KRB5_CONFIG # default "/etc/krb5.conf"
KRB5CCNAME # default "FILE:/tmp/krb5cc_<UserUID>"
```

The use of these variables follows the MIT Kerberos implementation:

* [KRB5\_CONFIG](https://web.mit.edu/kerberos/krb5-devel/doc/admin/conf_files/krb5_conf.html) - Kerberos configuration file
* [KRB5CCNAME](https://web.mit.edu/kerberos/krb5-1.12/doc/basic/ccache_def.html) - Kerberos credential cache
  * Currently, the only supported credential cache type (`ccache` types) is **FILE.**
  * It is important to note that the cache file cannot be updated by the CLI. This means that the cache file must be updated externally, for example, by running [kinit](https://web.mit.edu/kerberos/krb5-1.12/doc/user/user_commands/kinit.html).,

## Disable proxy authentication

To disable authentication, specify the following command line parameter:

```jsx
--proxy-noauth
```

## Troubleshooting

If you have connection problems, enable debug output `-d` for helpful insights.
