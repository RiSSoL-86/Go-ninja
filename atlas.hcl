variable "database_url" {
  type    = string
  default = getenv("DATABASE_URL")
}

env "local" {
  migration {
    dir = "file://migrations"
  }

  url = replace(var.database_url, "localhost", "host.docker.internal")
  src = "file://schema.sql"
  dev = urlsetpath(replace(var.database_url, "localhost", "host.docker.internal"), "atlas_dev")

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
