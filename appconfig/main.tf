provider "appconfig" {

}

variable "label" {
  default = ""
}
variable "value" {
  default = ""
}
variable "key" {
  default = ""
}
resource "appconfigkv_create" "set" {
  key = var.key
  value = var.value
  label = var.label
}