package models

import "time"

type ModelRequest struct {
	Id             string                `json:"id" bson:"id"`
	DisplayId      string                `json:"displayId" bson:"displayId"`
	CompanyId      string                `json:"companyId" bson:"companyId"`
	StoreId        string                `json:"storeId" bson:"storeId"`
	WarehouseId    *string               `json:"warehouseId,omitempty" bson:"warehouseId,omitempty"`
	Status         ModelRequestStatus    `json:"status" bson:"status"`
	RequestKind    ModelRequestKind      `json:"requestKind" bson:"requestKind"`
	CreatedAt      time.Time             `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt" bson:"updatedAt"`
	CreatedBy      ModelRequestUser      `json:"createdBy" bson:"createdBy"`
	Items          []ModelRequestItem    `json:"items" bson:"items"`
	RequestHistory []ModelRequestHistory `json:"requestHistory" bson:"requestHistory"`
	DocsTree       []ModelRequestDoc     `json:"docsTree" bson:"docsTree"`
}

type ModelRequestUser struct {
	Id        string `json:"id" bson:"id"`
	UserName  string `json:"userName" bson:"userName"`
	UserEmail string `json:"userEmail" bson:"userEmail"`
}

/*
Los estados de las solicitudes son:
- Creada
- Con conflicto
- Aprobada
- Cancelada
- Finalizada
*/
type ModelRequestStatus struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

/*
Los tipos de solicitudes son:
- Orden de Compra
- Orden de traspaso misma razon social
- Order de traspaso otra razon social
*/
type ModelRequestKind struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type RequestKind string

const (
	RequestKindPurchase RequestKind = "RequestForSupplier"
	RequestKindCompany  RequestKind = "RequestForCompany"
	RequestKindStore    RequestKind = "RequestForStore"
)

func (k RequestKind) ToString() string {
	return string(k)
}

func (k RequestKind) ToDisplayString() string {
	switch k {
	case RequestKindPurchase:
		return "Proveedor"
	case RequestKindCompany:
		return "Empresa"
	case RequestKindStore:
		return "Tienda"
	}
	return string(k)
}

func (k RequestKind) ToInt() int {
	switch k {
	case RequestKindPurchase:
		return 1
	case RequestKindCompany:
		return 2
	case RequestKindStore:
		return 3
	default:
		return 1 // Valor por defecto
	}
}

func RequestKindFromInt(value int) RequestKind {
	switch value {
	case 1:
		return RequestKindPurchase
	case 2:
		return RequestKindCompany
	case 3:
		return RequestKindStore
	default:
		return RequestKindPurchase
	}
}

type ModelRequestItem struct {
	Id                     string                  `json:"id" bson:"id"`
	Sku                    string                  `json:"sku" bson:"sku"`
	StoreProductId         string                  `json:"productId" bson:"productId"`
	RequestedQuantity      float32                 `json:"requestedQuantity" bson:"requestedQuantity"`
	RequestRestriction     ModelRequestRestriction `json:"restriction" bson:"restriction"`
	RequestMeasurementUnit ModelMeasurementUnique  `json:"requestMeasurementUnit" bson:"requestMeasurementUnit"`
}

type ModelRequestRestriction struct {
	MaxQuantity float32 `json:"maxQuantity" bson:"maxQuantity"` //0 sin restriccion, 1 o mas con restriccion
}

type ModelRequestHistory struct {
	Id          string                      `json:"id" bson:"id"`
	Status      ModelRequestStatus          `json:"status" bson:"status"`
	CreatedAt   time.Time                   `json:"changedAt" bson:"changedAt"`
	ChangedBy   ModelRequestHistoryChangeBy `json:"changedBy" bson:"changedBy"`
	Observation string                      `json:"observation" bson:"observation"`
}

type ModelRequestHistoryChangeBy struct {
	Name             string `json:"name" bson:"name"`                         //solo los usuarios llenan esto
	OrganizationName string `json:"organizationName" bson:"organizationName"` //todos llenan esto, company name para usuarios y supplierName para externos
}

type ModelRequestDoc struct {
	Id                   string                         `json:"id" bson:"id"`
	RequestId            string                         `json:"requestId" bson:"requestId"`
	DocType              ModelRequestDocType            `json:"docType" bson:"docType"`
	DocDisplayId         string                         `json:"docDisplayId" bson:"docDisplayId"`
	DocParentReferenceId *string                        `json:"docParentReferenceId" bson:"docParentReferenceId"`
	DocStatus            ModelRequestDocStatus          `json:"docStatus" bson:"docStatus"`
	Company              ModelCompany                   `json:"company" bson:"company"`
	Supplier             ModelRequestDocSupplier        `json:"supplier" bson:"supplier"`
	Items                []ModelRequestDocItem          `json:"docItems" bson:"docItems"`
	ReceptionDate        *time.Time                     `json:"receptionDate" bson:"receptionDate"`
	CreatedAt            time.Time                      `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time                      `json:"updatedAt" bson:"updatedAt"`
	ObservationHistory   []ModelRequestDocObservation   `json:"observationHistory" bson:"observationHistory"`
	StatusHistory        []ModelRequestDocStatusHistory `json:"statusHistory" bson:"statusHistory"`
}

/*
Los tipos de documentos son:
- Orden de Compra
- Orden de traspaso
*/
type ModelRequestDocType struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

/*
Los estados de los documentos son:
- Negociando
- Cancelada
- En camino
- Recepcionando
- Completada
- En disputa
- Finalizada
*/
type ModelRequestDocStatus struct {
	Id   string `json:"status" bson:"status"`
	Name string `json:"statusDisplay" bson:"statusDisplay"`
}

type ModelRequestDocStatusHistory struct {
	ChangedBy ModelRequestHistoryChangeBy `json:"changedBy" bson:"changedBy"`
	Status    ModelRequestDocStatus       `json:"status" bson:"status"`
	CreatedAt time.Time                   `json:"createdAt" bson:"createdAt"`
}

type ModelRequestDocObservation struct {
	ChangedBy           ModelRequestHistoryChangeBy `json:"changedBy" bson:"changedBy"`
	Observation         string                      `json:"observation" bson:"observation"`
	CreatedAt           time.Time                   `json:"createdAt" bson:"createdAt"`
	LeftDisplayPosition bool                        `json:"leftDisplayPosition" bson:"leftDisplayPosition"`
}

/*
El proveedor puede ser externo, o interno (de la empresa), incluso interno pero
de otra razon social
*/
type ModelRequestDocSupplier struct {
	Id              string `json:"id" bson:"id"`
	RawFiscalId     string `json:"rawFiscalId" bson:"rawFiscalId"`
	SupplierName    string `json:"supplierName" bson:"supplierName"`
	SupplierAddress string `json:"supplierAddress" bson:"supplierAddress"`
	SupplierCity    string `json:"supplierCity" bson:"supplierCity"`
	ContactEmail    string `json:"supplierEmail" bson:"supplierEmail"`
	ContactPhone    string `json:"supplierPhone" bson:"supplierPhone"`
}

type ModelRequestDocItem struct {
	ModelRequestItem
	SupplierItem  ModelRequestDocSupplierItem  `json:"supplierItem" bson:"supplierItem"`
	IsDeleted     bool                         `json:"isDeleted" bson:"isDeleted"`
	ReceptionData ModelRequestDocReceptionData `json:"receptionData" bson:"receptionData"`
}

type ModelRequestDocReceptionData struct {
	ReceptionedAt time.Time `json:"receptionedAt" bson:"receptionedAt"`
	ReceptionBy   string    `json:"receptionBy" bson:"receptionBy"`
	ReceptionNote string    `json:"receptionNote" bson:"receptionNote"`
	Quantity      float32   `json:"quantity" bson:"quantity"`
}

type ModelRequestDocSupplierItem struct {
	Id              string                         `json:"id" bson:"id"`
	Sku             *string                        `json:"sku" bson:"sku"`
	Quantity        float32                        `json:"quantity" bson:"quantity"`
	MeasurementUnit ModelMeasurementConversionUnit `json:"measurementUnit" bson:"measurementUnit"`
	UnitPrice       float32                        `json:"unitPrice" bson:"unitPrice"`
	TotalPrice      float32                        `json:"totalPrice" bson:"totalPrice"`
	Description     string                         `json:"description" bson:"description"`
	IsAccepted      bool                           `json:"acceptedStatus" bson:"acceptedStatus"`
}
