# Mounting secrets with Docker

Sometimes it is required to load sensitive configurations, the GitHub or Snyk token, from a file instead of from environment variables. Broker is using [dotenv](https://www.npmjs.com/package/dotenv) to load the config, so the process is relatively simple:

* Create a file named `.env` and put your sensitive configuration there.
* Mount this file, for example, using a [Kubernetes secret](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume)). Mount the file to be somewhere like `/broker`.
* Change the workdir of the docker image to be `/broker/`. An example of such file is located in your broker container at $HOME/.env.
