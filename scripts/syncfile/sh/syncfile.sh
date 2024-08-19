#!/bin/sh
set -Cefu

raw='https://raw.githubusercontent.com'
url="$raw/Neved4/homebrew-tap/main/README.md"

replace() {
	 src=$1
	out=$2
	synced="${out%.md}_synced.md"
	 start='<!-- START SYNC -->'
	   end='<!-- END SYNC -->'

	curl -s "$src" | sed -n "/$start/,/$end/p" |
		sed -e '1d;$d' >| sync_content.txt

	sed -e "/$start/,/$end/{//!d;}" \
		-e "/$start/r sync_content.txt" \
		-e "/$start/!b" \
		-e "/$end/b" "$out" >| "$synced"

	mv "$synced" "$out"
	rm sync_content.txt
}

replace "$url" 'README.md'
