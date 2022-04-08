# Snyk User Documentation

### Install snyk

Run the following command from a local terminal:

{% tabs %}
{% tab title="npm" %}
```
npm install -g snyk
```
{% endtab %}

{% tab title="Linux/Mac" %}
```
brew tap snyk/tap
brew install snyk
```
{% endtab %}

{% tab title="Windows" %}
```
scoop bucket add snyk https://github.com/snyk/scoop-snyk
scoop install snyk
```
{% endtab %}
{% endtabs %}

### **Authenticate your machine**

A Snyk account is required to authenticate, [sign-up for free](https://snyk.io/login?cta=sign-up\&loc=nav\&page=support\_docs\_page). Check our plans page for [free tier plan information](https://snyk.io/plans/).

```
snyk auth
```

### **Analyze and test your dependencies**

Navigate into your code’s directory and run:

```
snyk monitor
```

After scanning your project, you'll be given a URL where you can see the results.

### Next steps

Scan projects, prioritize vulnerabilities, fix issues and get back to development.

### Feedback

See an opportunity for us to improve our documentation? [Open an issue against our User Docs GitHub repo](https://github.com/snyk/user-docs/issues), and we'll be in touch, or [contact Snyk Support](https://support.snyk.io/hc/en-us/requests/new).
