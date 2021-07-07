# How are patches validated?

Each patch is custom made for a small and specific version range. We test every single minor version in that range during each CI build we have. A good example is the npm package qs, which has multiple patches for [different version ranges](https://snyk.io/vuln/npm:qs:20170213).

When we create a patch, our research team manually validates the functional correctness of the patch, and tests that it doesn't break the original testing suite that comes with the package.

