package documentdb

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)    

// DatabaseAccountKind enumerates the values for database account kind.
type DatabaseAccountKind string

const (
    // GlobalDocumentDB specifies the global document db state for database
    // account kind.
    GlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
    // MongoDB specifies the mongo db state for database account kind.
    MongoDB DatabaseAccountKind = "MongoDB"
    // Parse specifies the parse state for database account kind.
    Parse DatabaseAccountKind = "Parse"
)


// DatabaseAccountOfferType enumerates the values for database account offer
// type.
type DatabaseAccountOfferType string

const (
    // Standard specifies the standard state for database account offer type.
    Standard DatabaseAccountOfferType = "Standard"
)


// DefaultConsistencyLevel enumerates the values for default consistency level.
type DefaultConsistencyLevel string

const (
    // BoundedStaleness specifies the bounded staleness state for default
    // consistency level.
    BoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
    // Eventual specifies the eventual state for default consistency level.
    Eventual DefaultConsistencyLevel = "Eventual"
    // Session specifies the session state for default consistency level.
    Session DefaultConsistencyLevel = "Session"
    // Strong specifies the strong state for default consistency level.
    Strong DefaultConsistencyLevel = "Strong"
)


// KeyKind enumerates the values for key kind.
type KeyKind string

const (
    // Primary specifies the primary state for key kind.
    Primary KeyKind = "primary"
    // PrimaryReadonly specifies the primary readonly state for key kind.
    PrimaryReadonly KeyKind = "primaryReadonly"
    // Secondary specifies the secondary state for key kind.
    Secondary KeyKind = "secondary"
    // SecondaryReadonly specifies the secondary readonly state for key kind.
    SecondaryReadonly KeyKind = "secondaryReadonly"
)


// ConsistencyPolicy is the consistency policy for the DocumentDB database
// account.
type ConsistencyPolicy struct {
    DefaultConsistencyLevel DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`
    MaxStalenessPrefix *int64 `json:"maxStalenessPrefix,omitempty"`
    MaxIntervalInSeconds *int32 `json:"maxIntervalInSeconds,omitempty"`
}

// DatabaseAccount is a DocumentDB database account.
type DatabaseAccount struct {
    autorest.Response `json:"-"`
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    Kind DatabaseAccountKind `json:"kind,omitempty"`
    Properties *DatabaseAccountProperties `json:"properties,omitempty"`
}

// DatabaseAccountCreateUpdateParameters is parameters to create and update
// DocumentDB database accounts.
type DatabaseAccountCreateUpdateParameters struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
    Properties *DatabaseAccountCreateUpdateProperties `json:"properties,omitempty"`
    DatabaseAccountOfferType *string `json:"databaseAccountOfferType,omitempty"`
}

// DatabaseAccountCreateUpdateProperties is properties to create and update
// Azure DocumentDB database accounts.
type DatabaseAccountCreateUpdateProperties struct {
    ConsistencyPolicy *ConsistencyPolicy `json:"consistencyPolicy,omitempty"`
    Locations *[]Location `json:"locations,omitempty"`
}

// DatabaseAccountListKeysResult is the access keys for the given database
// account.
type DatabaseAccountListKeysResult struct {
    autorest.Response `json:"-"`
    PrimaryMasterKey *string `json:"primaryMasterKey,omitempty"`
    SecondaryMasterKey *string `json:"secondaryMasterKey,omitempty"`
    Properties *DatabaseAccountListReadOnlyKeysResult `json:"properties,omitempty"`
}

// DatabaseAccountListReadOnlyKeysResult is the read-only access keys for the
// given database account.
type DatabaseAccountListReadOnlyKeysResult struct {
    autorest.Response `json:"-"`
    PrimaryReadonlyMasterKey *string `json:"primaryReadonlyMasterKey,omitempty"`
    SecondaryReadonlyMasterKey *string `json:"secondaryReadonlyMasterKey,omitempty"`
}

// DatabaseAccountPatchParameters is parameters for patching Azure DocumentDB
// database account properties.
type DatabaseAccountPatchParameters struct {
    Tags *map[string]*string `json:"tags,omitempty"`
}

// DatabaseAccountProperties is properties for the database account.
type DatabaseAccountProperties struct {
    ProvisioningState *string `json:"provisioningState,omitempty"`
    DocumentEndpoint *string `json:"documentEndpoint,omitempty"`
    DatabaseAccountOfferType DatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`
    ConsistencyPolicy *ConsistencyPolicy `json:"consistencyPolicy,omitempty"`
    WriteLocations *[]Location `json:"writeLocations,omitempty"`
    ReadLocations *[]Location `json:"readLocations,omitempty"`
    FailoverPolicies *[]FailoverPolicy `json:"failoverPolicies,omitempty"`
}

// DatabaseAccountRegenerateKeyParameters is parameters to regenerate the keys
// within the database account.
type DatabaseAccountRegenerateKeyParameters struct {
    KeyKind KeyKind `json:"keyKind,omitempty"`
}

// DatabaseAccountsListResult is the List operation response, that contains
// the database accounts and their properties.
type DatabaseAccountsListResult struct {
    autorest.Response `json:"-"`
    Value *[]DatabaseAccount `json:"value,omitempty"`
}

// FailoverPolicy is the failover policy for a given region of a database
// account.
type FailoverPolicy struct {
    ID *string `json:"id,omitempty"`
    LocationName *string `json:"locationName,omitempty"`
    FailoverPriority *int32 `json:"failoverPriority,omitempty"`
}

// Location is a region in which the Azure DocumentDB database account is
// deployed.
type Location struct {
    ID *string `json:"id,omitempty"`
    LocationName *string `json:"locationName,omitempty"`
    DocumentEndpoint *string `json:"documentEndpoint,omitempty"`
    ProvisioningState *string `json:"provisioningState,omitempty"`
    FailoverPriority *int32 `json:"failoverPriority,omitempty"`
}

// Resource is a database account resource.
type Resource struct {
    ID *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
    Location *string `json:"location,omitempty"`
    Tags *map[string]*string `json:"tags,omitempty"`
}

