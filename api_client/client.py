import requests
import json

class Client(object):
    def __init__(self, url, out=None):
        if url in (None, ''):
            raise Exception('no url was provided')

        self.url = url.rstrip('/') + '/'
        self.out = out

    def call(self, api, data=None, method=None, multipart=False):
        try:
            response = requests.post(url=self.url + api, json=data)
            response.raise_for_status()
        except Exception as e:
            raise Exception((e, response.json()))
        return response.json()

    def create_game(self, init=True, rows=10, columns=10, **kwargs):
        self.call('games', data={ 'rows': rows, 'columns': columns })

    def reveal_cell(self, init=True, game_id='', row=0, col=0, **kwargs):
        self.call('games/'+game_id+'cells/'+row+'/'+col)
