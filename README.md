# Builder

Builder is a continuous integration tool with following design goals
- Text only configuration
- API / CLI as primary user interfaces
- Minimal functionality with maximal integrations to existing software
- Hackable, extensible web UI
- Seriously, executing a build should be nothing more than running a shell script in a controlled environment.

## Getting started

1. Get builder. If you are on AMD64/Linux, you can just grab the latest binary:
```shell
curl -O https://s3.eu-central-1.amazonaws.com/juhovuori/builder/builder
```

    If not, you must build it yourself. The build process is a tiny bit
    more complex than for a casual go project.
```shell
go get -d github.com/juhovuori/builder
cd $GOPATH/src/github.com/juhovuori/builder
make build
```

    Currently, builder only supports Unix-like operating systems.

2. Run builder server
```shell
./builder server -f https://raw.githubusercontent.com/juhovuori/builder/master/builder.hcl
```

    The above command runs builder with configuration that is used to build builder itself. You must adjust the configuration to suit your project's needs.

3. Trigger a build.
Find out what is the id of your configured project (the id is computed from the URL of your project configuration file):
```shell
curl http://localhost:8080/v1/projects
```

    And trigger a build
```shell
curl -d '' http://localhost:8080/v1/projects/<projectid>/trigger
```

## Configuration

Getting there...

### Builder configuration
...

### Project configuration
...
## TODO ideas
- websocket
- CLI
- Tokens
    - system token
        - Graceful shutdown
    - build token
- Better logging
- Migrations
- Notifications:
    - push notifications about build stage changes to other agents
    - http
    - something else?
- RESTify API. IDs => URLs, etc.
- Execution managers
    - executor pools
    - build queues.
    - environment creation here instead of app
- Query executors
- Configuration refresh during trigger
- ProjectID instead of string, etc.
- Confirm thread safety -- an issue with pipeOutput

- Organization manager
    - provides a web interface for creating organisations and running builder for each org

- API endpoint for stats
    - configurable?

- Basic web UI
- Web UI plugins for displaying standard build stages such as go coverprofile

- packer
- => prometheus
- => grafana
- => docker
- => nomad-dev-setup
