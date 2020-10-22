package chargebee

import (
	"github.com/pagely/chargebee-go/models/addon"
	"github.com/pagely/chargebee-go/models/address"
	"github.com/pagely/chargebee-go/models/advanceinvoiceschedule"
	"github.com/pagely/chargebee-go/models/card"
	"github.com/pagely/chargebee-go/models/comment"
	"github.com/pagely/chargebee-go/models/contact"
	"github.com/pagely/chargebee-go/models/contractterm"
	"github.com/pagely/chargebee-go/models/coupon"
	"github.com/pagely/chargebee-go/models/couponcode"
	"github.com/pagely/chargebee-go/models/couponset"
	"github.com/pagely/chargebee-go/models/creditnote"
	"github.com/pagely/chargebee-go/models/customer"
	"github.com/pagely/chargebee-go/models/download"
	"github.com/pagely/chargebee-go/models/estimate"
	"github.com/pagely/chargebee-go/models/event"
	"github.com/pagely/chargebee-go/models/export"
	"github.com/pagely/chargebee-go/models/gift"
	"github.com/pagely/chargebee-go/models/hierarchy"
	"github.com/pagely/chargebee-go/models/hostedpage"
	"github.com/pagely/chargebee-go/models/invoice"
	"github.com/pagely/chargebee-go/models/order"
	"github.com/pagely/chargebee-go/models/paymentintent"
	"github.com/pagely/chargebee-go/models/paymentsource"
	"github.com/pagely/chargebee-go/models/plan"
	"github.com/pagely/chargebee-go/models/portalsession"
	"github.com/pagely/chargebee-go/models/promotionalcredit"
	"github.com/pagely/chargebee-go/models/quote"
	"github.com/pagely/chargebee-go/models/quotedsubscription"
	"github.com/pagely/chargebee-go/models/quotelinegroup"
	"github.com/pagely/chargebee-go/models/resourcemigration"
	"github.com/pagely/chargebee-go/models/sitemigrationdetail"
	"github.com/pagely/chargebee-go/models/subscription"
	"github.com/pagely/chargebee-go/models/thirdpartypaymentmethod"
	"github.com/pagely/chargebee-go/models/timemachine"
	"github.com/pagely/chargebee-go/models/token"
	"github.com/pagely/chargebee-go/models/transaction"
	"github.com/pagely/chargebee-go/models/unbilledcharge"
	"github.com/pagely/chargebee-go/models/virtualbankaccount"
)

type ResultList struct {
	List       []*Result `json:"list"`
	NextOffset string    `json:"next_offset"`
}
type Result struct {
	Subscription            *subscription.Subscription                       `json:"subscription,omitempty"`
	ContractTerm            *contractterm.ContractTerm                       `json:"contract_term,omitempty"`
	AdvanceInvoiceSchedule  *advanceinvoiceschedule.AdvanceInvoiceSchedule   `json:"advance_invoice_schedule,omitempty"`
	Customer                *customer.Customer                               `json:"customer,omitempty"`
	Hierarchy               *hierarchy.Hierarchy                             `json:"hierarchy,omitempty"`
	Contact                 *contact.Contact                                 `json:"contact,omitempty"`
	Token                   *token.Token                                     `json:"token,omitempty"`
	PaymentSource           *paymentsource.PaymentSource                     `json:"payment_source,omitempty"`
	ThirdPartyPaymentMethod *thirdpartypaymentmethod.ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
	VirtualBankAccount      *virtualbankaccount.VirtualBankAccount           `json:"virtual_bank_account,omitempty"`
	Card                    *card.Card                                       `json:"card,omitempty"`
	PromotionalCredit       *promotionalcredit.PromotionalCredit             `json:"promotional_credit,omitempty"`
	Invoice                 *invoice.Invoice                                 `json:"invoice,omitempty"`
	CreditNote              *creditnote.CreditNote                           `json:"credit_note,omitempty"`
	UnbilledCharge          *unbilledcharge.UnbilledCharge                   `json:"unbilled_charge,omitempty"`
	Order                   *order.Order                                     `json:"order,omitempty"`
	Gift                    *gift.Gift                                       `json:"gift,omitempty"`
	Transaction             *transaction.Transaction                         `json:"transaction,omitempty"`
	HostedPage              *hostedpage.HostedPage                           `json:"hosted_page,omitempty"`
	Estimate                *estimate.Estimate                               `json:"estimate,omitempty"`
	Quote                   *quote.Quote                                     `json:"quote,omitempty"`
	QuotedSubscription      *quotedsubscription.QuotedSubscription           `json:"quoted_subscription,omitempty"`
	QuoteLineGroup          *quotelinegroup.QuoteLineGroup                   `json:"quote_line_group,omitempty"`
	Plan                    *plan.Plan                                       `json:"plan,omitempty"`
	Addon                   *addon.Addon                                     `json:"addon,omitempty"`
	Coupon                  *coupon.Coupon                                   `json:"coupon,omitempty"`
	CouponSet               *couponset.CouponSet                             `json:"coupon_set,omitempty"`
	CouponCode              *couponcode.CouponCode                           `json:"coupon_code,omitempty"`
	Address                 *address.Address                                 `json:"address,omitempty"`
	Event                   *event.Event                                     `json:"event,omitempty"`
	Comment                 *comment.Comment                                 `json:"comment,omitempty"`
	Download                *download.Download                               `json:"download,omitempty"`
	PortalSession           *portalsession.PortalSession                     `json:"portal_session,omitempty"`
	SiteMigrationDetail     *sitemigrationdetail.SiteMigrationDetail         `json:"site_migration_detail,omitempty"`
	ResourceMigration       *resourcemigration.ResourceMigration             `json:"resource_migration,omitempty"`
	TimeMachine             *timemachine.TimeMachine                         `json:"time_machine,omitempty"`
	Export                  *export.Export                                   `json:"export,omitempty"`
	PaymentIntent           *paymentintent.PaymentIntent                     `json:"payment_intent,omitempty"`
	UnbilledCharges         []*unbilledcharge.UnbilledCharge                 `json:"unbilled_charges,omitempty"`
	CreditNotes             []*creditnote.CreditNote                         `json:"credit_notes,omitempty"`
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	Hierarchies             []*hierarchy.Hierarchy                           `json:"hierarchies,omitempty"`
	Invoices                []*invoice.Invoice                               `json:"invoices,omitempty"`
}
