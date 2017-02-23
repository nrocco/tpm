#!/usr/bin/env python
import re
import io
import codecs

from setuptools import setup, find_packages


setup(
    name = 'team_password_cli',
    description = 'Team Password Manager cli',
    version = '1.0.1',
    author = 'Nico Di Rocco',
    author_email = 'dirocco.nico@gmail.com',
    url = 'https://github.com/nrocco/team_password_cli',
    license = 'GPLv3',
    long_description = codecs.open('README.rst', 'rb', 'utf-8').read(),
    download_url = 'https://github.com/nrocco/team_password_cli/tags',
    include_package_data = True,
    install_requires = [
        'click',
        'requests',
        'tabulate',
    ],
    entry_points = {
        'console_scripts': [
            'passctl = team_password_cli.__main__:cli',
        ]
    },
    packages = find_packages(),
    zip_safe = False,
    classifiers = [
        'Development Status :: 5 - Production/Stable',
        'Environment :: Console',
        'License :: OSI Approved :: GNU General Public License v3 (GPLv3)',
        'Operating System :: Unix',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.2',
        'Programming Language :: Python :: 3.3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Topic :: Software Development :: Libraries :: Python Modules',
        'Topic :: Utilities'
    ]
)
