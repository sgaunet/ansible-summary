[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/ansible-summary)](https://goreportcard.com/report/github.com/sgaunet/ansible-summary)
[![GitHub release](https://img.shields.io/github/release/sgaunet/ansible-summary.svg)](https://github.com/sgaunet/ansible-summary/releases/latest)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/ansible-summary/total)
![Test Coverage](https://raw.githubusercontent.com/wiki/sgaunet/ansible-summary/coverage-badge.svg)
[![Linter](https://github.com/sgaunet/ansible-summary/actions/workflows/linter.yml/badge.svg)](https://github.com/sgaunet/ansible-summary/actions/workflows/linter.yml)
[![Snapshot](https://github.com/sgaunet/ansible-summary/actions/workflows/snapshot.yml/badge.svg)](https://github.com/sgaunet/ansible-summary/actions/workflows/snapshot.yml)
[![Release](https://github.com/sgaunet/ansible-summary/actions/workflows/release.yml/badge.svg)](https://github.com/sgaunet/ansible-summary/actions/workflows/release.yml)
[![Coverage Badge](https://github.com/sgaunet/ansible-summary/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/ansible-summary/actions/workflows/coverage.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

ansible-summary is a tool to make a summary of the states of ansible tasks. I'm using it to get a report of how many tasks are in chaged state, ok, failures...

The input of the file is the JSON log of ansible.

# Getting started

Execute your playbook:

Set var env to configure ansible to write a json file:

```
export ANSIBLE_CALLBACKS_ENABLED=json
export ANSIBLE_STDOUT_CALLBACK=json 
```

And finaly get, the resue: 

```
ansible-summary -input  $TMPFILE 
```

Example output:

```
$ export ANSIBLE_CALLBACKS_ENABLED=json
$ export ANSIBLE_STDOUT_CALLBACK=json 
$ ansible-playbook -i inventories/yinventory playbook.yml --check --diff > /tmp//res.json
$ ansible-summary -input /tmp/res.json
Tasks not synchronised :
On Host prod_WWW task system/repositories : Configure yu repositories
On Host prod_WWW task website : install website
************************************
prod_WWW ok=229 changed=2 unreachable=0 failures=0 skipped=64 rescued=0 ignored=0
```

# Development

This project is using :

* golang 1.20+
* [task for development](https://taskfile.dev/#/)
* [goreleaser](https://goreleaser.com/)

## Project Status

🟨 **Maintenance Mode**: This project is in maintenance mode.

While we are committed to keeping the project's dependencies up-to-date and secure, please note the following:

- New features are unlikely to be added
- Bug fixes will be addressed, but not necessarily promptly
- Security updates will be prioritized

## Issues and Bug Reports

We still encourage you to use our issue tracker for:

- 🐛 Reporting critical bugs
- 🔒 Reporting security vulnerabilities
- 🔍 Asking questions about the project

Please check existing issues before creating a new one to avoid duplicates.

## Contributions

🤝 Limited contributions are still welcome.

While we're not actively developing new features, we appreciate contributions that:

- Fix bugs
- Update dependencies
- Improve documentation
- Enhance performance or security

## Support

As this project is in maintenance mode, support may be limited. We appreciate your understanding and patience.

Thank you for your interest in our project!
