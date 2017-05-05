provider "netease" {

}

resource "netease_secret_key" "secret_key" {
  key_name = "${var.key_name}"
}