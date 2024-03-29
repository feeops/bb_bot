// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailPassword holds the string denoting the emailpassword field in the database.
	FieldEmailPassword = "email_password"
	// FieldBindEmail holds the string denoting the bindemail field in the database.
	FieldBindEmail = "bind_email"
	// FieldPostCode holds the string denoting the postcode field in the database.
	FieldPostCode = "post_code"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldWindowName holds the string denoting the windowname field in the database.
	FieldWindowName = "window_name"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldIPUsed holds the string denoting the ipused field in the database.
	FieldIPUsed = "ip_used"
	// FieldRefURL holds the string denoting the refurl field in the database.
	FieldRefURL = "ref_url"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldEmailPassword,
	FieldBindEmail,
	FieldPostCode,
	FieldPassword,
	FieldFirstName,
	FieldLastName,
	FieldRegion,
	FieldIP,
	FieldWindowName,
	FieldBalance,
	FieldIPUsed,
	FieldRefURL,
	FieldRemark,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultEmailPassword holds the default value on creation for the "emailPassword" field.
	DefaultEmailPassword string
	// DefaultBindEmail holds the default value on creation for the "bindEmail" field.
	DefaultBindEmail string
	// DefaultPostCode holds the default value on creation for the "postCode" field.
	DefaultPostCode string
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultFirstName holds the default value on creation for the "firstName" field.
	DefaultFirstName string
	// DefaultLastName holds the default value on creation for the "lastName" field.
	DefaultLastName string
	// DefaultRegion holds the default value on creation for the "region" field.
	DefaultRegion string
	// DefaultIP holds the default value on creation for the "IP" field.
	DefaultIP string
	// DefaultWindowName holds the default value on creation for the "windowName" field.
	DefaultWindowName string
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance float64
	// DefaultIPUsed holds the default value on creation for the "IPUsed" field.
	DefaultIPUsed bool
	// DefaultRefURL holds the default value on creation for the "refURL" field.
	DefaultRefURL string
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
)

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByEmailPassword orders the results by the emailPassword field.
func ByEmailPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailPassword, opts...).ToFunc()
}

// ByBindEmail orders the results by the bindEmail field.
func ByBindEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBindEmail, opts...).ToFunc()
}

// ByPostCode orders the results by the postCode field.
func ByPostCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostCode, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByFirstName orders the results by the firstName field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the lastName field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByIP orders the results by the IP field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByWindowName orders the results by the windowName field.
func ByWindowName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWindowName, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByIPUsed orders the results by the IPUsed field.
func ByIPUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPUsed, opts...).ToFunc()
}

// ByRefURL orders the results by the refURL field.
func ByRefURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefURL, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}
