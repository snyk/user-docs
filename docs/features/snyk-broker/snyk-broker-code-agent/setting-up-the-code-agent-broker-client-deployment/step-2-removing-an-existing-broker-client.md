# Step 2: Removing an Existing Broker Client

**Note**: This section is applicable only if you already have a running Broker Client. If you do not have a Broker Client yet, move to the Step 3 - _Creating a network for the Broker Client and Code Agent communication_, page 26.

If you have a running Broker Client for the same Organization and the same Integration, you need to stop and remove it before you set up a new Broker Client for the Code Agent. Since the Broker Client needs to communicate with the Code Agent, and this communication is configured in the setup commands of the Broker Client container, you cannot use an existing Broker Client container for the Code Agent operation. &#x20;

**Note**: The new Broker Client that you will set up for the Code Agent, can also serve the Broker deployment for Snyk Open Source and Snyk IaC for the same Snyk Organization and the same SCM. For Snyk Container, you need to set up an additional Broker Client.&#x20;

To find out if you have a running Broker Client, use the Docker Process Status command, which provides a list of all the containers that are up and running on a machine. If you have a running Broker Client, this command will also provide you with the details you need in order to remove it. If you find a running Broker Client, you only need to remove it if it is configured for the same Snyk Organization and the same SCM as the Code Agent.

**To find out if you have a running Broker Client:**

* Run:

```
docker ps
```

If you have a running Broker Client, your output should look similar to the following:

```
CONTAINER ID   IMAGE                    COMMAND                  CREATED        STATUS        PORTS                    NAMES
6eab097879cc   snyk/broker:github-com   "broker --verbose"       18 hours ago   Up 18 hours   0.0.0.0:8001->8000/tcp   suspicious_banzai
```



**To remove the Broker Client container:**

* Run:

```
docker rm -f <Broker_Client_container_ID_or_name)
```

Where:

* `container_ID_or_name` – the ID or name of the container of the running Broker Client. You can obtain these details by using the `docker ps` command.&#x20;

For example:

<figure><img src="../../../../.gitbook/assets/Broker Client - removing.png" alt=""><figcaption></figcaption></figure>

#### &#x20;
