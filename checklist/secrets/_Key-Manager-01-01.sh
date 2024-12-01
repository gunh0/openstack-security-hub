#!/bin/bash

DESCRIPTION="Is user/group ownership of /etc/barbican/barbican.conf set to root:barbican?"
RESULT="[PASS]"
DETAILS=""

CONFIG_FILE="/etc/barbican/barbican.conf"
EXPECTED_USER="barbican"
EXPECTED_GROUP="barbican"

echo "[*] Checking Barbican configuration file ownership"
echo "[*] Target file: $CONFIG_FILE"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "[*] Configuration file not found"
    RESULT="[NA]"
    DETAILS="Config file $CONFIG_FILE does not exist"
else
    echo "[*] Checking file ownership"
    CURRENT_USER=$(stat -c '%U' $CONFIG_FILE)
    CURRENT_GROUP=$(stat -c '%G' $CONFIG_FILE)

    echo "[*] Current ownership: $CURRENT_USER:$CURRENT_GROUP"
    echo "[*] Expected ownership: $EXPECTED_USER:$EXPECTED_GROUP"

    if [ "$CURRENT_USER" == "$EXPECTED_USER" ] && [ "$CURRENT_GROUP" == "$EXPECTED_GROUP" ]; then
        DETAILS="Config file is owned by $CURRENT_USER:$CURRENT_GROUP as expected"
    else
        RESULT="[FAIL]"
        DETAILS="Config file is owned by $CURRENT_USER:$CURRENT_GROUP, expected $EXPECTED_USER:$EXPECTED_GROUP"
    fi
fi

# Get current timestamp
TIMESTAMP=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

# Output single-line JSON to stdout
echo "{\"description\":\"$DESCRIPTION\",\"result\":\"$RESULT\",\"details\":\"$DETAILS\",\"timestamp\":\"$TIMESTAMP\"}"
