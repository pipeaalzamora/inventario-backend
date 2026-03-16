package services

// Power permission constants
// This file centralizes all power constants used across services

// ==================== PROPERTY PREFIXES ====================
// These prefixes are used for ownership-based powers (ownPower)
// Format: "<service>:<entity_id>"

const (
	// PowerPrefixCompany is used for company ownership validation
	// Example: "company:123456" grants access to company with id 123456
	PowerPrefixCompany = "company:"

	// PowerPrefixStore is used for store ownership validation
	// Example: "store:789" grants access to store with id 789
	PowerPrefixStore = "store:"

	// PowerPrefixSupplier is used for supplier-related powers
	PowerPrefixSupplier = "supplier:"
)

// ==================== USER POWERS ====================

const (
	PowerUserCreate = "user:create"
	PowerUserUpdate = "user:update"
	PowerUserDelete = "user:enable-disable"
)

// ==================== PROFILE POWERS ====================

const (
	PowerProfileCreate = "profile:create"
	PowerProfileUpdate = "profile:update"
	PowerProfileDelete = "profile:delete"
)

// ==================== STORE POWERS ====================

const (
	PowerStoreCreate = PowerPrefixStore + "create"
	PowerStoreUpdate = PowerPrefixStore + "update"
)

// ==================== SUPPLIER POWERS ====================

const (
	PowerSupplierCreate = PowerPrefixSupplier + "create"
	PowerSupplierUpdate = PowerPrefixSupplier + "update"
	PowerSupplierDelete = PowerPrefixSupplier + "delete"
)

// ==================== PURCHASE POWERS ====================
/*
const (
	PowerPurchaseCreate  = "purchase:create"
	PowerPurchaseUpdate  = "purchase:update"
	PowerPurchaseApprove = "purchase:approve"
	PowerPurchaseDelete  = "purchase:delete"
)
*/

// ==================== INVENTORY REQUEST POWERS ====================

const (
	PowerRequestCreateForSupplier = "request:createForSupplier"
	PowerRequestCreateForStore    = "request:createForStore"
	PowerRequestCreateForCompany  = "request:createForCompany"
	//PowerRequestUpdate  = "request:update"
	//PowerRequestResolve = "request:resolve"
)

// ==================== WAREHOUSE POWERS ====================

const (
	PowerWarehouseCreate = "warehouse:create"
	PowerWarehouseUpdate = "warehouse:update"
)

// ==================== PRODUCT POWERS ====================

const (
	PowerProductCreate = "product:create"
	PowerProductUpdate = "product:update"
	PowerProductDelete = "product:delete"
)

// ==================== PRODUCT COMPANY POWERS ====================

const (
	PowerProductCompanyCreate = "product_company:create"
	PowerProductCompanyUpdate = "product_company:update"
)

// ==================== INVENTORY COUNT POWERS ====================

const (
	PowerInventoryCountCreate = "inventory_count:create"
	PowerInventoryCountUpdate = "inventory_count:update"
)

// ==================== PRODUCT MOVEMENT POWERS ====================

const (
	PowerProductMovementCreate = "product_movement:create"
	PowerProductMovementUpdate = "product_movement:update"
)

// ==================== DELIVERY PURCHASE NOTE POWERS ====================

const (
	PowerDeliveryPurchaseNoteCreate = "delivery_purchase_note:create"
	PowerDeliveryPurchaseNoteUpdate = "delivery_purchase_note:update"
)
