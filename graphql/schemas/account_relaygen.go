// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package schemas

import "github.com/oreqizer/go-relaygen/relay"

/*
AccountEdge is an interface holding a node and a cursor

https://facebook.github.io/relay/graphql/connections.htm#sec-Edge-Types
*/
type AccountEdge struct {
	Node   *Account `json:"node"`
	Cursor string   `json:"cursor"`
}

/*
AccountConnection holds information about a connection

https://facebook.github.io/relay/graphql/connections.htm#sec-Reserved-Types
*/
type AccountConnection struct {
	Edges    []*AccountEdge `json:"edges"`
	PageInfo relay.PageInfo `json:"pageInfo"`
}

/*
AccountConnectionFromArray creates a connection from an array of nodes
*/
func AccountConnectionFromArray(nodes []*Account, args *relay.ConnectionArgs) *AccountConnection {
	if args == nil {
		return nil
	}

	edges := make([]*AccountEdge, len(nodes))
	for i, n := range nodes {
		edges[i] = &AccountEdge{
			Node:   n,
			Cursor: n.ID(),
		}
	}

	return accountEdgesToReturn(edges, args.Before, args.After, args.First, args.Last)
}

/*
accountEdgesToReturn slices edges according to arguments, returning a connection

Consider returning an error like in
https://facebook.github.io/relay/graphql/connections.htm#sec-Pagination-algorithm
*/
func accountEdgesToReturn(all []*AccountEdge, before, after *string, first, last *int) *AccountConnection {
	edges := accountApplyCursorsToEdges(all, before, after)

	if first != nil && *first > 0 && *first < len(edges) {
		edges = edges[:*first]
	}

	if last != nil && *last > 0 && *last < len(edges) {
		edges = edges[len(edges)-*last:]
	}

	var startCursor, endCursor *string
	if len(edges) > 0 {
		if fst := edges[0]; fst != nil {
			str := fst.Cursor
			startCursor = &str
		}

		if lst := edges[len(edges)-1]; lst != nil {
			str := lst.Cursor
			endCursor = &str
		}
	}

	return &AccountConnection{
		Edges: edges,
		PageInfo: relay.PageInfo{
			HasPreviousPage: accountHasPreviousPage(all, before, after, last),
			HasNextPage:     accountHasNextPage(all, before, after, first),
			StartCursor:     startCursor,
			EndCursor:       endCursor,
		},
	}
}

/*
accountApplyCursorsToEdges slices edges according to cursors
*/
func accountApplyCursorsToEdges(all []*AccountEdge, before, after *string) []*AccountEdge {
	edges := all
	if after != nil {
		for i, e := range edges {
			if e.Cursor == *after {
				edges = edges[i:]
				break
			}
		}
	}

	if before != nil {
		for i, e := range edges {
			if e.Cursor == *before {
				edges = edges[:i+1]
				break
			}
		}
	}
	return edges
}

/*
accountHasPreviousPage determines whether there's a previous page according to cursors

https://facebook.github.io/relay/graphql/connections.htm#sec-undefined.PageInfo.Fields
*/
func accountHasPreviousPage(all []*AccountEdge, before, after *string, last *int) bool {
	if last != nil && *last > 0 {
		edges := accountApplyCursorsToEdges(all, before, after)
		return len(edges) > *last
	}

	return false
}

/*
accountHasNextPage determines whether there's another page according to cursors

https://facebook.github.io/relay/graphql/connections.htm#sec-undefined.PageInfo.Fields
*/
func accountHasNextPage(all []*AccountEdge, before, after *string, first *int) bool {
	if first != nil && *first > 0 {
		edges := accountApplyCursorsToEdges(all, before, after)
		return len(edges) > *first
	}

	return false
}