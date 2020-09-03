if "{{ .sortType }}" == "asc"; then
  echo '{{ .data }}' | jq 'sort_by({{ .sortBy }})'
else
  echo '{{ .data }}' | jq 'sort_by({{ .sortBy }})' | jq 'reverse'
fi
exit 0