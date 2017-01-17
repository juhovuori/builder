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

2. Run builder
```shell
./builder -f https://raw.githubusercontent.com/juhovuori/builder/master/builder.hcl
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

## Stale notes

## Configuration
TODO
- Modules
    - File configuration
    - Github configuration
- General configuration
- Projects configuration


## State storage
- Modules
    - File storage
    - PostgreSQL storage


## Operation
- Watch projects based on project configuration
- Web server to receive build stage commands
    - Mutate:
        - Create build
        - Enter stage x
        - Append to stage log x
        - Store stage data x
        - Finish stage
        - Finish build
    - Query:
        - List projects
        - List builds
        - List build stages
        - Show build log / data (streaming if possible)
    - Authentication
        - Github OAUTH2
- CLI for build stage commands
    - Everything that build stage server supports
- Run a build:
    - Takes in a project configuration and runs a build
    - Generate build token and store build start.
    - Acquire executor to run steps and pass token to that
    - Run steps one by one and notify builder while going on
- Web server that implements a UI
    - Browse projects and start builds
    - Dashboards with D3 or similar


## General config


## Notifications
- Push notifications about build stage changes to other agents.
- Modules:
- HTTP notification


## Execution manager
- Maintains a pool of executors to run build stages


## Organization manager
- provides a web interface for creating organisations and running builder for each organisation


## TODO
- packer

- websocket
- CLI
- Graceful shutdown
- Use secret build token to modify build
- Logging
- Migrations
- RESTify API. IDs => URLs, etc.
- Store executors / executor pools / queues. Environment passing  also from these
- Query executors
- Configuration refresh during trigger
- ProjectID instead of string, etc.
- Confirm thread safety -- an issue with pipeOutput

- Basic web UI
- Web UI plugins for displaying standard build stages such as go coverprofile

- => prometheus
- => grafana
- => docker
- => nomad-dev-setup
