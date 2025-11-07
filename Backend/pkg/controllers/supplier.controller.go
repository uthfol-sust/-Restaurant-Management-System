package controllers

import (
	"net/http"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
	"strconv"
)

type SupplierController interface {
	CreateSupplier(w http.ResponseWriter, r *http.Request)
	GetAllSupplier(w http.ResponseWriter, r *http.Request)
	GetBySupplierID(w http.ResponseWriter, r *http.Request)
	UpdateSupplier(w http.ResponseWriter, r *http.Request)
	DeleteSupplier(w http.ResponseWriter, r *http.Request)
}

type supplierController struct {
	supServices services.SupplierService
}

func NewSupplierController(supService services.SupplierService) SupplierController {
	return &supplierController{supServices: supService}
}

func (c *supplierController) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	supplier := &models.Supplier{}
	utils.ParseBody(r, supplier)

	newSupplier, err := c.supServices.CreateSupplier(supplier)
	if err != nil {
		http.Error(w, "Failed Create New Supplier", 500)
		return
	}

	utils.HTTPResponse(w, 201, newSupplier)
}

func (c *supplierController) GetAllSupplier(w http.ResponseWriter, r *http.Request) {
	supplierList, err := c.supServices.GetAllSupplier()
	if err != nil {
		http.Error(w, "Failed to load Suppliers List", 500)
		return
	}

	utils.HTTPResponse(w, 200, supplierList)
}

func (c *supplierController) GetBySupplierID(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
     http.Error(w,"Unvaild or Wrong ID",500)
	 return
	}

	supllier , err := c.supServices.GetBySupplierID(int64(ID))
    if err!=nil{
		http.Error(w,"Failed to load this supplier",404)
		return
	}

	utils.HTTPResponse(w,200,supllier)
}

func (c *supplierController) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
  supplier := &models.Supplier{}
  ID , err := strconv.Atoi(r.PathValue("id"))
  if err != nil {
     http.Error(w,"Unvaild or Wrong ID",500)
	 return
  }
  
  utils.ParseBody(r, supplier)

  updated, err := c.supServices.UpdateSupplier(int64(ID),supplier)
  if err != nil{
	http.Error(w,"failed to update",500)
	return
  }

  utils.HTTPResponse(w,200,updated)
}

func (c *supplierController) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
   ID , err := strconv.Atoi(r.PathValue("id"))
   if err != nil {
     http.Error(w,"Unvaild or Wrong ID",500)
	 return
   }
   
   err = c.supServices.DeleteSupplier(int64(ID))
   if err!=nil{
	http.Error(w,"Supplier Deletion Failed",500)
	return
   }

   utils.HTTPResponse(w,200,"Supplier Deleted Successfully")
}
