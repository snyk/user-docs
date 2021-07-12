# On-Prem Services: CSV export download gives a generic error

## On-Prem Services: CSV export download gives a generic error

In Snyk **OnPrem services**, if a **CSV export** does not download and gives you a generic error\(for eg a 500 response code\) it's highly likely a firewall or proxy issue.

The service handling this \(conveniently called **csv-service**\) actually calls the instance APIs to get the results in JSON and transform them into CSV. If the instance/service can't call itself, it will fail.

Sometimes firewalls on VMs are configured not to allow to hit its own IP address, or more often the call goes out of the perimeter via a proxy, hence the need to add the instance hostname to the **no\_proxy list**.

The `*.mycompany.com` is subdomain only. Subdomains need to be covered by `*.whateversubdomain.mycompany.com`

