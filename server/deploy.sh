#!/bin/bash

echo what should be the version
read VERSION

docker build -t haizza11/lireddit:$VERSION .
docker push haizza11/lireddit:$VERSION
ssh root@167.99.5.241 "docker pull haizza11/lireddit:$VERSION && docker tag haizza11/lireddit:$VERSION dokku/api:$VERSION && dokku tags:deploy api $VERSION "