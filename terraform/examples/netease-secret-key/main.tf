provider "netease" {
  app_key = "${var.app_key}"
  app_secret = "${var.app_secret}"
}

resource "netease_secret_key" "secret_key" {
  key_name = "${var.key_name}"
}