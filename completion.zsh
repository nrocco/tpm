#compdef tpm
_tpm() {
  local -a commands
  commands=(
    'help:Help about any command'
    'password:Manage passwords'
    'project:Manage projects'
    'version:Show version of the client and server'
  )

  local -a password_commands
  password_commands=(
    'help:Help about any command'
    'show:Show a single password'
    'generate:Generate a strong, random password'
    'list:List passwords'
  )

  local -a project_commands
  project_commands=(
    'help:Help about any command'
    'show:Show a single project'
    'archive:Archive a single project'
    'unarchive:Unarchive a single project'
    'list:List projects'
  )

  if (( CURRENT == 2 )); then
    _describe -t commands 'commands' commands
  elif (( CURRENT == 3)); then
    if [[ $words[2] == 'password' ]]; then
        _describe -t password_commands 'password_commands' password_commands
    elif [[ $words[2] == 'project' ]]; then
        _describe -t project_commands 'project_commands' project_commands
    fi
  fi

  return 0
}

_tpm
