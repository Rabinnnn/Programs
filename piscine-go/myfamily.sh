#!bin/bash
curl -s "https://learn.zone01kisumu.ke/assets/superhero/all.json" | jq --argjson hero_id "$HERO_ID" ' .[] | select( .id==$hero_id) | .connections.relatives '| sed 's/"//g'