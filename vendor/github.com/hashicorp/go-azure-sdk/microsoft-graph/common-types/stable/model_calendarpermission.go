package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CalendarPermission{}

type CalendarPermission struct {
	// List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead,
	// limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
	AllowedRoles *[]CalendarRoleType `json:"allowedRoles,omitempty"`

	// Represents a share recipient or delegate who has access to the calendar. For the 'My Organization' share recipient,
	// the address property is null. Read-only.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`

	// True if the user in context (recipient or delegate) is inside the same organization as the calendar owner.
	IsInsideOrganization nullable.Type[bool] `json:"isInsideOrganization,omitempty"`

	// True if the user can be removed from the list of recipients or delegates for the specified calendar, false otherwise.
	// The 'My organization' user determines the permissions other people within your organization have to the given
	// calendar. You can't remove 'My organization' as a share recipient to a calendar.
	IsRemovable nullable.Type[bool] `json:"isRemovable,omitempty"`

	// Current permission level of the calendar share recipient or delegate.
	Role *CalendarRoleType `json:"role,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CalendarPermission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CalendarPermission{}

func (s CalendarPermission) MarshalJSON() ([]byte, error) {
	type wrapper CalendarPermission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CalendarPermission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CalendarPermission: %+v", err)
	}

	delete(decoded, "emailAddress")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.calendarPermission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CalendarPermission: %+v", err)
	}

	return encoded, nil
}