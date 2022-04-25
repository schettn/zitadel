package command

import (
	"context"

	"github.com/caos/zitadel/internal/command/preparation"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/repository/instance"
)

func AddPrivacyPolicy(
	a *instance.Aggregate,
	tosLink,
	privacyLink,
	helpLink string,
) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			//TODO: check if already exists
			return []eventstore.Command{
				instance.NewPrivacyPolicyAddedEvent(ctx, &a.Aggregate, tosLink, privacyLink, helpLink),
			}, nil
		}, nil
	}
}