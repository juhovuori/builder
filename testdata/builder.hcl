# Server bind address
bind_addr = "0.0.0.0:8080"

url = "http://localhost:8080"

# List of projects to manage
projects = [
  {
    type = "git"
    repository = "testdata/repository.git",
    config = "success.hcl"
  },
  {
    type = "git"
    repository = "testdata/repository.git",
    config = "failure.hcl"
  },
  {
    type = "git"
    repository = "testdata/repository.git",
    config = "delay100s.hcl"
  },
  {
    type = "git"
    repository = "testdata/repository.git",
    config = "stages.hcl"
  },
  {
    type = "git"
    repository = "testdata/repository.git",
    config = "output.hcl"
  }
]

store = "sqlite:/tmp/builder-prod.db"
