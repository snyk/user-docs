# Step 3: Creating a network for the Broker Client and Code Agent communication

To run the Snyk Broker with the Code Agent, you need to establish a network connection between them. To perform this, you can create a Docker bridge network, as described below.

**Notes**:

* For more information on Docker bridge networks, see the Docker documentation: [https://docs.docker.com/network/bridge/](https://docs.docker.com/network/bridge/)
* To establish this network connection between the Broker Client and the Code Agent, you can also use other methods and tools, like Ngrok.

**To create a Docker bridge network:**

* Run:

```
docker network create <network_name>
```

Where:

* `network_name` – a name for the new network between the Broker Client and Code Agent.



For example:

```
docker network create mySnykBrokerNetwork
```

Where:

* `mySnykBrokerNetwork` – the name of the new network.



**To verify the network creation:**

* Run:

```
docker network ls
```

Your output should look similar to the following:

```
NETWORK ID     NAME                  DRIVER    SCOPE
cb352a803eb0   mySnykBrokerNetwork   bridge    local
```

