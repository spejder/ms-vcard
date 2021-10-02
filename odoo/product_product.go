package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	Active                               *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline                 *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration          *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                          *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                        *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                      *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                       *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                       *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttributeLineIds                     *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	Barcode                              *String    `xmlrpc:"barcode,omptempty"`
	CanBeExpensed                        *Bool      `xmlrpc:"can_be_expensed,omptempty"`
	CanImage1024BeZoomed                 *Bool      `xmlrpc:"can_image_1024_be_zoomed,omptempty"`
	CanImageVariant1024BeZoomed          *Bool      `xmlrpc:"can_image_variant_1024_be_zoomed,omptempty"`
	CategId                              *Many2One  `xmlrpc:"categ_id,omptempty"`
	Code                                 *String    `xmlrpc:"code,omptempty"`
	Color                                *Int       `xmlrpc:"color,omptempty"`
	CombinationIndices                   *String    `xmlrpc:"combination_indices,omptempty"`
	CompanyId                            *Many2One  `xmlrpc:"company_id,omptempty"`
	CostCurrencyId                       *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                           *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCode                          *String    `xmlrpc:"default_code,omptempty"`
	Description                          *String    `xmlrpc:"description,omptempty"`
	DescriptionPurchase                  *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale                      *String    `xmlrpc:"description_sale,omptempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omptempty"`
	EventOk                              *Bool      `xmlrpc:"event_ok,omptempty"`
	EventTicketIds                       *Relation  `xmlrpc:"event_ticket_ids,omptempty"`
	ExpensePolicy                        *Selection `xmlrpc:"expense_policy,omptempty"`
	HasConfigurableAttributes            *Bool      `xmlrpc:"has_configurable_attributes,omptempty"`
	Id                                   *Int       `xmlrpc:"id,omptempty"`
	Image1024                            *String    `xmlrpc:"image_1024,omptempty"`
	Image128                             *String    `xmlrpc:"image_128,omptempty"`
	Image1920                            *String    `xmlrpc:"image_1920,omptempty"`
	Image256                             *String    `xmlrpc:"image_256,omptempty"`
	Image512                             *String    `xmlrpc:"image_512,omptempty"`
	ImageVariant1024                     *String    `xmlrpc:"image_variant_1024,omptempty"`
	ImageVariant128                      *String    `xmlrpc:"image_variant_128,omptempty"`
	ImageVariant1920                     *String    `xmlrpc:"image_variant_1920,omptempty"`
	ImageVariant256                      *String    `xmlrpc:"image_variant_256,omptempty"`
	ImageVariant512                      *String    `xmlrpc:"image_variant_512,omptempty"`
	InvoicePolicy                        *Selection `xmlrpc:"invoice_policy,omptempty"`
	IsProductVariant                     *Bool      `xmlrpc:"is_product_variant,omptempty"`
	LastUpdate                           *Time      `xmlrpc:"__last_update,omptempty"`
	ListPrice                            *Float     `xmlrpc:"list_price,omptempty"`
	LstPrice                             *Float     `xmlrpc:"lst_price,omptempty"`
	MessageAttachmentCount               *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds                    *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds                   *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                      *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter               *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                   *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                           *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                    *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId              *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                    *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter             *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                    *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                        *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter                 *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                                 *String    `xmlrpc:"name,omptempty"`
	PackagingIds                         *Relation  `xmlrpc:"packaging_ids,omptempty"`
	PartnerRef                           *String    `xmlrpc:"partner_ref,omptempty"`
	Price                                *Float     `xmlrpc:"price,omptempty"`
	PriceExtra                           *Float     `xmlrpc:"price_extra,omptempty"`
	PricelistId                          *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	PricelistItemCount                   *Int       `xmlrpc:"pricelist_item_count,omptempty"`
	ProductTemplateAttributeValueIds     *Relation  `xmlrpc:"product_template_attribute_value_ids,omptempty"`
	ProductTmplId                        *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductVariantCount                  *Int       `xmlrpc:"product_variant_count,omptempty"`
	ProductVariantId                     *Many2One  `xmlrpc:"product_variant_id,omptempty"`
	ProductVariantIds                    *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	PropertyAccountExpenseId             *Many2One  `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeId              *Many2One  `xmlrpc:"property_account_income_id,omptempty"`
	PurchaseOk                           *Bool      `xmlrpc:"purchase_ok,omptempty"`
	Rental                               *Bool      `xmlrpc:"rental,omptempty"`
	SaleLineWarn                         *Selection `xmlrpc:"sale_line_warn,omptempty"`
	SaleLineWarnMsg                      *String    `xmlrpc:"sale_line_warn_msg,omptempty"`
	SaleOk                               *Bool      `xmlrpc:"sale_ok,omptempty"`
	SalesCount                           *Float     `xmlrpc:"sales_count,omptempty"`
	SellerIds                            *Relation  `xmlrpc:"seller_ids,omptempty"`
	Sequence                             *Int       `xmlrpc:"sequence,omptempty"`
	ServiceType                          *Selection `xmlrpc:"service_type,omptempty"`
	StandardPrice                        *Float     `xmlrpc:"standard_price,omptempty"`
	SupplierTaxesId                      *Relation  `xmlrpc:"supplier_taxes_id,omptempty"`
	TaxesId                              *Relation  `xmlrpc:"taxes_id,omptempty"`
	Type                                 *Selection `xmlrpc:"type,omptempty"`
	UomId                                *Many2One  `xmlrpc:"uom_id,omptempty"`
	UomName                              *String    `xmlrpc:"uom_name,omptempty"`
	UomPoId                              *Many2One  `xmlrpc:"uom_po_id,omptempty"`
	ValidProductTemplateAttributeLineIds *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omptempty"`
	VariantSellerIds                     *Relation  `xmlrpc:"variant_seller_ids,omptempty"`
	VisibleExpensePolicy                 *Bool      `xmlrpc:"visible_expense_policy,omptempty"`
	Volume                               *Float     `xmlrpc:"volume,omptempty"`
	VolumeUomName                        *String    `xmlrpc:"volume_uom_name,omptempty"`
	WebsiteMessageIds                    *Relation  `xmlrpc:"website_message_ids,omptempty"`
	Weight                               *Float     `xmlrpc:"weight,omptempty"`
	WeightUomName                        *String    `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	return c.Create(ProductProductModel, pp)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found")
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found")
}
