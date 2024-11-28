package keymanager

import (
	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/gunh0/openstack-security-hub/util"
	"golang.org/x/crypto/ssh"
)

func CheckKeyManager0101(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/keymanager/key-manager-01-01.sh",
		"Is the ownership of config files set to root/barbican?",
	)
}
