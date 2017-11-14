#!/bin/sh

echo "---" > .golab.yml
echo "url: \"http://localhost:8080\"" >> .golab.yml
echo "token: \"${GITLAB_ROOT_PRIVATE_TOKEN}\"" >> .golab.yml
