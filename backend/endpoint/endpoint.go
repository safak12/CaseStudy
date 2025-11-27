package endpoint

import (
	"context"
	"depixen/backend/database" 
	"depixen/backend/service"  
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	CreateCardEndpoint endpoint.Endpoint
	ListCardsEndpoint  endpoint.Endpoint
}


func New(svc service.CaseStudyService) Endpoints {
	return Endpoints{
		CreateCardEndpoint: makeCreateCardEndpoint(svc),
		ListCardsEndpoint:  makeListCardsEndpoint(svc),
	}
}

func makeCreateCardEndpoint(svc service.CaseStudyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCardRequest)
		card, err := svc.CreateCard(ctx, req.Title, req.Description, req.ImageBase64)
		if err != nil {
			return CreateCardResponse{Error: err.Error()}, nil
		}
		return CreateCardResponse{Card: card}, nil
	}
}


func makeListCardsEndpoint(svc service.CaseStudyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cards, err := svc.ListCards(ctx)
		if err != nil {
			return ListCardsResponse{Error: err.Error()}, nil
		}
		return ListCardsResponse{Cards: cards}, nil
	}
}




type CreateCardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageBase64 string `json:"imageBase64"` 
}


type CreateCardResponse struct {
	Card  database.CaseStudy `json:"card,omitempty"`
	Error string             `json:"error,omitempty"`
}

type ListCardsRequest struct{}


type ListCardsResponse struct {
	Cards []database.CaseStudy `json:"cards,omitempty"`
	Error string               `json:"error,omitempty"`
}