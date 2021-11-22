package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/adam-lavrik/go-imath/ix"
	"github.com/fremantle-industries/workshop/pkg/serving/nftactivity/graph/generated"
	"github.com/fremantle-industries/workshop/pkg/serving/nftactivity/graph/model"
	"github.com/fremantle-industries/workshop/pkg/serving/nftactivity/store"
)

// GetTokenSetBySlug is the resolver for the getTokenSetBySlug field.
func (r *queryResolver) GetTokenSetBySlug(ctx context.Context, slug string) (*model.TokenSet, error) {
	q := `
    SELECT
      c.slug,
      c.collection_id,
      c.contract_id,
      c.name,
      c.description,
      c.image,
      c.banner,
      c.discordurl,
      c.externalurl,
      c.twitterusername,
      c.isverified,
      c.total_tokens
	  FROM public.nonfungible_collection c
	  WHERE c.slug = $1
  `
	row := store.QueryRow(q, slug)
	ts := new(model.TokenSet)
	err := row.Scan(
		&ts.Slug,
		&ts.CollectionID,
		&ts.ContractID,
		&ts.Name,
		&ts.Description,
		&ts.Image,
		&ts.Banner,
		&ts.DiscordURL,
		&ts.ExternalURL,
		&ts.TwitterUsername,
		&ts.IsVerified,
		&ts.TotalTokens,
	)
	if err != nil {
		log.Print(err)
	}

	return ts, nil
}

// GetTokenSetStatsBySlug is the resolver for the getTokenSetStatsBySlug field.
func (r *queryResolver) GetTokenSetStatsBySlug(ctx context.Context, slug string) (*model.TokenSetStat, error) {
	var q strings.Builder
	q.WriteString(`
    SELECT
      c.slug,
      c.name,
      l.floor,
      NULL AS floor_24h_delta,
      0.0 AS volume_24h,
      NULL AS volume_24h_delta,
      NULL AS listing_percentage_24h_delta
    FROM public.nonfungible_collection c
      LEFT JOIN (
        SELECT
          c.slug,
          MIN(e.individual_value) AS floor
        FROM public.nonfungible_collection c
          JOIN public.nonfungible_market_event e ON c.collection_id = e.collection_id
        WHERE c.slug = $1
          AND e.event_type = 'new-order'
          AND e.order_type = 2
          AND e.valid_until > extract (epoch from now())
        GROUP BY
          c.slug
      ) l ON l.slug = c.slug
    WHERE
      c.slug = $1
  `)

	// Hydrate model structs with row results
	row := store.QueryRow(q.String(), slug)
	stat := new(model.TokenSetStat)
	err := row.Scan(
		&stat.Slug,
		&stat.Name,
		&stat.Floor,
		&stat.Floor24hDelta,
		&stat.Volume24h,
		&stat.Volume24hDelta,
		&stat.ListingPercentage24hDelta,
	)
	if err != nil {
		log.Print(err)
	}

	return stat, nil
}

// GetTokenSetStats is the resolver for the getTokenSetStats field.
func (r *queryResolver) GetTokenSetStats(ctx context.Context) ([]*model.TokenSetStat, error) {
	var q strings.Builder
	q.WriteString(`
    SELECT
      c.slug,
      c.name,
      l.floor,
      NULL AS floor_24h_delta,
      0.0 AS volume_24h,
      NULL AS volume_24h_delta,
      NULL AS listing_percentage_24h_delta
    FROM public.nonfungible_collection c
      LEFT JOIN (
        SELECT
          c.slug,
          MIN(e.individual_value) AS floor
        FROM public.nonfungible_collection c
          JOIN public.nonfungible_market_event e ON c.collection_id = e.collection_id
        WHERE e.event_type = 'new-order'
          AND e.order_type = 2
          AND e.valid_until > extract (epoch from now())
        GROUP BY
          c.slug
      ) l ON l.slug = c.slug
  `)

	rows, err := store.Query(q.String())
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	// Hydrate model structs with row results
	stats := make([]*model.TokenSetStat, 0)
	for rows.Next() {
		stat := new(model.TokenSetStat)
		err := rows.Scan(
			&stat.Slug,
			&stat.Name,
			&stat.Floor,
			&stat.Floor24hDelta,
			&stat.Volume24h,
			&stat.Volume24hDelta,
			&stat.ListingPercentage24hDelta,
		)
		if err != nil {
			log.Print(err)
		}
		stats = append(stats, stat)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return stats, nil
}

// GetTokenSetTradeActivityBySlug is the resolver for the getTokenSetTradeActivityBySlug field.
func (r *queryResolver) GetTokenSetTradeActivityBySlug(ctx context.Context, slug string) (*model.TokenSetTradeActivity, error) {
	ta := new(model.TokenSetTradeActivity)

	return ta, nil
}

// GetTokenSetGridItemAssetsBySlug is the resolver for the getTokenSetGridItemAssetsBySlug field.
func (r *queryResolver) GetTokenSetGridItemAssetsBySlug(ctx context.Context, slug string, offset int, limit int, tokenID *string, traits []*model.GridItemTrait, priceRange *model.PriceRange, rarityRange *model.RarityRange, buyNow bool, sort model.GridItemSort) ([]*model.TokenSetGridItemAsset, error) {
	// Build query to filter and sort items
	var q strings.Builder
	q.WriteString(`
	    SELECT
	      i.collection_id,
	      i.item_id,
	      i.item->>'token' AS "token",
	      i.collection->>'contract' AS "contract",
	      i.collection->>'slug' AS "slug",
	      i.item->>'image' AS "image",
	      i.item->'isFlagged' AS "is_flagged",
	      i.item->'rarity'->'rank' AS "rarity_rank",
        i.valuation,
	      i.marketability
	    FROM "payload"."nonfungible_items" i
	  WHERE
      i.collection->>'slug' = $1
  `)
	if traits != nil {
		fmt.Println("TODO> filter by traits...")
		// 	q.WriteString("AND (\n")
		// 	for i, t := range traits {
		// 		if i > 0 {
		// 			q.WriteString(" OR ")
		// 		}
		// 		q.WriteString(fmt.Sprintf("p.traits = %s, %s", t.Type, t.Value))
		// 	}
		// 	q.WriteString("\n)\n")
	}
	if priceRange != nil {
		fmt.Println("TODO> filter by price range...")
		// q.WriteString(`
		//       AND i.marketability->'floor'->'price'->'amount' >= $2::JSONB
		//       AND i.marketability->'floor'->'price'->'amount' <= $3::JSONB
		//     `)
	}
	if rarityRange != nil {
		q.WriteString(`
      AND i.item->'rarity'->'rank' >= $2::JSONB
      AND i.item->'rarity'->'rank' <= $3::JSONB
    `)
	}
	if buyNow == true {
		// q.WriteString("AND i.marketability->'floor'->'price'->'amount' > '0'::JSONB\n")
	}
	switch sort {
	case model.GridItemSortPriceHighToLow:
		q.WriteString("ORDER BY i.marketability->'floor'->'price'->'amount' DESC NULLS LAST\n")
	case model.GridItemSortPriceLowToHigh:
		q.WriteString("ORDER BY i.marketability->'floor'->'price'->'amount' ASC NULLS LAST\n")
	case model.GridItemSortHighestLastSale:
		q.WriteString("ORDER BY i.marketability->'last_sale'->'amount' DESC NULLS LAST\n")
	default:
		q.WriteString("ORDER BY i.valuation->'discounts'->'50' DESC NULLS LAST\n")
	}
	q.WriteString("OFFSET $4 LIMIT $5")

	// Execute query and return matching rows
	l := ix.Min(store.MaxLimit, limit)
	// rows, err := store.Query(q.String(), slug, priceRange.Min, priceRange.Max, rarityRange.Min, rarityRange.Max, offset, l)
	rows, err := store.Query(q.String(), slug, rarityRange.Min, rarityRange.Max, offset, l)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	// Hydrate model structs with row results
	assets := make([]*model.TokenSetGridItemAsset, 0)
	for rows.Next() {
		a := new(model.TokenSetGridItemAsset)
		err := rows.Scan(
			&a.CollectionID,
			&a.ItemID,
			&a.TokenID,
			&a.Contract,
			&a.Slug,
			&a.Image,
			&a.IsFlagged,
			&a.RarityRank,
			&a.Valuation,
			&a.Marketability,
		)
		if err != nil {
			log.Print(err)
		}
		assets = append(assets, a)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return assets, nil
}

// GetTokenSetItemsBySlug is the resolver for the getTokenSetItemsBySlug field.
func (r *queryResolver) GetTokenSetItemsBySlug(ctx context.Context, slug string, offset int, limit int, tokenID *string, traits []*model.TokenSetItemsTrait, priceRange *model.TokenSetItemsPriceRange, rarityRange *model.TokenSetItemsRarityRange, buyNow bool, sort model.TokenSetItemsSort) ([]*model.TokenSetItem, error) {
	// Build query to filter and sort items
	var q strings.Builder
	q.WriteString(`
	    SELECT
	      i.collection_id,
	      i.item_id,
	      i.item->>'token' AS "token",
	      i.collection->>'contract' AS "contract",
	      i.collection->>'slug' AS "slug",
	      i.item->>'image' AS "image",
	      i.item->'isFlagged' AS "is_flagged",
	      i.item->'rarity'->'rank' AS "rarity_rank",
        i.valuation,
	      i.marketability
	    FROM "payload"."nonfungible_items" i
	  WHERE
      i.collection->>'slug' = $1
  `)
	if traits != nil {
		fmt.Println("TODO> filter by traits...")
		// 	q.WriteString("AND (\n")
		// 	for i, t := range traits {
		// 		if i > 0 {
		// 			q.WriteString(" OR ")
		// 		}
		// 		q.WriteString(fmt.Sprintf("p.traits = %s, %s", t.Type, t.Value))
		// 	}
		// 	q.WriteString("\n)\n")
	}
	if priceRange != nil {
		fmt.Println("TODO> filter by price range...")
		// q.WriteString(`
		//       AND i.marketability->'floor'->'price'->'amount' >= $2::JSONB
		//       AND i.marketability->'floor'->'price'->'amount' <= $3::JSONB
		//     `)
	}
	if rarityRange != nil {
		q.WriteString(`
      AND i.item->'rarity'->'rank' >= $2::JSONB
      AND i.item->'rarity'->'rank' <= $3::JSONB
    `)
	}
	if buyNow == true {
		// q.WriteString("AND i.marketability->'floor'->'price'->'amount' > '0'::JSONB\n")
	}
	switch sort {
	case model.TokenSetItemsSortPriceHighToLow:
		q.WriteString("ORDER BY i.marketability->'floor'->'price'->'amount' DESC NULLS LAST\n")
	case model.TokenSetItemsSortPriceLowToHigh:
		q.WriteString("ORDER BY i.marketability->'floor'->'price'->'amount' ASC NULLS LAST\n")
	case model.TokenSetItemsSortHighestLastSale:
		q.WriteString("ORDER BY i.marketability->'last_sale'->'amount' DESC NULLS LAST\n")
	default:
		q.WriteString("ORDER BY i.valuation->'discounts'->'50' DESC NULLS LAST\n")
	}
	q.WriteString("OFFSET $4 LIMIT $5")

	// Execute query and return matching rows
	l := ix.Min(store.MaxLimit, limit)
	// rows, err := store.Query(q.String(), slug, priceRange.Min, priceRange.Max, rarityRange.Min, rarityRange.Max, offset, l)
	rows, err := store.Query(q.String(), slug, rarityRange.Min, rarityRange.Max, offset, l)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	// Hydrate model structs with row results
	assets := make([]*model.TokenSetItem, 0)
	for rows.Next() {
		a := new(model.TokenSetItem)
		err := rows.Scan(
			&a.CollectionID,
			&a.ItemID,
			&a.TokenID,
			&a.Contract,
			&a.Slug,
			&a.Image,
			&a.IsFlagged,
			&a.RarityRank,
			&a.Valuation,
			&a.Marketability,
		)
		if err != nil {
			log.Print(err)
		}
		assets = append(assets, a)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return assets, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
