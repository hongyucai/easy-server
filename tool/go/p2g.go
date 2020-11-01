package main

import (
	"flag"
	"fmt"
	"go-mod/common/helper"
	"go-mod/servers/broker"
	"os"
	"regexp"
	"strings"
)

type deal_f func(string) string

type class_struck struct {
	package_name string
	imports      []string
	funcs        []string
}

type lang_struck struct {
	php_word  deal_f
	php_fun   deal_f
	php_model deal_f
	go_word   deal_f
	go_fun    deal_f
	go_while  deal_f
	go_for    deal_f
}

type lang_life_struck struct {
	params []string
	code   string
	lang   lang_struck
}

type fast_dev struct {
	init_go func(string) lang_struck
}

var (
	IS_DEBUG           = true
	DEBUG_LEVEL_ALL    = 1
	DEBUG_LEVEL_INFO   = 2
	DEBUG_LEVEL_WARN   = 3
	DEBUG_LEVEL_ERROR  = 4
	DEBUG_LEVEL_IMPORT = 5
)

func packageName(str string, go_pack *string) {
	phpRow := splitRow(str)
	namespace := splitRoute(phpRow[0])
	*go_pack = namespace[len(namespace)-1]
}

func imports(str string, go_import *[]string) {

	phpRow := splitRow(str)
	for _, val := range phpRow {
		pag := ""
		if find := strings.Index(val, "use"); find > 0 {
			row := splitRoute(val)
			if find := strings.Index(val, "Services"); find > 0 {
				pag = "services"
			} else if find := strings.Index(val, "Models"); find > 0 {
				pag = "models"
			} else {

			}
			if pag != "" {
				new_go_import := append(*go_import, pag+"."+row[len(row)-1])
				*go_import = new_go_import
			}
		}
	}
}

func string_v(toke string) string {
	toke = helper.Trim_Word(toke)
	val, _ := regexp.Compile(`([\d]+)`)
	var_php := val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1])) == len(toke) {
		helper.Error(toke, DEBUG_LEVEL_ERROR)
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`\$([\w\d\_]+)['(.*)']`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1]))+1 == len(toke) {
		helper.Log(string(len(var_php)), string(var_php[1]), DEBUG_LEVEL_ERROR)
		helper.Error(string(var_php[1]), DEBUG_LEVEL_ERROR)
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`\$(.*)`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1]))+1 == len(toke) {
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`\$(.*)->(.*)`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1]))+3+len(string(var_php[2])) == len(toke) { //
		helper.Error(string(var_php[1]), DEBUG_LEVEL_ERROR)
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`(".*")`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1])) == len(toke) {
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`('.*')`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1])) == len(toke) {
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`(true)`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1])) == len(toke) {
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`(false)`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1])) == len(toke) {
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`\[(.*)\]`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1]))+2 == len(toke) {
		str := ""
		if "" == string(var_php[1]) {
			str = "[]interface{}{}\n"
		} else {
			helper.Error(string(var_php[1]), DEBUG_LEVEL_ERROR)
		}
		return str
	}
	val, _ = regexp.Compile(`\$(.*)\[(.*)\]`)
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php) > 0 && len(string(var_php[1]))+len(string(var_php[2]))+3 == len(toke) {
		str := ""
		if "" == string(var_php[2]) {
			str = string(var_php[2]) + ""
		} else {
			str = string(var_php[1]) + "." + string(var_php[2])
		}
		return str
	}
	return ""
}
func class_v(s string, lang *lang_life_struck) string {
	reg := regexp.MustCompile(`(.*)\s*=>\s*([\$]*.*),`)
	var_php := reg.FindAllStringSubmatch(s, -1)

	if len(var_php) > 0 {
		str := ""
		assign := strings.Split(s, " = ")
		val_index := -1
		val := ""
		for index, param := range lang.params {

			if strings.Index(assign[0], param) > 0 {
				val_index = index
			}
		}
		if val_index > 0 {
			val = lang.params[val_index]
		} else {
			str = " mp := make(map[string]string)\n"
			val = "mp"
		}

		for _, v := range var_php {
			va := class_v(v[2], lang)
			str = str + val + "[" + helper.Trim_Word(v[1]) + "] = " + va + "\n"
		}
		return str
	}
	reg = regexp.MustCompile(`(.*)\s*->\s*([\$]*.*)`)
	var_php = reg.FindAllStringSubmatch(s, -1)

	if len(var_php) > 0 {
		vals := strings.Split(s, "->")
		if strings.Index(vals[1], "->") == -1 {
			return string_v(vals[0]) + "." + vals[1]
		} else {
			//all_vals := strings.Count(assign[1],"->")
			return string_v(vals[0]) + "." + vals[1]
		}
	}
	return ""
}
func string_vstring(s string, lang lang_struck) string {

	if lang.php_word(s) != "" {
		// 嵌套参数调用

	}
	statics := []string{}
	// xxxxxxxx(pppp)
	methods, _ := regexp.Compile(`([\w\d\_\-\>\$]+)\s*\(([\s\n\[\]]+.*)\)`)
	sub := methods.FindSubmatch([]byte(s))
	str := ""
	if len(sub) == 0 {
		return lang.go_word(s)
	}
	if len(sub) > 0 && lang.php_fun(string(sub[1])) != "" {
		str = lang.go_fun(s)
	}
	for _, val := range strings.Split(string(sub[2]), ",") {
		// 多参数
		if find := strings.Index(val, "go:"); find > -1 {
			statics = append(statics, lang.go_word(val[3:len(val)-1]))
		} else {
			statics = append(statics, string(sub[1])+"("+string_v(val)+")")
		}
	}
	str = str + strings.Join(statics, ",")
	return str
}

func init_go(s string) *lang_struck {
	// 遇到php特性 取一个完整的代码
	php_word := func(s string) string {
		if strings.Index(s, "array") > -1 {
			val, _ := regexp.Compile(`(array\([\n\r\t\s\w\d\=\-\>\'\"\,\$]+\))`)
			var_php := val.FindSubmatch([]byte(s))
			if len(var_php) > 0 {
				return string(var_php[1])
			}
		}
		if strings.Index(s, "list") > 0 {

		}
		return ""
	}
	php_fun := func(s string) string {

		if find := strings.Index(s, "this->"); find > -1 {
			return s[find+6:]
		}
		if find := strings.Index(s, "request->"); find > -1 {
			return "request." + s[find+9:]
		}
		return ""
	}
	php_model := func(s string) string {

		if find := strings.Index(s, "where"); find > -1 {
			return s[find+6:]
		}
		return ""
	}
	go_word := func(s string) string {
		s = helper.Trim_Word(s)
		if strings.Index(s, "array") > -1 {
			s = s[6 : len(s)-2]
			if strings.Index(s, "=>") > -1 {

				str := "mapv := make(map[string]string)\n"
				s_a := strings.Split(s[0:len(s)-2], "|")

				for _, ele := range s_a {
					string_slice := strings.Split(ele, "=>")
					if len(string_slice) <= 1 {
						continue
					}
					str = str + "mapv[\"" + string_v(string_slice[0]) + "\"] = " + string_v(string_slice[1]) + "\n"
				}
				return str
			} else {

			}
		}
		return s
	}
	go_while := func(s string) string {
		val, _ := regexp.Compile(`\!\$(.*)`)
		var_php := val.FindSubmatch([]byte(s))
		if len(var_php) > 0 {
			return string(var_php[1])
		}
		val, _ = regexp.Compile(`\$(.*)`)
		var_php = val.FindSubmatch([]byte(s))
		if len(var_php) > 0 {
			return "len(" + string(var_php[1]) + ")>0"
		}

		return s
	}
	go_for := func(s string) string {
		val, _ := regexp.Compile(`\((.*)as(.*)\)`)
		var_php := val.FindSubmatch([]byte(s))
		if len(var_php) > 0 {

			v_k := strings.Split(string(var_php[2]), "=>")

			if len(v_k) > 1 {
				return "for " + string_v(v_k[0]) + "," + string_v(v_k[1]) + " range " + string_v(string(var_php[1])) + "{"
			} else {
				return "for " + string_v(v_k[0]) + " range " + string_v(string(var_php[1])) + "{"
			}

		} else {
			helper.Log("errort", s, DEBUG_LEVEL_ERROR)
		}
		return s
	}
	go_fun := func(s string) string {
		helper.Log("cal", "go_fun::::::", DEBUG_LEVEL_INFO)
		if find := strings.Index(s, "this->"); find > -1 {
			s = s[find+6:]
		}
		if find := strings.Index(s, "where"); find > -1 {

		}
		return s
	}
	return &lang_struck{php_word, php_fun, php_model, go_word, go_fun, go_while, go_for}
}

func core(life *lang_life_struck) string {
	code := life.code
	dcode := ""
	phpDeal := func(str string, f deal_f) bool {
		if dcode = f(str); dcode != "" {
			return true
		} else {
			return false
		}
	}

	f_input := func(s string) string {
		helper.Log("f_input::::::::", s, DEBUG_LEVEL_INFO)
		param, _ := regexp.Compile(`\$request->route\('(.*)'\)`)
		sub := param.FindSubmatch([]byte(code))
		if len(sub) > 0 {
			return "input:" + string(sub[1])
		}
		param, _ = regexp.Compile(`\$request->input\('(.*)'\)`)
		sub = param.FindSubmatch([]byte(code))
		if len(sub) > 0 {
			return "input:" + string(sub[1])
		}
		param, _ = regexp.Compile(`\$request->route\('(.*)',(.*)\)`)
		sub = param.FindSubmatch([]byte(code))
		if len(sub) > 0 {
			return "input:" + string(sub[1]) + string(sub[2])
		}
		return ""
	}
	f_p := func(s string) string {
		helper.Log("f_p::::::::", s, DEBUG_LEVEL_INFO)
		if strings.Index(s, "if (") > -1 || strings.Index(s, "if(") > -1 || strings.Index(s, "if  (") > -1 {
			methods, _ := regexp.Compile(`if\s*\((.*)\)`) // \s+\{([\s\n]+.*)
			sub := methods.FindSubmatch([]byte(s))
			return "if " + life.lang.go_while(string(sub[1])) + "{"
		} else if find := strings.Index(s, "for "); find > -1 {
			string_slice := strings.Split(s, "{")
			if len(string_slice) == 1 {
				return s
			}
			fmt.Println(s)
			os.Exit(1)
			return string_slice[1]
		} else if find := strings.Index(s, "foreach "); find > -1 {
			string_slice := strings.Split(s, "{")
			return life.lang.go_for(string_slice[0])
		} else if find := strings.Index(s, "switch"); find > -1 {
			string_slice := strings.Split(s, "{")
			return string_slice[1]
		}
		return ""
	}
	f_call := func(s string) string {
		helper.Log("f_call::::::::", s, DEBUG_LEVEL_INFO)
		class := []string{}
		if find := strings.Index(s, "::"); find > -1 {
			class = strings.Split(s, "::")
		}
		if find := strings.Index(s, "new "); find > -1 {
			class = strings.Split(s, "new")
		}

		if len(class) > 0 {
			op_php := splitOperate(class[1])
			go_class_do := ""
			if len(op_php) > 0 {
				for _, val := range op_php {
					params := splitArray(string(val[1]))

					goParams := ""
					for _, param := range params {

						goParams = goParams + string_vstring(param, life.lang)

					}

					go_class_do += string_vstring(string(val[0]), life.lang) + goParams
				}
				return class[0] + "." + go_class_do

			} else {
				// static
				fmt.Println("errort f_call:::", class[1])
			}
		}
		return ""
	}
	f_v := func(s string) string {
		helper.Log("f_v::::::::", s, DEBUG_LEVEL_INFO)
		find := strings.Index(s, " = ")

		if find > -1 {
			va := strings.Split(s, " = ")
			if f_p(va[0]) != "" {
				return ""
			}
			values := ""

			if phpDeal(va[1], string_v) {
				values = dcode
			} else if phpDeal(va[1], f_call) {
				values = dcode
			} else if phpDeal(va[1], f_input) {
				values = dcode
			} else {
				values = class_v(s, life)
			}

			if helper.IsValueInList(string_v(va[0]), life.params) {
				return string_v(va[0]) + " = " + values
			} else {
				life.params = append(life.params, string_v(va[0]))
				return string_v(va[0]) + " := " + values
			}
		}
		return ""
	}
	f_return := func(s string) string {
		helper.Log("f_return::::::::", s, DEBUG_LEVEL_INFO)
		if strings.Index(s, "return ") > -1 {
			va := strings.Split(s, "return")
			if find := strings.Index(va[1], "("); find > -1 {
				return "(" + string_vstring(va[1], life.lang) + ")"
			} else {
				return string_v(va[1])
			}
		}
		return ""
	}

	if phpDeal(code, f_input) ||
		phpDeal(code, f_v) ||
		phpDeal(code, f_p) ||
		phpDeal(code, f_call) ||
		phpDeal(code, f_return) {

		return dcode
	} else {
		helper.Log("errort", code, DEBUG_LEVEL_ERROR)
		return code
	}
}

func funcs(str string, go_func *[]string) {
	phpFuns := splitFunction(str)

	d_fun_name := func(code string) string {
		string_slice := strings.Split(code, "(")
		return string_slice[0]
	}
	d_fun_param := func(code string) string {
		code = strings.Replace(code, ")", "", -1)
		string_slice := splitFunHeader(code)
		var p []string
		for _, val := range string_slice {
			param := strings.Replace(val, "Request $request", "context iris.Context", -1)
			p = append(p, param)
		}
		return strings.Join(p, "")
	}
	is_block := func(code string) bool {

		if strings.Index(code, " {") > -1 && strings.Index(code, " }") > -1 {
			return true
		}
		return false
	}
	CodeT := lang_life_struck{}
	d_fun_block := func(code string) string {

		return CodeT.codeDeal(code)
	}
	type fun_c struct {
		fun_name string
		params   string
		returns  string
		block    []string
		class    string
	}
	//phpFuns[1] = phpFuns[1][strings.Index(phpFuns[1],"{") : len(phpFuns[1])]
	//错误拦截必须配合defer使用  通过匿名函数使用
	//defer func() {
	//	//恢复程序的控制权
	//	err := recover()
	//	if err != nil {
	//		fmt.Println(err)
	//		//helper.Error( "error",DEBUG_LEVEL_ERROR)
	//
	//	}
	//}()

	for i := 1; i < len(phpFuns); i++ {
		CodeT.params = nil
		fun_code := fun_c{}
		val := phpFuns[i]
		fmt.Println("function ::::", val)

		fun_code.fun_name = d_fun_name(val)
		fun_code.params = d_fun_param(val)
		CodeBlock := splitBlock(val[strings.Index(val, "{")+1 : strings.LastIndex(val, "}")])
		fmt.Println("CodeBlock::::::::::::::::::::::::::::::::::::", CodeBlock)
		os.Exit(1)
		CodeRow := []string{}
		for _, val := range CodeBlock {
			if len(helper.Trim_Word(val)) == 0 {
				continue
			}

			if is_block(val) {
				block_before := val[0:strings.Index(val, " {")]
				val = val[strings.Index(val, " {"):]
				fmt.Println(block_before)
				os.Exit(1)
				fun_code.block = append(fun_code.block, d_fun_block(block_before)+"\n")
			}

			CodeRow = splitRow(val)
			decode := ""
			for _, val := range CodeRow {

				fmt.Println("php code----------", val)
				decode = d_fun_block(val)
				fun_code.block = append(fun_code.block, helper.Copy_Word(decode, " ", -helper.Space_Word(val))+"\n")
				fmt.Println("go code ----------", decode)
			}
		}
		fmt.Println("go fun ::::", fun_code)

	}

}

func (do *lang_life_struck) codeDeal(str string) string {
	if str == "" {
		return str
	}
	//do.params = nil
	do.code = str
	do.lang = *init_go(str)
	return core(do)
}

func (code *class_struck) structDeal(str string) string {
	if str == "" {
		return str
	}
	str = strings.Replace(str, "<?php", "", -1)
	str = strings.Replace(str, "?>", "", -1)

	packageName(str, &code.package_name)
	imports(str, &code.imports)
	funcs(str, &code.funcs)
	fmt.Println("finish...")
	os.Exit(1)

	return str
}

func splitRow(str string) []string {
	string_slice := strings.Split(str, ";")
	return string_slice
}

func splitRoute(str string) []string {
	string_slice := strings.Split(str, "\\")
	return string_slice
}
func splitArray(str string) []string {
	p, _ := regexp.Compile(`\],[\s\n]*\[`)
	all_ix := p.FindAllIndex([]byte(str), -1)

	restr := []string{}
	strl := 0
	strr := 0
	for i := 0; i < len(all_ix); i++ {
		// 查找最近的 ] , [
		strr = all_ix[i][0] + 2
		restr = append(restr, str[strl:strr])
		strl = strr
	}
	restr = append(restr, str[strl:])
	return restr
}
func splitFunction(str string) []string {

	string_slice := strings.Split(str, "function")
	return string_slice
}
func splitFunHeader(str string) []string {
	string_slice := strings.Split(str, "{")
	string_slice = strings.Split(string_slice[0], "(")
	string_slice = strings.Split(string_slice[1], ",")
	return string_slice
}
func splitOperate(str string) [][]string {
	// ->xxxxxxxx(ppppppppp)->sssssssssss(dddddddddd) (.*)(\(.*\))
	ops := strings.Split(str, `)->`)

	opstrings := [][]string{}
	for _, val := range ops {
		val = val + ")"
		opstrings = append(opstrings, []string{val[0:strings.Index(val, "(")], val[strings.Index(val, "("):strings.LastIndex(val, ")")]})
		//if error != nil {
		//	panic(`regexp: Compile(` + quote(str) + `): ` + error.Error())
		//}

	}
	return opstrings
}
func splitBlock(str string) []string {

	p1, _ := regexp.Compile(`\{\s`)
	all_ix1 := p1.FindAllIndex([]byte(str), -1)
	p2, _ := regexp.Compile(`\s\}`)
	all_ix2 := p2.FindAllIndex([]byte(str), -1)
	if len(all_ix1) != len(all_ix2) {
		fmt.Println("errort", all_ix1, all_ix2, str)
	}
	if len(all_ix1) == 0 && len(all_ix2) == 0 {
		return []string{str}
	}
	restr := []string{}
	strl := 0
	strr := 0
	for i := 0; i < len(all_ix1); i++ {
		// 查找最近的 { }
		strr = strl + strings.Index(str, " }") + 2
		if strr == all_ix2[i][1] {
			// {} {}
			pos_block := strings.Index(str, "{ ")
			if pos_block > -1 {
				pos_block = strings.LastIndex(str[0:pos_block], ";")
			}
			restr = append(restr, str[0:pos_block+1])
			restr = append(restr, str[pos_block+1:all_ix2[i][1]-strl])
			fmt.Println("DDDDDDDDDDDDDDDDD", str[pos_block+1:all_ix2[i][1]-strl])
			str = str[all_ix2[i][1]-strl:]
			strl = all_ix2[i][1]
			//os.Exit(1)
		} else {
			// { { } }
			fmt.Println("dddddddddddd", strr, all_ix2[i][1], str[strr:])
			os.Exit(1)
		}
	}
	restr = append(restr, str)
	//for _,val := range restr{
	//	fmt.Println("row------------->",val+"\n")
	//}
	//fmt.Println(str,restr)
	//os.Exit(1)
	return restr
}
func load(file string) string {
	helper.Log("loading file...", file, DEBUG_LEVEL_INFO)
	if ok, _ := helper.PathExists(file); !ok {
		return ""
	}
	conent := helper.IoutilRead(file)
	return conent
}

func main() {
	flag.Parse()
	file := flag.Arg(0)
	//str_slices := splitRow(load(file))
	CodeBlock := &class_struck{"", []string{}, []string{}}
	CodeBlock.structDeal(load(file))
	fmt.Println(broker.Broker)
	//for _,val := range str_slices {
	//	val = structDeal(val)
	//	fmt.Println(val)
	//	//fmt.Println(splitRoute(val)[len(splitRoute(val))-1])
	//	os.Exit(1 )
	//}
	os.Exit(1)
}
