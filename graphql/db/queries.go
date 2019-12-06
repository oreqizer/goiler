package db

import . "github.com/volatiletech/sqlboiler/queries/qm"

var QueryNotDeleted = Where("deleted_at IS NULL")
