myweb
=====

A simple web and redis configuration in Docker environment.

Required packages:

```
docker-compose
docker-engine,  v >= 1.10
```

Configure Docker environment is not part of this. Please refer
to https://docs.docker.com/.

Build Docker image:

```
$ make build
```

Start app and redis container:

```
$ make start
```

Usage:

Below returns redis ping request:
```
$ curl <SERVER_IP>:8000/health
```
