# Fetch data from the JSON file using curl and store it in a variable
json=$(curl -s 'https://learn.zone01kisumu.ke/assets/superhero/all.json')

# Use jq to extract the name, power, and gender from the JSON and format the output
jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"' <<< "$json"
