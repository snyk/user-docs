# Best practices for Snyk Code custom rules

## **Naming custom rules**

Make sure to name custom rules in a way that reflects that they are custom rules and what the query results in. It is also advisable to have an author, last changed date, and who to contact for questions in the description.

## **Rules are code**

Custom queries and rules are code, meaning you can copy-paste, export, or import them and store them in any version management system. &#x20;

## **Test your query before making it a rule**

A query works at a point in time whenever you execute it, whereas a rule needs to work for all future cases. To ensure future-proof rules, testing them against foreign code like open source Projects is a good idea. This hardens the query toward different coding styles or libraries used. This could be useful if your business acquired another one and you must merge the codebases.

## **Sometimes the full qualified name of objects is hidden in syntactical sugar**

Either use autocomplete to find the name of an object usable in queries or use a regex that covers pre and postfixes (`~”.*[object_name].*”`).

\
