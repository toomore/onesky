# -*- coding: utf-8 -*-
import hashlib
import requests
import setting
import time
from pprint import pprint
from urlparse import urljoin

class Onesky(object):
    def __init__(self, api_key, api_secret):
        self.api_key = api_key
        self.api_secret = api_secret
        self.api_path = 'https://platform.api.onesky.io/1/'

    def render_auth(self):
        timestamp = str(int(time.time()))
        dev_hash = hashlib.md5()
        dev_hash.update(timestamp)
        dev_hash.update(self.api_secret)
        return {
                'api_key': self.api_key,
                'timestamp': timestamp,
                'dev_hash': dev_hash.hexdigest()
               }

    def api_get(self, path, params=None):
        if params and isinstance(params, dict):
            params.update(self.render_auth())
        else:
            params = self.render_auth()

        return requests.get(urljoin(self.api_path, path), params=params)

    def api_post(self, path, params=None, files=None):
        if params and isinstance(params, dict):
            params.update(self.render_auth())
        else:
            params = self.render_auth()

        return requests.post(urljoin(self.api_path, path), params=params,
                files=files)

    def upload_po(self, project_id, file):
        params = {'file_format': 'GNU_PO'}
        result = self.api_post('projects/%s/files' % project_id, params=params,
                files={'file': file})
        return result.json()

if __name__ == '__main__':
    onesky = Onesky(setting.API_KEY, setting.API_SECRET)
    pprint(onesky.api_get('project-groups').json())
    pprint(onesky.api_get('project-groups/%s/projects' % setting.PROJECT_GROUP_ID).json())
    #with open(setting.PO_FILES, 'rb') as pof:
    #    pprint(onesky.upload_po(setting.PROJECT_ID, pof))
