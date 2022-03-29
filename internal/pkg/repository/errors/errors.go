package errors

import errorutil "github.com/keithzetterstrom/faf-user-service/utils/error"

var ErrNotFound = errorutil.NewTypedError("db_record_not_found", "db record not found")
