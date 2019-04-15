package model

// FileShareSpec is a schema for fileshare API. Fileshare will be created on some backend
// and can be shared among multiple users.

type FileShareSpec struct {
	*BaseModel

	// The uuid of the project that the fileshare belongs to.
	TenantId string `json:"tenantId,omitempty"`

	// The uuid of the user that the fileshare belongs to.
	UserId string `json:"userId,omitempty"`

	// The name of the fileshare.
	Name string `json:"name,omitempty"`

	// Creation time of fileshare.
	CreatedAt string `json:"createdAt,omitempty"`

	// Updation time of fileshare.
	UpdatedAt string `json:"updatedAt,omitempty"`

	// The protocol of the fileshare. e.g NFS, SMB etc.
	Protocols []string `json:"protocols,omitempty"`

	// The description of the fileshare.
	// +optional
	Description string `json:"description,omitempty"`

	// The size of the fileshare requested by the user.
	// Default unit of fileshare Size is GB.
	Size int64 `json:"size,omitempty"`

	// The locality that fileshare belongs to.
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// The status of the fileshare.
	// One of: "available", "error" etc.
	Status string `json:"status,omitempty"`

	// The uuid of the pool which the fileshare belongs to.
	// +readOnly
	PoolId string `json:"poolId,omitempty"`

	// The uuid of the profile which the fileshare belongs to.
	ProfileId string `json:"profileId,omitempty"`

	// The uuid of the snapshot which the fileshare is created
	SnapshotId string `json:"snapshotId,omitempty"`

	// ExportLocations of the fileshare.
	ExportLocations []string `json:"exportLocations,omitempty"`
}

// FileShareSnapshotSpec is a description of fileshare snapshot resource.
type FileShareSnapshotSpec struct {
	*BaseModel

	// The uuid of the project that the fileshare snapshot belongs to.
	TenantId string `json:"tenantId,omitempty"`

	// The uuid of the user that the fileshare snapshot belongs to.
	// +optional
	UserId string `json:"userId,omitempty"`

	// The name of the fileshare snapshot.
	Name string `json:"name,omitempty"`

	// Creation time of snapshot.
	CreatedAt string `json:"createdAt,omitempty"`

	// Updation time of snapshot.
	UpdatedAt string `json:"updatedAt,omitempty"`

	// The protocol of the fileshare. e.g NFS, SMB etc.
	Protocols []string `json:"protocols,omitempty"`

	// The description of the fileshare snapshot.
	// +optional
	Description string `json:"description,omitempty"`

	// The size of the fileshare which the snapshot belongs to.
	// Default unit of filesahre Size is GB.
	ShareSize int64 `json:"shareSize,omitempty"`

	// The size of the snapshot. Default unit of files snapshot Size is GB.
	SnapshotSize int64 `json:"snapshotSize,omitempty"`

	// The status of the fileshare snapshot.
	// One of: "available", "error", etc.
	Status string `json:"status,omitempty"`

	// The uuid of the volume which the snapshot belongs to.
	SnapshotId string `json:"snapshotId,omitempty"`
}
