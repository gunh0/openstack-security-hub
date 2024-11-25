#!/bin/bash

DESCRIPTION="Is CSRF_COOKIE_SECURE parameter set to True?"
RESULT="[PASS]"
DETAILS=""

CONFIG_FILE="/etc/openstack-dashboard/local_settings.py"

echo "[*] Checking CSRF_COOKIE_SECURE setting in Dashboard configuration"
echo "[*] Target file: $CONFIG_FILE"

# Check if file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "[*] Configuration file not found"
    RESULT="[NA]"
    DETAILS="Config file $CONFIG_FILE does not exist"
else
    echo "[*] Checking CSRF_COOKIE_SECURE parameter"

    # Search for CSRF_COOKIE_SECURE setting
    CSRF_SETTING=$(grep -E "^CSRF_COOKIE_SECURE\s*=\s*(True|False)" "$CONFIG_FILE" 2>/dev/null)

    if [ -z "$CSRF_SETTING" ]; then
        echo "[*] CSRF_COOKIE_SECURE setting not found"
        RESULT="[FAIL]"
        DETAILS="CSRF_COOKIE_SECURE parameter not found in configuration file"
    else
        echo "[*] Found setting: $CSRF_SETTING"
        if echo "$CSRF_SETTING" | grep -q "True"; then
            DETAILS="CSRF_COOKIE_SECURE is properly set to True"
        else
            RESULT="[FAIL]"
            DETAILS="CSRF_COOKIE_SECURE is set to False, which is insecure"
        fi
    fi
fi

# Output JSON result
cat <<EOF
{
    "description": "$DESCRIPTION",
    "result": "$RESULT",
    "details": "$DETAILS"
}
EOF
