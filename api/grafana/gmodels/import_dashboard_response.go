// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImportDashboardResponse ImportDashboardResponse response object returned when importing a dashboard.
//
// swagger:model ImportDashboardResponse
type ImportDashboardResponse struct {
	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// folder Id
	FolderID int64 `json:"folderId,omitempty"`

	// imported
	Imported bool `json:"imported,omitempty"`

	// imported revision
	ImportedRevision int64 `json:"importedRevision,omitempty"`

	// imported Uri
	ImportedURI string `json:"importedUri,omitempty"`

	// imported Url
	ImportedURL string `json:"importedUrl,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// plugin Id
	PluginID string `json:"pluginId,omitempty"`

	// removed
	Removed bool `json:"removed,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// UID
	UID string `json:"uid,omitempty"`
}

// Validate validates this import dashboard response
func (m *ImportDashboardResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this import dashboard response based on context it is used
func (m *ImportDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportDashboardResponse) UnmarshalBinary(b []byte) error {
	var res ImportDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
