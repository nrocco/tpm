team_password_cli
=================

.. image:: https://travis-ci.org/nrocco/team_password_cli.svg?branch=master
    :target: https://travis-ci.org/nrocco/team_password_cli

A Team Password Manager CLI Application


installation
------------

`team_password_cli` is fully python 2 and python 3 compatible.

It is highly recommended to use virtualenv for this. After having your
virtualenv installed and activated run the following command to install the
`team_password_cli` package directly from pypi (using pip)::

    $ pip install team_password_cli


Alternatively you can manually clone `team_password_cli` and run setupttools `setup.py`::

    $ git clone https://github.com/nrocco/team_password_cli.git
    $ cd team_password_cli
    $ python setup.py install


This will install the needed python libraries.

If you don't want to install `team_password_cli` as a package you can run it directly
from the root directory of the git repository using the following command but
you are responsible for manually installing dependencies::

    $ python -m team_password_cli


To contribute install `team_password_cli` using setuptools develop::

    $ python setup.py develop


Usage
-----

Create a configuration file with the following format::

    % cat ~/.passctl.ini
    [passctl]
    baseurl = https://passwords.example.com/index.php
    username = my-username@example.com
    password = xxxxxxxxxxxxxxxxxxx


Contributing
------------

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Make sure that tests pass (`make test`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new Pull Request


Contributors
------------

- Nico Di Rocco (https://github.com/nrocco)
- Arnoud Vermeer (https://github.com/funzoneq)
