package driven_adapter_db

import (
	"net/http"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
)

var mysqlErrorNumberToHTTPStatusCode = map[uint16]int{
	// Erişim hataları
	1045: http.StatusUnauthorized, // ER_ACCESS_DENIED_ERROR
	1130: http.StatusForbidden,    // ER_HOST_NOT_PRIVILEGED

	// Tablo ve veritabanı hataları
	1049: http.StatusNotFound, // ER_BAD_DB_ERROR
	1051: http.StatusNotFound, // ER_BAD_TABLE_ERROR
	1146: http.StatusNotFound, // ER_NO_SUCH_TABLE

	// İşlem ve entegrasyon hataları
	1062: http.StatusConflict,       // ER_DUP_ENTRY
	1205: http.StatusRequestTimeout, // ER_LOCK_WAIT_TIMEOUT
	1213: http.StatusConflict,       // ER_LOCK_DEADLOCK
	1451: http.StatusConflict,       // ER_ROW_IS_REFERENCED_2 (bağlantılı veri)
	1452: http.StatusConflict,       // ER_NO_REFERENCED_ROW_2 (bağlantısız veri)

	// Sorgu ve veri hataları
	1054: http.StatusBadRequest, // ER_BAD_FIELD_ERROR (geçersiz sütun adı)
	1064: http.StatusBadRequest, // ER_PARSE_ERROR (Sorgu hatası)
	1364: http.StatusBadRequest, // ER_NO_DEFAULT_FOR_FIELD (varsayılan değer yok)
	1406: http.StatusBadRequest, // ER_DATA_TOO_LONG

	// Sunucu ve bağlantı hataları
	2002: http.StatusServiceUnavailable, // CR_CONNECTION_ERROR
	2003: http.StatusServiceUnavailable, // CR_CONN_HOST_ERROR
	2005: http.StatusServiceUnavailable, // CR_UNKNOWN_HOST
	2006: http.StatusServiceUnavailable, // CR_SERVER_GONE_ERROR
	2013: http.StatusServiceUnavailable, // CR_SERVER_LOST
}

var mysqlErrorNumberToFaultType = map[uint16]domainFaultPort.FaultType{
	1045: domainFaultPort.FaultTypeInaccessibleServer,
	1049: domainFaultPort.FaultTypeCollectionNotFound,
	1051: domainFaultPort.FaultTypeCollectionNotFound,
	1062: domainFaultPort.FaultTypeDuplicatedEntry,
	1130: domainFaultPort.FaultTypeDatabaseNotPrivileged,
	1146: domainFaultPort.FaultTypeCollectionNotFound,
	1452: domainFaultPort.FaultTypeAssociationViolation,
}
