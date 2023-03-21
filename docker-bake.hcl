// docker-bake.hcl
target "docker-metadata-action" {}

target "build" {
  inherits = ["docker-metadata-action"]
  context = "./"
  dockerfile = "Dockerfile.stage"
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
