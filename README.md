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

**Installed (Tested) Versions**

- Ubuntu 22.04 (Jammy)
- openstack 7.1.3

<br/>

### Openstack Security Guide

> <https://docs.openstack.org/security-guide/>

**This book provides best practices and conceptual information about securing an OpenStack cloud.**

- [identity-01] Is user/group ownership of config files set to keystone?
  - [x] [identity-01-01] `/etc/keystone/keystone.conf`
  - [x] [identity-01-02] `/etc/keystone/keystone-paste.ini`
  - [x] [identity-01-03] `/etc/keystone/policy.json`
  - [x] [identity-01-04] `/etc/keystone/logging.conf`
  - [x] [identity-01-05] `/etc/keystone/ssl/certs/signing_cert.pem`
  - [x] [identity-01-06] `/etc/keystone/ssl/private/signing_key.pem`
  - [x] [identity-01-07] `/etc/keystone/ssl/certs/ca.pem`
  - [x] [identity-01-08] `/etc/keystone`
- [identity-02] Are strict permissions set for Identity configuration files?
  - [x] [identity-02-01] `/etc/keystone/keystone.conf`
  - [x] [identity-02-02] `/etc/keystone/keystone-paste.ini`
  - [x] [identity-02-03] `/etc/keystone/policy.json`
  - [x] [identity-02-04] `/etc/keystone/logging.conf`
  - [x] [identity-02-05] `/etc/keystone/ssl/certs/signing_cert.pem`
  - [x] [identity-02-06] `/etc/keystone/ssl/private/signing_key.pem`
  - [x] [identity-02-07] `/etc/keystone/ssl/certs/ca.pem`
  - [x] [identity-02-08] `/etc/keystone`
- [x] [identity-03] is TLS enabled for Identity?
- [identity-04] (Obsolete)
- [x] [identity-05] Is max_request_body_size set to default (114688)?
- [x] [identity-06] Disable admin token in /etc/keystone/keystone.conf
- [key-manager-01] Is user/group ownership of config files set to barbican?
  - [x] [key-manager-01-01] `/etc/barbican/barbican.conf`