# Server bind address
bind_addr = "0.0.0.0:8080"

# List of projects to manage
projects = [
  "project.hcl"
]

# State store configuration
state_store {
  type = "file"
  directory = "/tmp/builder"
}
