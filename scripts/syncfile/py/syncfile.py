#!/usr/bin/env python3

import re
import requests

URL = 'https://raw.githubusercontent.com/Neved4/homebrew-tap/main/README.md'
START_SYNC = '<!-- START SYNC -->'
END_SYNC = '<!-- END SYNC -->'

def replace(input_url, output_file):
	content = requests.get(input_url).text
	synced_content = re.search(f'{START_SYNC}(.*?){END_SYNC}',
		content, re.DOTALL).group(1).strip()

	with open(output_file) as f:
		updated = re.sub(f'{START_SYNC}.*?{END_SYNC}',
			f'{START_SYNC}\n{synced_content}\n{END_SYNC}',
			f.read(), flags=re.DOTALL)

	with open(output_file, 'w') as f:
		f.write(updated)

replace(URL, 'README.md')
