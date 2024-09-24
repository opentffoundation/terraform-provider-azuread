package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRelatedResource = NetworkaccessRelatedTenant{}

type NetworkaccessRelatedTenant struct {
	TenantId *string `json:"tenantId,omitempty"`

	// Fields inherited from NetworkaccessRelatedResource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessRelatedTenant) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return BaseNetworkaccessRelatedResourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRelatedTenant{}

func (s NetworkaccessRelatedTenant) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRelatedTenant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRelatedTenant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedTenant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.relatedTenant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRelatedTenant: %+v", err)
	}

	return encoded, nil
}