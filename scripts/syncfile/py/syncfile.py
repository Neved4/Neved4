#!/usr/bin/env python3

import re
import requests

RAW = "https://raw.githubusercontent.com"
URL = RAW + "/Neved4/homebrew-tap/main/README.md"
START = "<!-- START SYNC -->"
END = "<!-- END SYNC -->"

def replace(input_url, output_file):
	content = requests.get(input_url).text
	synced_content = re.search(f"{START}(.*?){END}",
		content, re.DOTALL).group(1).strip()

	with open(output_file) as f:
		updated = re.sub(f"{START}.*?{END}",
			f"{START}\n{synced_content}\n{END}",
			f.read(), flags=re.DOTALL)

	with open(output_file, "w") as f:
		f.write(updated)

replace(URL, "README.md")
