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

    The above command runs builder with configuration that is used to build builder itself.
    You must adjust the configuration to suit your project's needs.

3. Trigger a build.
Find out what is the id of your configured project (the id is computed from the URL of your project configuration file):
 ```shell
curl http://localhost:8080/v1/projects
```

    And trigger a build
 ```shell
curl -d '' http://localhost:8080/v1/projects/<projectid>/trigger
```

    Now you should be good to go!

## Configuration

Getting there...

### Builder configuration
...

### Project configuration
...

## The build script

Before the build starts, the user defined build script is copied into build working directory as file `script`. That script is then started and it defines the build process.

The `builder` executable is guaranteed to be in path for the script, so that can be used to communicate with builder. E.g. a new build stage is added by `echo my data | builder add-stage my-stage`

Build finishes when the script finishes. If the script exits with status code 0, the build is considered a success. Otherwise its considered a failure.

### Environment
`BUILDER_URL` is the URL of builder server. This is used by builder client transparently.

`BUILDER_BUILD_ID` is the id of current build.

`BUILDER_TOKEN` may be set and it can be used to access privileged operations such as server shutdown.

`PATH` is a copy of `PATH` for builder server, prepended with the directory of builder executable. This way build script can just run `build add-stage my-stage` and so on to communicate with the server.

## TODO ideas
- websocket
- CLI
- Better logging
- stage data
- stage subtypes
- Confirm thread safety
- support project configuration in builder.hcl
- Configuration refresh during trigger
- Notifications:
    - push notifications about build stage changes to other agents
    - github state api
    - http
    - something else?
- RESTify API. IDs => URLs, etc.
- Execution managers
    - executor pools
    - build queues.
    - environment creation here instead of app
- Query executors
- ability to build without version-data
- Tokens
    - build token
- ProjectID instead of string, etc.

- Organization manager
    - provides a web interface for creating organisations and running builder for each org

- API endpoint for stats
    - configurable?

- Basic web UI
- Web UI plugins for displaying standard build stages such as go coverprofile
- Web UI plugins for displaying repository related info

- => docker
- => nomad-dev-setup
