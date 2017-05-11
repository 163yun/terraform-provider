output "namespace_id" {
  value = "${join(",", netease_namespace.namespace.*.id)}"
}