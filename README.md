tpm
===

A Team Password Manager Client written in go.

[![Build Status](https://travis-ci.org/nrocco/tpm.svg?branch=master)](https://travis-ci.org/nrocco/tpm)


Usage
-----

Create a configuration file with the following format:

    % cat $HOME/.tpm.yaml
    server: "https://tpm.example.com/index.php"
    username: "my-username"
    password: "my-dkj880s0sjd-password"


The following sub commands are available:

    % tpm help
    A Team Password Manager CLI Application

    Usage:
      tpm [command]

    Available Commands:
      help        Help about any command
      password    Manage passwords
      version     Show version of the client and server

    Flags:
          --config string     config file (default is $HOME/.tpm.yaml)
      -h, --help              help for tpm
      -p, --password string   Password
      -s, --server string     The base url of the Team Password Manager server
      -u, --username string   Username

    Use "tpm [command] --help" for more information about a command.


The `password` sub command has the following commands available:

    % tpm password --help
    Manage passwords

    Usage:
      tpm password [command]

    Available Commands:
      generate    Generate a strong, random password
      search      Search for passwords
      show        Show a single password

    Flags:
      -h, --help   help for password

    Global Flags:
          --config string     config file (default is $HOME/.tpm.yaml)
      -p, --password string   Password
      -s, --server string     The base url of the Team Password Manager server
      -u, --username string   Username

    Use "tpm password [command] --help" for more information about a command.


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
