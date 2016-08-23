from requests import Session

class FuuSession(Session):
    def __init__(self, base_url):
        self.base_url = base_url
        super(FuuSession, self).__init__()

    def request(self, method, url, data=None, headers={}, **kwargs):
        if not url.startswith('http'):
            url = '{}/{}'.format(self.base_url, url.lstrip('/'))
        return super(FuuSession, self).request(method, url, headers=headers, data=data, **kwargs)


def get_client(baseurl, username, password):
    s = FuuSession(baseurl)
    s.auth = (username, password)
    s.headers.update({'Content-Type': 'application/json; charset=utf-8'})
    return s
