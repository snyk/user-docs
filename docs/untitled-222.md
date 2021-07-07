# Unable to add Webhook to Gitlab brokered instance

##  Unable to add Webhook to Gitlab brokered instance

**Problem**:

When attempting to add a Webhook to a Gitlab project while using the Snyk broker, you may encounter an error.

In addition to this, in the Snyk logs we will see an 

```text
Invalid url given
```

error response from Gitlab

**Discussion**:

To prevent this type of exploitation from happening, starting with GitLab 10.6, all Webhook requests to the current GitLab instance server address and/or in a private network will be forbidden by default. That means that all requests made to 127.0.0.1, ::1 and 0.0.0.0, as well as IPv4 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16 and IPv6 site-local \(ffc0::/10\) addresses won't be allowed.

**Resolution**:

Add your broker IP to the list of whitelisted addresses in Gitlab to resolve this problem

