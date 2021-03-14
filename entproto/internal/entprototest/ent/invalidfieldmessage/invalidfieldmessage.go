// Code generated by entc, DO NOT EDIT.

package invalidfieldmessage

const (
	// Label holds the string label denoting the invalidfieldmessage type in the database.
	Label = "invalid_field_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHello holds the string denoting the hello field in the database.
	FieldHello = "hello"
	// Table holds the table name of the invalidfieldmessage in the database.
	Table = "invalid_field_messages"
)

// Columns holds all SQL columns for invalidfieldmessage fields.
var Columns = []string{
	FieldID,
	FieldHello,
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