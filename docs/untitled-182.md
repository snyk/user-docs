# Failing with exit code : Generic Error

##  Failing with exit code : Generic Error

On typical Unix and Linux systems, programs would be able to pass a value to their parent processes while terminating. Values like these are referred to as Exit codes

As part of Snyk output you must have seen Snyk process terminating with exit codes, we typically see

 **Exit code 0** This means Snyk did not find vulnerabilities in your code an exited the process without failing the job.

 **Exit code 1** This means Snyk found vulnerabilities in your code and have failed the build

**Exit code 2** This means Snyk exited with an error, please re-run with \`-d\` to see further information.

**Exit code 3** This means Snyk did not detect any supported projects/manifests to scan. Re-check the command or if the command should run in a different directory.

You can read more about exit codes [here](https://bencane.com/2014/09/02/understanding-exit-codes-and-how-to-use-them-in-bash-scripts/).

