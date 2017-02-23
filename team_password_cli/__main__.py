import click
import configparser
import os

from team_password_cli import __version__, __default_config__
from team_password_cli.cli import search, show
from team_password_cli.rest_client import get_client


#@click.group(invoke_without_command=True)
@click.group()
@click.version_option(__version__)
@click.option('--config', default=__default_config__)
@click.pass_context
def cli(context, config):
    '''
    Create a configuration file with the following format:

    \b
    % cat ~/.passctl.ini
    [passctl]
    baseurl = https://passwords.example.com/index.php
    username = my-username@example.com
    password = xxxxxxxxxxxxxxxxxxx
    '''
    parser = configparser.ConfigParser()
    parser.read(os.path.expanduser(config))

    client = get_client(
        parser['passctl']['baseurl'],
        parser['passctl']['username'],
        parser['passctl']['password']
    )

    context.obj = {
        'config': parser,
        'client': client,
    }


cli.add_command(search)
cli.add_command(show)


if '__main__' == __name__:
    cli()
