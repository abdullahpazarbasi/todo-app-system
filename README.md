# ToDo Application System

## Requirements:
- Docker Engine & Docker CLI
- Open SSL
- Kubernetes
- Helm
- Minikube

## How to Use

To boot-up the system:

```shell
make up
```

> root privileges may be needed for host configurations

To view the orchestration control panel:

```shell
make dash
```

To view Web UI:

[http://www.todo.local/](http://www.todo.local/)

> ```text
> Username: admin
> Password: admin
> ```

## How to Update After Changes

To build up and register up-to-date images:

```shell
make build
```

To undeploy from orchestration cluster:

```shell
make undeploy
```

To deploy into orchestration cluster:

```shell
make deploy
```

## How to Abandon

To shut down the system:

```shell
make down
```

To clean all related things:

```shell
make purge
```
