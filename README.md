# Builder


## Configuration
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
