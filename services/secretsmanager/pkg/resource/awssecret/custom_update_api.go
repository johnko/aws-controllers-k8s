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
	"fmt"

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
	if descriptionOrKmsKeyIdChanged(desired, latest) {
		updated, err = rm.updateDescription(ctx, updated)
		if err != nil {
			return nil, err
		}
	}
	return updated, nil
}

// descriptionOrKmsKeyIdChanged returns true if the description of kmskeyid of the
// supplied desired and latest Secret resources is different
func descriptionOrKmsKeyIdChanged(
  desired *resource,
  latest *resource,
) bool {
	descChanged := false
	kmsChanged := false

  dspec := desired.ko.Spec
	lspec := latest.ko.Spec

	// avoid nil pointer dereference
	if dspec.Description == nil {
		return lspec.Description != nil
	}
	if lspec.Description == nil {
		return true
	}
	dvaldesc := *dspec.Description
	lvaldesc := *lspec.Description
	descChanged = dvaldesc != lvaldesc

	// avoid nil pointer dereference
	if dspec.KMSKeyID == nil {
		return lspec.KMSKeyID != nil
	}
	if lspec.KMSKeyID == nil {
		return true
	}
	dvalkms := *dspec.KMSKeyID
	lvalkms := *lspec.KMSKeyID
	kmsChanged = dvalkms != lvalkms

	fmt.Printf("descChanged is %s\n", descChanged)
	fmt.Printf("kmsChanged is %s\n", descChanged)
	fmt.Printf("return is %s\n", descChanged || kmsChanged)
	return (descChanged || kmsChanged)
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
	if dspec.KMSKeyID != nil {
		input.SetKmsKeyId(*dspec.KMSKeyID)
	}
	_, err := rm.sdkapi.UpdateSecretWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return desired, nil
}
