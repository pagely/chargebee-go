package export

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	exportEnum "github.com/chargebee/chargebee-go/models/export/enum"
)

type Export struct {
	Id            string              `json:"id"`
	OperationType string              `json:"operation_type"`
	MimeType      exportEnum.MimeType `json:"mime_type"`
	Status        exportEnum.Status   `json:"status"`
	CreatedAt     int64               `json:"created_at"`
	Download      *Download           `json:"download"`
	Object        string              `json:"object"`
}
type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	Object      string `json:"object"`
}
type RevenueRecognitionRequestParams struct {
	ReportBy         enum.ReportBy                         `json:"report_by"`
	CurrencyCode     string                                `json:"currency_code,omitempty"`
	ReportFromMonth  *int32                                `json:"report_from_month"`
	ReportFromYear   *int32                                `json:"report_from_year"`
	ReportToMonth    *int32                                `json:"report_to_month"`
	ReportToYear     *int32                                `json:"report_to_year"`
	IncludeDiscounts *bool                                 `json:"include_discounts,omitempty"`
	Invoice          *RevenueRecognitionInvoiceParams      `json:"invoice,omitempty"`
	Subscription     *RevenueRecognitionSubscriptionParams `json:"subscription,omitempty"`
	Customer         *RevenueRecognitionCustomerParams     `json:"customer,omitempty"`
}
type RevenueRecognitionInvoiceParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	Recurring      *filter.BooleanFilter   `json:"recurring,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	PriceType      *filter.EnumFilter      `json:"price_type,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	PaidAt         *filter.TimestampFilter `json:"paid_at,omitempty"`
	Total          *filter.NumberFilter    `json:"total,omitempty"`
	AmountPaid     *filter.NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *filter.NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *filter.NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *filter.NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *filter.EnumFilter      `json:"dunning_status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type RevenueRecognitionSubscriptionParams struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type RevenueRecognitionCustomerParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type DeferredRevenueRequestParams struct {
	ReportBy         enum.ReportBy                      `json:"report_by"`
	CurrencyCode     string                             `json:"currency_code,omitempty"`
	ReportFromMonth  *int32                             `json:"report_from_month"`
	ReportFromYear   *int32                             `json:"report_from_year"`
	ReportToMonth    *int32                             `json:"report_to_month"`
	ReportToYear     *int32                             `json:"report_to_year"`
	IncludeDiscounts *bool                              `json:"include_discounts,omitempty"`
	Invoice          *DeferredRevenueInvoiceParams      `json:"invoice,omitempty"`
	Subscription     *DeferredRevenueSubscriptionParams `json:"subscription,omitempty"`
	Customer         *DeferredRevenueCustomerParams     `json:"customer,omitempty"`
}
type DeferredRevenueInvoiceParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	Recurring      *filter.BooleanFilter   `json:"recurring,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	PriceType      *filter.EnumFilter      `json:"price_type,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	PaidAt         *filter.TimestampFilter `json:"paid_at,omitempty"`
	Total          *filter.NumberFilter    `json:"total,omitempty"`
	AmountPaid     *filter.NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *filter.NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *filter.NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *filter.NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *filter.EnumFilter      `json:"dunning_status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type DeferredRevenueSubscriptionParams struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type DeferredRevenueCustomerParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type PlansRequestParams struct {
	Plan *PlansPlanParams `json:"plan,omitempty"`
}
type PlansPlanParams struct {
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	Name               *filter.StringFilter    `json:"name,omitempty"`
	Price              *filter.NumberFilter    `json:"price,omitempty"`
	Period             *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit         *filter.EnumFilter      `json:"period_unit,omitempty"`
	TrialPeriod        *filter.NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit    *filter.EnumFilter      `json:"trial_period_unit,omitempty"`
	AddonApplicability *filter.EnumFilter      `json:"addon_applicability,omitempty"`
	Giftable           *filter.BooleanFilter   `json:"giftable,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type AddonsRequestParams struct {
	Addon *AddonsAddonParams `json:"addon,omitempty"`
}
type AddonsAddonParams struct {
	Id         *filter.StringFilter    `json:"id,omitempty"`
	Name       *filter.StringFilter    `json:"name,omitempty"`
	ChargeType *filter.EnumFilter      `json:"charge_type,omitempty"`
	Price      *filter.NumberFilter    `json:"price,omitempty"`
	Period     *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit *filter.EnumFilter      `json:"period_unit,omitempty"`
	Status     *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt  *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type CouponsRequestParams struct {
	Coupon *CouponsCouponParams `json:"coupon,omitempty"`
}
type CouponsCouponParams struct {
	Id           *filter.StringFilter    `json:"id,omitempty"`
	Name         *filter.StringFilter    `json:"name,omitempty"`
	DiscountType *filter.EnumFilter      `json:"discount_type,omitempty"`
	DurationType *filter.EnumFilter      `json:"duration_type,omitempty"`
	Status       *filter.EnumFilter      `json:"status,omitempty"`
	ApplyOn      *filter.EnumFilter      `json:"apply_on,omitempty"`
	CreatedAt    *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt    *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type CustomersRequestParams struct {
	Customer *CustomersCustomerParams `json:"customer,omitempty"`
}
type CustomersCustomerParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type SubscriptionsRequestParams struct {
	Subscription *SubscriptionsSubscriptionParams `json:"subscription,omitempty"`
}
type SubscriptionsSubscriptionParams struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type InvoicesRequestParams struct {
	Invoice *InvoicesInvoiceParams `json:"invoice,omitempty"`
}
type InvoicesInvoiceParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	Recurring      *filter.BooleanFilter   `json:"recurring,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	PriceType      *filter.EnumFilter      `json:"price_type,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	PaidAt         *filter.TimestampFilter `json:"paid_at,omitempty"`
	Total          *filter.NumberFilter    `json:"total,omitempty"`
	AmountPaid     *filter.NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *filter.NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *filter.NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *filter.NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *filter.EnumFilter      `json:"dunning_status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type CreditNotesRequestParams struct {
	CreditNote *CreditNotesCreditNoteParams `json:"credit_note,omitempty"`
}
type CreditNotesCreditNoteParams struct {
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	CustomerId         *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId     *filter.StringFilter    `json:"subscription_id,omitempty"`
	ReferenceInvoiceId *filter.StringFilter    `json:"reference_invoice_id,omitempty"`
	Type               *filter.EnumFilter      `json:"type,omitempty"`
	ReasonCode         *filter.EnumFilter      `json:"reason_code,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	Date               *filter.TimestampFilter `json:"date,omitempty"`
	Total              *filter.NumberFilter    `json:"total,omitempty"`
	PriceType          *filter.EnumFilter      `json:"price_type,omitempty"`
	AmountAllocated    *filter.NumberFilter    `json:"amount_allocated,omitempty"`
	AmountRefunded     *filter.NumberFilter    `json:"amount_refunded,omitempty"`
	AmountAvailable    *filter.NumberFilter    `json:"amount_available,omitempty"`
	VoidedAt           *filter.TimestampFilter `json:"voided_at,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type TransactionsRequestParams struct {
	Transaction *TransactionsTransactionParams `json:"transaction,omitempty"`
}
type TransactionsTransactionParams struct {
	Id               *filter.StringFilter    `json:"id,omitempty"`
	CustomerId       *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId   *filter.StringFilter    `json:"subscription_id,omitempty"`
	PaymentSourceId  *filter.StringFilter    `json:"payment_source_id,omitempty"`
	PaymentMethod    *filter.EnumFilter      `json:"payment_method,omitempty"`
	Gateway          *filter.EnumFilter      `json:"gateway,omitempty"`
	GatewayAccountId *filter.StringFilter    `json:"gateway_account_id,omitempty"`
	IdAtGateway      *filter.StringFilter    `json:"id_at_gateway,omitempty"`
	ReferenceNumber  *filter.StringFilter    `json:"reference_number,omitempty"`
	Type             *filter.EnumFilter      `json:"type,omitempty"`
	Date             *filter.TimestampFilter `json:"date,omitempty"`
	Amount           *filter.NumberFilter    `json:"amount,omitempty"`
	AmountCapturable *filter.NumberFilter    `json:"amount_capturable,omitempty"`
	Status           *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt        *filter.TimestampFilter `json:"updated_at,omitempty"`
}
