# User Data API

The app fetch the data from postgres database that are build in docker, the container between application and database are connected using docker-compose.

# Run the Application

To run this application, clone the repository and run command:

```bash
$ docker-compose up
```

```bash
$ curl http://localhost:8080/displayallusers
[{"userid":1,"name":"Budi"},{"userid":2,"name":"Nano"}]
```
If docker has not been installed in your computer, you may reffer to this link [Get Docker](https://docs.docker.com/get-docker/).

# Endpoints

There are 2 available endpoints:
- displayallusers (GET)
- displayuser (POST) 