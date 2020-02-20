package dictionary

import (
	. "github.com/FactomProject/factomd/electionsCore/interpreter/common"
	. "github.com/FactomProject/factomd/electionsCore/interpreter/names"
)

type DictionaryEnrty struct {
	N Name
	FlagsStruct
	E interface{}
}

type Dictionary map[Name]DictionaryEnrty

func NewDictionary() Dictionary {
	return make(map[Name]DictionaryEnrty, 0)
}

func (d Dictionary) Add(n Name, e DictionaryEnrty) { d[n.GetRawName()] = e }
