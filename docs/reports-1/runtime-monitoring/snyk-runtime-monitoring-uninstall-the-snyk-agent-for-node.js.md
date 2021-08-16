# Snyk runtime monitoring: uninstall the Snyk agent for Node.js

1. From the project CLI, remove the agent as a dependency to your project by running the following command

   `npm uninstall @snyk/nodejs-runtime-agent`

2. Remove the `Require` code snippet that you inserted at installation \(and upgrade\).
3. Commit and push the changes to your manifest file \(for example `package.json`\).

