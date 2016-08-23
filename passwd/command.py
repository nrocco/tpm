from logging import getLogger
from requests import get, put
from tabulate import tabulate
from datetime import datetime
from urllib.parse import quote

from pycli_tools.parsers import get_argparser
from pycli_tools.commands import Command, arg

from passwd import __version__


log = getLogger(__name__)


class SearchCommand(Command):
    '''Search for passwords'''

    args = [
        arg('--no-headers', action="store_true"),
        arg('search', nargs="+")
    ]

    def run(self, args, parser):
        auth = (args.username, args.password)
        headers = {"Content-Type": "application/json; charset=utf-8"}

        if args.search:
            search = quote(' '.join(args.search))
            resource = '/api/v4/passwords/search/{}.json'.format(search)
        else:
            resource = '/api/v4/passwords.json'

        log.debug("Calling {}".format(resource))
        r = get(args.baseurl + resource, auth=auth, headers=headers)

        if not args.no_headers:
            headers = ['id', 'name', 'access info', 'username', 'tags']
            tablefmt = 'simple'
        else:
            headers = []
            tablefmt = 'plain'

        dataset = []
        for data in r.json():
            dataset.append([str(data['id']), data['name'], data['access_info'], data['username'], data['tags']])
        print(tabulate(dataset, headers, tablefmt=tablefmt))


def parse_and_run(args=None):
    parser = get_argparser(
        prog='passctl',
        version=__version__,
        logging_format='[%(asctime)-15s] %(levelname)s %(message)s',
        description='Team Password Manager cli',
        default_config=['.passctlrc']
    )

    parser.add_commands([
        SearchCommand()
    ])

    args = parser.parse_args()
    args.func(args, parser=parser)
