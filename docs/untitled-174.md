# Private Gem support for Ruby projects

##  Private Gem support for Ruby projects

Snyk does support Gems hosted in private GitHub repositories by default.

You can contact [support@snyk.io](mailto:support@snyk.io) to have this feature enabled for your organization.

After this feature has been enabled you can define Bundler environment variables in your organization or project settings.  


Note that this will only work for gems over HTTPS such as:

```text
gem "privvy", git: "<https://github.com/user/ruby-gem-for-private-source>"
```

