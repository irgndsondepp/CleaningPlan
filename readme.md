A simple REST-Webserver to display and update cleaning tasks for flatmates.

To build run 'make' or 'make build'.

This App depends on the gorilla/mux package, so you have to run 'go get github.com/gorilla/mux' first.

The list is saved in a json file ('./plan.json') and loaded if a file is present.

After starting the Server (currently listens hardcoded on port 8080), the following commands can be used:

localhost:8080
GET --> If nothing else is specified, this outputs the currently assigned Tasks and when they were last completed in a JSON format.

localhost:8080/tasks
PUT --> Marks the task as done for the given flatmate. If the task is not assigned to the given person, or the person does not exist an error message is displayed.