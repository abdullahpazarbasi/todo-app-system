package domain_fault_port

type FaultType string

const FaultTypeUnknown FaultType = "UNKNOWN"
const FaultTypeConnectionFailure FaultType = "CONNECTION_FAILURE"
const FaultTypeInaccessibleServer FaultType = "INACCESSIBLE_SERVER"
const FaultTypeDatabaseNotPrivileged FaultType = "DATABASE_NOT_PRIVILEGED"
const FaultTypeCollectionNotFound FaultType = "COLLECTION_NOT_FOUND"
const FaultTypeItemNotFound FaultType = "ITEM_NOT_FOUND"
const FaultTypeDuplicatedEntry FaultType = "DUPLICATED_ENTRY"
const FaultTypeAssociationViolation FaultType = "ASSOCIATION_VIOLATION"
const FaultTypeTimeout FaultType = "TIMEOUT"
const FaultTypeRaceCondition FaultType = "RACE_CONDITION"
const FaultTypeStructuralIncompatibility FaultType = "STRUCTURAL_INCOMPATIBILITY"
