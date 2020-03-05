package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
)

func getPersonId(ctx context.Context, c *mastodon.Client, req sdk.GetPeopleRequest) (rsp sdk.GetPeopleResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	err := func() (err error) {
		var account *mastodon.Account
		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			account, err = c.GetAccount(ctx, mastodon.ID(*req.Mode.Id.ServiceId.Value))
			if err != nil {
				return
			}

			break
		case sdk.IdKind.Email:
			break
		case sdk.IdKind.Name:
			break
		case sdk.IdKind.Username:
			break
		case sdk.IdKind.Ean:
			break
		case sdk.IdKind.Url:
			break
		case sdk.IdKind.Me:
			account, err = c.GetAccountCurrentUser(ctx)
			if err != nil {
				return
			}

			break
		}

		rsp.People = []sdk.Person{MapPersonFromMastodonAccount(*account)}

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: ptr.String(err.Error()),
			},
		})
	}

	return
}

func getPeopleSearch(ctx context.Context, c *mastodon.Client, req sdk.GetPeopleRequest) (rsp sdk.GetPeopleResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	err := func() (err error) {
		accounts, err := c.AccountsSearch(ctx, *req.Mode.Search.Term, 100)
		if err != nil {
			return
		}

		rsp.People = MapPeopleFromMastodonAccounts(accounts)

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: ptr.String(err.Error()),
			},
		})
	}

	return
}

func getPeopleRelation(ctx context.Context, c *mastodon.Client, req sdk.GetPeopleRequest) (rsp sdk.GetPeopleResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	var accounts []*mastodon.Account

	//var pagination *sdk.Pagination
	pg := &mastodon.Pagination{}

	//if len(req.Pages) > 0 {
	//	pagination = req.Pages[0].Page
	//}
	//
	//if pagination != nil &&
	//	pagination.Next != nil &&
	//	pagination.Next.CursorPage != nil {
	//	pg.MaxID = mastodon.ID(*pagination.Next.CursorPage.Value)
	//}
	//
	//if pagination != nil &&
	//	pagination.Previous != nil &&
	//	pagination.Previous.CursorPage != nil {
	//	pg.SinceID = mastodon.ID(*pagination.Previous.CursorPage.Value)
	//}

	err := func() (err error) {
		switch *req.Mode.Relation.Relation {
		case sdk.PersonRelationName.PersonBlocksPeople:
			accounts, err = c.GetBlocks(ctx, pg)
			if err != nil {
				return
			}
		case sdk.PersonRelationName.PersonFollowedByPeople:
			accounts, err = c.GetAccountFollowers(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case sdk.PersonRelationName.PersonFollowsPeople:
			accounts, err = c.GetAccountFollowing(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case sdk.PersonRelationName.PersonMutesPeople:
			accounts, err = c.GetMutes(ctx, pg)
			if err != nil {
				return
			}
		case sdk.PersonRelationName.PersonRequestedToBeFollowedByPeople:
			accounts, err = c.GetFollowRequests(ctx, pg)
			if err != nil {
				return
			}
		case sdk.StatusRelationName.StatusFavoredByPeople:
			accounts, err = c.GetFavouritedBy(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		}

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: ptr.String(err.Error()),
			},
		})
	}

	//if pg != nil {
	//	pagination := &sdk.Pagination{
	//		Previous: &sdk.Page{
	//			Kind: &sdk.PageKind.CursorPage,
	//			CursorPage: &sdk.CursorPage{
	//				Value: ptr.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &sdk.Page{
	//			Kind: &sdk.PageKind.CursorPage,
	//			CursorPage: &sdk.CursorPage{
	//				Value: ptr.String(string(pg.MaxID)),
	//			},
	//		},
	//	}
	//
	//	if pagination != nil {
	//		pagination.Current = pagination
	//	}
	//
	//	rsp.Meta.Pagination = pagination
	//}

	rsp.People = MapPeopleFromMastodonAccounts(accounts)

	return
}

// todo v1: support IdUnion
func putPeopleRelation(ctx context.Context, c *mastodon.Client, req sdk.PutPeopleRequest) (rsp sdk.PutPeopleResponse) {
	rsp.Meta = &sdk.ResponseMeta{}

	add := func() (errs []error) {
		switch *req.Mode.Relation.Relation {
		case sdk.PersonRelationName.PersonBlocksPeople:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.AccountBlock(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		case sdk.PersonRelationName.PersonFollowsPeople:
			//for _, id := range req.Mode.Relation.Ids {
			//	switch *id.Kind {
			//	case sdk.IdKind.URL:
			//		_, err := c.FollowRemoteUser(ctx, *id.Url.Value)
			//		if err != nil {
			//			errs = append(errs, err)
			//		}
			//	case sdk.IdKind.SERVICE_ID:
			//		_, err := c.AccountFollow(ctx, mastodon.ID(*id.Value))
			//		if err != nil {
			//			errs = append(errs, err)
			//		}
			//	}
			//}
		case sdk.PersonRelationName.PersonMutesPeople:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.AccountMute(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		default:
		}

		return
	}

	remove := func() (errs []error) {
		switch *req.Mode.Relation.Relation {
		case sdk.PersonRelationName.PersonBlocksPeople:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.AccountUnblock(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		case sdk.PersonRelationName.PersonFollowsPeople:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.AccountUnfollow(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		case sdk.PersonRelationName.PersonMutesPeople:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.AccountUnmute(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		default:
		}

		return
	}

	var errs []error
	if *req.Mode.Relation.Operation == sdk.RelationOperation.Add {
		errs = add()
	} else {
		errs = remove()
	}

	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: ptr.String(err.Error()),
			},
		})
	}

	return
}