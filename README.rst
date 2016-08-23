team-password-cli
======

A Team Password Manager CLI Application


installation
------------

`team-password-cli` is fully python 2.7 and python 3.4 compatible.

It is highly recommended to use virtualenv for this.
After having your virtualenv installed and activated run the following command to install
the `team-password-cli` package directly from pypi (using pip)::

    $ pip install team-password-cli


Alternatively you can manually clone `team-password-cli` and run setupttools `setup.py`::

    $ git clone https://github.com/funzoneq/team-password-cli.git
    $ cd team-password-cli
    $ python setup.py install


This will install the needed python libraries.

If you don't want to install `team-password-cli` as a package you can run it directly
from the root directory of the git repository using the following command but
you are responsible for manually installing dependencies::

    $ python -m team-password-cli


To install the required dependencies manually see `requirements.txt` 
or simply run::

    $ pip install -r requirements.txt
