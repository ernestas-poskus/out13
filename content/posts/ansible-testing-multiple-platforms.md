+++
Description = "Testing Ansible playbooks on multiple platforms with Continuous Integration (Travis)."
Categories = ["ansible"]
Tags = [
  "personal",
  "ansible",
  "testing",
]
date = "2016-09-25T19:12:59+03:00"
title = "Testing Ansible on multiple platforms"

+++

It is very uncommon to find tests on Ansible playbooks. However when they exist it means playbook was created with care.

When playbook is created with Ansible command ```ansible-galaxy init myplaybook``` it creates number of directories and files, includes basic Travis CI .travis.yml file.

```yaml
script:
  # Basic role syntax check
  - ansible-playbook tests/test.yml -i tests/inventory --syntax-check
```

Basic Travis script brings standard virtual environment operating system which is Ubuntu 12.04 LTS Server Edition 64 bit by default. 
This limits testing playbooks on multiple operating systems. 
In fact Travis CI virtual environment is limited to Debian operating systems and includes BETA Ubuntu 14.04 LTS Server Edition 64 bit container configurable via.

```yaml
dist: trusty
```

---

#### Fortunately Travis supports docker.

I have created bare docker containers that include ansible, can be found here https://github.com/ansible-docker-images and in docker hub https://hub.docker.com/r/ernestasposkus.

Currently available containers are:

 - ubuntu1404
 - ubuntu1604
 - centos6
 - centos7
 
If you need extra platforms let me now or join organization and contribute.
 

## Drop in replacement .travis.yml script for testing ansible on multiple platforms

Template can be found here: https://github.com/ansible-docker-images/template.

This includes four files:

 - .travis.yml for starting CI itself
 - tests/dependencies.yml for extra dependencies to be installed before testing playbook
 - playbook.yml for syntax check, first install and idempotence tests
 - test.yml where tests are defined basically Ansible tasks with exit status checking
 
If you are looking for real examples:

 - https://github.com/AnsibleShipyard/ansible-zookeeper
 - https://github.com/AnsibleShipyard/ansible-mesos
 - https://github.com/ernestas-poskus/ansible-nsq
 - https://github.com/ernestas-poskus/ansible-prometheus

Sample .travis.yml script below. 
Includes testing on 4 platforms, dependencies installation, Ansible syntax check, idempotence test and actual tests.

```yaml
sudo: required

env:
  global:
    # https://github.com/travis-ci/travis-ci/issues/6461#issuecomment-239577306
    DOCKER_VERSION: "1.9.1-0~trusty"
  matrix:
    - distro: ernestasposkus/centos7
      init: /usr/lib/systemd/systemd
      run_opts: "--privileged --volume=/sys/fs/cgroup:/sys/fs/cgroup:ro"
    - distro: ernestasposkus/centos6
      init: /sbin/init
      run_opts: ""
    - distro: ernestasposkus/ubuntu1604
      init: /lib/systemd/systemd
      run_opts: "--privileged --volume=/sys/fs/cgroup:/sys/fs/cgroup:ro"
    - distro: ernestasposkus/ubuntu1404
      init: /sbin/init
      run_opts: ""

services:
  - docker

before_install:
  # Downgrade to specific version of Docker engine.
  - sudo apt-get update
  - sudo apt-get remove docker-engine -yq
  - sudo apt-get install docker-engine=$DOCKER_VERSION -yq --no-install-suggests --no-install-recommends --force-yes -o Dpkg::Options::="--force-confnew"

  # Pull container.
  - 'sudo docker pull ${distro}:latest'

script:
  - container_id=$(mktemp)
    # Run container in detached state.
  - 'sudo docker run --detach --volume="${PWD}":/etc/ansible/roles/role_under_test:ro ${run_opts} ${distro}:latest "${init}" > "${container_id}"'

  # Inspect docker container
  - 'sudo docker inspect $(cat ${container_id})'

  # Print ansible version
  - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible --version'

  # Check Ansible host setup
  - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible all -i "localhost," -c local -m setup'

  # Install dependencies
  # Uncomment to install dependencies
  # - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible-galaxy install geerlingguy.java'
  # - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible-playbook /etc/ansible/roles/role_under_test/tests/dependencies.yml'

  # Ansible syntax check.
  - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible-playbook /etc/ansible/roles/role_under_test/tests/playbook.yml --syntax-check'

  # Test role.
  - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible-playbook /etc/ansible/roles/role_under_test/tests/playbook.yml'

  # Test role idempotence.
  - idempotence=$(mktemp)
  - sudo docker exec "$(cat ${container_id})" ansible-playbook /etc/ansible/roles/role_under_test/tests/playbook.yml | tee -a ${idempotence}
  - >
    tail ${idempotence}
    | grep -q 'changed=0.*failed=0'
    && (echo 'Idempotence test: pass' && exit 0)
    || (echo 'Idempotence test: fail' && exit 1)

  # Test role.
  - 'sudo docker exec --tty "$(cat ${container_id})" env TERM=xterm ansible-playbook /etc/ansible/roles/role_under_test/tests/test.yml'

  # View container logs
  - 'sudo docker logs "$(cat ${container_id})"'

  # Clean up.
  - 'sudo docker stop "$(cat ${container_id})"'

notifications:
  webhooks: https://galaxy.ansible.com/api/v1/notifications/
```
