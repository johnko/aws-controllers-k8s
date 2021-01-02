// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package awssecret

import (
	"context"

	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/secretsmanager"
)

// customUpdateSecret implements specialized logic for handling Secret
// resource updates. The SecretsManager API has 1 separate API calls to update a
// Secret, depending on the Secret attribute that has changed:
//
// * UpdateSecret for when the KmsKeyId, Description changed
func (rm *resourceManager) customUpdateSecret(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	var err error
	var updated *resource
	updated = desired
	if descriptionChanged(desired, latest) {
		updated, err = rm.updateDescription(ctx, updated)
		if err != nil {
			return nil, err
		}
	}
	return updated, nil
}

// descriptionChanged returns true if the image tag mutability of the
// supplied desired and latest Repository resources is different
func descriptionChanged(
	desired *resource,
	latest *resource,
) bool {
	dspec := desired.ko.Spec
	lspec := latest.ko.Spec
	if dspec.Description == nil {
		return lspec.Description != nil
	}
	if lspec.Description == nil {
		return true
	}
	dval := *dspec.Description
	lval := *lspec.Description
	return dval != lval
}

// updateDescription calls the UpdateSecret SecretsManager API call for a
// specific secret
func (rm *resourceManager) updateDescription(
	ctx context.Context,
	desired *resource,
) (*resource, error) {
	dspec := desired.ko.Spec
	input := &svcsdk.UpdateSecretInput{
		SecretId: aws.String(*dspec.Name),
	}
	if dspec.Description != nil {
		input.SetDescription(*dspec.Description)
	}
	_, err := rm.sdkapi.UpdateSecretWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return desired, nil
}
