echo "$(units {{ .from }} {{ .to }} | sed --silent -e '1p' | awk '{print $2}')*10" | bc
exit 0
