# Step 1: Obtaining the required Tokens for the setup procedure

To set up the Code Agent and the Client Broker components, you need the following tokens:

* **Snyk API token** - this token is **required for the Code Agent component setup**. It is used in the `-e SNYK_TOKEN` parameter to authenticate the Code Agent component with your Snyk Account. See [Obtaining your Snyk API token](obtaining-your-snyk-api-token.md).
* **Broker token** - this token is **required for the Broker Client setup**. It is used in the `-e BROKER_TOKEN` parameter, to enable the Broker deployment for a specific Organization and a specific SCM. See [Obtaining your Broker token](obtaining-your-broker-token.md).
* **Integrated SCM token** - this token is required for the Broker Client setup. It is used in the `-e <SCM>_TOKEN` parameter, to enable access with certain permissions to the integrated SCM. See [Obtaining your SCM token](obtaining-your-scm-token.md).

After you have obtained the required tokens, save them in a safe and accessible place. When you start setting up the Code Agent and the Client Broker components, you will need to use these tokens.
