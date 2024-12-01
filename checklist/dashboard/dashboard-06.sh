#!/bin/bash

DESCRIPTION="Is SESSION_COOKIE_HTTPONLY parameter set to True?"
RESULT="[PASS]"
DETAILS=""

CONFIG_FILE="/etc/openstack-dashboard/local_settings.py"
ALTERNATE_CONFIG="/etc/openstack-dashboard/local_settings"

echo "[*] Checking SESSION_COOKIE_HTTPONLY setting in Dashboard configuration"
echo "[*] Primary target file: $CONFIG_FILE"
echo "[*] Alternate target file: $ALTERNATE_CONFIG"

# Check if either file exists
if [ ! -f "$CONFIG_FILE" ] && [ ! -f "$ALTERNATE_CONFIG" ]; then
    echo "[*] Neither configuration file exists"
    RESULT="[NA]"
    DETAILS="Configuration files not found"
else
    # Determine which config file to use
    if [ -f "$CONFIG_FILE" ]; then
        USE_CONFIG="$CONFIG_FILE"
        echo "[*] Using primary configuration file"
    else
        USE_CONFIG="$ALTERNATE_CONFIG"
        echo "[*] Using alternate configuration file"
    fi

    echo "[*] Checking SESSION_COOKIE_HTTPONLY parameter"

    # Search for SESSION_COOKIE_HTTPONLY setting
    SESSION_SETTING=$(grep -E "^SESSION_COOKIE_HTTPONLY\s*=\s*(True|False)" "$USE_CONFIG" 2>/dev/null)

    if [ -z "$SESSION_SETTING" ]; then
        echo "[*] SESSION_COOKIE_HTTPONLY setting not found"
        RESULT="[FAIL]"
        DETAILS="SESSION_COOKIE_HTTPONLY parameter not found in configuration file"
    else
        echo "[*] Found setting: $SESSION_SETTING"
        if echo "$SESSION_SETTING" | grep -q "True"; then
            DETAILS="SESSION_COOKIE_HTTPONLY is properly set to True"
        else
            RESULT="[FAIL]"
            DETAILS="SESSION_COOKIE_HTTPONLY is set to False, which is insecure and vulnerable to XSS attacks"
        fi
    fi
fi

# Get current timestamp
TIMESTAMP=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

# Output single-line JSON to stdout
echo "{\"description\":\"$DESCRIPTION\",\"result\":\"$RESULT\",\"details\":\"$DETAILS\",\"timestamp\":\"$TIMESTAMP\"}"
