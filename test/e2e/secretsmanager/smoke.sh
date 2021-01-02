#!/usr/bin/env bash

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$THIS_DIR/../../.."
SCRIPTS_DIR="$ROOT_DIR/scripts"

source "$SCRIPTS_DIR/lib/common.sh"
source "$SCRIPTS_DIR/lib/aws.sh"
source "$SCRIPTS_DIR/lib/aws/secretsmanager.sh"
source "$SCRIPTS_DIR/lib/k8s.sh"
source "$SCRIPTS_DIR/lib/testutil.sh"

wait_seconds=10
test_name="$( filenoext "${BASH_SOURCE[0]}" )"
service_name="secretsmanager"
ack_ctrl_pod_id=$( controller_pod_id )
debug_msg "executing test: $service_name/$test_name"

# This smoke test creates and deletes a set of SecretsManager secrets. It creates
# more than 1 secret in order to ensure that the ReadMany code paths and
# the associated generated code that do object lookups for a single object work
# when >1 object are returned in various List operations.

# PRE-CHECKS

for x in a b c; do

    secret_name="ack-test-smoke-$service_name-$x"
    resource_name="awssecrets/$secret_name"

    if secretsmanager_secret_exists "$secret_name"; then
        echo "FAIL: expected $secret_name to not exist in SecretsManager. Did previous test run cleanup?"
        exit 1
    fi

    if k8s_resource_exists "$resource_name"; then
        echo "FAIL: expected $resource_name to not exist. Did previous test run cleanup?"
        exit 1
    fi

done

# TEST ACTIONS and ASSERTIONS

# Create the secrets
for x in a b c; do

    secret_name="ack-test-smoke-$service_name-$x"

    cat <<EOF | kubectl apply -f -
apiVersion: secretsmanager.services.k8s.aws/v1alpha1
kind: AWSSecret
metadata:
  name: $secret_name
spec:
  name: $secret_name
EOF

done

sleep $wait_seconds

# Check the secrets were created
for x in a b c; do

    secret_name="ack-test-smoke-$service_name-$x"
    resource_name="awssecrets/$secret_name"

    debug_msg "checking secret $secret_name created in SecretsManager"
    if ! secretsmanager_secret_exists "$secret_name"; then
        echo "FAIL: expected $secret_name to have been created in SecretsManager"
        kubectl logs -n ack-system "$ack_ctrl_pod_id"
        exit 1
    fi
done

sleep $wait_seconds

# Update one of the secret's Description
# attributes to check the update code path
updated_secret_name="ack-test-smoke-$service_name-b"

# isc_json=$( secretsmanager_secret_jq $updated_secret_name '.Description')
# assert_equal "false" "$isc_json" "Expected description to be '' but got '$isc_json'" || exit 1

cat <<EOF | kubectl apply -f -
apiVersion: secretsmanager.services.k8s.aws/v1alpha1
kind: AWSSecret
metadata:
  name: $updated_secret_name
spec:
  name: $updated_secret_name
  description: b
EOF

sleep $wait_seconds

debug_msg "checking secret $updated_secret_name updated description in SecretsManager"
isc_json=$( secretsmanager_secret_jq $updated_secret_name '.Description')
assert_equal "b" "$isc_json" "Expected description to be 'b' but got '$isc_json'" || exit 1

sleep $wait_seconds

# Delete each of the created secrets

for x in a b c; do

    secret_name="ack-test-smoke-$service_name-$x"
    resource_name="awssecrets/$secret_name"

    kubectl delete "$resource_name" 2>/dev/null
    assert_equal "0" "$?" "Expected success from kubectl delete but got $?" || exit 1
done

sleep $wait_seconds

for x in a b c; do

    secret_name="ack-test-smoke-$service_name-$x"

    if secretsmanager_secret_exists "$secret_name"; then
        echo "FAIL: expected $secret_name to be deleted in SecretsManager"
        kubectl logs -n ack-system "$ack_ctrl_pod_id"
        exit 1
    fi
done

assert_pod_not_restarted $ack_ctrl_pod_id
