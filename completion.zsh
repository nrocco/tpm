#compdef tpm
_tpm() {
  local -a commands
  commands=(
    'help:Help about any command'
    'password:Manage passwords'
    'version:Show version of the client and server'
  )

  local -a password_commands
  password_commands=(
    'help:Help about any command'
    'show:Show a single password'
    'generate:Generate a strong, random password'
    'search:Search for passwords'
  )

  if (( CURRENT == 2 )); then
    _describe -t commands 'commands' commands
  elif (( CURRENT == 3)); then
    if [[ $words[2] == 'password' ]]; then
        _describe -t password_commands 'password_commands' password_commands
    fi
  fi

  return 0
}

_tpm