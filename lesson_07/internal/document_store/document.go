package document_store

import (
	"encoding/json"
	"errors"
)

var (
	ErrDocumentNotFound         = errors.New("document not found")
	ErrUnsupportedDocumentField = errors.New("unsupported focument field")
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

func MarshalDocument(input any) (*Document, error) {
	jsonStr, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	var tempMap map[string]any

	if err := json.Unmarshal(jsonStr, &tempMap); err != nil {
		return nil, err
	}

	fields := make(map[string]DocumentField, len(tempMap))

	for key, value := range tempMap {
		fieldType, err := defineType(value)

		if err != nil {
			return nil, err
		}

		fields[key] = DocumentField{Type: fieldType, Value: value}
	}

	return &Document{Fields: fields}, nil
}

func defineType(input any) (DocumentFieldType, error) {
	switch input.(type) {
	case string:
		return DocumentFieldTypeString, nil
	case bool:
		return DocumentFieldTypeBool, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return DocumentFieldTypeNumber, nil
	case []interface{}:
		return DocumentFieldTypeArray, nil
	case interface{}:
		return DocumentFieldTypeObject, nil
	default:
		return "", ErrUnsupportedDocumentField
	}
}

func UnmarshalDocument(doc *Document, output any) error {
	tempMap := make(map[string]any)
	for key, value := range doc.Fields {
		tempMap[key] = value.Value
	}

	docJson, err := json.Marshal(tempMap)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(docJson, output); err != nil {
		return err
	}

	return nil
}
