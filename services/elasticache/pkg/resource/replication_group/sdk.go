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

// Code generated by ack-generate. DO NOT EDIT.

package replication_group

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"strings"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/elasticache/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.ReplicationGroup{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeReplicationGroupsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeReplicationGroups", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "ReplicationGroupNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.ReplicationGroups {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AtRestEncryptionEnabled != nil {
			ko.Spec.AtRestEncryptionEnabled = elem.AtRestEncryptionEnabled
		}
		if elem.AuthTokenEnabled != nil {
			ko.Status.AuthTokenEnabled = elem.AuthTokenEnabled
		}
		if elem.AuthTokenLastModifiedDate != nil {
			ko.Status.AuthTokenLastModifiedDate = &metav1.Time{*elem.AuthTokenLastModifiedDate}
		}
		if elem.AutomaticFailover != nil {
			ko.Status.AutomaticFailover = elem.AutomaticFailover
		}
		if elem.CacheNodeType != nil {
			ko.Spec.CacheNodeType = elem.CacheNodeType
		}
		if elem.ClusterEnabled != nil {
			ko.Status.ClusterEnabled = elem.ClusterEnabled
		}
		if elem.ConfigurationEndpoint != nil {
			f7 := &svcapitypes.Endpoint{}
			if elem.ConfigurationEndpoint.Address != nil {
				f7.Address = elem.ConfigurationEndpoint.Address
			}
			if elem.ConfigurationEndpoint.Port != nil {
				f7.Port = elem.ConfigurationEndpoint.Port
			}
			ko.Status.ConfigurationEndpoint = f7
		}
		if elem.Description != nil {
			ko.Status.Description = elem.Description
		}
		if elem.GlobalReplicationGroupInfo != nil {
			f9 := &svcapitypes.GlobalReplicationGroupInfo{}
			if elem.GlobalReplicationGroupInfo.GlobalReplicationGroupId != nil {
				f9.GlobalReplicationGroupID = elem.GlobalReplicationGroupInfo.GlobalReplicationGroupId
			}
			if elem.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole != nil {
				f9.GlobalReplicationGroupMemberRole = elem.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole
			}
			ko.Status.GlobalReplicationGroupInfo = f9
		}
		if elem.KmsKeyId != nil {
			ko.Spec.KMSKeyID = elem.KmsKeyId
		}
		if elem.MemberClusters != nil {
			f11 := []*string{}
			for _, f11iter := range elem.MemberClusters {
				var f11elem string
				f11elem = *f11iter
				f11 = append(f11, &f11elem)
			}
			ko.Status.MemberClusters = f11
		}
		if elem.MultiAZ != nil {
			ko.Status.MultiAZ = elem.MultiAZ
		}
		if elem.NodeGroups != nil {
			f13 := []*svcapitypes.NodeGroup{}
			for _, f13iter := range elem.NodeGroups {
				f13elem := &svcapitypes.NodeGroup{}
				if f13iter.NodeGroupId != nil {
					f13elem.NodeGroupID = f13iter.NodeGroupId
				}
				if f13iter.NodeGroupMembers != nil {
					f13elemf1 := []*svcapitypes.NodeGroupMember{}
					for _, f13elemf1iter := range f13iter.NodeGroupMembers {
						f13elemf1elem := &svcapitypes.NodeGroupMember{}
						if f13elemf1iter.CacheClusterId != nil {
							f13elemf1elem.CacheClusterID = f13elemf1iter.CacheClusterId
						}
						if f13elemf1iter.CacheNodeId != nil {
							f13elemf1elem.CacheNodeID = f13elemf1iter.CacheNodeId
						}
						if f13elemf1iter.CurrentRole != nil {
							f13elemf1elem.CurrentRole = f13elemf1iter.CurrentRole
						}
						if f13elemf1iter.PreferredAvailabilityZone != nil {
							f13elemf1elem.PreferredAvailabilityZone = f13elemf1iter.PreferredAvailabilityZone
						}
						if f13elemf1iter.ReadEndpoint != nil {
							f13elemf1elemf4 := &svcapitypes.Endpoint{}
							if f13elemf1iter.ReadEndpoint.Address != nil {
								f13elemf1elemf4.Address = f13elemf1iter.ReadEndpoint.Address
							}
							if f13elemf1iter.ReadEndpoint.Port != nil {
								f13elemf1elemf4.Port = f13elemf1iter.ReadEndpoint.Port
							}
							f13elemf1elem.ReadEndpoint = f13elemf1elemf4
						}
						f13elemf1 = append(f13elemf1, f13elemf1elem)
					}
					f13elem.NodeGroupMembers = f13elemf1
				}
				if f13iter.PrimaryEndpoint != nil {
					f13elemf2 := &svcapitypes.Endpoint{}
					if f13iter.PrimaryEndpoint.Address != nil {
						f13elemf2.Address = f13iter.PrimaryEndpoint.Address
					}
					if f13iter.PrimaryEndpoint.Port != nil {
						f13elemf2.Port = f13iter.PrimaryEndpoint.Port
					}
					f13elem.PrimaryEndpoint = f13elemf2
				}
				if f13iter.ReaderEndpoint != nil {
					f13elemf3 := &svcapitypes.Endpoint{}
					if f13iter.ReaderEndpoint.Address != nil {
						f13elemf3.Address = f13iter.ReaderEndpoint.Address
					}
					if f13iter.ReaderEndpoint.Port != nil {
						f13elemf3.Port = f13iter.ReaderEndpoint.Port
					}
					f13elem.ReaderEndpoint = f13elemf3
				}
				if f13iter.Slots != nil {
					f13elem.Slots = f13iter.Slots
				}
				if f13iter.Status != nil {
					f13elem.Status = f13iter.Status
				}
				f13 = append(f13, f13elem)
			}
			ko.Status.NodeGroups = f13
		}
		if elem.PendingModifiedValues != nil {
			f14 := &svcapitypes.ReplicationGroupPendingModifiedValues{}
			if elem.PendingModifiedValues.AuthTokenStatus != nil {
				f14.AuthTokenStatus = elem.PendingModifiedValues.AuthTokenStatus
			}
			if elem.PendingModifiedValues.AutomaticFailoverStatus != nil {
				f14.AutomaticFailoverStatus = elem.PendingModifiedValues.AutomaticFailoverStatus
			}
			if elem.PendingModifiedValues.PrimaryClusterId != nil {
				f14.PrimaryClusterID = elem.PendingModifiedValues.PrimaryClusterId
			}
			if elem.PendingModifiedValues.Resharding != nil {
				f14f3 := &svcapitypes.ReshardingStatus{}
				if elem.PendingModifiedValues.Resharding.SlotMigration != nil {
					f14f3f0 := &svcapitypes.SlotMigration{}
					if elem.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage != nil {
						f14f3f0.ProgressPercentage = elem.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage
					}
					f14f3.SlotMigration = f14f3f0
				}
				f14.Resharding = f14f3
			}
			ko.Status.PendingModifiedValues = f14
		}
		if elem.ReplicationGroupId != nil {
			ko.Spec.ReplicationGroupID = elem.ReplicationGroupId
		}
		if elem.SnapshotRetentionLimit != nil {
			ko.Spec.SnapshotRetentionLimit = elem.SnapshotRetentionLimit
		}
		if elem.SnapshotWindow != nil {
			ko.Spec.SnapshotWindow = elem.SnapshotWindow
		}
		if elem.SnapshottingClusterId != nil {
			ko.Status.SnapshottingClusterID = elem.SnapshottingClusterId
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		}
		if elem.TransitEncryptionEnabled != nil {
			ko.Spec.TransitEncryptionEnabled = elem.TransitEncryptionEnabled
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)

	// custom set output from response
	ko, err = rm.CustomDescribeReplicationGroupsSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeReplicationGroupsInput, error) {
	res := &svcsdk.DescribeReplicationGroupsInput{}

	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateReplicationGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateReplicationGroup", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ReplicationGroup.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ReplicationGroup.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ReplicationGroup.AuthTokenEnabled != nil {
		ko.Status.AuthTokenEnabled = resp.ReplicationGroup.AuthTokenEnabled
	}
	if resp.ReplicationGroup.AuthTokenLastModifiedDate != nil {
		ko.Status.AuthTokenLastModifiedDate = &metav1.Time{*resp.ReplicationGroup.AuthTokenLastModifiedDate}
	}
	if resp.ReplicationGroup.AutomaticFailover != nil {
		ko.Status.AutomaticFailover = resp.ReplicationGroup.AutomaticFailover
	}
	if resp.ReplicationGroup.ClusterEnabled != nil {
		ko.Status.ClusterEnabled = resp.ReplicationGroup.ClusterEnabled
	}
	if resp.ReplicationGroup.ConfigurationEndpoint != nil {
		f7 := &svcapitypes.Endpoint{}
		if resp.ReplicationGroup.ConfigurationEndpoint.Address != nil {
			f7.Address = resp.ReplicationGroup.ConfigurationEndpoint.Address
		}
		if resp.ReplicationGroup.ConfigurationEndpoint.Port != nil {
			f7.Port = resp.ReplicationGroup.ConfigurationEndpoint.Port
		}
		ko.Status.ConfigurationEndpoint = f7
	}
	if resp.ReplicationGroup.Description != nil {
		ko.Status.Description = resp.ReplicationGroup.Description
	}
	if resp.ReplicationGroup.GlobalReplicationGroupInfo != nil {
		f9 := &svcapitypes.GlobalReplicationGroupInfo{}
		if resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupId != nil {
			f9.GlobalReplicationGroupID = resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupId
		}
		if resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole != nil {
			f9.GlobalReplicationGroupMemberRole = resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole
		}
		ko.Status.GlobalReplicationGroupInfo = f9
	}
	if resp.ReplicationGroup.MemberClusters != nil {
		f11 := []*string{}
		for _, f11iter := range resp.ReplicationGroup.MemberClusters {
			var f11elem string
			f11elem = *f11iter
			f11 = append(f11, &f11elem)
		}
		ko.Status.MemberClusters = f11
	}
	if resp.ReplicationGroup.MultiAZ != nil {
		ko.Status.MultiAZ = resp.ReplicationGroup.MultiAZ
	}
	if resp.ReplicationGroup.NodeGroups != nil {
		f13 := []*svcapitypes.NodeGroup{}
		for _, f13iter := range resp.ReplicationGroup.NodeGroups {
			f13elem := &svcapitypes.NodeGroup{}
			if f13iter.NodeGroupId != nil {
				f13elem.NodeGroupID = f13iter.NodeGroupId
			}
			if f13iter.NodeGroupMembers != nil {
				f13elemf1 := []*svcapitypes.NodeGroupMember{}
				for _, f13elemf1iter := range f13iter.NodeGroupMembers {
					f13elemf1elem := &svcapitypes.NodeGroupMember{}
					if f13elemf1iter.CacheClusterId != nil {
						f13elemf1elem.CacheClusterID = f13elemf1iter.CacheClusterId
					}
					if f13elemf1iter.CacheNodeId != nil {
						f13elemf1elem.CacheNodeID = f13elemf1iter.CacheNodeId
					}
					if f13elemf1iter.CurrentRole != nil {
						f13elemf1elem.CurrentRole = f13elemf1iter.CurrentRole
					}
					if f13elemf1iter.PreferredAvailabilityZone != nil {
						f13elemf1elem.PreferredAvailabilityZone = f13elemf1iter.PreferredAvailabilityZone
					}
					if f13elemf1iter.ReadEndpoint != nil {
						f13elemf1elemf4 := &svcapitypes.Endpoint{}
						if f13elemf1iter.ReadEndpoint.Address != nil {
							f13elemf1elemf4.Address = f13elemf1iter.ReadEndpoint.Address
						}
						if f13elemf1iter.ReadEndpoint.Port != nil {
							f13elemf1elemf4.Port = f13elemf1iter.ReadEndpoint.Port
						}
						f13elemf1elem.ReadEndpoint = f13elemf1elemf4
					}
					f13elemf1 = append(f13elemf1, f13elemf1elem)
				}
				f13elem.NodeGroupMembers = f13elemf1
			}
			if f13iter.PrimaryEndpoint != nil {
				f13elemf2 := &svcapitypes.Endpoint{}
				if f13iter.PrimaryEndpoint.Address != nil {
					f13elemf2.Address = f13iter.PrimaryEndpoint.Address
				}
				if f13iter.PrimaryEndpoint.Port != nil {
					f13elemf2.Port = f13iter.PrimaryEndpoint.Port
				}
				f13elem.PrimaryEndpoint = f13elemf2
			}
			if f13iter.ReaderEndpoint != nil {
				f13elemf3 := &svcapitypes.Endpoint{}
				if f13iter.ReaderEndpoint.Address != nil {
					f13elemf3.Address = f13iter.ReaderEndpoint.Address
				}
				if f13iter.ReaderEndpoint.Port != nil {
					f13elemf3.Port = f13iter.ReaderEndpoint.Port
				}
				f13elem.ReaderEndpoint = f13elemf3
			}
			if f13iter.Slots != nil {
				f13elem.Slots = f13iter.Slots
			}
			if f13iter.Status != nil {
				f13elem.Status = f13iter.Status
			}
			f13 = append(f13, f13elem)
		}
		ko.Status.NodeGroups = f13
	}
	if resp.ReplicationGroup.PendingModifiedValues != nil {
		f14 := &svcapitypes.ReplicationGroupPendingModifiedValues{}
		if resp.ReplicationGroup.PendingModifiedValues.AuthTokenStatus != nil {
			f14.AuthTokenStatus = resp.ReplicationGroup.PendingModifiedValues.AuthTokenStatus
		}
		if resp.ReplicationGroup.PendingModifiedValues.AutomaticFailoverStatus != nil {
			f14.AutomaticFailoverStatus = resp.ReplicationGroup.PendingModifiedValues.AutomaticFailoverStatus
		}
		if resp.ReplicationGroup.PendingModifiedValues.PrimaryClusterId != nil {
			f14.PrimaryClusterID = resp.ReplicationGroup.PendingModifiedValues.PrimaryClusterId
		}
		if resp.ReplicationGroup.PendingModifiedValues.Resharding != nil {
			f14f3 := &svcapitypes.ReshardingStatus{}
			if resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration != nil {
				f14f3f0 := &svcapitypes.SlotMigration{}
				if resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage != nil {
					f14f3f0.ProgressPercentage = resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage
				}
				f14f3.SlotMigration = f14f3f0
			}
			f14.Resharding = f14f3
		}
		ko.Status.PendingModifiedValues = f14
	}
	if resp.ReplicationGroup.SnapshottingClusterId != nil {
		ko.Status.SnapshottingClusterID = resp.ReplicationGroup.SnapshottingClusterId
	}
	if resp.ReplicationGroup.Status != nil {
		ko.Status.Status = resp.ReplicationGroup.Status
	}

	rm.setStatusDefaults(ko)

	// custom set output from response
	ko, err = rm.CustomCreateReplicationGroupSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateReplicationGroupInput, error) {
	res := &svcsdk.CreateReplicationGroupInput{}

	if r.ko.Spec.AtRestEncryptionEnabled != nil {
		res.SetAtRestEncryptionEnabled(*r.ko.Spec.AtRestEncryptionEnabled)
	}
	if r.ko.Spec.AuthToken != nil {
		res.SetAuthToken(*r.ko.Spec.AuthToken)
	}
	if r.ko.Spec.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	}
	if r.ko.Spec.AutomaticFailoverEnabled != nil {
		res.SetAutomaticFailoverEnabled(*r.ko.Spec.AutomaticFailoverEnabled)
	}
	if r.ko.Spec.CacheNodeType != nil {
		res.SetCacheNodeType(*r.ko.Spec.CacheNodeType)
	}
	if r.ko.Spec.CacheParameterGroupName != nil {
		res.SetCacheParameterGroupName(*r.ko.Spec.CacheParameterGroupName)
	}
	if r.ko.Spec.CacheSecurityGroupNames != nil {
		f6 := []*string{}
		for _, f6iter := range r.ko.Spec.CacheSecurityGroupNames {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		res.SetCacheSecurityGroupNames(f6)
	}
	if r.ko.Spec.CacheSubnetGroupName != nil {
		res.SetCacheSubnetGroupName(*r.ko.Spec.CacheSubnetGroupName)
	}
	if r.ko.Spec.Engine != nil {
		res.SetEngine(*r.ko.Spec.Engine)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.GlobalReplicationGroupID != nil {
		res.SetGlobalReplicationGroupId(*r.ko.Spec.GlobalReplicationGroupID)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.MultiAZEnabled != nil {
		res.SetMultiAZEnabled(*r.ko.Spec.MultiAZEnabled)
	}
	if r.ko.Spec.NodeGroupConfiguration != nil {
		f13 := []*svcsdk.NodeGroupConfiguration{}
		for _, f13iter := range r.ko.Spec.NodeGroupConfiguration {
			f13elem := &svcsdk.NodeGroupConfiguration{}
			if f13iter.NodeGroupID != nil {
				f13elem.SetNodeGroupId(*f13iter.NodeGroupID)
			}
			if f13iter.PrimaryAvailabilityZone != nil {
				f13elem.SetPrimaryAvailabilityZone(*f13iter.PrimaryAvailabilityZone)
			}
			if f13iter.ReplicaAvailabilityZones != nil {
				f13elemf2 := []*string{}
				for _, f13elemf2iter := range f13iter.ReplicaAvailabilityZones {
					var f13elemf2elem string
					f13elemf2elem = *f13elemf2iter
					f13elemf2 = append(f13elemf2, &f13elemf2elem)
				}
				f13elem.SetReplicaAvailabilityZones(f13elemf2)
			}
			if f13iter.ReplicaCount != nil {
				f13elem.SetReplicaCount(*f13iter.ReplicaCount)
			}
			if f13iter.Slots != nil {
				f13elem.SetSlots(*f13iter.Slots)
			}
			f13 = append(f13, f13elem)
		}
		res.SetNodeGroupConfiguration(f13)
	}
	if r.ko.Spec.NotificationTopicARN != nil {
		res.SetNotificationTopicArn(*r.ko.Spec.NotificationTopicARN)
	}
	if r.ko.Spec.NumCacheClusters != nil {
		res.SetNumCacheClusters(*r.ko.Spec.NumCacheClusters)
	}
	if r.ko.Spec.NumNodeGroups != nil {
		res.SetNumNodeGroups(*r.ko.Spec.NumNodeGroups)
	}
	if r.ko.Spec.Port != nil {
		res.SetPort(*r.ko.Spec.Port)
	}
	if r.ko.Spec.PreferredCacheClusterAZs != nil {
		f18 := []*string{}
		for _, f18iter := range r.ko.Spec.PreferredCacheClusterAZs {
			var f18elem string
			f18elem = *f18iter
			f18 = append(f18, &f18elem)
		}
		res.SetPreferredCacheClusterAZs(f18)
	}
	if r.ko.Spec.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*r.ko.Spec.PreferredMaintenanceWindow)
	}
	if r.ko.Spec.PrimaryClusterID != nil {
		res.SetPrimaryClusterId(*r.ko.Spec.PrimaryClusterID)
	}
	if r.ko.Spec.ReplicasPerNodeGroup != nil {
		res.SetReplicasPerNodeGroup(*r.ko.Spec.ReplicasPerNodeGroup)
	}
	if r.ko.Spec.ReplicationGroupDescription != nil {
		res.SetReplicationGroupDescription(*r.ko.Spec.ReplicationGroupDescription)
	}
	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f24 := []*string{}
		for _, f24iter := range r.ko.Spec.SecurityGroupIDs {
			var f24elem string
			f24elem = *f24iter
			f24 = append(f24, &f24elem)
		}
		res.SetSecurityGroupIds(f24)
	}
	if r.ko.Spec.SnapshotARNs != nil {
		f25 := []*string{}
		for _, f25iter := range r.ko.Spec.SnapshotARNs {
			var f25elem string
			f25elem = *f25iter
			f25 = append(f25, &f25elem)
		}
		res.SetSnapshotArns(f25)
	}
	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}
	if r.ko.Spec.SnapshotRetentionLimit != nil {
		res.SetSnapshotRetentionLimit(*r.ko.Spec.SnapshotRetentionLimit)
	}
	if r.ko.Spec.SnapshotWindow != nil {
		res.SetSnapshotWindow(*r.ko.Spec.SnapshotWindow)
	}
	if r.ko.Spec.Tags != nil {
		f29 := []*svcsdk.Tag{}
		for _, f29iter := range r.ko.Spec.Tags {
			f29elem := &svcsdk.Tag{}
			if f29iter.Key != nil {
				f29elem.SetKey(*f29iter.Key)
			}
			if f29iter.Value != nil {
				f29elem.SetValue(*f29iter.Value)
			}
			f29 = append(f29, f29elem)
		}
		res.SetTags(f29)
	}
	if r.ko.Spec.TransitEncryptionEnabled != nil {
		res.SetTransitEncryptionEnabled(*r.ko.Spec.TransitEncryptionEnabled)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	customResp, customRespErr := rm.CustomModifyReplicationGroup(ctx, desired, latest, diffReporter)
	if customResp != nil || customRespErr != nil {
		return customResp, customRespErr
	}

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.ModifyReplicationGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyReplicationGroup", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ReplicationGroup.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ReplicationGroup.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ReplicationGroup.AuthTokenEnabled != nil {
		ko.Status.AuthTokenEnabled = resp.ReplicationGroup.AuthTokenEnabled
	}
	if resp.ReplicationGroup.AuthTokenLastModifiedDate != nil {
		ko.Status.AuthTokenLastModifiedDate = &metav1.Time{*resp.ReplicationGroup.AuthTokenLastModifiedDate}
	}
	if resp.ReplicationGroup.AutomaticFailover != nil {
		ko.Status.AutomaticFailover = resp.ReplicationGroup.AutomaticFailover
	}
	if resp.ReplicationGroup.ClusterEnabled != nil {
		ko.Status.ClusterEnabled = resp.ReplicationGroup.ClusterEnabled
	}
	if resp.ReplicationGroup.ConfigurationEndpoint != nil {
		f7 := &svcapitypes.Endpoint{}
		if resp.ReplicationGroup.ConfigurationEndpoint.Address != nil {
			f7.Address = resp.ReplicationGroup.ConfigurationEndpoint.Address
		}
		if resp.ReplicationGroup.ConfigurationEndpoint.Port != nil {
			f7.Port = resp.ReplicationGroup.ConfigurationEndpoint.Port
		}
		ko.Status.ConfigurationEndpoint = f7
	}
	if resp.ReplicationGroup.Description != nil {
		ko.Status.Description = resp.ReplicationGroup.Description
	}
	if resp.ReplicationGroup.GlobalReplicationGroupInfo != nil {
		f9 := &svcapitypes.GlobalReplicationGroupInfo{}
		if resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupId != nil {
			f9.GlobalReplicationGroupID = resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupId
		}
		if resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole != nil {
			f9.GlobalReplicationGroupMemberRole = resp.ReplicationGroup.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole
		}
		ko.Status.GlobalReplicationGroupInfo = f9
	}
	if resp.ReplicationGroup.MemberClusters != nil {
		f11 := []*string{}
		for _, f11iter := range resp.ReplicationGroup.MemberClusters {
			var f11elem string
			f11elem = *f11iter
			f11 = append(f11, &f11elem)
		}
		ko.Status.MemberClusters = f11
	}
	if resp.ReplicationGroup.MultiAZ != nil {
		ko.Status.MultiAZ = resp.ReplicationGroup.MultiAZ
	}
	if resp.ReplicationGroup.NodeGroups != nil {
		f13 := []*svcapitypes.NodeGroup{}
		for _, f13iter := range resp.ReplicationGroup.NodeGroups {
			f13elem := &svcapitypes.NodeGroup{}
			if f13iter.NodeGroupId != nil {
				f13elem.NodeGroupID = f13iter.NodeGroupId
			}
			if f13iter.NodeGroupMembers != nil {
				f13elemf1 := []*svcapitypes.NodeGroupMember{}
				for _, f13elemf1iter := range f13iter.NodeGroupMembers {
					f13elemf1elem := &svcapitypes.NodeGroupMember{}
					if f13elemf1iter.CacheClusterId != nil {
						f13elemf1elem.CacheClusterID = f13elemf1iter.CacheClusterId
					}
					if f13elemf1iter.CacheNodeId != nil {
						f13elemf1elem.CacheNodeID = f13elemf1iter.CacheNodeId
					}
					if f13elemf1iter.CurrentRole != nil {
						f13elemf1elem.CurrentRole = f13elemf1iter.CurrentRole
					}
					if f13elemf1iter.PreferredAvailabilityZone != nil {
						f13elemf1elem.PreferredAvailabilityZone = f13elemf1iter.PreferredAvailabilityZone
					}
					if f13elemf1iter.ReadEndpoint != nil {
						f13elemf1elemf4 := &svcapitypes.Endpoint{}
						if f13elemf1iter.ReadEndpoint.Address != nil {
							f13elemf1elemf4.Address = f13elemf1iter.ReadEndpoint.Address
						}
						if f13elemf1iter.ReadEndpoint.Port != nil {
							f13elemf1elemf4.Port = f13elemf1iter.ReadEndpoint.Port
						}
						f13elemf1elem.ReadEndpoint = f13elemf1elemf4
					}
					f13elemf1 = append(f13elemf1, f13elemf1elem)
				}
				f13elem.NodeGroupMembers = f13elemf1
			}
			if f13iter.PrimaryEndpoint != nil {
				f13elemf2 := &svcapitypes.Endpoint{}
				if f13iter.PrimaryEndpoint.Address != nil {
					f13elemf2.Address = f13iter.PrimaryEndpoint.Address
				}
				if f13iter.PrimaryEndpoint.Port != nil {
					f13elemf2.Port = f13iter.PrimaryEndpoint.Port
				}
				f13elem.PrimaryEndpoint = f13elemf2
			}
			if f13iter.ReaderEndpoint != nil {
				f13elemf3 := &svcapitypes.Endpoint{}
				if f13iter.ReaderEndpoint.Address != nil {
					f13elemf3.Address = f13iter.ReaderEndpoint.Address
				}
				if f13iter.ReaderEndpoint.Port != nil {
					f13elemf3.Port = f13iter.ReaderEndpoint.Port
				}
				f13elem.ReaderEndpoint = f13elemf3
			}
			if f13iter.Slots != nil {
				f13elem.Slots = f13iter.Slots
			}
			if f13iter.Status != nil {
				f13elem.Status = f13iter.Status
			}
			f13 = append(f13, f13elem)
		}
		ko.Status.NodeGroups = f13
	}
	if resp.ReplicationGroup.PendingModifiedValues != nil {
		f14 := &svcapitypes.ReplicationGroupPendingModifiedValues{}
		if resp.ReplicationGroup.PendingModifiedValues.AuthTokenStatus != nil {
			f14.AuthTokenStatus = resp.ReplicationGroup.PendingModifiedValues.AuthTokenStatus
		}
		if resp.ReplicationGroup.PendingModifiedValues.AutomaticFailoverStatus != nil {
			f14.AutomaticFailoverStatus = resp.ReplicationGroup.PendingModifiedValues.AutomaticFailoverStatus
		}
		if resp.ReplicationGroup.PendingModifiedValues.PrimaryClusterId != nil {
			f14.PrimaryClusterID = resp.ReplicationGroup.PendingModifiedValues.PrimaryClusterId
		}
		if resp.ReplicationGroup.PendingModifiedValues.Resharding != nil {
			f14f3 := &svcapitypes.ReshardingStatus{}
			if resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration != nil {
				f14f3f0 := &svcapitypes.SlotMigration{}
				if resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage != nil {
					f14f3f0.ProgressPercentage = resp.ReplicationGroup.PendingModifiedValues.Resharding.SlotMigration.ProgressPercentage
				}
				f14f3.SlotMigration = f14f3f0
			}
			f14.Resharding = f14f3
		}
		ko.Status.PendingModifiedValues = f14
	}
	if resp.ReplicationGroup.SnapshottingClusterId != nil {
		ko.Status.SnapshottingClusterID = resp.ReplicationGroup.SnapshottingClusterId
	}
	if resp.ReplicationGroup.Status != nil {
		ko.Status.Status = resp.ReplicationGroup.Status
	}

	rm.setStatusDefaults(ko)

	// custom set output from response
	ko, err = rm.CustomModifyReplicationGroupSetOutput(ctx, desired, resp, ko)
	if err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.ModifyReplicationGroupInput, error) {
	res := &svcsdk.ModifyReplicationGroupInput{}

	res.SetApplyImmediately(true)
	if r.ko.Spec.AuthToken != nil {
		res.SetAuthToken(*r.ko.Spec.AuthToken)
	}
	if r.ko.Spec.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	}
	if r.ko.Spec.AutomaticFailoverEnabled != nil {
		res.SetAutomaticFailoverEnabled(*r.ko.Spec.AutomaticFailoverEnabled)
	}
	if r.ko.Spec.CacheNodeType != nil {
		res.SetCacheNodeType(*r.ko.Spec.CacheNodeType)
	}
	if r.ko.Spec.CacheParameterGroupName != nil {
		res.SetCacheParameterGroupName(*r.ko.Spec.CacheParameterGroupName)
	}
	if r.ko.Spec.CacheSecurityGroupNames != nil {
		f7 := []*string{}
		for _, f7iter := range r.ko.Spec.CacheSecurityGroupNames {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		res.SetCacheSecurityGroupNames(f7)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.MultiAZEnabled != nil {
		res.SetMultiAZEnabled(*r.ko.Spec.MultiAZEnabled)
	}
	if r.ko.Spec.NotificationTopicARN != nil {
		res.SetNotificationTopicArn(*r.ko.Spec.NotificationTopicARN)
	}
	if r.ko.Spec.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*r.ko.Spec.PreferredMaintenanceWindow)
	}
	if r.ko.Spec.PrimaryClusterID != nil {
		res.SetPrimaryClusterId(*r.ko.Spec.PrimaryClusterID)
	}
	if r.ko.Spec.ReplicationGroupDescription != nil {
		res.SetReplicationGroupDescription(*r.ko.Spec.ReplicationGroupDescription)
	}
	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f17 := []*string{}
		for _, f17iter := range r.ko.Spec.SecurityGroupIDs {
			var f17elem string
			f17elem = *f17iter
			f17 = append(f17, &f17elem)
		}
		res.SetSecurityGroupIds(f17)
	}
	if r.ko.Spec.SnapshotRetentionLimit != nil {
		res.SetSnapshotRetentionLimit(*r.ko.Spec.SnapshotRetentionLimit)
	}
	if r.ko.Spec.SnapshotWindow != nil {
		res.SetSnapshotWindow(*r.ko.Spec.SnapshotWindow)
	}
	if r.ko.Status.SnapshottingClusterID != nil {
		res.SetSnapshottingClusterId(*r.ko.Status.SnapshottingClusterID)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteReplicationGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteReplicationGroup", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteReplicationGroupInput, error) {
	res := &svcsdk.DeleteReplicationGroupInput{}

	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ReplicationGroup,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	// custom update conditions
	customUpdate := rm.CustomUpdateConditions(ko, r, err)
	if terminalCondition != nil || customUpdate {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameter",
		"InvalidParameterValue",
		"InvalidParameterCombination",
		"InsufficientCacheClusterCapacity",
		"CacheSecurityGroupNotFound",
		"CacheSubnetGroupNotFoundFault",
		"ClusterQuotaForCustomerExceeded",
		"NodeQuotaForClusterExceeded",
		"NodeQuotaForCustomerExceeded",
		"InvalidVPCNetworkStateFault",
		"TagQuotaPerResourceExceeded",
		"NodeGroupsPerReplicationGroupQuotaExceeded",
		"InvalidCacheSecurityGroupState",
		"CacheParameterGroupNotFound",
		"InvalidKMSKeyFault":
		return true
	default:
		return false
	}
}
