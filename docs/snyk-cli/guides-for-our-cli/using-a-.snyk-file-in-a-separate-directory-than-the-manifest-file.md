# Using a .snyk file in a separate directory than the manifest file

When scanning a file within the CLI that utilize a .snyk file, they may not be present within the same directory due to the structure of the project or the project holding multiple manifest files which wish to use the same policy file.

To do this within the CLI, the use of `--policy-path` can be used to specify a policy path manually to be applied with the current project scanned or monitored.

Below is an examples use cases of this.

├── .snyk  
└── another-project  
├── node\_modules  
│ └── ...  
├── package-lock.json  
└── package.json **&lt;----**

Current directory - manifest level

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

