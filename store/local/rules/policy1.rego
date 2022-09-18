package example.authz

import future.keywords

default allow := false

allow {
    is_admin
}

is_admin {
    input.username == "admin"
}