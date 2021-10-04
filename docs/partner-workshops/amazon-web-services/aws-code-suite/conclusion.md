---
description: >-
  Congratulations. You have completed this DevSecOps workshop on shifting
  security testing left.
---

# Conclusion

## Recap on what you have learned

* Learned how to deploy CloudFormation stacks
* Learned about a modern CI/CD pipeline 
* Learned how to test in AWS CodeBuild
* Learned about a couple of open source tools for security testing
* Learned why it is important to test as early as possible in the pipeline

## Final Thoughts

Due to the time and scope of the workshop there are several things that can and should be instrumented to improve security testing within the CI/CD pipeline. Here are a few suggestions to go above and beyond what you have learned in this workshop.

* Add notifications that provide feedback to developers using technology that developers are already familiar with and using.  For build failures send to a slack channel, SMS, or email notifications using Amazon Simple Notification Service \(SNS\) and AWS Lambda.
* Use a branching method such as gitflow and test branches awaiting a pull request review.
* Utilize blue/green deployments to instrument additional security testing prior to production deployment.
* Enable git hooks to automate testing right when a developer commits code on her/his local machine. 
* Add additional testing such as language specific linters, SAST, DAST, dependency CVE scanning, IAST, and RASP.  Implementation should be similiar to what we accomplished in this workshop.

The sky's the limit on adding additional features and functionality to DevSecOps. The point is to monitor your pipeline and continually make improvements to accelerate the release of features and functionality to your end customers.

## Next Steps

Try and implement some of the learnings from the workshop on your company's development process. Don't try and do too much at one time and use an agile iterative approach. Remember that you are also trying to change culture by baking in security so don't try and do too much too fast.

