tpm
===

A Team Password Manager Client written in go.

[![Build Status](https://travis-ci.org/nrocco/tpm.svg?branch=master)](https://travis-ci.org/nrocco/tpm)
[![GoDoc](https://godoc.org/github.com/nrocco/tpm/pkg/client?status.svg)](https://godoc.org/github.com/nrocco/tpm/pkg/client)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrocco/tpm)](https://goreportcard.com/report/github.com/nrocco/tpm)


Usage
-----

Create a configuration file with the following format:

    % cat $HOME/.tpm.yaml
    server: "https://tpm.example.com/index.php"
    username: "my-username"
    password: "my-dkj880s0sjd-password"


Alternatively, if you do not like to store passwords in plain text files you
can `gpg` encrypt it:

    echo -n my-dkj880s0sjd-password | gpg --encrypt --armor

And this is how your configuration file looks:

    % cat $HOME/.tpm.yaml
    server: "https://tpm.example.com/index.php"
    username: "my-username"
    password: |
      -----BEGIN PGP MESSAGE-----

      hQEMA4tSPIGBQJQPAQf/Sf52JbFCYctlXl4jB9k60m6XXfs5WO7PVDBoOL55EOrK
      OUWhFm25SyhAqfdsSuBvhzXyszgmA0XJqCvy5y+kT95SP9vvQfuj26kpfeIcyalL
      Io+xr+trD0pdg0C7XpqYExtoFTvBe5XyHSQxQgtAbxVy8I5+MQj0xF9XdGDf5rPf
      pSzJ8QIGSsbd0ybxKkFABGbc7hdgma3lgt+zcHTnA3FNcyTWBg84gtZmTQArhuMU
      4bGnBc8QIGSsbd0ybxKkFABGbc6rlyfqsdOjB9Dt1phubPw3AI8b0hmgoNfITW18
      OLITv/5cti6HSV689YQuG9JcAc8QIGSsbd0ybxKkFABGbcPqYa1n94ZkuGGMp+xX
      azRPHF5lxCJNGW/AsPUwOP2mZNZSR3kaHypX2xAfq8QIGSsbd0ybxKkFABGbcZQV
      oDlA3ZvmiuEPvu+DHPxXQ=
      =YJhV
      -----END PGP MESSAGE-----


The following sub commands are available:

    % tpm help
    A Team Password Manager CLI Application

    Usage:
      tpm [command]

    Available Commands:
      help        Help about any command
      password    Manage passwords
      project     Manage projects
      version     Show version of the client and server

    Flags:
          --config string     config file (default is $HOME/.tpm.yaml)
      -h, --help              help for tpm
      -p, --password string   Password
      -s, --server string     The base url of the Team Password Manager server
      -u, --username string   Username

    Use "tpm [command] --help" for more information about a command.


The `password` sub command has the following commands available:

    % tpm password
    Manage passwords

    Usage:
      tpm password [command]

    Available Commands:
      generate    Generate a strong, random password
      list        List passwords
      lock        Lock a password
      permissions Show the permissions of a password
      show        Show a password
      unlock      Unlock a password

    Flags:
      -h, --help   help for password

    Global Flags:
          --config string     config file (default is $HOME/.tpm.yaml)
      -p, --password string   Password
      -s, --server string     The base url of the Team Password Manager server
      -u, --username string   Username

    Use "tpm password [command] --help" for more information about a command.


The `project` sub command has the following commands available:

    % tpm project
    Manage projects

    Usage:
      tpm project [command]

    Available Commands:
      archive     Archive a project
      list        List projects
      show        Show a single project
      unarchive   Unarchive a project

    Flags:
      -h, --help   help for project

    Global Flags:
          --config string     config file (default is $HOME/.tpm.yaml)
      -p, --password string   Password
      -s, --server string     The base url of the Team Password Manager server
      -u, --username string   Username

    Use "tpm project [command] --help" for more information about a command.


The `version` sub command outputs both client and server versions:

    % tpm version
    Client:
      Version:    ab62aba
      OS/Arch:    darwin/amd64
      Shell:      /bin/zsh
      User:       nrocco

    Server:
      Url:        https://passwords.example.com
      Version:    6.68.138
      Date:       2016-02-26
      ApiVersion: 4


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
