package secrets

import (
	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/gunh0/openstack-security-hub/util"
	"golang.org/x/crypto/ssh"
)

func CheckKeyManager0101(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/secrets/key-manager-01-01.sh",
		"Is the ownership of config files set to root/barbican? (/etc/barbican/barbican.conf)",
	)
}

func CheckKeyManager0102(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/secrets/key-manager-01-02.sh",
		"Is the ownership of config files set to root/barbican? (/etc/barbican/barbican-api-paste.ini)",
	)
}

func CheckKeyManager03(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/secrets/key-manager-03.sh",
		"Is OpenStack Identity used for authentication?",
	)
}
