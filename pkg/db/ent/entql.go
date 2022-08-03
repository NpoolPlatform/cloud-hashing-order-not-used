// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/goodpaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   compensate.Table,
			Columns: compensate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compensate.FieldID,
			},
		},
		Type: "Compensate",
		Fields: map[string]*sqlgraph.FieldSpec{
			compensate.FieldOrderID:  {Type: field.TypeUUID, Column: compensate.FieldOrderID},
			compensate.FieldStart:    {Type: field.TypeUint32, Column: compensate.FieldStart},
			compensate.FieldEnd:      {Type: field.TypeUint32, Column: compensate.FieldEnd},
			compensate.FieldMessage:  {Type: field.TypeString, Column: compensate.FieldMessage},
			compensate.FieldCreateAt: {Type: field.TypeUint32, Column: compensate.FieldCreateAt},
			compensate.FieldUpdateAt: {Type: field.TypeUint32, Column: compensate.FieldUpdateAt},
			compensate.FieldDeleteAt: {Type: field.TypeUint32, Column: compensate.FieldDeleteAt},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   gaspaying.Table,
			Columns: gaspaying.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gaspaying.FieldID,
			},
		},
		Type: "GasPaying",
		Fields: map[string]*sqlgraph.FieldSpec{
			gaspaying.FieldOrderID:         {Type: field.TypeUUID, Column: gaspaying.FieldOrderID},
			gaspaying.FieldFeeTypeID:       {Type: field.TypeUUID, Column: gaspaying.FieldFeeTypeID},
			gaspaying.FieldPaymentID:       {Type: field.TypeUUID, Column: gaspaying.FieldPaymentID},
			gaspaying.FieldDurationMinutes: {Type: field.TypeUint32, Column: gaspaying.FieldDurationMinutes},
			gaspaying.FieldCreateAt:        {Type: field.TypeUint32, Column: gaspaying.FieldCreateAt},
			gaspaying.FieldUpdateAt:        {Type: field.TypeUint32, Column: gaspaying.FieldUpdateAt},
			gaspaying.FieldDeleteAt:        {Type: field.TypeUint32, Column: gaspaying.FieldDeleteAt},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   goodpaying.Table,
			Columns: goodpaying.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodpaying.FieldID,
			},
		},
		Type: "GoodPaying",
		Fields: map[string]*sqlgraph.FieldSpec{
			goodpaying.FieldOrderID:   {Type: field.TypeUUID, Column: goodpaying.FieldOrderID},
			goodpaying.FieldPaymentID: {Type: field.TypeUUID, Column: goodpaying.FieldPaymentID},
			goodpaying.FieldCreateAt:  {Type: field.TypeUint32, Column: goodpaying.FieldCreateAt},
			goodpaying.FieldUpdateAt:  {Type: field.TypeUint32, Column: goodpaying.FieldUpdateAt},
			goodpaying.FieldDeleteAt:  {Type: field.TypeUint32, Column: goodpaying.FieldDeleteAt},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
		Type: "Order",
		Fields: map[string]*sqlgraph.FieldSpec{
			order.FieldGoodID:                 {Type: field.TypeUUID, Column: order.FieldGoodID},
			order.FieldAppID:                  {Type: field.TypeUUID, Column: order.FieldAppID},
			order.FieldUserID:                 {Type: field.TypeUUID, Column: order.FieldUserID},
			order.FieldParentID:               {Type: field.TypeUUID, Column: order.FieldParentID},
			order.FieldUnits:                  {Type: field.TypeUint32, Column: order.FieldUnits},
			order.FieldPromotionID:            {Type: field.TypeUUID, Column: order.FieldPromotionID},
			order.FieldDiscountCouponID:       {Type: field.TypeUUID, Column: order.FieldDiscountCouponID},
			order.FieldUserSpecialReductionID: {Type: field.TypeUUID, Column: order.FieldUserSpecialReductionID},
			order.FieldStart:                  {Type: field.TypeUint32, Column: order.FieldStart},
			order.FieldEnd:                    {Type: field.TypeUint32, Column: order.FieldEnd},
			order.FieldCouponID:               {Type: field.TypeUUID, Column: order.FieldCouponID},
			order.FieldOrderType:              {Type: field.TypeString, Column: order.FieldOrderType},
			order.FieldCreateAt:               {Type: field.TypeUint32, Column: order.FieldCreateAt},
			order.FieldUpdateAt:               {Type: field.TypeUint32, Column: order.FieldUpdateAt},
			order.FieldDeleteAt:               {Type: field.TypeUint32, Column: order.FieldDeleteAt},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
		Type: "OutOfGas",
		Fields: map[string]*sqlgraph.FieldSpec{
			outofgas.FieldOrderID:  {Type: field.TypeUUID, Column: outofgas.FieldOrderID},
			outofgas.FieldStart:    {Type: field.TypeUint32, Column: outofgas.FieldStart},
			outofgas.FieldEnd:      {Type: field.TypeUint32, Column: outofgas.FieldEnd},
			outofgas.FieldCreateAt: {Type: field.TypeUint32, Column: outofgas.FieldCreateAt},
			outofgas.FieldUpdateAt: {Type: field.TypeUint32, Column: outofgas.FieldUpdateAt},
			outofgas.FieldDeleteAt: {Type: field.TypeUint32, Column: outofgas.FieldDeleteAt},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		},
		Type: "Payment",
		Fields: map[string]*sqlgraph.FieldSpec{
			payment.FieldAppID:                 {Type: field.TypeUUID, Column: payment.FieldAppID},
			payment.FieldUserID:                {Type: field.TypeUUID, Column: payment.FieldUserID},
			payment.FieldGoodID:                {Type: field.TypeUUID, Column: payment.FieldGoodID},
			payment.FieldOrderID:               {Type: field.TypeUUID, Column: payment.FieldOrderID},
			payment.FieldAccountID:             {Type: field.TypeUUID, Column: payment.FieldAccountID},
			payment.FieldStartAmount:           {Type: field.TypeUint64, Column: payment.FieldStartAmount},
			payment.FieldAmount:                {Type: field.TypeUint64, Column: payment.FieldAmount},
			payment.FieldPayWithBalanceAmount:  {Type: field.TypeFloat64, Column: payment.FieldPayWithBalanceAmount},
			payment.FieldFinishAmount:          {Type: field.TypeUint64, Column: payment.FieldFinishAmount},
			payment.FieldCoinUsdCurrency:       {Type: field.TypeUint64, Column: payment.FieldCoinUsdCurrency},
			payment.FieldLocalCoinUsdCurrency:  {Type: field.TypeUint64, Column: payment.FieldLocalCoinUsdCurrency},
			payment.FieldLiveCoinUsdCurrency:   {Type: field.TypeUint64, Column: payment.FieldLiveCoinUsdCurrency},
			payment.FieldCoinInfoID:            {Type: field.TypeUUID, Column: payment.FieldCoinInfoID},
			payment.FieldState:                 {Type: field.TypeEnum, Column: payment.FieldState},
			payment.FieldChainTransactionID:    {Type: field.TypeString, Column: payment.FieldChainTransactionID},
			payment.FieldPlatformTransactionID: {Type: field.TypeUUID, Column: payment.FieldPlatformTransactionID},
			payment.FieldUserSetPaid:           {Type: field.TypeBool, Column: payment.FieldUserSetPaid},
			payment.FieldUserSetCanceled:       {Type: field.TypeBool, Column: payment.FieldUserSetCanceled},
			payment.FieldUserPaymentTxid:       {Type: field.TypeString, Column: payment.FieldUserPaymentTxid},
			payment.FieldFakePayment:           {Type: field.TypeBool, Column: payment.FieldFakePayment},
			payment.FieldCreateAt:              {Type: field.TypeUint32, Column: payment.FieldCreateAt},
			payment.FieldUpdateAt:              {Type: field.TypeUint32, Column: payment.FieldUpdateAt},
			payment.FieldDeleteAt:              {Type: field.TypeUint32, Column: payment.FieldDeleteAt},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CompensateQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CompensateQuery builder.
func (cq *CompensateQuery) Filter() *CompensateFilter {
	return &CompensateFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CompensateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CompensateMutation builder.
func (m *CompensateMutation) Filter() *CompensateFilter {
	return &CompensateFilter{config: m.config, predicateAdder: m}
}

// CompensateFilter provides a generic filtering capability at runtime for CompensateQuery.
type CompensateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CompensateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CompensateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(compensate.FieldID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *CompensateFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(compensate.FieldOrderID))
}

// WhereStart applies the entql uint32 predicate on the start field.
func (f *CompensateFilter) WhereStart(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldStart))
}

// WhereEnd applies the entql uint32 predicate on the end field.
func (f *CompensateFilter) WhereEnd(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldEnd))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *CompensateFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(compensate.FieldMessage))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *CompensateFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *CompensateFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *CompensateFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(compensate.FieldDeleteAt))
}

// addPredicate implements the predicateAdder interface.
func (gpq *GasPayingQuery) addPredicate(pred func(s *sql.Selector)) {
	gpq.predicates = append(gpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GasPayingQuery builder.
func (gpq *GasPayingQuery) Filter() *GasPayingFilter {
	return &GasPayingFilter{config: gpq.config, predicateAdder: gpq}
}

// addPredicate implements the predicateAdder interface.
func (m *GasPayingMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GasPayingMutation builder.
func (m *GasPayingMutation) Filter() *GasPayingFilter {
	return &GasPayingFilter{config: m.config, predicateAdder: m}
}

// GasPayingFilter provides a generic filtering capability at runtime for GasPayingQuery.
type GasPayingFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GasPayingFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *GasPayingFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(gaspaying.FieldID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *GasPayingFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(gaspaying.FieldOrderID))
}

// WhereFeeTypeID applies the entql [16]byte predicate on the fee_type_id field.
func (f *GasPayingFilter) WhereFeeTypeID(p entql.ValueP) {
	f.Where(p.Field(gaspaying.FieldFeeTypeID))
}

// WherePaymentID applies the entql [16]byte predicate on the payment_id field.
func (f *GasPayingFilter) WherePaymentID(p entql.ValueP) {
	f.Where(p.Field(gaspaying.FieldPaymentID))
}

// WhereDurationMinutes applies the entql uint32 predicate on the duration_minutes field.
func (f *GasPayingFilter) WhereDurationMinutes(p entql.Uint32P) {
	f.Where(p.Field(gaspaying.FieldDurationMinutes))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *GasPayingFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(gaspaying.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *GasPayingFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(gaspaying.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *GasPayingFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(gaspaying.FieldDeleteAt))
}

// addPredicate implements the predicateAdder interface.
func (gpq *GoodPayingQuery) addPredicate(pred func(s *sql.Selector)) {
	gpq.predicates = append(gpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GoodPayingQuery builder.
func (gpq *GoodPayingQuery) Filter() *GoodPayingFilter {
	return &GoodPayingFilter{config: gpq.config, predicateAdder: gpq}
}

// addPredicate implements the predicateAdder interface.
func (m *GoodPayingMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GoodPayingMutation builder.
func (m *GoodPayingMutation) Filter() *GoodPayingFilter {
	return &GoodPayingFilter{config: m.config, predicateAdder: m}
}

// GoodPayingFilter provides a generic filtering capability at runtime for GoodPayingQuery.
type GoodPayingFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GoodPayingFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *GoodPayingFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(goodpaying.FieldID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *GoodPayingFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(goodpaying.FieldOrderID))
}

// WherePaymentID applies the entql [16]byte predicate on the payment_id field.
func (f *GoodPayingFilter) WherePaymentID(p entql.ValueP) {
	f.Where(p.Field(goodpaying.FieldPaymentID))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *GoodPayingFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(goodpaying.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *GoodPayingFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(goodpaying.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *GoodPayingFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(goodpaying.FieldDeleteAt))
}

// addPredicate implements the predicateAdder interface.
func (oq *OrderQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrderQuery builder.
func (oq *OrderQuery) Filter() *OrderFilter {
	return &OrderFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrderMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrderMutation builder.
func (m *OrderMutation) Filter() *OrderFilter {
	return &OrderFilter{config: m.config, predicateAdder: m}
}

// OrderFilter provides a generic filtering capability at runtime for OrderQuery.
type OrderFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrderFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OrderFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(order.FieldID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *OrderFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(order.FieldGoodID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *OrderFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(order.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *OrderFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(order.FieldUserID))
}

// WhereParentID applies the entql [16]byte predicate on the parent_id field.
func (f *OrderFilter) WhereParentID(p entql.ValueP) {
	f.Where(p.Field(order.FieldParentID))
}

// WhereUnits applies the entql uint32 predicate on the units field.
func (f *OrderFilter) WhereUnits(p entql.Uint32P) {
	f.Where(p.Field(order.FieldUnits))
}

// WherePromotionID applies the entql [16]byte predicate on the promotion_id field.
func (f *OrderFilter) WherePromotionID(p entql.ValueP) {
	f.Where(p.Field(order.FieldPromotionID))
}

// WhereDiscountCouponID applies the entql [16]byte predicate on the discount_coupon_id field.
func (f *OrderFilter) WhereDiscountCouponID(p entql.ValueP) {
	f.Where(p.Field(order.FieldDiscountCouponID))
}

// WhereUserSpecialReductionID applies the entql [16]byte predicate on the user_special_reduction_id field.
func (f *OrderFilter) WhereUserSpecialReductionID(p entql.ValueP) {
	f.Where(p.Field(order.FieldUserSpecialReductionID))
}

// WhereStart applies the entql uint32 predicate on the start field.
func (f *OrderFilter) WhereStart(p entql.Uint32P) {
	f.Where(p.Field(order.FieldStart))
}

// WhereEnd applies the entql uint32 predicate on the end field.
func (f *OrderFilter) WhereEnd(p entql.Uint32P) {
	f.Where(p.Field(order.FieldEnd))
}

// WhereCouponID applies the entql [16]byte predicate on the coupon_id field.
func (f *OrderFilter) WhereCouponID(p entql.ValueP) {
	f.Where(p.Field(order.FieldCouponID))
}

// WhereOrderType applies the entql string predicate on the order_type field.
func (f *OrderFilter) WhereOrderType(p entql.StringP) {
	f.Where(p.Field(order.FieldOrderType))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *OrderFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *OrderFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *OrderFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(order.FieldDeleteAt))
}

// addPredicate implements the predicateAdder interface.
func (oogq *OutOfGasQuery) addPredicate(pred func(s *sql.Selector)) {
	oogq.predicates = append(oogq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OutOfGasQuery builder.
func (oogq *OutOfGasQuery) Filter() *OutOfGasFilter {
	return &OutOfGasFilter{config: oogq.config, predicateAdder: oogq}
}

// addPredicate implements the predicateAdder interface.
func (m *OutOfGasMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OutOfGasMutation builder.
func (m *OutOfGasMutation) Filter() *OutOfGasFilter {
	return &OutOfGasFilter{config: m.config, predicateAdder: m}
}

// OutOfGasFilter provides a generic filtering capability at runtime for OutOfGasQuery.
type OutOfGasFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OutOfGasFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OutOfGasFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(outofgas.FieldID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *OutOfGasFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(outofgas.FieldOrderID))
}

// WhereStart applies the entql uint32 predicate on the start field.
func (f *OutOfGasFilter) WhereStart(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldStart))
}

// WhereEnd applies the entql uint32 predicate on the end field.
func (f *OutOfGasFilter) WhereEnd(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldEnd))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *OutOfGasFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *OutOfGasFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *OutOfGasFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(outofgas.FieldDeleteAt))
}

// addPredicate implements the predicateAdder interface.
func (pq *PaymentQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PaymentQuery builder.
func (pq *PaymentQuery) Filter() *PaymentFilter {
	return &PaymentFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PaymentMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PaymentMutation builder.
func (m *PaymentMutation) Filter() *PaymentFilter {
	return &PaymentFilter{config: m.config, predicateAdder: m}
}

// PaymentFilter provides a generic filtering capability at runtime for PaymentQuery.
type PaymentFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PaymentFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PaymentFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *PaymentFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *PaymentFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldUserID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *PaymentFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldGoodID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *PaymentFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldOrderID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *PaymentFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldAccountID))
}

// WhereStartAmount applies the entql uint64 predicate on the start_amount field.
func (f *PaymentFilter) WhereStartAmount(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldStartAmount))
}

// WhereAmount applies the entql uint64 predicate on the amount field.
func (f *PaymentFilter) WhereAmount(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldAmount))
}

// WherePayWithBalanceAmount applies the entql float64 predicate on the pay_with_balance_amount field.
func (f *PaymentFilter) WherePayWithBalanceAmount(p entql.Float64P) {
	f.Where(p.Field(payment.FieldPayWithBalanceAmount))
}

// WhereFinishAmount applies the entql uint64 predicate on the finish_amount field.
func (f *PaymentFilter) WhereFinishAmount(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldFinishAmount))
}

// WhereCoinUsdCurrency applies the entql uint64 predicate on the coin_usd_currency field.
func (f *PaymentFilter) WhereCoinUsdCurrency(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldCoinUsdCurrency))
}

// WhereLocalCoinUsdCurrency applies the entql uint64 predicate on the local_coin_usd_currency field.
func (f *PaymentFilter) WhereLocalCoinUsdCurrency(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldLocalCoinUsdCurrency))
}

// WhereLiveCoinUsdCurrency applies the entql uint64 predicate on the live_coin_usd_currency field.
func (f *PaymentFilter) WhereLiveCoinUsdCurrency(p entql.Uint64P) {
	f.Where(p.Field(payment.FieldLiveCoinUsdCurrency))
}

// WhereCoinInfoID applies the entql [16]byte predicate on the coin_info_id field.
func (f *PaymentFilter) WhereCoinInfoID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldCoinInfoID))
}

// WhereState applies the entql string predicate on the state field.
func (f *PaymentFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(payment.FieldState))
}

// WhereChainTransactionID applies the entql string predicate on the chain_transaction_id field.
func (f *PaymentFilter) WhereChainTransactionID(p entql.StringP) {
	f.Where(p.Field(payment.FieldChainTransactionID))
}

// WherePlatformTransactionID applies the entql [16]byte predicate on the platform_transaction_id field.
func (f *PaymentFilter) WherePlatformTransactionID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldPlatformTransactionID))
}

// WhereUserSetPaid applies the entql bool predicate on the user_set_paid field.
func (f *PaymentFilter) WhereUserSetPaid(p entql.BoolP) {
	f.Where(p.Field(payment.FieldUserSetPaid))
}

// WhereUserSetCanceled applies the entql bool predicate on the user_set_canceled field.
func (f *PaymentFilter) WhereUserSetCanceled(p entql.BoolP) {
	f.Where(p.Field(payment.FieldUserSetCanceled))
}

// WhereUserPaymentTxid applies the entql string predicate on the user_payment_txid field.
func (f *PaymentFilter) WhereUserPaymentTxid(p entql.StringP) {
	f.Where(p.Field(payment.FieldUserPaymentTxid))
}

// WhereFakePayment applies the entql bool predicate on the fake_payment field.
func (f *PaymentFilter) WhereFakePayment(p entql.BoolP) {
	f.Where(p.Field(payment.FieldFakePayment))
}

// WhereCreateAt applies the entql uint32 predicate on the create_at field.
func (f *PaymentFilter) WhereCreateAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldCreateAt))
}

// WhereUpdateAt applies the entql uint32 predicate on the update_at field.
func (f *PaymentFilter) WhereUpdateAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldUpdateAt))
}

// WhereDeleteAt applies the entql uint32 predicate on the delete_at field.
func (f *PaymentFilter) WhereDeleteAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldDeleteAt))
}
