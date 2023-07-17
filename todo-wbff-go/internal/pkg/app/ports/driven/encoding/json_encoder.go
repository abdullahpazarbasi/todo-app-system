package driven_app_ports_encoding

import "context"

type JSONEncoderKey struct{}

type JSONEncoder interface {
	NewContextWith(parentContext context.Context) context.Context
	EncodeModel(sourceModel map[string]interface{}) ([]byte, error)
	DecodeModel(sourceSerial []byte) (map[string]interface{}, error)
	EncodeCollection(sourceCollection []map[string]interface{}) ([]byte, error)
	DecodeCollection(sourceSerial []byte) ([]map[string]interface{}, error)
}
