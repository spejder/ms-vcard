package odoo

import (
	"fmt"
)

// MemberProfile represents member.profile model.
type MemberProfile struct {
	AcceptSpecials                         *Selection `xmlrpc:"accept_specials,omptempty"`
	AccNumber                              *String    `xmlrpc:"acc_number,omptempty"`
	AccountInvoiceIds                      *Relation  `xmlrpc:"account_invoice_ids,omptempty"`
	AccountInvoiceLineIds                  *Relation  `xmlrpc:"account_invoice_line_ids,omptempty"`
	AccountInvoiceNotDraftIds              *Relation  `xmlrpc:"account_invoice_not_draft_ids,omptempty"`
	AccountInvoiceOpenIds                  *Relation  `xmlrpc:"account_invoice_open_ids,omptempty"`
	AccountTotalResiduals                  *Float     `xmlrpc:"account_total_residuals,omptempty"`
	Active                                 *Bool      `xmlrpc:"active,omptempty"`
	ActiveFunctionIds                      *Relation  `xmlrpc:"active_function_ids,omptempty"`
	ActiveFunctionsInCurrentOrganization   *Relation  `xmlrpc:"active_functions_in_current_organization,omptempty"`
	ActiveFunctionsInProfile               *Relation  `xmlrpc:"active_functions_in_profile,omptempty"`
	ActiveMembershipIds                    *Relation  `xmlrpc:"active_membership_ids,omptempty"`
	ActiveMembershipsInCurrentOrganization *Relation  `xmlrpc:"active_memberships_in_current_organization,omptempty"`
	ActiveMembershipsInProfile             *Relation  `xmlrpc:"active_memberships_in_profile,omptempty"`
	ActiveProfileForUserIds                *Relation  `xmlrpc:"active_profile_for_user_ids,omptempty"`
	ActiveProfileIds                       *Relation  `xmlrpc:"active_profile_ids,omptempty"`
	AddressCo                              *String    `xmlrpc:"address_co,omptempty"`
	Age                                    *Int       `xmlrpc:"age,omptempty"`
	AllFunctionsInCurrentOrganization      *Relation  `xmlrpc:"all_functions_in_current_organization,omptempty"`
	AllFunctionsInProfile                  *Relation  `xmlrpc:"all_functions_in_profile,omptempty"`
	AllMembershipsInCurrentOrganization    *Relation  `xmlrpc:"all_memberships_in_current_organization,omptempty"`
	AllMembershipsInProfile                *Relation  `xmlrpc:"all_memberships_in_profile,omptempty"`
	Anonymized                             *Bool      `xmlrpc:"anonymized,omptempty"`
	AnonymizeWarning                       *Bool      `xmlrpc:"anonymize_warning,omptempty"`
	BankIds                                *Relation  `xmlrpc:"bank_ids,omptempty"`
	BelongsToCompany                       *Bool      `xmlrpc:"belongs_to_company,omptempty"`
	Birthdate                              *Time      `xmlrpc:"birthdate,omptempty"`
	BirthdateShort                         *String    `xmlrpc:"birthdate_short,omptempty"`
	Bmhash                                 *String    `xmlrpc:"bmhash,omptempty"`
	BmId                                   *String    `xmlrpc:"bm_id,omptempty"`
	BmPin                                  *String    `xmlrpc:"bm_pin,omptempty"`
	BmPwdhash                              *String    `xmlrpc:"bm_pwdhash,omptempty"`
	Bmref                                  *String    `xmlrpc:"bmref,omptempty"`
	BmSsoId                                *Int       `xmlrpc:"bm_sso_id,omptempty"`
	CalendarLastNotifAck                   *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanAccessContactInfo                   *Bool      `xmlrpc:"can_access_contact_info,omptempty"`
	CanApprove                             *Bool      `xmlrpc:"can_approve,omptempty"`
	CanDelete                              *Bool      `xmlrpc:"can_delete,omptempty"`
	CanEdit                                *Bool      `xmlrpc:"can_edit,omptempty"`
	CanExpense                             *Bool      `xmlrpc:"can_expense,omptempty"`
	CanFullAll                             *Bool      `xmlrpc:"can_full_all,omptempty"`
	CanRead                                *Bool      `xmlrpc:"can_read,omptempty"`
	CategoryId                             *Relation  `xmlrpc:"category_id,omptempty"`
	ChildIds                               *Relation  `xmlrpc:"child_ids,omptempty"`
	City                                   *String    `xmlrpc:"city,omptempty"`
	CkrCheckIds                            *Relation  `xmlrpc:"ckr_check_ids,omptempty"`
	CkrPane                                *Bool      `xmlrpc:"ckr_pane,omptempty"`
	Color                                  *Int       `xmlrpc:"color,omptempty"`
	Comment                                *String    `xmlrpc:"comment,omptempty"`
	CommercialPartnerId                    *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omptempty"`
	CompleteAddress                        *String    `xmlrpc:"complete_address,omptempty"`
	ContactAddress                         *String    `xmlrpc:"contact_address,omptempty"`
	ContextAge                             *Int       `xmlrpc:"context_age,omptempty"`
	ContractIds                            *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractsCount                         *Int       `xmlrpc:"contracts_count,omptempty"`
	CountryId                              *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                              *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                                 *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                            *Float     `xmlrpc:"credit_limit,omptempty"`
	CurEventId                             *Many2One  `xmlrpc:"cur_event_id,omptempty"`
	Customer                               *Bool      `xmlrpc:"customer,omptempty"`
	Date                                   *Time      `xmlrpc:"date,omptempty"`
	DateLocalization                       *Time      `xmlrpc:"date_localization,omptempty"`
	Debit                                  *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                             *Float     `xmlrpc:"debit_limit,omptempty"`
	Diseases                               *String    `xmlrpc:"diseases,omptempty"`
	DisplayName                            *String    `xmlrpc:"display_name,omptempty"`
	Ean13                                  *String    `xmlrpc:"ean13,omptempty"`
	EditBirthdate                          *Bool      `xmlrpc:"edit_birthdate,omptempty"`
	EditMemberNumber                       *Bool      `xmlrpc:"edit_member_number,omptempty"`
	EditName                               *Bool      `xmlrpc:"edit_name,omptempty"`
	EditOrganizationId                     *Bool      `xmlrpc:"edit_organization_id,omptempty"`
	EditRestrictedFields                   *Bool      `xmlrpc:"edit_restricted_fields,omptempty"`
	Email                                  *String    `xmlrpc:"email,omptempty"`
	Employee                               *Bool      `xmlrpc:"employee,omptempty"`
	EventRegistrationIds                   *Relation  `xmlrpc:"event_registration_ids,omptempty"`
	Externalid                             *String    `xmlrpc:"externalid,omptempty"`
	Fax                                    *String    `xmlrpc:"fax,omptempty"`
	Firstname                              *String    `xmlrpc:"firstname,omptempty"`
	Function                               *String    `xmlrpc:"function,omptempty"`
	FunctionIds                            *Relation  `xmlrpc:"function_ids,omptempty"`
	FunctionsText                          *String    `xmlrpc:"functions_text,omptempty"`
	Gender                                 *Selection `xmlrpc:"gender,omptempty"`
	Handicap                               *String    `xmlrpc:"handicap,omptempty"`
	HasImage                               *Bool      `xmlrpc:"has_image,omptempty"`
	HelpInfo                               *String    `xmlrpc:"help_info,omptempty"`
	Id                                     *Int       `xmlrpc:"id,omptempty"`
	Image                                  *String    `xmlrpc:"image,omptempty"`
	ImageMedium                            *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                             *String    `xmlrpc:"image_small,omptempty"`
	ImportStatus                           *String    `xmlrpc:"import_status,omptempty"`
	InvoiceIds                             *Relation  `xmlrpc:"invoice_ids,omptempty"`
	IsActiveLeader                         *Bool      `xmlrpc:"is_active_leader,omptempty"`
	IsCompany                              *Bool      `xmlrpc:"is_company,omptempty"`
	JournalItemCount                       *Int       `xmlrpc:"journal_item_count,omptempty"`
	Lang                                   *Selection `xmlrpc:"lang,omptempty"`
	LastActiveDate                         *Time      `xmlrpc:"last_active_date,omptempty"`
	LastContactConfirm                     *Time      `xmlrpc:"last_contact_confirm,omptempty"`
	LastImport                             *Time      `xmlrpc:"last_import,omptempty"`
	LastKnownCompleteAddress               *String    `xmlrpc:"last_known_complete_address,omptempty"`
	LastKnownContactInfo                   *String    `xmlrpc:"last_known_contact_info,omptempty"`
	LastKnownEmail                         *String    `xmlrpc:"last_known_email,omptempty"`
	LastKnownPhoneCombo                    *String    `xmlrpc:"last_known_phone_combo,omptempty"`
	Lastname                               *String    `xmlrpc:"lastname,omptempty"`
	LastReconciliationDate                 *Time      `xmlrpc:"last_reconciliation_date,omptempty"`
	LastUpdate                             *Time      `xmlrpc:"__last_update,omptempty"`
	LeaderFunctionIds                      *Relation  `xmlrpc:"leader_function_ids,omptempty"`
	LocalOrgApproverIds                    *Relation  `xmlrpc:"local_org_approver_ids,omptempty"`
	LockingCkrCheckIds                     *Relation  `xmlrpc:"locking_ckr_check_ids,omptempty"`
	LocksNameIds                           *Relation  `xmlrpc:"locks_name_ids,omptempty"`
	MeetingCount                           *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                             *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MemberAnonymized                       *Bool      `xmlrpc:"member_anonymized,omptempty"`
	MemberId                               *Many2One  `xmlrpc:"member_id,omptempty"`
	MemberMagazineOptionId                 *Many2One  `xmlrpc:"member_magazine_option_id,omptempty"`
	MemberNumber                           *String    `xmlrpc:"member_number,omptempty"`
	MembershipIds                          *Relation  `xmlrpc:"membership_ids,omptempty"`
	MessageFollowerIds                     *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                             *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                      *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost                        *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageSummary                         *String    `xmlrpc:"message_summary,omptempty"`
	MessageUnread                          *Bool      `xmlrpc:"message_unread,omptempty"`
	Mobile                                 *String    `xmlrpc:"mobile,omptempty"`
	MobileClean                            *String    `xmlrpc:"mobile_clean,omptempty"`
	MunicipalityId                         *Many2One  `xmlrpc:"municipality_id,omptempty"`
	Name                                   *String    `xmlrpc:"name,omptempty"`
	NotifyEmail                            *Selection `xmlrpc:"notify_email,omptempty"`
	OpportunityCount                       *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                         *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	OptOut                                 *Bool      `xmlrpc:"opt_out,omptempty"`
	OrganizationId                         *Many2One  `xmlrpc:"organization_id,omptempty"`
	OrganizationStructureParentId          *Many2One  `xmlrpc:"organization_structure_parent_id,omptempty"`
	OrganizationTypeId                     *Many2One  `xmlrpc:"organization_type_id,omptempty"`
	OtherInfo                              *String    `xmlrpc:"other_info,omptempty"`
	OwnEventRegistrationIds                *Relation  `xmlrpc:"own_event_registration_ids,omptempty"`
	ParentId                               *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                             *String    `xmlrpc:"parent_name,omptempty"`
	ParishId                               *Many2One  `xmlrpc:"parish_id,omptempty"`
	PartnerId                              *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude                        *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude                       *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerPayerId                         *Many2One  `xmlrpc:"partner_payer_id,omptempty"`
	PayerForProfileIds                     *Relation  `xmlrpc:"payer_for_profile_ids,omptempty"`
	PayerForProfileThisOrganizationIds     *Relation  `xmlrpc:"payer_for_profile_this_organization_ids,omptempty"`
	Pbmhash                                *String    `xmlrpc:"pbmhash,omptempty"`
	Pbmref                                 *String    `xmlrpc:"pbmref,omptempty"`
	PermissionPhoto                        *Selection `xmlrpc:"permission_photo,omptempty"`
	Phone                                  *String    `xmlrpc:"phone,omptempty"`
	PhonecallCount                         *Int       `xmlrpc:"phonecall_count,omptempty"`
	PhonecallIds                           *Relation  `xmlrpc:"phonecall_ids,omptempty"`
	PhoneCombo                             *String    `xmlrpc:"phone_combo,omptempty"`
	PlastImport                            *Time      `xmlrpc:"plast_import,omptempty"`
	PreliminaryOrganizationId              *Many2One  `xmlrpc:"preliminary_organization_id,omptempty"`
	PrimaryMembershipOrganizationId        *Many2One  `xmlrpc:"primary_membership_organization_id,omptempty"`
	ProfileAnonymized                      *Bool      `xmlrpc:"profile_anonymized,omptempty"`
	ProfileIds                             *Relation  `xmlrpc:"profile_ids,omptempty"`
	PromotionIds                           *Relation  `xmlrpc:"promotion_ids,omptempty"`
	PropertyAccountPayable                 *Many2One  `xmlrpc:"property_account_payable,omptempty"`
	PropertyAccountPosition                *Many2One  `xmlrpc:"property_account_position,omptempty"`
	PropertyAccountReceivable              *Many2One  `xmlrpc:"property_account_receivable,omptempty"`
	PropertyPaymentTerm                    *Many2One  `xmlrpc:"property_payment_term,omptempty"`
	PropertyProductPricelist               *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyStockCustomer                  *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier                  *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTerm            *Many2One  `xmlrpc:"property_supplier_payment_term,omptempty"`
	Ref                                    *String    `xmlrpc:"ref,omptempty"`
	RefCompanies                           *Relation  `xmlrpc:"ref_companies,omptempty"`
	Registered                             *Bool      `xmlrpc:"registered,omptempty"`
	RegNumber                              *String    `xmlrpc:"reg_number,omptempty"`
	RelationAllIds                         *Relation  `xmlrpc:"relation_all_ids,omptempty"`
	RelationAllMemberIds                   *Relation  `xmlrpc:"relation_all_member_ids,omptempty"`
	RelationCount                          *Int       `xmlrpc:"relation_count,omptempty"`
	RelationIds                            *Relation  `xmlrpc:"relation_ids,omptempty"`
	RelationPartnerList                    *String    `xmlrpc:"relation_partner_list,omptempty"`
	RelationPrimaryMemberIds               *Relation  `xmlrpc:"relation_primary_member_ids,omptempty"`
	RelativeForProfileId                   *Many2One  `xmlrpc:"relative_for_profile_id,omptempty"`
	RelativeMemberId                       *Many2One  `xmlrpc:"relative_member_id,omptempty"`
	RelativePartnerId                      *Many2One  `xmlrpc:"relative_partner_id,omptempty"`
	RelativeTypeId                         *Many2One  `xmlrpc:"relative_type_id,omptempty"`
	SaleOrderCount                         *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                           *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	School                                 *String    `xmlrpc:"school,omptempty"`
	SchoolClassLetter                      *String    `xmlrpc:"school_class_letter,omptempty"`
	SchoolClassNumber                      *String    `xmlrpc:"school_class_number,omptempty"`
	SchoolStartYear                        *Int       `xmlrpc:"school_start_year,omptempty"`
	ScoutName                              *String    `xmlrpc:"scout_name,omptempty"`
	SearchRelationDate                     *Time      `xmlrpc:"search_relation_date,omptempty"`
	SearchRelationId                       *Many2One  `xmlrpc:"search_relation_id,omptempty"`
	SearchRelationPartnerCategoryId        *Many2One  `xmlrpc:"search_relation_partner_category_id,omptempty"`
	SearchRelationPartnerId                *Many2One  `xmlrpc:"search_relation_partner_id,omptempty"`
	SectionId                              *Many2One  `xmlrpc:"section_id,omptempty"`
	Self                                   *Many2One  `xmlrpc:"self,omptempty"`
	SelfRelationPartnerList                *String    `xmlrpc:"self_relation_partner_list,omptempty"`
	SignupExpiration                       *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                            *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                             *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                              *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                            *Bool      `xmlrpc:"signup_valid,omptempty"`
	Speaker                                *Bool      `xmlrpc:"speaker,omptempty"`
	SpecialConsiderations                  *String    `xmlrpc:"special_considerations,omptempty"`
	State                                  *Selection `xmlrpc:"state,omptempty"`
	StateId                                *Many2One  `xmlrpc:"state_id,omptempty"`
	StoredPartnerId                        *Many2One  `xmlrpc:"stored_partner_id,omptempty"`
	Street                                 *String    `xmlrpc:"street,omptempty"`
	Street2                                *String    `xmlrpc:"street2,omptempty"`
	StreetFloor                            *String    `xmlrpc:"street_floor,omptempty"`
	StreetLetter                           *String    `xmlrpc:"street_letter,omptempty"`
	StreetName                             *String    `xmlrpc:"street_name,omptempty"`
	StreetNumber                           *String    `xmlrpc:"street_number,omptempty"`
	StreetPlacement                        *String    `xmlrpc:"street_placement,omptempty"`
	SubscriptionCard                       *String    `xmlrpc:"subscription_card,omptempty"`
	SubscriptionFeeDateEnd                 *Time      `xmlrpc:"subscription_fee_date_end,omptempty"`
	SubscriptionFeeDateStart               *Time      `xmlrpc:"subscription_fee_date_start,omptempty"`
	SubscriptionFeeHasDraftLines           *Bool      `xmlrpc:"subscription_fee_has_draft_lines,omptempty"`
	SubscriptionFeeReceivable              *Float     `xmlrpc:"subscription_fee_receivable,omptempty"`
	SubscriptionLastChargedEndDate         *Time      `xmlrpc:"subscription_last_charged_end_date,omptempty"`
	SubscriptionLastWarningDate            *Time      `xmlrpc:"subscription_last_warning_date,omptempty"`
	SubscriptionProductId                  *Many2One  `xmlrpc:"subscription_product_id,omptempty"`
	SubscriptionTransaction                *Many2One  `xmlrpc:"subscription_transaction,omptempty"`
	Supplier                               *Bool      `xmlrpc:"supplier,omptempty"`
	TaskCount                              *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                                *Relation  `xmlrpc:"task_ids,omptempty"`
	Title                                  *Many2One  `xmlrpc:"title,omptempty"`
	TotalInvoiced                          *Float     `xmlrpc:"total_invoiced,omptempty"`
	Type                                   *Selection `xmlrpc:"type,omptempty"`
	Tz                                     *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                               *String    `xmlrpc:"tz_offset,omptempty"`
	UseParentAddress                       *Bool      `xmlrpc:"use_parent_address,omptempty"`
	UserFullIds                            *Relation  `xmlrpc:"user_full_ids,omptempty"`
	UserId                                 *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                                *Relation  `xmlrpc:"user_ids,omptempty"`
	UserReadIds                            *Relation  `xmlrpc:"user_read_ids,omptempty"`
	UserReadLimitedIds                     *Relation  `xmlrpc:"user_read_limited_ids,omptempty"`
	Vat                                    *String    `xmlrpc:"vat,omptempty"`
	VatSubjected                           *Bool      `xmlrpc:"vat_subjected,omptempty"`
	WaitinglistDate                        *Time      `xmlrpc:"waitinglist_date,omptempty"`
	Website                                *String    `xmlrpc:"website,omptempty"`
	WebsiteDescription                     *String    `xmlrpc:"website_description,omptempty"`
	WebsiteMessageIds                      *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription                 *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords                    *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaTitle                       *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished                       *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription                *String    `xmlrpc:"website_short_description,omptempty"`
	WriteDate                              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                               *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                                    *String    `xmlrpc:"zip,omptempty"`
}

// MemberProfiles represents array of member.profile model.
type MemberProfiles []MemberProfile

// MemberProfileModel is the odoo model name.
const MemberProfileModel = "member.profile"

// Many2One convert MemberProfile to *Many2One.
func (mp *MemberProfile) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMemberProfile creates a new member.profile model and returns its id.
func (c *Client) CreateMemberProfile(mp *MemberProfile) (int64, error) {
	return c.Create(MemberProfileModel, mp)
}

// UpdateMemberProfile updates an existing member.profile record.
func (c *Client) UpdateMemberProfile(mp *MemberProfile) error {
	return c.UpdateMemberProfiles([]int64{mp.Id.Get()}, mp)
}

// UpdateMemberProfiles updates existing member.profile records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMemberProfiles(ids []int64, mp *MemberProfile) error {
	return c.Update(MemberProfileModel, ids, mp)
}

// DeleteMemberProfile deletes an existing member.profile record.
func (c *Client) DeleteMemberProfile(id int64) error {
	return c.DeleteMemberProfiles([]int64{id})
}

// DeleteMemberProfiles deletes existing member.profile records.
func (c *Client) DeleteMemberProfiles(ids []int64) error {
	return c.Delete(MemberProfileModel, ids)
}

// GetMemberProfile gets member.profile existing record.
func (c *Client) GetMemberProfile(id int64) (*MemberProfile, error) {
	mps, err := c.GetMemberProfiles([]int64{id})
	if err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of member.profile not found", id)
}

// GetMemberProfiles gets member.profile existing records.
func (c *Client) GetMemberProfiles(ids []int64) (*MemberProfiles, error) {
	mps := &MemberProfiles{}
	if err := c.Read(MemberProfileModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMemberProfile finds member.profile record by querying it with criteria.
func (c *Client) FindMemberProfile(criteria *Criteria) (*MemberProfile, error) {
	mps := &MemberProfiles{}
	if err := c.SearchRead(MemberProfileModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("member.profile was not found")
}

// FindMemberProfiles finds member.profile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberProfiles(criteria *Criteria, options *Options) (*MemberProfiles, error) {
	mps := &MemberProfiles{}
	if err := c.SearchRead(MemberProfileModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMemberProfileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberProfileIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MemberProfileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMemberProfileId finds record id by querying it with criteria.
func (c *Client) FindMemberProfileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MemberProfileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("member.profile was not found")
}
