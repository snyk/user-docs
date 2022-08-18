# Setup requirements for AWS CodePipeline

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
