# go-rest
Basic RESTful API with Go, Gorilla, GORM, MySQL and Docker.

## Installation
1. Install Docker and Docker Compose (https://docs.docker.com/install/linux/docker-ce/ubuntu).
2. Clone Github repository to you machine: `git clone https://github.com/devops787/go-rest.git`
3. Build Docker images: `docker-compose build`
4. Run Docker containers: `docker-compose up -d`

After successful installation process you should get three running containers:
```
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
1acf23029264        devops787/go-rest   "./app"                  16 minutes ago      Up 16 minutes       0.0.0.0:80->3000/tcp                gorest_web_1
562bddcbc302        adminer             "entrypoint.sh docke…"   16 minutes ago      Up 16 minutes       0.0.0.0:8080->8080/tcp              gorest_adminer_1
2477134424bb        mysql               "docker-entrypoint.s…"   16 minutes ago      Up 16 minutes       0.0.0.0:3306->3306/tcp, 33060/tcp   gorest_database_1
```

Containers:
1. `devops787/go-rest` - Go application
2. `mysql` - MySQL database (inside container with `db_data` volume)
3. `adminer` - MySQL administration panel

The ports are mapped to localhost machine, so you can access them via 127.0.0.1 address. To see changes in database you can use admin panel which is available
in browser `http://127.0.0.1:8080`.

## API Endpoints
```
/users - GET
/users - POST
/users/{id} - GET
/users/{id} - PUT
/users/{id} - DELETE
```

## Examples
1. Create new user:
```
$ curl -X POST -d 'name=John&email=john@mail.com' http://127.0.0.1/users
{"id":5,"name":"John","email":"john@mail.com"}
```
2. Get user with id:
```
$ curl -X GET http://127.0.0.1/users/5
{"id":5,"name":"John","email":"john@mail.com"}
```
3. Update user with id:
```
$ curl -X PUT -d 'name=John&email=john2@mail.com' http://127.0.0.1/users/5
{"id":5,"name":"John","email":"john2@mail.com"}
```
4. Delete user with id:
```
$ curl -X DELETE http://127.0.0.1/users/5
{"id":5,"name":"John","email":"john2@mail.com"}
```
5. Get all users:
```
curl -X GET http://127.0.0.1/users
[{"id":2,"name":"John","email":"john@mail3.com"},{"id":3,"name":"John","email":"john@mail3.com"},{"id":4,"name":"John","email":"john@mail3.com"}]
```