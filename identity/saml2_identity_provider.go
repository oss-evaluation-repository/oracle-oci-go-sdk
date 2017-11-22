// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"encoding/json"
)

// Saml2IdentityProvider. A special type of IdentityProvider that
// supports the SAML 2.0 protocol. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
type Saml2IdentityProvider struct {

	// The OCID of the `IdentityProvider`.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation. The name
	// must be unique across all `IdentityProvider` objects in the tenancy and
	// cannot be changed. This is the name federated users see when choosing
	// which identity provider to use when signing in to the Oracle Cloud Infrastructure
	// Console.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation. Does
	// not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Allowed values are:
	// - `ADFS`
	// - `IDCS`
	// Example: `IDCS`
	ProductType *string `mandatory:"true" json:"productType,omitempty"`

	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The URL for retrieving the identity provider's metadata, which
	// contains information required for federating.
	MetadataURL *string `mandatory:"true" json:"metadataUrl,omitempty"`

	// The identity provider's signing certificate used by the IAM Service
	// to validate the SAML2 token.
	SigningCertificate *string `mandatory:"true" json:"signingCertificate,omitempty"`

	// The URL to redirect federated users to for authentication with the
	// identity provider.
	RedirectURL *string `mandatory:"true" json:"redirectUrl,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`

	// The current state. After creating an `IdentityProvider`, make sure its
	// `lifecycleState` changes from CREATING to ACTIVE before using it.
	LifecycleState IdentityProviderLifecycleStateEnum
}

func (model Saml2IdentityProvider) GetID() *string {
	return model.ID
}
func (model Saml2IdentityProvider) GetCompartmentID() *string {
	return model.CompartmentID
}
func (model Saml2IdentityProvider) GetName() *string {
	return model.Name
}
func (model Saml2IdentityProvider) GetDescription() *string {
	return model.Description
}
func (model Saml2IdentityProvider) GetProductType() *string {
	return model.ProductType
}
func (model Saml2IdentityProvider) GetTimeCreated() *common.SDKTime {
	return model.TimeCreated
}
func (model Saml2IdentityProvider) GetLifecycleState() IdentityProviderLifecycleStateEnum {
	return model.LifecycleState
}
func (model Saml2IdentityProvider) GetInactiveStatus() *int {
	return model.InactiveStatus
}

func (m Saml2IdentityProvider) String() string {
	return common.PointerString(m)
}

func (m Saml2IdentityProvider) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSaml2IdentityProvider Saml2IdentityProvider
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeSaml2IdentityProvider
	}{
		"SAML2",
		(MarshalTypeSaml2IdentityProvider)(m),
	}

	return json.Marshal(&s)
}
