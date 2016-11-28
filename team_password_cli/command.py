from datetime import datetime
from logging import getLogger
from os.path import expanduser
from pprint import pprint
from pycli_tools.commands import Command, arg
from pycli_tools.parsers import get_argparser
from tabulate import tabulate
from team_password_cli import __version__
from team_password_cli.rest_client import get_client
from urllib.parse import quote


log = getLogger(__name__)


def get_credentials(args):
    if not args.username:
        args.username = input('Username: ')
    if not args.password:
        args.password = getpass('Password: ')


class HelpCommand(Command):
    '''Get information on how to use this tool'''

    def run(self, args, parser):
        print('Searching:')
        print('''
    tag:string
            Search passwords that have a tag that matches the string.

    access:string
            Search passwords that have the string in the access field.

    username:string
            Search passwords that have the string in the username field.

    name:string
            Search passwords that have the string in the name field
        ''')

        print('Configuration:')
        print('''
    Create a configuration file with the following format:

    % cat ~/.passctlrc
    [passctl]
    baseurl = https://passwords.example.com/index.php
    username = my-username@example.com
    password = xxxxxxxxxxxxxxxxxxx
        ''')


class SearchCommand(Command):
    '''Search for passwords

    When searching for passwords in Team Password Manager you can use special
    operators that can help you refine your results. Search operators are
    special words that allow you to find passwords quickly and accurately.
    '''

    args = [
        arg('--no-headers', action="store_true"),
        arg('search')
    ]

    def run(self, args, parser):
        resource = '/api/v4/passwords/search/{}.json'.format(quote(args.search))

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
        default_config=[expanduser('~/.passctlrc')]
    )

    parser.add_commands([
        SearchCommand(),
        ShowCommand(),
        HelpCommand(),
    ])

    args = parser.parse_args()
    args.client = get_client(args.baseurl, args.username, args.password)

    args.func(args, parser=parser)
