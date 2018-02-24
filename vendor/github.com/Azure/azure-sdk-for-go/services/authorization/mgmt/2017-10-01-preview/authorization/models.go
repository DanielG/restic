package authorization

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ClassicAdministrator classic Administrators
type ClassicAdministrator struct {
	// ID - The ID of the administrator.
	ID *string `json:"id,omitempty"`
	// Name - The name of the administrator.
	Name *string `json:"name,omitempty"`
	// Type - The type of the administrator.
	Type *string `json:"type,omitempty"`
	// ClassicAdministratorProperties - Properties for the classic administrator.
	*ClassicAdministratorProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ClassicAdministrator struct.
func (ca *ClassicAdministrator) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		ca.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		ca.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		ca.Type = &typeVar
	}

	v = m["properties"]
	if v != nil {
		var properties ClassicAdministratorProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		ca.ClassicAdministratorProperties = &properties
	}

	return nil
}

// ClassicAdministratorListResult classicAdministrator list result information.
type ClassicAdministratorListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of administrators.
	Value *[]ClassicAdministrator `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ClassicAdministratorListResultIterator provides access to a complete listing of ClassicAdministrator values.
type ClassicAdministratorListResultIterator struct {
	i    int
	page ClassicAdministratorListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ClassicAdministratorListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ClassicAdministratorListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ClassicAdministratorListResultIterator) Response() ClassicAdministratorListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ClassicAdministratorListResultIterator) Value() ClassicAdministrator {
	if !iter.page.NotDone() {
		return ClassicAdministrator{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (calr ClassicAdministratorListResult) IsEmpty() bool {
	return calr.Value == nil || len(*calr.Value) == 0
}

// classicAdministratorListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (calr ClassicAdministratorListResult) classicAdministratorListResultPreparer() (*http.Request, error) {
	if calr.NextLink == nil || len(to.String(calr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(calr.NextLink)))
}

// ClassicAdministratorListResultPage contains a page of ClassicAdministrator values.
type ClassicAdministratorListResultPage struct {
	fn   func(ClassicAdministratorListResult) (ClassicAdministratorListResult, error)
	calr ClassicAdministratorListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ClassicAdministratorListResultPage) Next() error {
	next, err := page.fn(page.calr)
	if err != nil {
		return err
	}
	page.calr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ClassicAdministratorListResultPage) NotDone() bool {
	return !page.calr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ClassicAdministratorListResultPage) Response() ClassicAdministratorListResult {
	return page.calr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ClassicAdministratorListResultPage) Values() []ClassicAdministrator {
	if page.calr.IsEmpty() {
		return nil
	}
	return *page.calr.Value
}

// ClassicAdministratorProperties classic Administrator properties.
type ClassicAdministratorProperties struct {
	// EmailAddress - The email address of the administrator.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Role - The role of the administrator.
	Role *string `json:"role,omitempty"`
}

// Permission role definition permissions.
type Permission struct {
	// Actions - Allowed actions.
	Actions *[]string `json:"actions,omitempty"`
	// NotActions - Denied actions.
	NotActions *[]string `json:"notActions,omitempty"`
}

// PermissionGetResult permissions information.
type PermissionGetResult struct {
	autorest.Response `json:"-"`
	// Value - An array of permissions.
	Value *[]Permission `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// PermissionGetResultIterator provides access to a complete listing of Permission values.
type PermissionGetResultIterator struct {
	i    int
	page PermissionGetResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PermissionGetResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PermissionGetResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PermissionGetResultIterator) Response() PermissionGetResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PermissionGetResultIterator) Value() Permission {
	if !iter.page.NotDone() {
		return Permission{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (pgr PermissionGetResult) IsEmpty() bool {
	return pgr.Value == nil || len(*pgr.Value) == 0
}

// permissionGetResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (pgr PermissionGetResult) permissionGetResultPreparer() (*http.Request, error) {
	if pgr.NextLink == nil || len(to.String(pgr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(pgr.NextLink)))
}

// PermissionGetResultPage contains a page of Permission values.
type PermissionGetResultPage struct {
	fn  func(PermissionGetResult) (PermissionGetResult, error)
	pgr PermissionGetResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PermissionGetResultPage) Next() error {
	next, err := page.fn(page.pgr)
	if err != nil {
		return err
	}
	page.pgr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PermissionGetResultPage) NotDone() bool {
	return !page.pgr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PermissionGetResultPage) Response() PermissionGetResult {
	return page.pgr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PermissionGetResultPage) Values() []Permission {
	if page.pgr.IsEmpty() {
		return nil
	}
	return *page.pgr.Value
}

// ProviderOperation operation
type ProviderOperation struct {
	// Name - The operation name.
	Name *string `json:"name,omitempty"`
	// DisplayName - The operation display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The operation description.
	Description *string `json:"description,omitempty"`
	// Origin - The operation origin.
	Origin *string `json:"origin,omitempty"`
	// Properties - The operation properties.
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// ProviderOperationsMetadata provider Operations metadata
type ProviderOperationsMetadata struct {
	autorest.Response `json:"-"`
	// ID - The provider id.
	ID *string `json:"id,omitempty"`
	// Name - The provider name.
	Name *string `json:"name,omitempty"`
	// Type - The provider type.
	Type *string `json:"type,omitempty"`
	// DisplayName - The provider display name.
	DisplayName *string `json:"displayName,omitempty"`
	// ResourceTypes - The provider resource types
	ResourceTypes *[]ResourceType `json:"resourceTypes,omitempty"`
	// Operations - The provider operations.
	Operations *[]ProviderOperation `json:"operations,omitempty"`
}

// ProviderOperationsMetadataListResult provider operations metadata list
type ProviderOperationsMetadataListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of providers.
	Value *[]ProviderOperationsMetadata `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ProviderOperationsMetadataListResultIterator provides access to a complete listing of ProviderOperationsMetadata
// values.
type ProviderOperationsMetadataListResultIterator struct {
	i    int
	page ProviderOperationsMetadataListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ProviderOperationsMetadataListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ProviderOperationsMetadataListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ProviderOperationsMetadataListResultIterator) Response() ProviderOperationsMetadataListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ProviderOperationsMetadataListResultIterator) Value() ProviderOperationsMetadata {
	if !iter.page.NotDone() {
		return ProviderOperationsMetadata{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (pomlr ProviderOperationsMetadataListResult) IsEmpty() bool {
	return pomlr.Value == nil || len(*pomlr.Value) == 0
}

// providerOperationsMetadataListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (pomlr ProviderOperationsMetadataListResult) providerOperationsMetadataListResultPreparer() (*http.Request, error) {
	if pomlr.NextLink == nil || len(to.String(pomlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(pomlr.NextLink)))
}

// ProviderOperationsMetadataListResultPage contains a page of ProviderOperationsMetadata values.
type ProviderOperationsMetadataListResultPage struct {
	fn    func(ProviderOperationsMetadataListResult) (ProviderOperationsMetadataListResult, error)
	pomlr ProviderOperationsMetadataListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ProviderOperationsMetadataListResultPage) Next() error {
	next, err := page.fn(page.pomlr)
	if err != nil {
		return err
	}
	page.pomlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ProviderOperationsMetadataListResultPage) NotDone() bool {
	return !page.pomlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ProviderOperationsMetadataListResultPage) Response() ProviderOperationsMetadataListResult {
	return page.pomlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ProviderOperationsMetadataListResultPage) Values() []ProviderOperationsMetadata {
	if page.pomlr.IsEmpty() {
		return nil
	}
	return *page.pomlr.Value
}

// ResourceType resource Type
type ResourceType struct {
	// Name - The resource type name.
	Name *string `json:"name,omitempty"`
	// DisplayName - The resource type display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Operations - The resource type operations.
	Operations *[]ProviderOperation `json:"operations,omitempty"`
}

// RoleAssignment role Assignments
type RoleAssignment struct {
	autorest.Response `json:"-"`
	// ID - The role assignment ID.
	ID *string `json:"id,omitempty"`
	// Name - The role assignment name.
	Name *string `json:"name,omitempty"`
	// Type - The role assignment type.
	Type *string `json:"type,omitempty"`
	// RoleAssignmentPropertiesWithScope - Role assignment properties.
	*RoleAssignmentPropertiesWithScope `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for RoleAssignment struct.
func (ra *RoleAssignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		ra.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		ra.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		ra.Type = &typeVar
	}

	v = m["properties"]
	if v != nil {
		var properties RoleAssignmentPropertiesWithScope
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		ra.RoleAssignmentPropertiesWithScope = &properties
	}

	return nil
}

// RoleAssignmentCreateParameters role assignment create parameters.
type RoleAssignmentCreateParameters struct {
	// RoleAssignmentProperties - Role assignment properties.
	*RoleAssignmentProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for RoleAssignmentCreateParameters struct.
func (racp *RoleAssignmentCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties RoleAssignmentProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		racp.RoleAssignmentProperties = &properties
	}

	return nil
}

// RoleAssignmentFilter role Assignments filter
type RoleAssignmentFilter struct {
	// PrincipalID - Returns role assignment of the specific principal.
	PrincipalID *string `json:"principalId,omitempty"`
	// CanDelegate - The Delegation flag for the roleassignment
	CanDelegate *bool `json:"canDelegate,omitempty"`
}

// RoleAssignmentListResult role assignment list operation result.
type RoleAssignmentListResult struct {
	autorest.Response `json:"-"`
	// Value - Role assignment list.
	Value *[]RoleAssignment `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// RoleAssignmentListResultIterator provides access to a complete listing of RoleAssignment values.
type RoleAssignmentListResultIterator struct {
	i    int
	page RoleAssignmentListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RoleAssignmentListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RoleAssignmentListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RoleAssignmentListResultIterator) Response() RoleAssignmentListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RoleAssignmentListResultIterator) Value() RoleAssignment {
	if !iter.page.NotDone() {
		return RoleAssignment{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ralr RoleAssignmentListResult) IsEmpty() bool {
	return ralr.Value == nil || len(*ralr.Value) == 0
}

// roleAssignmentListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ralr RoleAssignmentListResult) roleAssignmentListResultPreparer() (*http.Request, error) {
	if ralr.NextLink == nil || len(to.String(ralr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ralr.NextLink)))
}

// RoleAssignmentListResultPage contains a page of RoleAssignment values.
type RoleAssignmentListResultPage struct {
	fn   func(RoleAssignmentListResult) (RoleAssignmentListResult, error)
	ralr RoleAssignmentListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RoleAssignmentListResultPage) Next() error {
	next, err := page.fn(page.ralr)
	if err != nil {
		return err
	}
	page.ralr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RoleAssignmentListResultPage) NotDone() bool {
	return !page.ralr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RoleAssignmentListResultPage) Response() RoleAssignmentListResult {
	return page.ralr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RoleAssignmentListResultPage) Values() []RoleAssignment {
	if page.ralr.IsEmpty() {
		return nil
	}
	return *page.ralr.Value
}

// RoleAssignmentProperties role assignment properties.
type RoleAssignmentProperties struct {
	// RoleDefinitionID - The role definition ID used in the role assignment.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// PrincipalID - The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
	PrincipalID *string `json:"principalId,omitempty"`
	// CanDelegate - The delgation flag used for creating a role assignment
	CanDelegate *bool `json:"canDelegate,omitempty"`
}

// RoleAssignmentPropertiesWithScope role assignment properties with scope.
type RoleAssignmentPropertiesWithScope struct {
	// Scope - The role assignment scope.
	Scope *string `json:"scope,omitempty"`
	// RoleDefinitionID - The role definition ID.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// PrincipalID - The principal ID.
	PrincipalID *string `json:"principalId,omitempty"`
	// CanDelegate - The Delegation flag for the roleassignment
	CanDelegate *bool `json:"canDelegate,omitempty"`
}

// RoleDefinition role definition.
type RoleDefinition struct {
	autorest.Response `json:"-"`
	// ID - The role definition ID.
	ID *string `json:"id,omitempty"`
	// Name - The role definition name.
	Name *string `json:"name,omitempty"`
	// Type - The role definition type.
	Type *string `json:"type,omitempty"`
	// RoleDefinitionProperties - Role definition properties.
	*RoleDefinitionProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for RoleDefinition struct.
func (rd *RoleDefinition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		rd.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		rd.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		rd.Type = &typeVar
	}

	v = m["properties"]
	if v != nil {
		var properties RoleDefinitionProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		rd.RoleDefinitionProperties = &properties
	}

	return nil
}

// RoleDefinitionFilter role Definitions filter
type RoleDefinitionFilter struct {
	// RoleName - Returns role definition with the specific name.
	RoleName *string `json:"roleName,omitempty"`
}

// RoleDefinitionListResult role definition list operation result.
type RoleDefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - Role definition list.
	Value *[]RoleDefinition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// RoleDefinitionListResultIterator provides access to a complete listing of RoleDefinition values.
type RoleDefinitionListResultIterator struct {
	i    int
	page RoleDefinitionListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RoleDefinitionListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RoleDefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RoleDefinitionListResultIterator) Response() RoleDefinitionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RoleDefinitionListResultIterator) Value() RoleDefinition {
	if !iter.page.NotDone() {
		return RoleDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (rdlr RoleDefinitionListResult) IsEmpty() bool {
	return rdlr.Value == nil || len(*rdlr.Value) == 0
}

// roleDefinitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rdlr RoleDefinitionListResult) roleDefinitionListResultPreparer() (*http.Request, error) {
	if rdlr.NextLink == nil || len(to.String(rdlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rdlr.NextLink)))
}

// RoleDefinitionListResultPage contains a page of RoleDefinition values.
type RoleDefinitionListResultPage struct {
	fn   func(RoleDefinitionListResult) (RoleDefinitionListResult, error)
	rdlr RoleDefinitionListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RoleDefinitionListResultPage) Next() error {
	next, err := page.fn(page.rdlr)
	if err != nil {
		return err
	}
	page.rdlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RoleDefinitionListResultPage) NotDone() bool {
	return !page.rdlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RoleDefinitionListResultPage) Response() RoleDefinitionListResult {
	return page.rdlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RoleDefinitionListResultPage) Values() []RoleDefinition {
	if page.rdlr.IsEmpty() {
		return nil
	}
	return *page.rdlr.Value
}

// RoleDefinitionProperties role definition properties.
type RoleDefinitionProperties struct {
	// RoleName - The role name.
	RoleName *string `json:"roleName,omitempty"`
	// Description - The role definition description.
	Description *string `json:"description,omitempty"`
	// RoleType - The role type.
	RoleType *string `json:"type,omitempty"`
	// Permissions - Role definition permissions.
	Permissions *[]Permission `json:"permissions,omitempty"`
	// AssignableScopes - Role definition assignable scopes.
	AssignableScopes *[]string `json:"assignableScopes,omitempty"`
}