# Step 1: Obtaining the required Tokens for the setup procedure

To set up the Code Agent and the Client Broker components, you need 3 tokens:

**For the Code Agent**:

* **Snyk API token** - this token is required for the Code Agent component setup. It is used in the `-e SNYK_TOKEN` parameter, to authenticate the Code Agent component with your Snyk Account. See [Obtaining your Snyk API token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-snyk-api-token).

**For the Broker Client**:

* **Broker token** - this token is required for the Broker Client setup. It is used in the `-e BROKER_TOKEN` parameter, to enable the Broker deployment for a specific Organization and a specific SCM. See [Obtaining your Broker token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token).
* **Integrated SCM token** - this token is required for the Broker Client setup. It is used in the `-e <SCM>_TOKEN` parameter, to enable access with certain permissions to the integrated SCM. See [Obtaining your SCM token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token).&#x20;

After you obtained the required tokens, save them in a safe and accessible place. When you will start setting up the Code Agent and the Client Broker components, you will need to use these tokens as part of the setup commands.  &#x20;
