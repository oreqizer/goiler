// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package schemas

import (
	"testing"

	"github.com/oreqizer/go-relaygen/relay"
)

var (
	accountonestr   = "1"
	accounttwostr   = "2"
	accountthreestr = "3"
	accountfourstr  = "4"
	accountfivestr  = "5"

	// Set later
	accountoneid   = ""
	accounttwoid   = ""
	accountthreeid = ""
	accountfourid  = ""
	accountfiveid  = ""
)

var (
	accountzero = 0
	accountone  = 1
	accounttwo  = 2
	// accountthree = 3
	accountfour = 4
	// accountfive  = 5
	accountsix = 6
)

var nodesAccount = []*Account{
	&Account{},
	&Account{},
	&Account{},
	&Account{},
	&Account{},
}

var edgesAccount = []*AccountEdge{
	{Node: &Account{}, Cursor: ""},
	{Node: &Account{}, Cursor: ""},
	{Node: &Account{}, Cursor: ""},
	{Node: &Account{}, Cursor: ""},
	{Node: &Account{}, Cursor: ""},
}

func init() {
	nodesAccount[0].Account.ID = accountonestr
	nodesAccount[1].Account.ID = accounttwostr
	nodesAccount[2].Account.ID = accountthreestr
	nodesAccount[3].Account.ID = accountfourstr
	nodesAccount[4].Account.ID = accountfivestr

	accountoneid = nodesAccount[0].ID()
	accounttwoid = nodesAccount[1].ID()
	accountthreeid = nodesAccount[2].ID()
	accountfourid = nodesAccount[3].ID()
	accountfiveid = nodesAccount[4].ID()

	edgesAccount[0].Node.Account.ID = accountonestr
	edgesAccount[0].Cursor = accountoneid
	edgesAccount[1].Node.Account.ID = accounttwostr
	edgesAccount[1].Cursor = accounttwoid
	edgesAccount[2].Node.Account.ID = accountthreestr
	edgesAccount[2].Cursor = accountthreeid
	edgesAccount[3].Node.Account.ID = accountfourstr
	edgesAccount[3].Cursor = accountfourid
	edgesAccount[4].Node.Account.ID = accountfivestr
	edgesAccount[4].Cursor = accountfiveid
}

var tableAccountConnectionFromArray = []struct {
	nodes []*Account
	args  *relay.ConnectionArgs
	out   *AccountConnection
}{
	{
		nodes: []*Account{},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &accountzero, Last: &accountzero},
		out: &AccountConnection{
			Edges:    []*AccountEdge{},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: nil, EndCursor: nil},
		},
	},
	{
		nodes: nodesAccount,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &accountzero, Last: &accountzero},
		out: &AccountConnection{
			Edges:    edgesAccount,
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &accountoneid, EndCursor: &accountfiveid},
		},
	},
	{
		nodes: nodesAccount,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &accounttwo, Last: &accountzero},
		out: &AccountConnection{
			Edges:    edgesAccount[:2],
			PageInfo: relay.PageInfo{HasNextPage: true, HasPreviousPage: false, StartCursor: &accountoneid, EndCursor: &accounttwoid},
		},
	},
	{
		nodes: nodesAccount,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &accountzero, Last: &accounttwo},
		out: &AccountConnection{
			Edges:    edgesAccount[3:],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: true, StartCursor: &accountfourid, EndCursor: &accountfiveid},
		},
	},
	{
		nodes: nodesAccount,
		args:  &relay.ConnectionArgs{Before: &accountfourid, After: &accounttwoid, First: &accountzero, Last: &accountzero},
		out: &AccountConnection{
			Edges:    edgesAccount[1:4],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &accounttwoid, EndCursor: &accountfourid},
		},
	},
}

func TestAccountConnectionFromArray(t *testing.T) {
	empty := AccountConnectionFromArray(tableAccountConnectionFromArray[0].nodes, nil)
	if empty != nil {
		t.Errorf("Expected nil output for nil args")
	}

	for i, e := range tableAccountConnectionFromArray {
		out := AccountConnectionFromArray(e.nodes, e.args)
		if out == nil {
			t.Errorf("Unexpected nil output")
			return
		}

		// PageInfo
		if out.PageInfo.HasNextPage != e.out.PageInfo.HasNextPage {
			t.Errorf("%d: Has next page: got %v, want %v", i, out.PageInfo.HasNextPage, e.out.PageInfo.HasNextPage)
		}

		if out.PageInfo.HasPreviousPage != e.out.PageInfo.HasPreviousPage {
			t.Errorf("%d: Has previous page: got %v, want %v", i, out.PageInfo.HasPreviousPage, e.out.PageInfo.HasPreviousPage)
		}

		if out.PageInfo.StartCursor != nil || e.out.PageInfo.StartCursor != nil {
			if *out.PageInfo.StartCursor != *e.out.PageInfo.StartCursor {
				t.Errorf("%d: Start cursor: got %v, want %v", i, *out.PageInfo.StartCursor, *e.out.PageInfo.StartCursor)
			}
		}

		if out.PageInfo.EndCursor != nil || e.out.PageInfo.EndCursor != nil {
			if *out.PageInfo.EndCursor != *e.out.PageInfo.EndCursor {
				t.Errorf("%d: End cursor: got %v, want %v", i, *out.PageInfo.EndCursor, *e.out.PageInfo.EndCursor)
			}
		}

		// Edges
		if len(out.Edges) != len(e.out.Edges) {
			t.Errorf("%d: Edges length: got %d, want %d", i, len(out.Edges), len(e.out.Edges))
			return
		}

		for j, eedge := range e.out.Edges {
			oedge := out.Edges[j]
			if eedge.Cursor != oedge.Cursor {
				t.Errorf("%d: Edge #%d: Cursor: got %s, want %s", i, j, oedge.Cursor, eedge.Cursor)
			}
		}
	}
}

var tableAccountEdgesToReturn = []struct {
	edges  []*AccountEdge
	before *string
	after  *string
	first  *int
	last   *int
	out    *AccountConnection
}{
	{edges: []*AccountEdge{}, before: nil, after: nil, first: &accountzero, last: &accountzero, out: tableAccountConnectionFromArray[0].out},
	{edges: edgesAccount, before: nil, after: nil, first: &accountzero, last: &accountzero, out: tableAccountConnectionFromArray[1].out},
	{edges: edgesAccount, before: nil, after: nil, first: &accounttwo, last: &accountzero, out: tableAccountConnectionFromArray[2].out},
	{edges: edgesAccount, before: nil, after: nil, first: &accountzero, last: &accounttwo, out: tableAccountConnectionFromArray[3].out},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, first: &accountzero, last: &accountzero, out: tableAccountConnectionFromArray[4].out},
}

func TestAccountEdgesToReturn(t *testing.T) {
	for i, e := range tableAccountEdgesToReturn {
		out := accountEdgesToReturn(e.edges, e.before, e.after, e.first, e.last)
		// PageInfo
		if out.PageInfo.HasNextPage != e.out.PageInfo.HasNextPage {
			t.Errorf("%d: Has next page: got %v, want %v", i, out.PageInfo.HasNextPage, e.out.PageInfo.HasNextPage)
		}

		if out.PageInfo.HasPreviousPage != e.out.PageInfo.HasPreviousPage {
			t.Errorf("%d: Has previous page: got %v, want %v", i, out.PageInfo.HasPreviousPage, e.out.PageInfo.HasPreviousPage)
		}

		if out.PageInfo.StartCursor != nil || e.out.PageInfo.StartCursor != nil {
			if *out.PageInfo.StartCursor != *e.out.PageInfo.StartCursor {
				t.Errorf("%d: Start cursor: got %v, want %v", i, *out.PageInfo.StartCursor, *e.out.PageInfo.StartCursor)
			}
		}

		if out.PageInfo.EndCursor != nil || e.out.PageInfo.EndCursor != nil {
			if *out.PageInfo.EndCursor != *e.out.PageInfo.EndCursor {
				t.Errorf("%d: End cursor: got %v, want %v", i, *out.PageInfo.EndCursor, *e.out.PageInfo.EndCursor)
			}
		}

		// Edges
		if len(out.Edges) != len(e.out.Edges) {
			t.Errorf("%d: Edges length: got %d, want %d", i, len(out.Edges), len(e.out.Edges))
			return
		}

		for j, eedge := range e.out.Edges {
			oedge := out.Edges[j]
			if eedge.Cursor != oedge.Cursor {
				t.Errorf("%d: Edge #%d: Cursor: got %s, want %s", i, j, oedge.Cursor, eedge.Cursor)
			}
		}
	}
}

var tableAccountApplyCursorsToEdges = []struct {
	edges  []*AccountEdge
	before *string
	after  *string
	len    int
}{
	{edges: []*AccountEdge{}, before: nil, after: nil, len: 0},
	{edges: edgesAccount, before: nil, after: nil, len: 5},
	{edges: edgesAccount, before: nil, after: &accounttwoid, len: 4},
	{edges: edgesAccount, before: &accountfourid, after: nil, len: 4},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, len: 3},
}

func TestAccountApplyCursorsToEdges(t *testing.T) {
	for i, e := range tableAccountApplyCursorsToEdges {
		out := accountApplyCursorsToEdges(e.edges, e.before, e.after)
		if len(out) != e.len {
			t.Errorf("%d: Length: got %d, want %d", i, len(out), e.len)
		}

		if len(out) > 0 {
			if cursor := out[len(out)-1].Cursor; e.before != nil && cursor != *e.before {
				t.Errorf("%d: Before: got %s, want %s", i, cursor, *e.before)
			}

			if cursor := out[0].Cursor; e.after != nil && cursor != *e.after {
				t.Errorf("%d: After: got %s, want %s", i, cursor, *e.after)
			}
		}
	}
}

var tableAccountHasPreviousPage = []struct {
	edges  []*AccountEdge
	before *string
	after  *string
	last   *int
	out    bool
}{
	{edges: []*AccountEdge{}, before: nil, after: nil, last: nil, out: false},
	{edges: edgesAccount, before: nil, after: nil, last: nil, out: false},
	{edges: edgesAccount, before: nil, after: nil, last: &accountzero, out: false},
	{edges: edgesAccount, before: nil, after: nil, last: &accountsix, out: false},
	{edges: edgesAccount, before: nil, after: nil, last: &accountfour, out: true},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, last: &accountfour, out: false},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, last: &accountone, out: true},
}

func TestAccountHasPreviousPage(t *testing.T) {
	for i, e := range tableAccountHasPreviousPage {
		out := accountHasPreviousPage(e.edges, e.before, e.after, e.last)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}

var tableAccountHasNextPage = []struct {
	edges  []*AccountEdge
	before *string
	after  *string
	first  *int
	out    bool
}{
	{edges: []*AccountEdge{}, before: nil, after: nil, first: nil, out: false},
	{edges: edgesAccount, before: nil, after: nil, first: nil, out: false},
	{edges: edgesAccount, before: nil, after: nil, first: &accountzero, out: false},
	{edges: edgesAccount, before: nil, after: nil, first: &accountsix, out: false},
	{edges: edgesAccount, before: nil, after: nil, first: &accountfour, out: true},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, first: &accountfour, out: false},
	{edges: edgesAccount, before: &accountfourid, after: &accounttwoid, first: &accountone, out: true},
}

func TestAccountHasNextPage(t *testing.T) {
	for i, e := range tableAccountHasNextPage {
		out := accountHasNextPage(e.edges, e.before, e.after, e.first)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}