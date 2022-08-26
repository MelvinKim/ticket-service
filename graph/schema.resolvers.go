package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MelvinKim/ticket-app-golang-graphql/graph/generated"
	"github.com/MelvinKim/ticket-app-golang-graphql/graph/model"
)

// CreateTicket is the resolver for the createTicket field.
func (r *mutationResolver) CreateTicket(ctx context.Context, input model.NewTicket) (*model.Ticket, error) {
	ticket := model.Ticket{
		Type:   "Support ticket",
		Status: false,
	}

	_, err := r.DB.Model(&ticket).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new ticket: %v", err)
	}

	return &ticket, nil
}

// Tickets is the resolver for the tickets field.
func (r *queryResolver) Tickets(ctx context.Context) ([]*model.Ticket, error) {
	var tickets []*model.Ticket

	err := r.DB.Model(&tickets).Select()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
