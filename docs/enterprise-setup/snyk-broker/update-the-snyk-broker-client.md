# Update the Snyk Broker client

Snyk regularly releases updated versions of the Broker client in order to provide new features, bug fixes, and more. Snyk recommends that you install the updated version of your Broker client regularly to take advantage of bug fixes and new features.

The full list of versions with release notes is available on [Snyk Broker GitHub](https://github.com/snyk/broker/releases). Snyk encourages you to [subscribe to the RSS feed](https://github.com/snyk/broker/releases.atom) for that page to receive information about versions as they are released. The Docker images are released more slowly than the GitHub releases. The `docker pull` command always pulls the most recent image.

Update the Broker client in the same way you installed it initially. Stopping and starting your Broker by itself will not update the version of Broker client that is used

## Updating using the Docker method

Run `docker pull snyk/broker:<brokerClienttype>` to pull the latest version of that tag, then stop and start your Broker client or clients.&#x20;

## Updating using the Helm method

The specifics will be determined by how you initially installed and ran your Snyk Broker, but running a Helm upgrade\` instead of Helm install\` with the same arguments will pull the latest chart, unless you target a specific chart version. Ensure that you still supply the correct values for your Broker, either by using the `--set <variable>:<value>` commands, or by using a YAML file.&#x20;

The Helm Chart targets the latest version of the applicable Broker client and has `imagePullPolicy: Always`, so upgrading to the latest chart will reinitialize the pod and pull the latest version of the client.
