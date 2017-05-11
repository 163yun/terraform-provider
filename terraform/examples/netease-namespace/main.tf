provider "netease" {

}

resource "netease_namespace" "namespace" {
  name = "${var.name}"
}