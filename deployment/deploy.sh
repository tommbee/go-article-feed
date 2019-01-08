#!/bin/bash

echo 'Deploying...'

helm upgrade -i article-feed ./article-feed \n
--set image.tag=${CIRCLE_BUILD_NUM} \n
--set port=${PORT} --set server=${SERVER} \n
--set db=${DB} \n
--set articleCollection=${ARTICLE_COLLECTION} \n
--set dbUser=${DB_USER} \n
--set dbPassword=${DB_PASSWORD} \n
--set authDb=${AUTH_DB} \n
--set dbSsl=${DB_SSL}
