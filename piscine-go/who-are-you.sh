#!/bin/bash

# Fetch data from the JSON file using curl and store it in a variable
json=$(curl -s 'https://learn.zone01kisumu.ke/assets/superhero/all.json')  

# Use jq to extract the name from the JSON
name=$(echo "$json" | jq -r '.[] | select(.id == 70) | .name')

# Display the name
echo "\"$name\""
