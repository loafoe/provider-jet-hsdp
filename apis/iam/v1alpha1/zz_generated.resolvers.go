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
	rconfig "github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Application.
func (mg *Application) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PropositionID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PropositionRef,
		Selector:     mg.Spec.ForProvider.PropositionIDSelector,
		To: reference.To{
			List:    &PropositionList{},
			Managed: &Proposition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PropositionID")
	}
	mg.Spec.ForProvider.PropositionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PropositionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagingOrganization),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.ManagingOrganizationSelector,
		To: reference.To{
			List:    &OrgList{},
			Managed: &Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagingOrganization")
	}
	mg.Spec.ForProvider.ManagingOrganization = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Proposition.
func (mg *Proposition) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrganizationID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationIDSelector,
		To: reference.To{
			List:    &OrgList{},
			Managed: &Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrganizationID")
	}
	mg.Spec.ForProvider.OrganizationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Role.
func (mg *Role) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagingOrganization),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.ManagingOrganizationSelector,
		To: reference.To{
			List:    &OrgList{},
			Managed: &Org{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagingOrganization")
	}
	mg.Spec.ForProvider.ManagingOrganization = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationRef = rsp.ResolvedReference

	return nil
}
