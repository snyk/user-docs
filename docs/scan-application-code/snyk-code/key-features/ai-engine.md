# Snyk Code AI Engine

Snyk Code is powered by a semantic, AI-based analysis Engine, and it can analyze the following in your code:

### **Hardcoded secrets**

Snyk Code includes secret detection capabilities that scan and highlight secrets like keys, credentials, and sensitive information in your source code. Unlike tools that use entropy checks or regular expressions, Snyk Code uses machine learning to improve the accuracy of detecting secrets and minimizing the occurrence of false positives.

![](<../../../.gitbook/assets/Introduction - AI Engine - Hardcoded secrets.png>)

### **Coding issues**

Problems such as dead code, branches that are predefined, and branches having the same code on each side.

### **Type inference**

Determining the initial type and its changes - this is of special interest for dynamically typed languages.

### **Value ranges**

Infers possible values for variables used to call functions to track off-by-one errors in arrays, division-by-zero, and null dereferences.

### **Data flow**

Follows the flow of data within the application from the Source to the Sink. Combined with AI-based learning of external insecure data source, data Sinks, and sanitation functions, this enables a strong taint analysis.

### **API usage**

Identifies multiple potential issues including API misuses, null dereferences, and type mismatches by modeling the usage of memory in variables and references. This mechanism can also identify use of insecure functions.

### **Control flow**

Identifies null dereference or race conditions by modeling each possible control flow in the application.

### **Point-to analysis**

Identifies multiple potential issues including buffer overruns, null dereferences, and type mismatches by modeling the usage of memory in variables and references.
