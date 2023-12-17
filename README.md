
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/ansible-summary)](https://goreportcard.com/report/github.com/sgaunet/ansible-summary)

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
