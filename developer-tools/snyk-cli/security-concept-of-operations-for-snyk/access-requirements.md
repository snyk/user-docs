# Access requirements

When you are using Snyk applications like the [CLI ](../getting-started-with-the-snyk-cli.md)or [IDE integrations](../../snyk-ide-plugins-and-extensions/), certain local and remote resources must be accessible. This documentation explains how to harden your system without affecting Snyk functionality.

## Local filesystem

The required filesystem access may vary by product.

* CACHE\_PATH (Read,Write,Execute)
  * Windows: `C:\\Users\\<USERNAME>\\AppData\\Local\\snyk`
  * Linux/Alpine: `/home/<USERNAME>/.cache/snyk`
  * macOS: `/Users/<USERNAME>/Library/Caches/snyk`
* CONFIG\_PATH (Read,Write)
  * Windows: `%USERPROFILE%\\.config\\configstore`
  * Linux: `$HOME/.config/configstore`
  * macOs: `$HOME/.config/configstore`

## Network resources

### Required

* api.\<SNYK\_INSTANCE>.io:443
* app.\<SNYK\_INSTANCE>.io:443
* 127.0.0.1:\<EPHEMERAL\_PORT\_RANGE>
  * required for Inter Process Communication
  * [Ephemeral Port Range](https://www.rfc-editor.org/rfc/rfc6335.html#section-6) might vary by Operating System and Settings

### Optional

* deeproxy.snyk.io:443 (for Snyk Code)
* downloads.snyk.io:443 (depending on features used, such as downloading the CLI)
* learn.snyk.io:443 (to be able to display Snyk Learn links in issue details)
* static.snyk.io:443 (depending on features used, such as downloading the CLI)
* snyk.io:443 (depending on features used)
* \*.sentry.io:443 (error reporting)
