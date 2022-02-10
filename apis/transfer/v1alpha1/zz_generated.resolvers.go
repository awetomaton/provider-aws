/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta11 "github.com/crossplane/provider-aws/apis/acm/v1beta1"
	v1alpha1 "github.com/crossplane/provider-aws/apis/ec2/v1alpha1"
	v1beta1 "github.com/crossplane/provider-aws/apis/ec2/v1beta1"
	v1beta12 "github.com/crossplane/provider-aws/apis/iam/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails != nil {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDs),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDRefs,
			Selector:      mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDSelector,
			To: reference.To{
				List:    &v1beta1.AddressList{},
				Managed: &v1beta1.Address{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDs")
		}
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDs = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.AddressAllocationIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails != nil {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDs),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDSelector,
			To: reference.To{
				List:    &v1beta1.SecurityGroupList{},
				Managed: &v1beta1.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDs")
		}
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDs = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails != nil {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDs),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDs")
		}
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDs = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.SubnetIDRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointIDRef,
			Selector:     mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointIDSelector,
			To: reference.To{
				List:    &v1alpha1.VPCEndpointList{},
				Managed: &v1alpha1.VPCEndpoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointID")
		}
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCEndpointIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCIDRef,
			Selector:     mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCIDSelector,
			To: reference.To{
				List:    &v1beta1.VPCList{},
				Managed: &v1beta1.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCID")
		}
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomServerParameters.CustomEndpointDetails.VPCIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomServerParameters.Certificate),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomServerParameters.CertificateRef,
		Selector:     mg.Spec.ForProvider.CustomServerParameters.CertificateSelector,
		To: reference.To{
			List:    &v1beta11.CertificateList{},
			Managed: &v1beta11.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.Certificate")
	}
	mg.Spec.ForProvider.CustomServerParameters.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomServerParameters.CertificateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomServerParameters.LoggingRole),
		Extract:      v1beta12.RoleARN(),
		Reference:    mg.Spec.ForProvider.CustomServerParameters.LoggingRoleRef,
		Selector:     mg.Spec.ForProvider.CustomServerParameters.LoggingRoleSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomServerParameters.LoggingRole")
	}
	mg.Spec.ForProvider.CustomServerParameters.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomServerParameters.LoggingRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomUserParameters.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomUserParameters.ServerIDRef,
		Selector:     mg.Spec.ForProvider.CustomUserParameters.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomUserParameters.ServerID")
	}
	mg.Spec.ForProvider.CustomUserParameters.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomUserParameters.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomUserParameters.Role),
		Extract:      v1beta12.RoleARN(),
		Reference:    mg.Spec.ForProvider.CustomUserParameters.RoleRef,
		Selector:     mg.Spec.ForProvider.CustomUserParameters.RoleSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomUserParameters.Role")
	}
	mg.Spec.ForProvider.CustomUserParameters.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomUserParameters.RoleRef = rsp.ResolvedReference

	return nil
}