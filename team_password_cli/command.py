from pprint import pprint
from logging import getLogger
from tabulate import tabulate
from datetime import datetime
from urllib.parse import quote
from pycli_tools.parsers import get_argparser
from pycli_tools.commands import Command, arg
from team_password_cli import __version__
from team_password_cli.rest_client import get_client


log = getLogger(__name__)


class SearchCommand(Command):
    '''Search for passwords'''

    args = [
        arg('--no-headers', action="store_true"),
        arg('search', nargs="?")
    ]

    def run(self, args, parser):
        if args.search:
            search = quote(' '.join(args.search))
            resource = '/api/v4/passwords/search/{}.json'.format(search)
        else:
            resource = '/api/v4/passwords.json'

        log.debug("Calling {}".format(resource))
        r = args.client.get(resource)

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


class ShowCommand(Command):
    '''Show details for a password'''

    args = [
        arg('--raw', help='Echo the raw results', action='store_true'),
        arg('id', help='The id of a password')
    ]

    def run(self, args, parser):
        log.debug('Fetching password for enty {}'.format(args.id))

        response = args.client.get('/api/v4/passwords/{}.json'.format(args.id))

        if response.status_code != 200:
            log.error("Could not find password with this id")
            return 1

        entry = response.json()

        if args.raw:
            pprint(entry)
        else:
            print('Name:     {}'.format(entry['name']))
            print('Id:       {}'.format(entry['id']))
            print('Group:    {}'.format(entry['project']['name']))
            print('Access:   {}'.format(entry['access_info']))
            print('Username: {}'.format(entry['username']))
            print('Password: {}'.format(entry['password']))
            print('Tags:     {}'.format(entry['tags']))


def parse_and_run(args=None):
    parser = get_argparser(
        prog='passctl',
        version=__version__,
        logging_format='[%(asctime)-15s] %(levelname)s %(message)s',
        description='Team Password Manager cli',
        default_config=['.passctlrc']
    )

    parser.add_commands([
        SearchCommand(),
        ShowCommand()
    ])

    args = parser.parse_args()
    args.client = get_client(args.baseurl, args.username, args.password)

    args.func(args, parser=parser)
