package service

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/nbasker/api_research/sports/proto/sports"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

// Sports interface
type Sports interface {
	// ListEvents will return a collection of events.
	ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService() Sports {
	return &sportsService{}
}

func (s *sportsService) ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error) {
	var events []*sports.Event
	log.Debugf("sportsService::ListEvents() called...")
	ts1, err := ptypes.TimestampProto(time.Now().UTC())
	if err != nil {
		return nil, err
	}
	events = append(events, &sports.Event{
		Id:                  1,
		Name:                "JavelinThrow",
		AdvertisedStartTime: ts1,
	})
	ts2, err := ptypes.TimestampProto(time.Now().UTC())
	if err != nil {
		return nil, err
	}
	events = append(events, &sports.Event{
		Id:                  2,
		Name:                "SyncSwim",
		AdvertisedStartTime: ts2,
	})
	log.Debugf("sportsService::ListEvents() returning...events = %v", events)
	return &sports.ListEventsResponse{Events: events}, nil
}
