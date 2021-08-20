# Continuous Integration: language support

Configure the CI for different languages as follows.

####  For Node.js

1. Install the Snyk utility using `npm install -g snyk`.
2. Run \#`snyk wizard` in the directory of your project following the prompts which will also generate a `.snyk` policy file.
3. Ensure the `.snyk` file you generated was added to your source control \(`git add .snyk`\).
4. If you selected to, Snyk will include `snyk test` as part of your `npm test` command, so if there are new vulnerabilities in the future, your CI will fail, protecting you from introducing vulnerabilities to production. Alternatively, you can add `snyk test` to any other CI test platform you use.
5. Configure your CI environment to include the `SNYK_TOKEN` environment variable. You can find your API token in your [account settings on snyk.io](https://app.snyk.io/account/)

####  For Ruby, Python, Java, Go, and .NET

1. Install the Snyk utility using `npm install -g snyk`.
2. Add `snyk test` to your CI test platform.
3. Configure your environment to include the `SNYK_TOKEN` environment variable. You can find your API token in your account settings on snyk.io.

####  For Scala

1. Install the Snyk utility using `npm install -g snyk`.
2. Install the [sbt-dependency-graph plugin](https://github.com/jrudolph/sbt-dependency-graph).
3. Add `snyk test` to your CI test platform.

