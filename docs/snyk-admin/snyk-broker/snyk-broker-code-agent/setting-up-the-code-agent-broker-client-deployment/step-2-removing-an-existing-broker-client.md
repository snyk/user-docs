# Step 2: Removing an existing Broker Client

{% hint style="info" %}
This section applies only if you already have a running Broker Client. If you do not have a Broker Client yet, proceed to [Step 3 - Creating a network for the Broker Client and Code Agent communication](step-3-creating-a-network-for-the-broker-client-and-code-agent-communication.md).
{% endhint %}

If you have a running Broker Client for the same Organization and the same Integration, stop and remove it before you set up a new Broker Client for the Code Agent. Since the Broker Client needs to communicate with the Code Agent, and this communication is configured in the setup commands of the Broker Client container, you cannot use an existing Broker Client container for the Code Agent operation.

{% hint style="info" %}
The new Broker Client that you set up for the Code Agent can also serve the Broker deployment for Snyk Open Source and Snyk IaC for the same Snyk Organization and the same SCM. For [Snyk Container](../../snyk-broker-container-registry-agent/), you must set up an additional Broker Client.
{% endhint %}

**To find out if you have a running Broker Client**, use the Docker Process Status command, which provides a list of all the containers that are up and running on a machine.

{% hint style="info" %}
If you have a running Broker Client, this command also provides you with the details you need in order to remove the running Broker Client. If you find a running Broker Client, you only need to remove it if it is configured for the same Snyk Organization and the same SCM as the Code Agent.
{% endhint %}

Run the following:

```
docker ps
```

If you have a running Broker Client, your output is similar to the following:

```
CONTAINER ID   IMAGE                    COMMAND                  CREATED        STATUS        PORTS                    NAMES
6eab097879cc   snyk/broker:github-com   "broker --verbose"       18 hours ago   Up 18 hours   0.0.0.0:8001->8000/tcp   suspicious_banzai
```

**To remove the Broker Client container**, run the following:

```
docker rm -f <Broker_Client_container_ID_or_name)
```

where `container_ID_or_name` is the ID or name of the container of the running Broker Client. You can obtain these details by using the `docker ps` command, for example:

<figure><img src="../../../../.gitbook/assets/Broker Client - removing.png" alt="Example of docker ps command"><figcaption><p>Example of docker ps command</p></figcaption></figure>
