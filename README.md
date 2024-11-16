# openstack-security-hub

### Environment Setup

**Setup from `openstack/devstack`**

> <https://opendev.org/openstack/devstack>
>
> DevStack is a series of extensible scripts used to quickly bring up a complete OpenStack environment based on the latest versions of everything from git master. It is used interactively as a development environment and as the basis for much of the OpenStack project’s functional testing.
>
> ...
>
> Install Linux
>
> Start with a clean and minimal install of a Linux system. DevStack attempts to support the two latest LTS releases of Ubuntu, Rocky Linux 9 and openEuler.
>
> If you do not have a preference, **Ubuntu 22.04 (Jammy)** is the most tested, and will probably go the smoothest.

<br/>

### Openstack Security Guide

> <https://docs.openstack.org/security-guide/>

**This book provides best practices and conceptual information about securing an OpenStack cloud.**

- [identity-01] Is user/group ownership of config files set to keystone?
  - [x] [identity-01-01] `/etc/keystone/keystone.conf`
  - [ ] `/etc/keystone/keystone-paste.ini`
  - [ ] `/etc/keystone/policy.json`
  - [ ] `/etc/keystone/logging.conf`
  - [ ] `/etc/keystone/ssl/certs/signing_cert.pem`
  - [ ] `/etc/keystone/ssl/private/signing_key.pem`
  - [ ] `/etc/keystone/ssl/certs/ca.pem`
  - [ ] `/etc/keystone`
