package sale_order_dates

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.SaleOrder().DeclareModel()

	h.SaleOrder().AddFields(map[string]models.FieldDefinition{
		"CommitmentDate": models.DateTimeField{
			Compute: h.SaleOrder().Methods().ComputeCommitmentDate(),
			String:  "Commitment Date",
			Stored:  true,
			Help: "Date by which the products are sure to be delivered. This" +
				"is a date that you can promise to the customer, based on" +
				"the Product Lead Times.",
		},
		"RequestedDate": models.DateTimeField{
			String:   "Requested Date",
			ReadOnly: true,
			//states={'draft': [('readonly', False)],'sent': [('readonly', False)]}
			NoCopy: true,
			Help: "Date by which the customer has requested the items to be delivered." +
				"When this Order gets confirmed, the Delivery Order's expected" +
				"date will be computed based on this date and the Company's" +
				"Security Delay." +
				"Leave this field empty if you want the Delivery Order to" +
				"be processed as soon as possible. In that case the expected" +
				"date will be computed using the default method: based on" +
				"the Product Lead Times and the Company's Security Delay.",
		},
		"EffectiveDate": models.DateField{
			Compute: h.SaleOrder().Methods().ComputePickingIds(),
			String:  "Effective Date",
			Stored:  true,
			Help:    "Date on which the first Delivery Order was created.",
		},
	})
	h.SaleOrder().Methods().ComputeCommitmentDate().DeclareMethod(
		`Compute the commitment date`,
		func(rs h.SaleOrderSet) h.SaleOrderData {
			//        for order in self:
			//            dates_list = []
			//            order_datetime = fields.Datetime.from_string(order.date_order)
			//            for line in order.order_line.filtered(lambda x: x.state != 'cancel'):
			//                dt = order_datetime + timedelta(days=line.customer_lead or 0.0)
			//                dates_list.append(dt)
			//            if dates_list:
			//                commit_date = min(
			//                    dates_list) if order.picking_policy == 'direct' else max(dates_list)
			//                order.commitment_date = fields.Datetime.to_string(commit_date)
		})
	h.SaleOrder().Methods().ComputePickingIds().DeclareMethod(
		`ComputePickingIds`,
		func(rs h.SaleOrderSet) h.SaleOrderData {
			//        super(SaleOrder, self)._compute_picking_ids()
			//        for order in self:
			//            dates_list = []
			//            for pick in order.picking_ids:
			//                dates_list.append(fields.Datetime.from_string(pick.date))
			//            if dates_list:
			//                order.effective_date = fields.Datetime.to_string(
			//                    min(dates_list))
		})
	h.SaleOrder().Methods().OnchangeRequestedDate().DeclareMethod(
		`Warn if the requested dates is sooner than the commitment date`,
		func(rs m.SaleOrderSet) {
			//        if (self.requested_date and self.commitment_date and self.requested_date < self.commitment_date):
			//            return {'warning': {
			//                'title': _('Requested date is too soon!'),
			//                'message': _("The date requested by the customer is "
			//                             "sooner than the commitment date. You may be "
			//                             "unable to honor the customer's request.")
			//            }
			//            }
		})
	h.SaleOrderLine().DeclareModel()

	h.SaleOrderLine().Methods().PrepareOrderLineProcurement().DeclareMethod(
		`PrepareOrderLineProcurement`,
		func(rs m.SaleOrderLineSet, group_id interface{}) {
			//        vals = super(SaleOrderLine, self)._prepare_order_line_procurement(
			//            group_id=group_id)
			//        for line in self.filtered("order_id.requested_date"):
			//            date_planned = fields.Datetime.from_string(
			//                line.order_id.requested_date) - timedelta(days=line.order_id.company_id.security_lead)
			//            vals.update({
			//                'date_planned': fields.Datetime.to_string(date_planned),
			//            })
			//        return vals
		})
}
