#!/bin/bash

DESCRIPTION="Is OpenStack Identity used for authentication?"
RESULT="[PASS]"
DETAILS=""

CONFIG_FILE="/etc/barbican/barbican-api-paste.ini"

echo "[*] Checking Keystone authentication configuration in Barbican"
echo "[*] Target file: $CONFIG_FILE"

# Check if configuration file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "[*] Configuration file not found"
    RESULT="[NA]"
    DETAILS="Configuration file $CONFIG_FILE not found"
else
    echo "[*] Checking for authtoken in barbican-api-keystone pipeline"

    # Read the pipeline section
    PIPELINE_SECTION=$(awk '/^\[pipeline:barbican-api-keystone\]/{flag=1; next} /^\[/{flag=0} flag' "$CONFIG_FILE" 2>/dev/null)

    if [ -z "$PIPELINE_SECTION" ]; then
        echo "[*] barbican-api-keystone pipeline section not found"
        RESULT="[FAIL]"
        DETAILS="pipeline:barbican-api-keystone section not found in configuration file"
    else
        echo "[*] Found pipeline section. Checking for authtoken"

        # Check if authtoken is present in the pipeline
        if echo "$PIPELINE_SECTION" | grep -q "authtoken"; then
            DETAILS="authtoken is properly configured in barbican-api-keystone pipeline"
            echo "[*] authtoken found in pipeline configuration"
        else
            RESULT="[FAIL]"
            DETAILS="authtoken is missing in barbican-api-keystone pipeline, authentication may be insecure"
            echo "[*] authtoken not found in pipeline configuration"
        fi
    fi
fi

# Get current timestamp
TIMESTAMP=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

# Output single-line JSON to stdout
echo "{\"description\":\"$DESCRIPTION\",\"result\":\"$RESULT\",\"details\":\"$DETAILS\",\"timestamp\":\"$TIMESTAMP\"}"
