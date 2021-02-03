package ex12_10

import (
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	kindListStart = '('
	kindListEnd   = ')'
)

var interfaceTypes map[string]reflect.Type

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func init() {
	interfaceTypes = make(map[string]reflect.Type)
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic("invalid rune")
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.Int:
		intVal, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(intVal))
		lex.next()
		return
	case scanner.Float:
		floatVal, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(floatVal)
		lex.next()
		return
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case kindListStart:
		lex.next()
		listing(lex, v)
		lex.next()
		return
	}
	panic("unexpected token")
}

func listing(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !end(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !end(lex) {
			elem := reflect.New(v.Type()).Elem().Elem()
			read(lex, elem)
			v.Set(reflect.Append(v, elem))
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !end(lex) {
			lex.consume(kindListStart)
			key := reflect.New(v.Type().Key().Elem())
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, key)
			v.SetMapIndex(key, value)
			lex.consume(kindListEnd)
		}
	case reflect.Struct:
		for !end(lex) {
			lex.consume(kindListStart)
			if lex.token != scanner.Ident {
				panic("unexpected token")
			}
			n := lex.text()
			lex.text()
			read(lex, v.FieldByName(n))
			lex.consume(kindListEnd)
		}
	case reflect.Interface:
		name := strings.Trim(lex.text(), "\"")
		lex.next()
		if _, ok := interfaceTypes[name]; !ok {
			panic("unknown types interface value")
		}
		v := reflect.New(interfaceTypes[name])
		read(lex, reflect.Indirect(v))
		v.Set(reflect.Indirect(v))
	default:
		panic("decode failed")
	}
}

func end(lex *lexer) bool {
	switch lex.token {
	case kindListEnd:
		panic("reached to end of file")
	case scanner.EOF:
		return true
	}
	return false
}
