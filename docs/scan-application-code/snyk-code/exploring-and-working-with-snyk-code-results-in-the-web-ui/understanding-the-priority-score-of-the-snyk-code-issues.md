# Understanding the Priority Score of the Snyk Code issues

Snyk Code assigns a Priority Score to each discovered issue that is displayed on the Web UI. This Priority Score can assist you in filtering and prioritizing the issues that are discovered in your source code, according to their level of importance, risk, frequency, and ease of fix:

![](<../../../.gitbook/assets/Snyk Code - Results - Priority Score.png>)

### **Exploring the Priority Score factors**

A Priority Score for each issue can be between 0-1000, and it will automatically change if one of its factors changes. For example, if the Severity Level of an issue has increased or decreased, the Priority Score of the issue would change accordingly.

The calculation of the Priority Score is based on the following factors:

* **Severity Level**: the higher the severity level, the higher is the security risk of the issue. Therefore, each severity level adds a different Score to the issue. The Score can be at the most 500 points.\
  **Note**: Snyk Code currently does not use the **Critical** severity level.
* **Availability of a Fix Example for the issue**: when an issue has a Fix Example, it is easier to fix, and therefore it has a higher Priority Score. The Score can be at the most 200 points.\
  **Note**: When a Fix Example is available, it is displayed in the **Full Details** pane of the issue > **Fix analysis** tab.
* **Issue occurrence in a Project**: the number of times a specific issue appears in the **Code Analysis** Project. The higher the number of the issue occurrences in the Project, the higher is the risk, and therefore the Score. The Score can be at the most 100 points.
* **Issue occurrence in a File:** the number of times an issue appears in a specific file. The higher the number of the issue occurrences in a file, the higher is the risk and the Score. The Score can be at the most 100 points.
* **Community projects** – the number of times an issue was fixed in external open source projects that were examined by Snyk. The Score can be at the most 100 points.
* **Internal tags:** when an issue has an internal tag, this tag decreases the Priority Score by 100 points. This internal tag can be one of the following:
  * Test - the issue was found in a test file.
  * Beta - the vulnerability type of the issue is in Beta status.\
    **Note**: These internal tags are automatically assigned by Snyk Code analysis, and they are not visible on the Web UI.

### **Filtering discovered issues according to their Priority Score**

You can filter the discovered issues according to their Priority Score, thus displaying only the issues that are between a certain Priority Score range.

**To filter issues according to their Priority Score:**

* On the Web UI > **Code Analysis** page > Filter pane, locate the **PRIORITY SCORE** slider. Then, drag one or two of the handles to set the range of the Priority Score of issues that you want to display:

![](<../../../.gitbook/assets/Snyk Code - Results - Priority Score - Filtering.png>)

The issue display changes automatically according to your selection, and only issues that have a Priority Score in the selected range are displayed.
