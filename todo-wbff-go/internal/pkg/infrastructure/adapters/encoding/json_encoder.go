package infrastructure_adapters_encoding

import (
	"context"
	"encoding/json"
	drivenAppPortsEncoding "todo-app-wbff/internal/pkg/app/ports/driven/encoding"
)

type jsonEncoder struct{}

func NewJSONEncoder() *jsonEncoder {
	return &jsonEncoder{}
}

func (e *jsonEncoder) NewContextWith(parentContext context.Context) context.Context {
	return context.WithValue(parentContext, drivenAppPortsEncoding.JSONEncoderKey{}, e)
}

func (e *jsonEncoder) EncodeModel(sourceModel map[string]interface{}) ([]byte, error) {
	return json.Marshal(sourceModel)
}

func (e *jsonEncoder) DecodeModel(sourceSerial []byte) (map[string]interface{}, error) {
	var targetModel map[string]interface{}
	err := json.Unmarshal(sourceSerial, &targetModel)
	if err != nil {
		return nil, err
	}

	return targetModel, nil
}

func (e *jsonEncoder) EncodeCollection(sourceCollection []map[string]interface{}) ([]byte, error) {
	return json.Marshal(sourceCollection)
}

func (e *jsonEncoder) DecodeCollection(sourceSerial []byte) ([]map[string]interface{}, error) {
	var targetCollection []map[string]interface{}
	err := json.Unmarshal(sourceSerial, &targetCollection)
	if err != nil {
		return nil, err
	}

	return targetCollection, nil
}
