#!/bin/sh
set -Cefu

source='https://raw.githubusercontent.com/Neved4/homebrew-tap/main/README.md'

replace() {
     input=$1
	output=$2
	synced="${output%.md}_synced.md"
     start='<!-- START SYNC -->'
       end='<!-- END SYNC -->'

    sed -n "/$start/,/$end/p" "$input" |
        sed -e '1d;$d' >| sync_content.txt

    sed -e "/$start/,/$end/{//!d;}" \
        -e "/$start/r sync_content.txt" \
        -e "/$start/!b" \
        -e "/$end/b" "$output" >| "$synced"

    mv "$synced" "$output"
	rm sync_content.txt
}

replace "$source" 'README.md'
