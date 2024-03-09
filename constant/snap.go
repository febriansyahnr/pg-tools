package constant

const (
	SnapDateFormatLayout = "2006-01-02T15:04:05-07:00"
	// va type
	VATypeCloseDynamic = "closed_dynamic"
	VATypeCloseStatic  = "closed_static"
	VATypeOpenStatic   = "open_static"
	// response code

	// 200XX00
	SNAP_SUCCESS = "success" // 200XX00
	// 202XX00
	SNAP_INPROGRESS = "inprogress" // 202XX00
	// 400XX00
	SNAP_BAD_REQUEST       = "bad_request"       // 400XX00
	SNAP_INVALID_FIELD     = "invalid_field"     // 400XX01
	SNAP_INVALID_MANDATORY = "invalid_mandatory" // 400XX02
	// 401XX00
	SNAP_UNAUTHORIZED             = "unauthorized"             // 401XX00
	SNAP_INVALID_SIGNATURE        = "invalid_signature"        // 401XX01
	SNAP_ACCESS_TOKEN_INVALID     = "access_token_invalid"     // 401XX02
	SNAP_INVALID_TOKEN_B2B        = "invalid_token_b2b"        // 401XX01
	SNAP_INVALID_CUSTOMER_TOKEN   = "invalid_customer_token"   // 401XX02
	SNAP_TOKEN_NOT_FOUND          = "token_not_found"          // 401XX03
	SNAP_CUSTOMER_TOKEN_NOT_FOUND = "customer_token_not_found" // 401XX04
	//403XX00
	SNAP_TRANSACTION_EXPIRED              = "transaction_expired"              // 403XX00
	SNAP_FEATURE_NOT_ALLOWED              = "feature_not_allowed"              // 403XX01
	SNAP_EXCEEDS_TRANSACTION_AMOUNT_LIMIT = "exceeds_transaction_amount_limit" // 403XX02
	SNAP_SUSPECTED_FRAUD                  = "suspected_fraud"                  // 403XX03
	SNAP_ACTIVITY_LIMIT_EXCEEDED          = "activity_limit_exceeded"          // 403XX04
	SNAP_DO_NOT_HONOR                     = "do_not_honor"                     // 403XX05
	SNAP_FEATURE_NOT_ALLOWED_THIS_TIME    = "feature_not_allowed_this_time"    // 403XX06
	// 404XX00
	SNAP_INVALID_TRANSACTION_STATUS = "invalid_transaction_status" // 404XX00
	SNAP_TRANSACTION_NOT_FOUND      = "transaction_not_found"      // 404XX01
	SNAP_INVALID_ROUTING            = "invalid_routing"            // 404XX02
	SNAP_BANK_NOT_SUPPORTED         = "bank_not_supported"         // 404XX03
	SNAP_TRANSACTION_CANCELLED      = "transaction_cancelled"      // 404XX04
	SNAP_INVALID_VA                 = "invalid_va"                 // 404XX12
	SNAP_INVALID_AMOUNT             = "invalid_amount"             // 404XX13
	SNAP_INVALID_ALREADY_PAID       = "invalid_already_paid"       // 404XX14
	SNAP_INVALID_BILL_EXPIRED       = "invalid_bill_expired"       // 404XX19
	// 409XX00
	SNAP_CONFLICT = "conflict" // 409XX00
	// 429XX00
	SNAP_TO_MANY_REQUEST = "to_many_request" // 00
	//500XX00
	SNAP_GENERAL_ERROR         = "general_error"         // 500XX00
	SNAP_INTERNAL_SERVER_ERROR = "internal_server_error" // 500XX01
	SNAP_EXTERNAL_SERVER_ERROR = "external_server_error" // 500XX02
	// 504XX00
	SNAP_TIMEOUT = "timeout" // 00

	// snap service

	SNAP_SERVICE_INQUIRY = "24"
	SNAP_SERVICE_PAYMENT = "25"
	SNAP_SERVICE_B2B     = "73"
)
