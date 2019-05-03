# variable "client_certificate" {}
# variable "client_key" {}
# variable "cluster_ca_certificate" {}
# variable "host" {}
# variable "token" {}
variable "kubeconfig" {}
variable "helm_service_account" {}
variable "helm_namespace" {}

## app specific
variable "image_repository" {}
variable "image_tag" {}
variable "server" {}
variable "port" {}
variable "db" {}
variable "article_collection" {}
variable "db_user" {}
variable "db_password" {}
variable "auth_db" {}
variable "db_ssl" {}
variable "namespace" {}
