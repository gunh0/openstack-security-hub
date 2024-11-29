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

- **Identity**
- [ ] [identity-01] Is user/group ownership of config files set to keystone?
  - [x] [identity-01-01] `/etc/keystone/keystone.conf`
  - [ ] [identity-01-02] `/etc/keystone/keystone-paste.ini`
  - [ ] [identity-01-03] `/etc/keystone/policy.json`
  - [ ] [identity-01-04] `/etc/keystone/logging.conf`
  - [ ] [identity-01-05] `/etc/keystone/ssl/certs/signing_cert.pem`
  - [ ] [identity-01-06] `/etc/keystone/ssl/private/signing_key.pem`
  - [ ] [identity-01-07] `/etc/keystone/ssl/certs/ca.pem`
  - [ ] [identity-01-08] `/etc/keystone`
- [ ] [identity-02] Are strict permissions set for Identity configuration files?
  - [ ] [identity-02-01] `/etc/keystone/keystone.conf`
  - [ ] [identity-02-02] `/etc/keystone/keystone-paste.ini`
  - [ ] [identity-02-03] `/etc/keystone/policy.json`
  - [ ] [identity-02-04] `/etc/keystone/logging.conf`
  - [ ] [identity-02-05] `/etc/keystone/ssl/certs/signing_cert.pem`
  - [ ] [identity-02-06] `/etc/keystone/ssl/private/signing_key.pem`
  - [ ] [identity-02-07] `/etc/keystone/ssl/certs/ca.pem`
  - [ ] [identity-02-08] `/etc/keystone`
- [ ] [identity-03] is TLS enabled for Identity?
- [identity-04] (Obsolete)
- [ ] [identity-05] Is max_request_body_size set to default (114688)?
- [ ] [identity-06] Disable admin token in /etc/keystone/keystone.conf
- **Dashboard**
- [x] [dashboard-01] Is user/group of config files set to root/horizon?
- [ ] [dashboard-02] Are strict permissions set for horizon configuration files?
- [ ] [dashboard-03] Is DISALLOW_IFRAME_EMBED parameter set to True?
- [x] [dashboard-04] Is CSRF_COOKIE_SECURE parameter set to True?
- [x] [dashboard-05] Is SESSION_COOKIE_SECURE parameter set to True?
- [x] [dashboard-06] Is SESSION_COOKIE_HTTPONLY parameter set to True?
- [ ] [dashboard-07] Is PASSWORD_AUTOCOMPLETE set to False?
- [ ] [dashboard-08] Is DISABLE_PASSWORD_REVEAL set to True?
- **Networking**
- [ ] [networking-01] Is user/group ownership of config files set to root/neutron?
  - [ ] [networking-01-01] `/etc/neutron/neutron.conf`
  - [ ] [networking-01-02] `/etc/neutron/api-paste.ini`
  - [ ] [networking-01-03] `/etc/neutron/policy.json`
  - [ ] [networking-01-04] `/etc/neutron/rootwrap.conf`
  - [ ] [networking-01-05] `/etc/neutron`
- [ ] [networking-02] Are strict permissions set for configuration files?
- [ ] [networking-03] Is keystone used for authentication?
- [ ] [networking-04] Is secure protocol used for authentication?
- [ ] [networking-05] Is TLS enabled on Neutron API server?¶
- **Secrets Management**
- [ ] [key-manager-01] Is user/group ownership of config files set to barbican?
  - [x] [key-manager-01-01] `/etc/barbican/barbican.conf`
- [ ] [key-manager-02] Are strict permissions set for configuration files?
- [x] [key-manager-03] Is OpenStack Identity used for authentication?
- [ ] [key-manager-04] Is TLS enabled for authentication?¶
