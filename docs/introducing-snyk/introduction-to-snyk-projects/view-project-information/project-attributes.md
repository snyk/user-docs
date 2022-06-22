# Project attributes

Project attributes are static and non-configurable fields which allow you to add additional metadata (also known as “values”) to a project. Attributes have a predefined list of values you can select from. Values can be assigned and removed from attributes, and can be used to group, prioritize, and filter projects.

## **Available attributes and their values**

### **Lifecycle stage**

* Production
* Development
* Sandbox

### **Business criticality**

* Critical
* High
* Medium
* Low

### **Environment**

* Frontend
* Backend
* Internal
* External
* Mobile
* SaaS
* Hosted
* Distributed

The ability to assign and remove values from attributes, is available in the [API](https://snyk.docs.apiary.io/#reference/projects/project-attributes) and the [CLI](../../../getting-started/introduction-to-snyk-projects/view-project-information/broken-reference/). The ability to filter projects by their values is also available in the [API](https://snyk.docs.apiary.io/#reference/projects/project-attributes).

## **Assigning values to attributes**

1. On the project page, click the “+” icon below the attribute you want to assign the value
2. Select a value from the list available

![](../../../.gitbook/assets/gs1.png)

After you have assigned a value to the attribute, it appears on the project listing page. Each attribute can have multiple values assigned to it, and you can assign values to multiple attributes.

![](../../../.gitbook/assets/gs2.png)

## **Removing values from attributes**

1. Select the attribute you want to remove a value from.
2. Click the “x” icon on the value.

![](../../../.gitbook/assets/gs3.png)

The value has been removed from the attribute.

![](../../../.gitbook/assets/gs4.png)

## **Filtering values in the project listing page**

1. On the left of the project listing page, select the values for the attributes that you want to filter projects by.
2. When you filter by multiple values on a single attribute, you return projects that have been assigned one or more of the values in the filter.
3. When you filter by multiple attributes, you return projects which have been assigned values of both attributes in the filter.

![](../../../.gitbook/assets/gs5.png)
