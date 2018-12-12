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
	TypeNodeArray       = "ArrayTypeName"
	TypeNodeMapping     = "Mapping"
	TypeBool            = "bool"
	TypeUint8           = "uint8"
	TypeUint64          = "uint64"
	TypeUint256         = "uint256"
	TypeAddresss        = "address"
	TypeString          = "string"
	TypeBytes           = "bytes"
	TypeBytesMem        = "bytes memory"
	TypeBytes4          = "bytes4"
	TypeBytes4Arr       = "bytes4[]"
	TypeBytes8          = "bytes8"
	TypeBytes8Arr       = "bytes8[]"
	TypeBytes16         = "bytes16"
	TypeBytes16Arr      = "bytes16[]"
	TypeBytes32         = "bytes32"
	TypeBytes32Arr      = "bytes32[]"
	TypeStructPrefix    = "struct "
	TypeEnumPrefix      = "enum "
)

func parseType(typeName *fastjson.Value) string {
	nodeType := typeName.GetStringBytes("nodeType")
	switch string(nodeType) {
	case TypeNodeElementary:
		return parseElementaryType(typeName)
	case TypeNodeArray:
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
	//log.Println("req :", string(typeBytes))
	var str string
	switch string(typeBytes) {
	case TypeBool:
		str = "bool"
	case TypeUint8:
		str = "uint8"
	case TypeUint64:
		str = "uint64"
	case TypeUint256:
		str = "*big.Int"
	case TypeAddresss:
		str = "common.Address"
	case TypeString:
		str = "string"
	case TypeBytes:
		str = "[]byte"
	case TypeBytesMem:
		str = "[]byte"
	case TypeBytes4:
		str = "[4]byte"
	case TypeBytes4Arr:
		str = "[][4]byte"
	case TypeBytes8:
		str = "ablCommon.ID"
	case TypeBytes8Arr:
		str = "[]ablCommon.ID"
	case TypeBytes16:
		str = "[16]byte"
	case TypeBytes16Arr:
		str = "[][16]byte"
	case TypeBytes32:
		str = "common.Hash"
	case TypeBytes32Arr:
		str = "[]common.Hash"
	default:
		str = ""
	}
	//log.Println("res :", str)
	return str
}

func parseUserDefinedType(typeName *fastjson.Value) string {
	typeBytes := string(typeName.GetStringBytes("typeDescriptions", "typeString"))
	switch {
	case strings.HasPrefix(string(typeBytes), TypeStructPrefix):
		canonical := strings.TrimPrefix(string(typeBytes), TypeStructPrefix)
		splits := strings.Split(canonical, ".")
		return splits[len(splits)-1]

	case strings.HasPrefix(string(typeBytes), TypeEnumPrefix):
		canonical := strings.TrimPrefix(string(typeBytes), TypeEnumPrefix)
		return strings.Replace(canonical, ".", "", -1)
	}
	return ""
}

func parseMappingType(typeName *fastjson.Value) string {
	keyType := typeName.Get("keyType")
	keyT := parseType(keyType)
	valType := typeName.Get("valueType")
	valT := parseType(valType)
	return fmt.Sprintf("map[%s]%s", keyT, valT)
}
