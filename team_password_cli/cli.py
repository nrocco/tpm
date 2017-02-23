import click
from datetime import datetime
from tabulate import tabulate
from urllib.parse import quote
from pprint import pprint


@click.command()
@click.argument('search')
@click.pass_context
def search(context, search):
    '''
    Search for passwords.

    When searching for passwords in Team Password Manager you can use special
    operators that can help you refine your results. Search operators are
    special words that allow you to find passwords quickly and accurately.

    \b
    tag:string
            Search passwords that have a tag that matches the string.

    \b
    access:string
            Search passwords that have the string in the access field.

    \b
    username:string
            Search passwords that have the string in the username field.

    \b
    name:string
            Search passwords that have the string in the name field
    '''
    page = 1
    passwords = []

    while True:
        uri = '/api/v4/passwords/search/{}/page/{}.json'.format(quote(search), page)
        response = context.obj['client'].get(uri)

        if response.status_code != 200:
            click.echo("Error!")
            return 1

        for data in response.json():
            passwords.append([str(data['id']), data['name'], data['access_info'], data['username'], data['tags']])

        if 'Link' not in response.headers:
            break

        page += 1

    click.echo(tabulate(passwords,
                        ['id', 'name', 'access info', 'username', 'tags'],
                        tablefmt='simple'))


@click.command()
@click.option('--raw', is_flag=True)
@click.argument('id')
@click.pass_context
def show(context, raw, id):
    response = context.obj['client'].get('/api/v4/passwords/{}.json'.format(id))

    if response.status_code != 200:
        log.error("Could not find password with this id")
        return 1

    entry = response.json()

    if raw:
        pprint(entry)
    else:
        click.echo('Name:     {}'.format(entry['name']))
        click.echo('Id:       {}'.format(entry['id']))
        click.echo('Group:    {}'.format(entry['project']['name']))
        click.echo('Access:   {}'.format(entry['access_info']))
        click.echo('Username: {}'.format(entry['username']))
        click.echo('Password: ' + click.style(entry['password'], bg='red', fg='red'))
        click.echo('Tags:     {}'.format(entry['tags']))
