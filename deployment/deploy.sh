#!/bin/bash

echo 'Deploying...'

helm upgrade -i article-feed ./article-feed --set image.tag=${CIRCLE_BUILD_NUM} --set port=${PORT} --set server=${SERVER} --set db=${DB} --set articleCollection=${ARTICLE_COLLECTION} --set dbUser=${DB_USER} --set dbPassword=${DB_PASSWORD} --set authDb=${AUTH_DB} --set dbSsl=${DB_SSL}
