# Server bind address
bind_addr = "0.0.0.0:8080"

# List of projects to manage
projects = [
  "testdata/success.hcl",
  "testdata/failure.hcl",
  "testdata/delay100s.hcl"
]

# State store configuration
state_store {
  type = "file"
  directory = "/tmp/builder"
}
