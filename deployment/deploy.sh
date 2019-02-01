#!/bin/bash

echo 'Deploying...'

## install helm
echo "Check Helm is installed"
if [[ $((helm) 2>&1 | grep "command not found" ) ]]; then
    echo "You must install Helm. https://github.com/helm/helm/blob/master/docs/install.md"
    exit 1
fi

echo "Installing Helm"

## authenticate with GKE

apt-get install -qq -y gettext
echo $GCLOUD_SERVICE_KEY > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
gcloud --quiet config set project ${GOOGLE_PROJECT_ID}
gcloud --quiet config set compute/zone ${GOOGLE_COMPUTE_ZONE}
gcloud --quiet container clusters get-credentials ${GOOGLE_CLUSTER_NAME}

## deploy helm chart

helm upgrade -i article-feed ./article-feed --set image.tag=${CIRCLE_BUILD_NUM} --set port=${PORT} --set server=${SERVER} --set db=${DB} --set articleCollection=${ARTICLE_COLLECTION} --set dbUser=${DB_USER} --set dbPassword=${DB_PASSWORD} --set authDb=${AUTH_DB} --set dbSsl=${DB_SSL}
