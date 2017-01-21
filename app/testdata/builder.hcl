# Server bind address
bind_addr = "0.0.0.0:8080"

# List of projects to manage
projects = [
    {
        type = "git"
        url = "http://example.com"
        config = "project.hcl"
    }
]

store = "memory"
