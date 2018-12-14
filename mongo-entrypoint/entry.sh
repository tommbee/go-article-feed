#!/usr/bin/env bash

echo 'Creating application user and article db'

mongo article \
        --host localhost \
        --port 27017 \
        -u root_username \
        -p very_secure_root_pass \
        --authenticationDatabase admin \
        --eval "db.createUser({user: 'article_user', pwd: 'article123', roles:[{role:'dbOwner', db: 'article'}]});"
