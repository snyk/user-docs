# Best Practices

**Naming Custom Rules:** Make sure to name custom rules in a way that reflects that they are custom rules and what the query results in. It is also advisable to have an author, last changed date, and who to contact for questions in the description.

**Rules are code:** Custom Queries and Custom Rules are code in that you can copy-paste, export, or import them, and store them in any version management system. &#x20;

**Test your query before making it a rule:** A query works at a point in time (whenever you execute it), but a rule needs to work for all future. To ensure future-proof rules, it is a good idea to test it against foreign code like open source projects. This hardens the query toward different coding styles or libraries used. You will be thankful you did it when your business acquired another one, and you have to merge the codebases.

**Sometimes the full qualified name of objects is hidden in syntactical sugar:** Either use autocomplete to find the name of an object usable in queries or use a regex that covers pre and postfixes (e.g. `~”.*[object_name].*”`).

\
