# -*- coding: utf-8 -*-
import hashlib
import requests
import setting
import time
from pprint import pprint


def render_auth():
    timestamp = str(int(time.time()))
    dev_hash = hashlib.md5()
    dev_hash.update(timestamp)
    dev_hash.update(setting.API_SECRET)
    return {
            'api_key': setting.API_KEY,
            'timestamp': timestamp,
            'dev_hash': dev_hash.hexdigest()
           }

def api(api_path):
    r = requests.get('https://platform.api.onesky.io/1/%s' % api_path,
            params=render_auth())
    return r.json()

def upload_po(project_id, file):
    params = render_auth()
    params['file_format'] = 'GNU_PO'
    r = requests.post('https://platform.api.onesky.io/1/projects/%s/files' % project_id,
            params=params, files={'file': file})
    return r.json()

if __name__ == '__main__':
    #pprint(api('project-groups'))
    #pprint(api('project-groups/%s/projects' % setting.PROJECT_GROUP_ID))
    with open(setting.PO_FILES, 'rb') as pof:
        pprint(upload_po(setting.PROJECT_ID, pof))
