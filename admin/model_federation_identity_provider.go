// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// FederationIdentityProvider struct for FederationIdentityProvider
type FederationIdentityProvider struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains *[]string `json:"associatedDomains,omitempty"`
	// List that contains the connected organization configurations associated with the identity provider.
	AssociatedOrgs *[]ConnectedOrgConfig `json:"associatedOrgs,omitempty"`
	// The description of the identity provider.
	Description *string `json:"description,omitempty"`
	// Human-readable label that identifies the identity provider.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the identity provider.
	// Read only field.
	Id string `json:"id"`
	// Unique string that identifies the issuer of the SAML Assertion.
	IssuerUri *string `json:"issuerUri,omitempty"`
	// Unique 20-hexadecimal digit string that identifies the identity provider.
	OktaIdpId string `json:"oktaIdpId"`
	// The protocol of the identity provider.
	Protocol *string `json:"protocol,omitempty"`
	// Identifier of the intended recipient of the token.
	AudienceClaim *[]string `json:"audienceClaim,omitempty"`
	// Client identifier that is assigned to an application by the Identity Provider.
	ClientId *string `json:"clientId,omitempty"`
	// Identifier of the claim which contains IdP Group IDs in the token.
	GroupsClaim *string `json:"groupsClaim,omitempty"`
	// The description of an IdP, configured by customers.
	RequestedScopes *[]string `json:"requestedScopes,omitempty"`
	// Identifier of the claim which contains the user ID in the token.
	UserClaim *string `json:"userClaim,omitempty"`
	// URL that points to where to send the SAML response.
	AcsUrl *string `json:"acsUrl,omitempty"`
	// Unique string that identifies the intended audience of the SAML assertion.
	AudienceUri *string      `json:"audienceUri,omitempty"`
	PemFileInfo *PemFileInfo `json:"pemFileInfo,omitempty"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.
	RequestBinding *string `json:"requestBinding,omitempty"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty"`
	// URL that points to the receiver of the SAML authentication request.
	SsoUrl *string `json:"ssoUrl,omitempty"`
	// String enum that indicates whether the identity provider is active.
	Status *string `json:"status,omitempty"`
}

// NewFederationIdentityProvider instantiates a new FederationIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationIdentityProvider(id string, oktaIdpId string) *FederationIdentityProvider {
	this := FederationIdentityProvider{}
	this.Id = id
	this.OktaIdpId = oktaIdpId
	return &this
}

// NewFederationIdentityProviderWithDefaults instantiates a new FederationIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationIdentityProviderWithDefaults() *FederationIdentityProvider {
	this := FederationIdentityProvider{}
	return &this
}

// GetAssociatedDomains returns the AssociatedDomains field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetAssociatedDomains() []string {
	if o == nil || IsNil(o.AssociatedDomains) {
		var ret []string
		return ret
	}
	return *o.AssociatedDomains
}

// GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetAssociatedDomainsOk() (*[]string, bool) {
	if o == nil || IsNil(o.AssociatedDomains) {
		return nil, false
	}

	return o.AssociatedDomains, true
}

// HasAssociatedDomains returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasAssociatedDomains() bool {
	if o != nil && !IsNil(o.AssociatedDomains) {
		return true
	}

	return false
}

// SetAssociatedDomains gets a reference to the given []string and assigns it to the AssociatedDomains field.
func (o *FederationIdentityProvider) SetAssociatedDomains(v []string) {
	o.AssociatedDomains = &v
}

// GetAssociatedOrgs returns the AssociatedOrgs field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetAssociatedOrgs() []ConnectedOrgConfig {
	if o == nil || IsNil(o.AssociatedOrgs) {
		var ret []ConnectedOrgConfig
		return ret
	}
	return *o.AssociatedOrgs
}

// GetAssociatedOrgsOk returns a tuple with the AssociatedOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetAssociatedOrgsOk() (*[]ConnectedOrgConfig, bool) {
	if o == nil || IsNil(o.AssociatedOrgs) {
		return nil, false
	}

	return o.AssociatedOrgs, true
}

// HasAssociatedOrgs returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasAssociatedOrgs() bool {
	if o != nil && !IsNil(o.AssociatedOrgs) {
		return true
	}

	return false
}

// SetAssociatedOrgs gets a reference to the given []ConnectedOrgConfig and assigns it to the AssociatedOrgs field.
func (o *FederationIdentityProvider) SetAssociatedOrgs(v []ConnectedOrgConfig) {
	o.AssociatedOrgs = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FederationIdentityProvider) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}

	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FederationIdentityProvider) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value
func (o *FederationIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FederationIdentityProvider) SetId(v string) {
	o.Id = v
}

// GetIssuerUri returns the IssuerUri field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetIssuerUri() string {
	if o == nil || IsNil(o.IssuerUri) {
		var ret string
		return ret
	}
	return *o.IssuerUri
}

// GetIssuerUriOk returns a tuple with the IssuerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetIssuerUriOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerUri) {
		return nil, false
	}

	return o.IssuerUri, true
}

// HasIssuerUri returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasIssuerUri() bool {
	if o != nil && !IsNil(o.IssuerUri) {
		return true
	}

	return false
}

// SetIssuerUri gets a reference to the given string and assigns it to the IssuerUri field.
func (o *FederationIdentityProvider) SetIssuerUri(v string) {
	o.IssuerUri = &v
}

// GetOktaIdpId returns the OktaIdpId field value
func (o *FederationIdentityProvider) GetOktaIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaIdpId
}

// GetOktaIdpIdOk returns a tuple with the OktaIdpId field value
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetOktaIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaIdpId, true
}

// SetOktaIdpId sets field value
func (o *FederationIdentityProvider) SetOktaIdpId(v string) {
	o.OktaIdpId = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}

	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FederationIdentityProvider) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAudienceClaim returns the AudienceClaim field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetAudienceClaim() []string {
	if o == nil || IsNil(o.AudienceClaim) {
		var ret []string
		return ret
	}
	return *o.AudienceClaim
}

// GetAudienceClaimOk returns a tuple with the AudienceClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetAudienceClaimOk() (*[]string, bool) {
	if o == nil || IsNil(o.AudienceClaim) {
		return nil, false
	}

	return o.AudienceClaim, true
}

// HasAudienceClaim returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasAudienceClaim() bool {
	if o != nil && !IsNil(o.AudienceClaim) {
		return true
	}

	return false
}

// SetAudienceClaim gets a reference to the given []string and assigns it to the AudienceClaim field.
func (o *FederationIdentityProvider) SetAudienceClaim(v []string) {
	o.AudienceClaim = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *FederationIdentityProvider) SetClientId(v string) {
	o.ClientId = &v
}

// GetGroupsClaim returns the GroupsClaim field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetGroupsClaim() string {
	if o == nil || IsNil(o.GroupsClaim) {
		var ret string
		return ret
	}
	return *o.GroupsClaim
}

// GetGroupsClaimOk returns a tuple with the GroupsClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetGroupsClaimOk() (*string, bool) {
	if o == nil || IsNil(o.GroupsClaim) {
		return nil, false
	}

	return o.GroupsClaim, true
}

// HasGroupsClaim returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasGroupsClaim() bool {
	if o != nil && !IsNil(o.GroupsClaim) {
		return true
	}

	return false
}

// SetGroupsClaim gets a reference to the given string and assigns it to the GroupsClaim field.
func (o *FederationIdentityProvider) SetGroupsClaim(v string) {
	o.GroupsClaim = &v
}

// GetRequestedScopes returns the RequestedScopes field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetRequestedScopes() []string {
	if o == nil || IsNil(o.RequestedScopes) {
		var ret []string
		return ret
	}
	return *o.RequestedScopes
}

// GetRequestedScopesOk returns a tuple with the RequestedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetRequestedScopesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RequestedScopes) {
		return nil, false
	}

	return o.RequestedScopes, true
}

// HasRequestedScopes returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasRequestedScopes() bool {
	if o != nil && !IsNil(o.RequestedScopes) {
		return true
	}

	return false
}

// SetRequestedScopes gets a reference to the given []string and assigns it to the RequestedScopes field.
func (o *FederationIdentityProvider) SetRequestedScopes(v []string) {
	o.RequestedScopes = &v
}

// GetUserClaim returns the UserClaim field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetUserClaim() string {
	if o == nil || IsNil(o.UserClaim) {
		var ret string
		return ret
	}
	return *o.UserClaim
}

// GetUserClaimOk returns a tuple with the UserClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetUserClaimOk() (*string, bool) {
	if o == nil || IsNil(o.UserClaim) {
		return nil, false
	}

	return o.UserClaim, true
}

// HasUserClaim returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasUserClaim() bool {
	if o != nil && !IsNil(o.UserClaim) {
		return true
	}

	return false
}

// SetUserClaim gets a reference to the given string and assigns it to the UserClaim field.
func (o *FederationIdentityProvider) SetUserClaim(v string) {
	o.UserClaim = &v
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetAcsUrl() string {
	if o == nil || IsNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetAcsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcsUrl) {
		return nil, false
	}

	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasAcsUrl() bool {
	if o != nil && !IsNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *FederationIdentityProvider) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAudienceUri returns the AudienceUri field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetAudienceUri() string {
	if o == nil || IsNil(o.AudienceUri) {
		var ret string
		return ret
	}
	return *o.AudienceUri
}

// GetAudienceUriOk returns a tuple with the AudienceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetAudienceUriOk() (*string, bool) {
	if o == nil || IsNil(o.AudienceUri) {
		return nil, false
	}

	return o.AudienceUri, true
}

// HasAudienceUri returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasAudienceUri() bool {
	if o != nil && !IsNil(o.AudienceUri) {
		return true
	}

	return false
}

// SetAudienceUri gets a reference to the given string and assigns it to the AudienceUri field.
func (o *FederationIdentityProvider) SetAudienceUri(v string) {
	o.AudienceUri = &v
}

// GetPemFileInfo returns the PemFileInfo field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetPemFileInfo() PemFileInfo {
	if o == nil || IsNil(o.PemFileInfo) {
		var ret PemFileInfo
		return ret
	}
	return *o.PemFileInfo
}

// GetPemFileInfoOk returns a tuple with the PemFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetPemFileInfoOk() (*PemFileInfo, bool) {
	if o == nil || IsNil(o.PemFileInfo) {
		return nil, false
	}

	return o.PemFileInfo, true
}

// HasPemFileInfo returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasPemFileInfo() bool {
	if o != nil && !IsNil(o.PemFileInfo) {
		return true
	}

	return false
}

// SetPemFileInfo gets a reference to the given PemFileInfo and assigns it to the PemFileInfo field.
func (o *FederationIdentityProvider) SetPemFileInfo(v PemFileInfo) {
	o.PemFileInfo = &v
}

// GetRequestBinding returns the RequestBinding field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetRequestBinding() string {
	if o == nil || IsNil(o.RequestBinding) {
		var ret string
		return ret
	}
	return *o.RequestBinding
}

// GetRequestBindingOk returns a tuple with the RequestBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetRequestBindingOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBinding) {
		return nil, false
	}

	return o.RequestBinding, true
}

// HasRequestBinding returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasRequestBinding() bool {
	if o != nil && !IsNil(o.RequestBinding) {
		return true
	}

	return false
}

// SetRequestBinding gets a reference to the given string and assigns it to the RequestBinding field.
func (o *FederationIdentityProvider) SetRequestBinding(v string) {
	o.RequestBinding = &v
}

// GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetResponseSignatureAlgorithm() string {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.ResponseSignatureAlgorithm
}

// GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetResponseSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		return nil, false
	}

	return o.ResponseSignatureAlgorithm, true
}

// HasResponseSignatureAlgorithm returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasResponseSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.ResponseSignatureAlgorithm) {
		return true
	}

	return false
}

// SetResponseSignatureAlgorithm gets a reference to the given string and assigns it to the ResponseSignatureAlgorithm field.
func (o *FederationIdentityProvider) SetResponseSignatureAlgorithm(v string) {
	o.ResponseSignatureAlgorithm = &v
}

// GetSsoDebugEnabled returns the SsoDebugEnabled field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetSsoDebugEnabled() bool {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		var ret bool
		return ret
	}
	return *o.SsoDebugEnabled
}

// GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetSsoDebugEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		return nil, false
	}

	return o.SsoDebugEnabled, true
}

// HasSsoDebugEnabled returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasSsoDebugEnabled() bool {
	if o != nil && !IsNil(o.SsoDebugEnabled) {
		return true
	}

	return false
}

// SetSsoDebugEnabled gets a reference to the given bool and assigns it to the SsoDebugEnabled field.
func (o *FederationIdentityProvider) SetSsoDebugEnabled(v bool) {
	o.SsoDebugEnabled = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}

	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *FederationIdentityProvider) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *FederationIdentityProvider) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationIdentityProvider) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FederationIdentityProvider) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FederationIdentityProvider) SetStatus(v string) {
	o.Status = &v
}

func (o FederationIdentityProvider) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o FederationIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedDomains) {
		toSerialize["associatedDomains"] = o.AssociatedDomains
	}
	if !IsNil(o.AssociatedOrgs) {
		toSerialize["associatedOrgs"] = o.AssociatedOrgs
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IssuerUri) {
		toSerialize["issuerUri"] = o.IssuerUri
	}
	toSerialize["oktaIdpId"] = o.OktaIdpId
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.AudienceClaim) {
		toSerialize["audienceClaim"] = o.AudienceClaim
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.GroupsClaim) {
		toSerialize["groupsClaim"] = o.GroupsClaim
	}
	if !IsNil(o.RequestedScopes) {
		toSerialize["requestedScopes"] = o.RequestedScopes
	}
	if !IsNil(o.UserClaim) {
		toSerialize["userClaim"] = o.UserClaim
	}
	if !IsNil(o.AcsUrl) {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if !IsNil(o.AudienceUri) {
		toSerialize["audienceUri"] = o.AudienceUri
	}
	if !IsNil(o.PemFileInfo) {
		toSerialize["pemFileInfo"] = o.PemFileInfo
	}
	if !IsNil(o.RequestBinding) {
		toSerialize["requestBinding"] = o.RequestBinding
	}
	if !IsNil(o.ResponseSignatureAlgorithm) {
		toSerialize["responseSignatureAlgorithm"] = o.ResponseSignatureAlgorithm
	}
	if !IsNil(o.SsoDebugEnabled) {
		toSerialize["ssoDebugEnabled"] = o.SsoDebugEnabled
	}
	if !IsNil(o.SsoUrl) {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}
