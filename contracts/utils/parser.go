package utils

import (
	"strings"

	"fmt"

	"github.com/valyala/fastjson"
)

const (
	NodeTypeContract  = "ContractDefinition"
	NodeTypeVariable  = "VariableDeclaration"
	NodeTypeStruct    = "StructDefinition"
	NodeTypeEnum      = "EnumDefinition"
	NodeTypeEnumValue = "EnumValue"

	TypeNodeUserDefined = "UserDefinedTypeName"
	TypeNodeElementary  = "ElementaryTypeName"
	TypeNodeMapping     = "Mapping"
	TypeBool            = "bool"
	TypeUint8           = "uint8"
	TypeUint64          = "uint64"
	TypeUint256         = "uint256"
	TypeAddresss        = "address"
	TypeString          = "string"
	TypeBytes           = "bytes memory"
	TypeBytes8          = "bytes8"
	TypeBytes32         = "bytes32"
	TypeStructPrefix    = "struct "
	TypeEnumPrefix      = "enum "
)

func parseType(typeName *fastjson.Value) string {
	nodeType := typeName.GetStringBytes("nodeType")
	switch string(nodeType) {
	case TypeNodeElementary:
		return parseElementaryType(typeName)
	case TypeNodeUserDefined:
		return parseUserDefinedType(typeName)
	case TypeNodeMapping:
		return parseMappingType(typeName)
	}
	return ""
}

func parseElementaryType(typeName *fastjson.Value) string {
	typeBytes := typeName.GetStringBytes("typeDescriptions", "typeString")
	switch string(typeBytes) {
	case TypeBool:
		return "bool"
	case TypeUint8:
		return "uint8"
	case TypeUint64:
		return "uint64"
	case TypeUint256:
		return "*big.Int"
	case TypeAddresss:
		return "ethCommon.Address"
	case TypeString:
		return "string"
	case TypeBytes:
		return "[]byte"
	case TypeBytes8:
		return "ablCommon.ID"
	case TypeBytes32:
		return "ethCommon.Hash"
	}
	return ""
}

func parseUserDefinedType(typeName *fastjson.Value) string {
	typeBytes := string(typeName.GetStringBytes("typeDescriptions", "typeString"))
	canonical := ""
	switch {
	case strings.HasPrefix(string(typeBytes), TypeStructPrefix):
		canonical = strings.TrimPrefix(string(typeBytes), TypeStructPrefix)
	case strings.HasPrefix(string(typeBytes), TypeEnumPrefix):
		canonical = strings.TrimPrefix(string(typeBytes), TypeEnumPrefix)
	}
	return strings.Map(func(r rune) rune {
		if r == '.' {
			return -1
		}
		return r
	}, canonical)
}

func parseMappingType(typeName *fastjson.Value) string {
	keyType := typeName.Get("keyType")
	keyT := parseType(keyType)
	valType := typeName.Get("valueType")
	valT := parseType(valType)
	return fmt.Sprintf("map[%s]%s", keyT, valT)
}
