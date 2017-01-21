# Server bind address
bind_addr = "0.0.0.0:8080"

# Server URL - passed to builds t
url = "http://builder.juhovuori.net"

# List of projects to manage
projects = [
  {
    type = "git"
    repository = "https://github.com/juhovuori/builder.git"
    config = "project.hcl"
  }
]

store = "sqlite:/tmp/builder-prod.db"
