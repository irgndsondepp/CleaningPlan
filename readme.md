# A simple REST-Webserver to host a cleaning plan

## What is this?

This application starts a webservice with a REST interface to manage a simple rotating task list. The idea behind this is to replace the oldfashioned blackboard or whatever you use and still be able to blame all your flatmates for not taking care of their tasks.
You of course will know how to manipulate everything in your favor ;-).

## How to build

To build run one of these targets
```
make
make build
```

The default make target points to `make build`.

## Dependencies

This App depends on the gorilla/mux package, so you have to 

```
go get github.com/gorilla/mux
```

first.

## How does this work?

### Persistence

If a 'plan.json' file is present in your working directory, the application will try to load a saved plan from there. If this fails a default plan is created and saved to your working directory.

### REST

After starting the Server (currently listens hardcoded on port 8080), the following commands can be used:

#### GET

```
localhost:8080
localhost:8080/tasks
```

If nothing else is specified, this outputs the currently assigned Tasks and when they were last completed in a JSON format.

```
localhost:8080/tasks/<assigneename>
```

Filters the list of tasks by assignee name.

All other requests will recieve a 404 error.

#### PUT

```
localhost:8080/tasks/<taskname>/<assignee>
```

Marks the task as done for the given flatmate. If the task is not assigned to the given person, or the person does not exist a http error code is returned.

All other requests return a 404 not found error.
