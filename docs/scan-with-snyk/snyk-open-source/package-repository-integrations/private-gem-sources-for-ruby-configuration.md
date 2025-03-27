# Private gem sources for Ruby configuration

{% hint style="info" %}
**Feature availability**\
This functionality is in beta. [Contact Snyk Support](https://support.snyk.io) to have it enabled for your Organization.

This guide is relevant for Snyk UI integrations only. The CLI supports Ruby Projects with private registries without extra configuration.
{% endhint %}

You can add a configuration to tell Snyk where your private gems are hosted. This is the same information you would normally add as a Bundler environment variable.

After you have added this configuration, Snyk uses the information to access private dependencies when creating Pull/Merge Requests, by allowing Bundler to reach those dependencies in order to regenerate the lockfile.

## Configuration of private gem sources for Ruby

1. Navigate to **Settings** > **General**.
2. Find the `RubyGems Bundler environment variables` section; see the screen illustration.
3. Add environment variable names and values to define credentials for gem sources.\
   These are generally the same as the values you set on your developer machine, in your CI environments, or both.\
   Example name: `BUNDLE_GITHUB__COM`, Value: `abcd0123generatedtoken:x-oauth-basic`
4. To test the configuration, open a Pull/Merge Request on a Project that contains gems from your private registries to see a lockfile updated and included in the Snyk Fix Pull Request.

<figure><img src="../../../.gitbook/assets/94445628-8fdd3980-019f-11eb-816e-2c61c5b99c5c.png" alt="RubyGems Bundler environment variables"><figcaption><p>RubyGems Bundler environment variables</p></figcaption></figure>

## Requirements for configuration of private gem sources for Ruby

A list of requirements follows:

* Variable values must be CGI escaped.
* Gem sources must use `https` URLs.\
  Example: Supported: `gem "privvy", git: "https://github.com/testexample/ruby-gem-for-private-source"`\
  Not supporte&#x64;**:** `gem "privvy", git: "git@github.com:testexample/ruby-gem-for-private-source"`
* Gem sources must be publicly resolvable, that is, not behind a firewall.
* Variables must be configured according to the [Bundler Credentials for Gem Sources documentation](https://bundler.io/v1.16/bundle_config.html#CREDENTIALS-FOR-GEM-SOURCES).
