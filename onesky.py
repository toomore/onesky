# -*- coding: utf-8 -*-
import hashlib
import requests
import time
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

    def upload(self, project_id, file, file_format='GNU_PO'):
        params = {'file_format': file_format}
        result = self.api_post('projects/%s/files' % project_id, params=params,
                files={'file': file})
        return result.json()

    def download(self, project_id, locale, source_file_name, export_file_name=None):
        save_filename = source_file_name
        params = {'locale': locale,
                  'source_file_name': source_file_name}

        if export_file_name:
            params.update(export_file_name=export_file_name)
            save_filename = export_file_name

        result = self.api_get('projects/%s/translations' % project_id, params=params)
        if result.status_code == 200:
            with open('./%s_%s' % (locale, save_filename), 'w') as pof:
                pof.write(result.content)

        return result.status_code

if __name__ == '__main__':
    import setting
    from pprint import pprint
    onesky = Onesky(setting.API_KEY, setting.API_SECRET)
    pprint(onesky.api_get('project-groups').json())
    pprint(onesky.api_get('project-groups/%s/projects' % setting.PROJECT_GROUP_ID).json())
    #with open(setting.PO_FILES, 'rb') as pof:
    #    pprint(onesky.upload(setting.PROJECT_ID, pof))
    print onesky.download(setting.PROJECT_ID, 'zh-Hant-TW', '...')
