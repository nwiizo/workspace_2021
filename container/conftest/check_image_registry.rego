package main

deny[msg] {

  input.kind == "Pod"
  image := input.spec.containers[_].image
  not startswith(image, "my-registry.com/")
  msg := sprintf("image '%v' doesn't come from my-company.com repository", [image])
}
