package services

import (
	"context"
	"errors"

	"github.com/endrio-maciel/rockeat-go.git/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BidsService struct {
	poll    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewBidsService(poll *pgxpool.Pool) BidsService {
	return BidsService{
		poll:    poll,
		queries: pgstore.New(poll),
	}
}

var ErrBidIsTooLow = errors.New("the bid value is too low")

func (bs *BidsService) PlaceBid(ctx context.Context, product_id, bidder_id uuid.UUID, amount float64) (pgstore.Bid, error) {
	product, err := bs.queries.GetProductById(ctx, product_id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	highesBid, err := bs.queries.GetHighestBidByProductId(ctx, product_id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	if product.Baseprice >= amount || highesBid.BidAmount >= amount {
		return pgstore.Bid{}, ErrBidIsTooLow
	}

	highesBid, err = bs.queries.CreateBid(ctx, pgstore.CreateBidParams{
		ProductID: product_id,
		BidderID:  bidder_id,
		BidAmount: amount,
	})

	if err != nil {
		return pgstore.Bid{}, err
	}

	return highesBid, nil

}
