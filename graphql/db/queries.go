package db

import "github.com/volatiletech/sqlboiler/queries/qm"

// QueryNotDeleted is a query that filters results marked as deleted
var QueryNotDeleted = qm.Where("deleted_at IS NULL")
