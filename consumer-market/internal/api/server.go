package api

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/gloonch/log-aggregator/consumer-market/internal/proto/github.com/gloonch/log-aggregator/api/candlepb"
	"github.com/gloonch/log-aggregator/consumer-market/internal/store"
)

type CandleServer struct {
	store *store.RedisStore
	pb.UnimplementedCandleServiceServer
}

func NewCandleServer(store *store.RedisStore) *CandleServer {
	return &CandleServer{store: store}
}

func (s *CandleServer) GetCandles(ctx context.Context, req *pb.CandleQuery) (*pb.CandleList, error) {
	key := fmt.Sprintf("market:%s:%s", req.Symbol, req.Timeframe)

	zrange, err := s.store.RangeByScore(ctx, key, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	var result []*pb.Candle
	for _, z := range zrange {
		var candle pb.Candle
		if err := json.Unmarshal([]byte(z.Member.(string)), &candle); err == nil {
			result = append(result, &candle)
		}
	}

	return &pb.CandleList{Candles: result}, nil
}
