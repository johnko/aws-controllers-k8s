#!/usr/bin/env bash

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$THIS_DIR/../../.."
SCRIPTS_DIR="$ROOT_DIR/scripts"

. $SCRIPTS_DIR/lib/common.sh
. $SCRIPTS_DIR/lib/aws.sh

# secretsmanager_secret_exists() returns 0 if an SecretsManager Secret with the supplied name
# exists, 1 otherwise.
#
# Usage:
#
#   if ! secretsmanager_secret_exists "$secret_name"; then
#       echo "Secret $secret_name does not exist!"
#   fi
secretsmanager_secret_exists() {
    __secret_id="$1"
    daws secretsmanager describe-secret --secret-id "$__secret_id" --output json >/dev/null 2>&1
    if [[ $? -eq 254 ]]; then
        return 1
    else
        return 0
    fi
}

secretsmanager_secret_jq() {
    __secret_id="$1"
    __jq_query="$2"
    json=$( daws secretsmanager describe-secret --secret-id "$__secret_id" --output json || exit 1 )
    echo "$json" | jq --raw-output $__jq_query
}
