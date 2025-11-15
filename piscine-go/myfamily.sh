#! /bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
| jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives" \
| perl -pe 's/\r?\n/\\n/g' | sed 's/\\n$//'
