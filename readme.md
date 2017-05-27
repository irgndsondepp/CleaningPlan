A simple REST-Webserver to display and update cleaning tasks for flatmates.

The list is saved in an xml file and loaded if a file is present.

After starting the Server (currently listens hardcoded on port 8080), the following commands can be used:

localhost:8080/
--> Whatever comes after the slash is ignored, except if it starts with "done". This outputs the currently assigned Tasks and when they were last completed.

localhost:8080/done/name/taskname
--> Marks the task as done for the given flatmate. If the task is not assigned to the given person, or the person does not exist an error message is displayed.