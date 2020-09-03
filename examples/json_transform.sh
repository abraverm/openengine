echo '{{ .data }}' | jq '.[] | {{ .template }}' | jq -s '.'
exit 0