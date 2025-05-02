# Setup requirements for AWS CodePipeline

{% hint style="warning" %}
**The Snyk integration for AWS CodePipeline will be discontinued**

\
**Action Required**

In order to safeguard the security of our services and our customers, Snyk has begun the deprecation of its integration with **AWS CodePipeline**. To minimize disruption, we recommend that you transition to using **AWS CodeBuild** and the Snyk CLI as an alternative which will support the same use case and functionality.&#x20;



**Migration Timeline**

Effective **Oct 30th, 2024**, you will no longer be able to add or modify the Snyk plug-in for new or existing pipelines. Existing pipelines will continue to work as-is for 7 months, though we recommend migrating to the new process as soon as possible. To avoid disrupting your CI/CD workflows, you must transition to the Snyk CLI before **May 30, 2025**. Please refer to the steps in this [migration guide](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/migrating-to-aws-codebuild) to use Snyk CLI with AWS CodeBuild.\


We are confident that AWS CodeBuild and the Snyk CLI will meet your requirements.&#x20;
{% endhint %}

Check if your project must be built before the scan in the CodePipeline. If the project needs to be built, you must add a CodeBuild step before the Snyk Step.

|      Language     | Project Type | Build Required |                                            Notes                                           |
| :---------------: | :----------: | -------------- | :----------------------------------------------------------------------------------------: |
|     Javascript    |      npm     | No\*           |   Build only required if no `package-lock.json` file present; run npm install to generate  |
|     Javascript    |     Yarn     | No\*           |      Build only required if no `yarn.lock` file present; run yarn install to generate      |
|        Java       |     Maven    | Yes            |                              Run `mvn install` before testing                              |
|        Java       |    Gradle    | No             |                                                                                            |
|        .NET       |     Nuget    | No\*           |                  Build only required if no `packages.config` file present                  |
|       Python      |      Pip     | No\*           |     Build only required if missing a Snyk config file with the language-settings param     |
|       Python      |   Setup.py   | Yes            |                            Run `pip install -e .` before testing                           |
|       Python      |    Poetry    | No\*           |     Build only required if no `poetry.lock` file present; run `poetry lock` to generate    |
|        Ruby       |    Bundler   | No\*           |   Build only required if no `Gemfile.lock` file present; run `bundle install` to generate  |
|        PHP        |   Composer   | No\*           | Build only required if no `composer.lock` file present; run `composer install` to generate |
|       Scala       |      SBT     | No             |                                                                                            |
|         Go        |  Go Modules  | No             |                                                                                            |
| Swift/Objective-C |   Cocoapods  | No\*           |     Build only required if no `Podfile.lock` file present; run pod install to generate     |
