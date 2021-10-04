# Application scanning

Our next job begins by attaching the workspace and calling `snyk/scan` with a few parameters. We've made a few choices here such as setting `fail-on-issues` to `false` and setting our `severity-threshold` to `high`. 

```yaml
  scan_app:
    <<: *defaults
    steps:
      - attach_workspace:
          at: ~/repo
      - snyk/scan:
          fail-on-issues: false
          monitor-on-build: true
          project: '${CIRCLE_PROJECT_REPONAME}/${CIRCLE_BRANCH}-app'
          severity-threshold: high
          token-variable: SNYK_TOKEN
          target-file: ./submodules/goof/package.json
```

{% hint style="info" %}
A detailed list of all supported parameters is available in the [Snyk orb documentation](https://circleci.com/orbs/registry/orb/snyk/snyk#commands-scan) page.
{% endhint %}

