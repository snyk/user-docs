# How can I use Snyk behind a proxy?

* [ How efficient is our Vulnerability Database ?](/hc/en-us/articles/360003968978-How-efficient-is-our-Vulnerability-Database-)
* [ What counts as a test?](/hc/en-us/articles/360000925418-What-counts-as-a-test-)
* [ How can I use Snyk behind a proxy?](/hc/en-us/articles/360000925358-How-can-I-use-Snyk-behind-a-proxy-)
* [ How are patches validated?](/hc/en-us/articles/360000925338-How-are-patches-validated-)
* [ Does Snyk have an API?](/hc/en-us/articles/360000914857-Does-Snyk-have-an-API-)
* [ I can't import projects from one of my GitHub organisations](/hc/en-us/articles/360000914837-I-can-t-import-projects-from-one-of-my-GitHub-organisations)

##  How can I use Snyk behind a proxy?

Some common issues when working behind a proxy:

* Snyk CLI requests are not going through the proxy
* Requests to app.snyk.io are blocked by the proxy
* Snyk CLI does not recognize an authorized certificate and prevents generating insecure requests

### Snyk CLI requests are not going through the proxy

First check what are the details of your proxy. In Windows, run a `ipconfig /all` command and look for your proxy settings. You must retrieve the following information:

* http or https
* private IP of the proxy
* port \(80, 8080 or 443 typically\).

To run Snyk from behind a proxy, use an environment variable to point to your proxy.

Snyk supports a set of environment variables to enforce specific Proxy settings:

* HTTP\_PROXY / http\_proxy
* HTTPS\_PROXY / https\_proxy
* NO\_PROXY / no\_proxy

For example, to configure a proxy and run a Snyk test with the new proxy value, run:

```text
$ export https_proxy=https://my.corporate.proxy:8080
$ snyk test
```

Notes:

* The command to add an environment variable could differ between Windows CMD, Powershell, Mac terminal and Linux. 
* `export https_proxy=value` sets a global env variable. If you prefer to set a local env variable you can use `set https_proxy=value` instead

### Requests to app.snyk.io are blocked by the proxy

To check if this is the case, add a debugging option to the CLI, using:

```text
$ snyk test -d
```

If you see a **timeout error** or a **econnrefused** error in the request, then **https://app.snyk.io** may be blocked by Snyk. 

### Snyk CLI does not recognize an authorized certificate and prevents generating insecure requests

To check if this is the case, you can tell Snyk to ignore insecure certificates, knowing that your proxy could be the reason for presenting a custom certificate to Snyk.

You can test this by running:

```text
$ snyk auth --insecure
$ snyk test --insecure
```

If the **--insecure** flag helps, you may want to point Snyk towards your custom CA certs more permanently by exporting them \(Snyk CLI is a Node application and honors environmental variables\):

```text
$ export NODE_EXTRA_CA_CERTS=/path-to-the-ca-cert.crt
```

You can also try a combination of the above to resolve the issues:

```text
$ set https_proxy=<PROXY>:PORT
$ snyk auth --insecure
$ snyk test -d --insecure
```

### Also see

[How can we whtelist Snyk IP addresses?](https://support.snyk.io/hc/en-us/articles/360002153077-How-can-we-whitelist-Snyk-IP-addresses-)

