package db

import (
	"context"
	"time"

	"github.com/Nextdrive-AaronWu/elec-trade-data/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertTradeDataBatch(ctx context.Context, pool *pgxpool.Pool, data []model.TradeData) error {
	query := `
		INSERT INTO elec_trade_data_tw (
			tran_date, tran_hour, marginal_price, reg_bid, reg_bid_qse, reg_bid_nontrade,
			reg_demand, reg_offering, reg_price, reg_registered,
			sr_bid, sr_bid_qse, sr_bid_nontrade, sr_demand, sr_offering, sr_price, sr_registered,
			sup_bid, sup_bid_qse, sup_bid_nontrade, sup_demand, sup_offering, sup_price, sup_registered,
			edreg_bid, edreg_price
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17,
			$18, $19, $20, $21, $22, $23, $24,
			$25, $26
		)
		ON CONFLICT (tran_date, tran_hour) DO UPDATE SET
			marginal_price = EXCLUDED.marginal_price,
			reg_bid = EXCLUDED.reg_bid,
			reg_bid_qse = EXCLUDED.reg_bid_qse,
			reg_bid_nontrade = EXCLUDED.reg_bid_nontrade,
			reg_demand = EXCLUDED.reg_demand,
			reg_offering = EXCLUDED.reg_offering,
			reg_price = EXCLUDED.reg_price,
			reg_registered = EXCLUDED.reg_registered,
			sr_bid = EXCLUDED.sr_bid,
			sr_bid_qse = EXCLUDED.sr_bid_qse,
			sr_bid_nontrade = EXCLUDED.sr_bid_nontrade,
			sr_demand = EXCLUDED.sr_demand,
			sr_offering = EXCLUDED.sr_offering,
			sr_price = EXCLUDED.sr_price,
			sr_registered = EXCLUDED.sr_registered,
			sup_bid = EXCLUDED.sup_bid,
			sup_bid_qse = EXCLUDED.sup_bid_qse,
			sup_bid_nontrade = EXCLUDED.sup_bid_nontrade,
			sup_demand = EXCLUDED.sup_demand,
			sup_offering = EXCLUDED.sup_offering,
			sup_price = EXCLUDED.sup_price,
			sup_registered = EXCLUDED.sup_registered,
			edreg_bid = EXCLUDED.edreg_bid,
			edreg_price = EXCLUDED.edreg_price
	`
	batch := &pgx.Batch{}
	for _, d := range data {
		tranDate, _ := time.Parse("2006-01-02", d.TranDate)
		tranHour, _ := time.Parse("15:04", d.TranHour)
		batch.Queue(query,
			tranDate.Format("2006-01-02"),
			tranHour.Format("15:04:00"),
			d.MarginalPrice, d.RegBid, d.RegBidQse, d.RegBidNontrade,
			d.RegDemand, d.RegOffering, d.RegPrice, d.RegRegistered,
			d.SrBid, d.SrBidQse, d.SrBidNontrade, d.SrDemand, d.SrOffering, d.SrPrice, d.SrRegistered,
			d.SupBid, d.SupBidQse, d.SupBidNontrade, d.SupDemand, d.SupOffering, d.SupPrice, d.SupRegistered,
			d.EdregBid, d.EdregPrice,
		)
	}
	br := pool.SendBatch(ctx, batch)
	defer br.Close()
	for range data {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}
	return nil
}
