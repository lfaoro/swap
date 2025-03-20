package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	pb "github.com/lfaoro/swap/gen/go/swap/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SwapAPI struct {
	c      pb.CoinServiceClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSwapAPI(client pb.CoinServiceClient) SwapAPI {
	ctx, cancel := context.WithCancel(context.Background())
	return SwapAPI{c: client, ctx: ctx, cancel: cancel}
}

func (s *SwapAPI) Close() {
	s.cancel()
	s.ctx, s.cancel = context.WithCancel(context.Background())
}

type CoinReqRespMsg struct {
	Resp []*pb.Coin
}

func (s SwapAPI) ListCoins() tea.Cmd {
	return func() tea.Msg {
		resp, err := s.c.ListCoins(s.ctx, &emptypb.Empty{})
		if err != nil {
			return tea.Batch(
				AddLog("api: error (ListCoins): %v", err),
				AddError(fmt.Errorf("connection error: unable to load coins table")),
			)()
		}
		return tea.Batch(
			AddLog("api: success (ListCoins)"),
			func() tea.Msg {
				return CoinReqRespMsg{Resp: resp.GetCoins()}
			},
		)()
	}
}

type SwapRateReqMsg struct {
	AmountFrom  string
	AmountTo    string
	TickerFrom  string
	NetworkFrom string
	TickerTo    string
	NetworkTo   string
	Payment     bool
}
type SwapRateRespMsg struct {
	Resp *pb.SwapRateResponse
}

func (s SwapAPI) SwapRate(req SwapRateReqMsg) tea.Cmd {
	return func() tea.Msg {
		resp, err := s.c.SwapRate(s.ctx, &pb.SwapRateRequest{
			AmountFrom:  parseFloat(req.AmountFrom),
			AmountTo:    parseFloat(req.AmountTo),
			TickerFrom:  req.TickerFrom,
			NetworkFrom: req.NetworkFrom,
			TickerTo:    req.TickerTo,
			NetworkTo:   req.NetworkTo,
			Payment:     req.Payment,
		})
		if err != nil {
			// Extract error message from gRPC error
			errMsg := err.Error()
			if start := strings.Index(errMsg, "desc = "); start != -1 {
				errMsg = errMsg[start+len("desc = "):]
				// Try to parse as JSON to get clean error message
				var errObj struct {
					Error string `json:"error"`
				}
				if err := json.Unmarshal([]byte(errMsg), &errObj); err == nil {
					errMsg = errObj.Error
				}
			}

			return tea.Batch(
				AddLog("api: error (SwapRate): %v", err),
				AddError(fmt.Errorf(errMsg)),
			)()
		}
		return tea.Batch(
			AddLog("api: success (SwapRate): %v", resp.GetStatus()),
			func() tea.Msg {
				return SwapRateRespMsg{Resp: resp}
			},
		)()
	}
}

type SwapTradeReq struct {
	*CoinData
}
type SwapTradeRespMsg struct {
	*pb.SwapTradeResponse
}

func (s SwapAPI) SwapTrade(req SwapTradeReq) tea.Cmd {
	return func() tea.Msg {
		req := &pb.SwapTradeRequest{
			Id:          req.TradeID,
			TickerFrom:  req.From.Ticker,
			TickerTo:    req.To.Ticker,
			NetworkFrom: req.From.Network,
			NetworkTo:   req.To.Network,
			AmountFrom:  parseFloat(req.Amount.Value()),
			AmountTo:    parseFloat(req.Amount.Value()),
			Payment:     req.Payment,
			Address:     req.Address.Value(),
			Refund:      "",
			Provider:    req.Exchange,
		}
		resp, err := s.c.SwapTrade(s.ctx, req)
		if err != nil {
			return tea.Batch(
				AddLog("api: error (SwapTrade): %v", err),
				AddError(err),
			)()
		}
		return tea.Batch(
			AddLog("api: success (SwapTrade) %v", resp.GetStatus()),
			func() tea.Msg {
				return SwapTradeRespMsg{resp}
			},
		)()
	}
}

type SwapStatusReqMsg struct {
	TradeID string
}
type SwapStatusRespMsg struct {
	*pb.SwapStatusResponse
}

func (s SwapAPI) SwapStatus(req SwapStatusReqMsg) tea.Cmd {
	return func() tea.Msg {
		resp, err := s.c.SwapStatus(s.ctx, &pb.SwapStatusRequest{TradeId: req.TradeID})
		if err != nil {
			return tea.Batch(
				AddLog("api: error (SwapStatus): %v", err),
				AddError(err),
			)()
		}

		// This section uses Recursion to handle the stream.
		var recvCmd tea.Cmd
		recvCmd = func() tea.Msg {
			status, err := resp.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				if s.ctx.Err() != nil {
					return nil
				}
				return tea.Batch(
					AddLog("api: error (SwapStatus): %v", err),
					AddError(err),
				)()
			}

			return tea.Batch(
				AddLog("api: success (SwapStatus): %v", status.GetStatus()),
				func() tea.Msg { return SwapStatusRespMsg{status} },
				recvCmd, // Chain the next receive on the same stream
			)()
		}
		return recvCmd()
	}
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
