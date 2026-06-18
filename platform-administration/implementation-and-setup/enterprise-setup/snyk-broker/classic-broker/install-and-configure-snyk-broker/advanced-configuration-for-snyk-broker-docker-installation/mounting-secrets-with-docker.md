# Mounting secrets with Docker

Sometimes you must load sensitive configurations, such as the GitHub or Snyk token, from a file instead of from environment variables. Broker uses [dotenv](https://www.npmjs.com/package/dotenv) to load the config, so the process is straightforward:

* Create a file named `.env` and put your sensitive configuration there.
* Mount this file, for example, using a [Kubernetes secret](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume). Mount the file to somewhere like `/broker`.
* Change the workdir of the Docker image to `/broker/`. An example of such a file is located in your Broker container at $HOME/.env.
