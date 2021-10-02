package odoo

import (
	"fmt"
)

// ResUsers represents res.users model.
type ResUsers struct {
	AccessesCount                 *Int       `xmlrpc:"accesses_count,omptempty"`
	ActionId                      *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omptempty"`
	ActivePartner                 *Bool      `xmlrpc:"active_partner,omptempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalInfo                *String    `xmlrpc:"additional_info,omptempty"`
	AdditionalNote                *String    `xmlrpc:"additional_note,omptempty"`
	AddressHomeId                 *Many2One  `xmlrpc:"address_home_id,omptempty"`
	AddressId                     *Many2One  `xmlrpc:"address_id,omptempty"`
	AliasContact                  *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasId                       *Many2One  `xmlrpc:"alias_id,omptempty"`
	AllocationCount               *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay             *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationUsedCount           *Float     `xmlrpc:"allocation_used_count,omptempty"`
	AllocationUsedDisplay         *String    `xmlrpc:"allocation_used_display,omptempty"`
	AttendanceState               *Selection `xmlrpc:"attendance_state,omptempty"`
	BadgeIds                      *Relation  `xmlrpc:"badge_ids,omptempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankAccountId                 *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omptempty"`
	Barcode                       *String    `xmlrpc:"barcode,omptempty"`
	Birthday                      *Time      `xmlrpc:"birthday,omptempty"`
	BronzeBadge                   *Int       `xmlrpc:"bronze_badge,omptempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanEdit                       *Bool      `xmlrpc:"can_edit,omptempty"`
	CanPublish                    *Bool      `xmlrpc:"can_publish,omptempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omptempty"`
	CategoryIds                   *Relation  `xmlrpc:"category_ids,omptempty"`
	Certificate                   *Selection `xmlrpc:"certificate,omptempty"`
	CertificationsCompanyCount    *Int       `xmlrpc:"certifications_company_count,omptempty"`
	CertificationsCount           *Int       `xmlrpc:"certifications_count,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	Children                      *Int       `xmlrpc:"children,omptempty"`
	City                          *String    `xmlrpc:"city,omptempty"`
	CoachId                       *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                         *Int       `xmlrpc:"color,omptempty"`
	Comment                       *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompaniesCount                *Int       `xmlrpc:"companies_count,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyIds                    *Relation  `xmlrpc:"company_ids,omptempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omptempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omptempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omptempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omptempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omptempty"`
	CountryOfBirth                *Many2One  `xmlrpc:"country_of_birth,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                        *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomerRank                  *Int       `xmlrpc:"customer_rank,omptempty"`
	Date                          *Time      `xmlrpc:"date,omptempty"`
	Debit                         *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omptempty"`
	DepartmentId                  *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	Email                         *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omptempty"`
	EmailNormalized               *String    `xmlrpc:"email_normalized,omptempty"`
	EmergencyContact              *String    `xmlrpc:"emergency_contact,omptempty"`
	EmergencyPhone                *String    `xmlrpc:"emergency_phone,omptempty"`
	Employee                      *Bool      `xmlrpc:"employee,omptempty"`
	EmployeeBankAccountId         *Many2One  `xmlrpc:"employee_bank_account_id,omptempty"`
	EmployeeCount                 *Int       `xmlrpc:"employee_count,omptempty"`
	EmployeeCountryId             *Many2One  `xmlrpc:"employee_country_id,omptempty"`
	EmployeeId                    *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeIds                   *Relation  `xmlrpc:"employee_ids,omptempty"`
	EmployeeParentId              *Many2One  `xmlrpc:"employee_parent_id,omptempty"`
	EmployeePhone                 *String    `xmlrpc:"employee_phone,omptempty"`
	EventCount                    *Int       `xmlrpc:"event_count,omptempty"`
	ExpenseManagerId              *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	Function                      *String    `xmlrpc:"function,omptempty"`
	Gender                        *Selection `xmlrpc:"gender,omptempty"`
	GoalIds                       *Relation  `xmlrpc:"goal_ids,omptempty"`
	GoldBadge                     *Int       `xmlrpc:"gold_badge,omptempty"`
	GroupsCount                   *Int       `xmlrpc:"groups_count,omptempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omptempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	HoursLastMonth                *Float     `xmlrpc:"hours_last_month,omptempty"`
	HoursLastMonthDisplay         *String    `xmlrpc:"hours_last_month_display,omptempty"`
	HrPresenceState               *Selection `xmlrpc:"hr_presence_state,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	IdentificationId              *String    `xmlrpc:"identification_id,omptempty"`
	Image1024                     *String    `xmlrpc:"image_1024,omptempty"`
	Image128                      *String    `xmlrpc:"image_128,omptempty"`
	Image1920                     *String    `xmlrpc:"image_1920,omptempty"`
	Image256                      *String    `xmlrpc:"image_256,omptempty"`
	Image512                      *String    `xmlrpc:"image_512,omptempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omptempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omptempty"`
	Instructor                    *Bool      `xmlrpc:"instructor,omptempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsAbsent                      *Bool      `xmlrpc:"is_absent,omptempty"`
	IsAddressHomeACompany         *Bool      `xmlrpc:"is_address_home_a_company,omptempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omptempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omptempty"`
	IsModerator                   *Bool      `xmlrpc:"is_moderator,omptempty"`
	IsPublished                   *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	JobTitle                      *String    `xmlrpc:"job_title,omptempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omptempty"`
	Karma                         *Int       `xmlrpc:"karma,omptempty"`
	KmHomeWork                    *Int       `xmlrpc:"km_home_work,omptempty"`
	L10NSgUniqueEntityNumber      *String    `xmlrpc:"l10n_sg_unique_entity_number,omptempty"`
	Lang                          *Selection `xmlrpc:"lang,omptempty"`
	LastActivity                  *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime              *String    `xmlrpc:"last_activity_time,omptempty"`
	LastCheckIn                   *Time      `xmlrpc:"last_check_in,omptempty"`
	LastCheckOut                  *Time      `xmlrpc:"last_check_out,omptempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	LeaveDateTo                   *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId                *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LivechatUsername              *String    `xmlrpc:"livechat_username,omptempty"`
	LogIds                        *Relation  `xmlrpc:"log_ids,omptempty"`
	Login                         *String    `xmlrpc:"login,omptempty"`
	LoginDate                     *Time      `xmlrpc:"login_date,omptempty"`
	Marital                       *Selection `xmlrpc:"marital,omptempty"`
	MedicExam                     *Time      `xmlrpc:"medic_exam,omptempty"`
	MeetingCount                  *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                    *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mobile                        *String    `xmlrpc:"mobile,omptempty"`
	MobilePhone                   *String    `xmlrpc:"mobile_phone,omptempty"`
	ModerationChannelIds          *Relation  `xmlrpc:"moderation_channel_ids,omptempty"`
	ModerationCounter             *Int       `xmlrpc:"moderation_counter,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	NewPassword                   *String    `xmlrpc:"new_password,omptempty"`
	NextRankId                    *Many2One  `xmlrpc:"next_rank_id,omptempty"`
	NotificationType              *Selection `xmlrpc:"notification_type,omptempty"`
	OdoobotState                  *Selection `xmlrpc:"odoobot_state,omptempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	OutOfOfficeMessage            *String    `xmlrpc:"out_of_office_message,omptempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omptempty"`
	PartnerGid                    *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omptempty"`
	Passport                      *String    `xmlrpc:"passport,omptempty"`
	PassportId                    *String    `xmlrpc:"passport_id,omptempty"`
	Password                      *String    `xmlrpc:"password,omptempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	PermitNo                      *String    `xmlrpc:"permit_no,omptempty"`
	Phone                         *String    `xmlrpc:"phone,omptempty"`
	PhoneBlacklisted              *Bool      `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneSanitized                *String    `xmlrpc:"phone_sanitized,omptempty"`
	Pin                           *String    `xmlrpc:"pin,omptempty"`
	PlaceOfBirth                  *String    `xmlrpc:"place_of_birth,omptempty"`
	PrivateEmail                  *String    `xmlrpc:"private_email,omptempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	RankId                        *Many2One  `xmlrpc:"rank_id,omptempty"`
	Ref                           *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	ResourceCalendarId            *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceIds                   *Relation  `xmlrpc:"resource_ids,omptempty"`
	RulesCount                    *Int       `xmlrpc:"rules_count,omptempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleTeamId                    *Many2One  `xmlrpc:"sale_team_id,omptempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omptempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	Self                          *Many2One  `xmlrpc:"self,omptempty"`
	SessionIds                    *Relation  `xmlrpc:"session_ids,omptempty"`
	Share                         *Bool      `xmlrpc:"share,omptempty"`
	ShowLeaves                    *Bool      `xmlrpc:"show_leaves,omptempty"`
	Signature                     *String    `xmlrpc:"signature,omptempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omptempty"`
	SilverBadge                   *Int       `xmlrpc:"silver_badge,omptempty"`
	SlideChannelCompanyCount      *Int       `xmlrpc:"slide_channel_company_count,omptempty"`
	SlideChannelCount             *Int       `xmlrpc:"slide_channel_count,omptempty"`
	SlideChannelIds               *Relation  `xmlrpc:"slide_channel_ids,omptempty"`
	SpouseBirthdate               *Time      `xmlrpc:"spouse_birthdate,omptempty"`
	SpouseCompleteName            *String    `xmlrpc:"spouse_complete_name,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                        *String    `xmlrpc:"street,omptempty"`
	Street2                       *String    `xmlrpc:"street2,omptempty"`
	StudyField                    *String    `xmlrpc:"study_field,omptempty"`
	StudySchool                   *String    `xmlrpc:"study_school,omptempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omptempty"`
	TargetSalesDone               *Int       `xmlrpc:"target_sales_done,omptempty"`
	TargetSalesInvoiced           *Int       `xmlrpc:"target_sales_invoiced,omptempty"`
	TargetSalesWon                *Int       `xmlrpc:"target_sales_won,omptempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omptempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omptempty"`
	Title                         *Many2One  `xmlrpc:"title,omptempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omptempty"`
	Trust                         *Selection `xmlrpc:"trust,omptempty"`
	Type                          *Selection `xmlrpc:"type,omptempty"`
	Tz                            *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                           *String    `xmlrpc:"vat,omptempty"`
	Vehicle                       *String    `xmlrpc:"vehicle,omptempty"`
	VisaExpire                    *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                        *String    `xmlrpc:"visa_no,omptempty"`
	VisitorIds                    *Relation  `xmlrpc:"visitor_ids,omptempty"`
	Website                       *String    `xmlrpc:"website,omptempty"`
	WebsiteDescription            *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                     *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription        *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords           *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg              *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle              *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription       *String    `xmlrpc:"website_short_description,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	WorkEmail                     *String    `xmlrpc:"work_email,omptempty"`
	WorkLocation                  *String    `xmlrpc:"work_location,omptempty"`
	WorkPhone                     *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                           *String    `xmlrpc:"zip,omptempty"`
}

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	return c.Create(ResUsersModel, ru)
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsersNoResetPassword(ru *ResUsers) (int64, error) {
	return c.CreateWithOption(ResUsersModel, ru, NewOptions().NoResetPassword(true))
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) ResetPasswordUsers(ids []int64) error {
	_ , err := c.ExecuteKw("action_reset_password", "res.users", []interface{}{ids}, NewOptions().CreateUserOption(1))

	if err != nil {
		return err
	}
	return nil
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users not found", id)
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("res.users was not found")
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users was not found")
}
