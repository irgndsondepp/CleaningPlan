A simple REST-Webserver to display and update cleaning tasks for flatmates.

The list is saved in a json file ('./plan.json') and loaded if a file is present.

After starting the Server (currently listens hardcoded on port 8080), the following commands can be used:

localhost:8080/
--> If nothing else is specified, this outputs the currently assigned Tasks and when they were last completed in a simple string format.

localhost:8080/done/name/taskname
--> Marks the task as done for the given flatmate. If the task is not assigned to the given person, or the person does not exist an error message is displayed.

localhost:8080/xml/
--> Returns the current state in xml formatting

localhost:8080/json/
--> Returns the current state in json formatting