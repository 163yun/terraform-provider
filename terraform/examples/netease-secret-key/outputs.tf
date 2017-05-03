output "secret_key_ids" {
  value = "${join(",", netease_secret_key.secret_key.*.id)}"
}