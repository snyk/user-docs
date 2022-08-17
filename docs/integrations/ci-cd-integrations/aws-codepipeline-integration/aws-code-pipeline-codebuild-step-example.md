# AWS Code Pipeline CodeBuild step example

Note the Scan's input artifact must be provided with the build's output artifact as shown in the configuration.

Example of Javascript CodeBuild (`buildspec.yml`):

```
version: 0.2
phases:
  build:
    commands:
      - npm install
artifacts:
  files:
    - '**/*'
```

Example of Maven CodeBuild (`buildspec.yml`):

```
version: 0.2
phases:
  build:
    commands:
      - mvn install
artifacts:
  files:
    - '**/*'
```

### S
