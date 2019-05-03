terraform {
  backend "gcs" {
    bucket = "article-app-storage"
    prefix = "terraform/state"
  }
}

module "article-app-cluster" {
  source = "git@github.com:tommbee/k8s-prometheus-terraform-module.git"

  config_file = "${var.config_file}"
  region = "europe-west1-c"
  projet_name = "temporal-parser-229715"
  cluster_name = "article-app"
  machine_type = "n1-standard-2"
}

module "deploy" {
    source = "./deploy"

    helm_service_account = "${module.article-app-cluster.helm_service_account}"
    helm_namespace = "${module.article-app-cluster.helm_namespace}"
    kubeconfig = "${module.article-app-cluster.kubeconfig}"
    #token = "${module.article-app-cluster.token}"

    ## app specific
    image_repository = "${var.image_repository}"
    image_tag = "${var.image_tag}"
    server = "${var.server}"
    port = "${var.port}"
    db = "${var.db}"
    article_collection = "${var.article_collection}"
    db_user = "${var.db_user}"
    db_password = "${var.db_password}"
    auth_db = "${var.auth_db}"
    db_ssl = "${var.db_ssl}"
    namespace = "${var.namespace}"
}
