package main

import (
	"flag"
	"fmt"
	"go-mod/common/helper"
	"os"
	"regexp"
	"strings"
)

type deal_f func(string)(string)

type class_struck struct {
	package_name string
	imports []string
	funcs []string
}

type lang_struck struct {
	php_word deal_f
	php_fun deal_f
	go_word deal_f
	go_fun deal_f
	go_while deal_f
}

type lang_life_struck struct {

	code string
	lang lang_struck
}

type fast_dev struct {
	init_go func(string)(lang_struck)
}

func packageName(str string,go_pack *string)  {
	phpRow := splitRow(str)
	namespace := splitRoute(phpRow[0])
	*go_pack = namespace[len(namespace)-1]
}

func imports(str string,go_import *[]string)  {

	phpRow := splitRow(str)
	for _,val := range phpRow {
		pag := ""
		if find := strings.Index(val,"use");find>0{
			row := splitRoute(val)
			if find := strings.Index(val,"Services");find>0{
				pag = "services"
			} else if find := strings.Index(val,"Models");find>0{
				pag = "models"
			}else{

			}
			if pag!="" {
				new_go_import := append( *go_import,pag+"."+row[len(row)-1] )
				*go_import = new_go_import
			}
		}
	}
}

func s_v(toke string) string {
	toke = helper.Trim_Word(toke)
	val, _ := regexp.Compile(`'([\d]+)'` )
	var_php := val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return  string(var_php[1])
	}
	val, _ = regexp.Compile(`\$([\w\d\_]+)['(.*)']` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return  string(var_php[1])
	}
	val, _ = regexp.Compile(`\$(.*)` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`\$(.*)->(.*)` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`"(.*)"` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`'(.*)'` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`(true)` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	val, _ = regexp.Compile(`(false)` )
	var_php = val.FindSubmatch([]byte(toke))
	if len(var_php)>0{
		return string(var_php[1])
	}
	//val, _ = regexp.Compile(`(.*)->(.*)` )
	//var_php = val.FindSubmatch([]byte(toke))
	//if len(var_php)>0{
	//	return string(var_php[1])+ string(var_php[2])
	//}
	return ""
}
func s_vstring(s string,lang lang_struck) string {
	php_string := lang.php_word(s)
	if strings.Index(s,",")>-1{
		s = strings.Replace( s,php_string,"go:"+strings.Replace(php_string,",","|",-1),-1 )
	}
	statics := []string{}
	//methods, _ :=regexp.Compile(`([\w\d\_\-\>\$]+)\(([\t\s\n\r\w\d\,\'\"\(\)\$\:\=\>\|]+)` )
	methods, _ :=regexp.Compile(`([\w\d\_\-\>\$]+)\s*\((.*)\)` )
	sub := methods.FindSubmatch([]byte(s))
	str := ""
	if len(sub)>0 && lang.php_fun(string(sub[1]))!=""{
		str = lang.go_fun(string(sub[1]))
		if find := strings.Index(str,"where");find>-1{

		}
	}
	for _,val := range strings.Split( string(sub[2]),"," ){
		if find:=strings.Index(val,"go:");find>-1{
			statics = append(statics,lang.go_word( val[3:len(val)-1] ) )
		}else {
			statics = append(statics,string(sub[1])+"("+s_v( val )+")")
		}
	}
	str = str + strings.Join( statics,"," )
	return str
}

func init_go(s string) *lang_struck {
	//php_while := func(s string) bool {
	//	val, _ := regexp.Compile(`\!\$(.*)` )
	//	var_php := val.FindSubmatch([]byte(s))
	//	if len(var_php)>0{
	//		return true
	//	}
	//	return flase
	//}
	php_word := func(s string) string {
		if strings.Index(s, "array")>-1{
			val, _ := regexp.Compile(`(array\([\n\r\t\s\w\d\=\-\>\'\"\,\$]+\))` )
			var_php := val.FindSubmatch([]byte(s))
			if len(var_php)>0{
				return  string(var_php[1])
			}
		}
		if strings.Index(s, "list")>0{

		}
		return s
	}
	php_fun := func(s string) string {

		if find:=strings.Index(s,"this->");find>-1{
			return s[find+6:]
		}
		return ""
	}
	go_word := func(s string) string {
		fmt.Println("goword::::::",s)
		s = helper.Trim_Word(s)
		if strings.Index(s, "array")>-1{
			s = s[6:len(s)-2]
			if strings.Index(s,"=>")>-1{

				str := "mapv := make(map[string]string)\n"
				s_a := strings.Split(s[0:len(s)-2],"|")

				for _,ele := range s_a{
					string_slice := strings.Split(ele,"=>")
					if len(string_slice)<=1{
						continue;
					}
					str = str+"mapv[\""+s_v(string_slice[0])+"\"] = "+s_v(string_slice[1])+"\n"
				}
				return str
			}else{

			}
		}
		return s
	}
	go_while := func(s string) string {
		val, _ := regexp.Compile(`\!\$(.*)` )
		var_php := val.FindSubmatch([]byte(s))
		if len(var_php)>0{
			return  string( var_php[1] )
		}
		val, _ = regexp.Compile(`\$(.*)` )
		var_php = val.FindSubmatch([]byte(s))
		if len(var_php)>0{
			return  "len("+string( var_php[1] )+")>0"
		}

		return s
	}
	go_fun := func(s string) string {
		fmt.Println("go_fun::::::",s)
		if find:=strings.Index(s,"this->");find>-1{
			s=s[find+6:]
		}
		return ""
	}
	return &lang_struck{ php_word,php_fun,go_word,go_while,go_fun }
}

func core(life *lang_life_struck)  string {
	code := life.code
	f_input := func (s string) string {
		fmt.Println("f_input::::::::")
		param, _ := regexp.Compile(`\$request->route\('(.*)'\)` )
		sub := param.FindSubmatch([]byte(code))
		if len(sub)>0{
			return "input:" + string(sub[1])
		}
		param, _ = regexp.Compile(`\$request->input\('(.*)'\)` )
		sub = param.FindSubmatch([]byte(code))
		if len(sub)>0{
			return "input:" + string(sub[1])
		}
		return s
	}
	f_p := func(s string) string {
		fmt.Println("f_p::::::::",s)
		if strings.Index(s,"if (")>-1 || strings.Index(s,"if(")>-1 || strings.Index(s,"if  (")>-1{
			methods, _ :=regexp.Compile(`if\s*\((.*)\)` ) // \s+\{([\s\n]+.*)
			sub := methods.FindSubmatch([]byte(s))
			return "if " + life.lang.go_while( string( sub[1] ) )
		} else
		if find := strings.Index(s,"for");find>-1{
			string_slice := strings.Split(s,"{")
			return string_slice[1]
		} else
		if find := strings.Index(s,"foreach");find>-1{
			string_slice := strings.Split(s,"{")
			return string_slice[1]
		} else
		if find := strings.Index(s,"switch");find>-1{
			string_slice := strings.Split(s,"{")
			return string_slice[1]
		}
		return s
	}
	f_call := func(s string) string {
		fmt.Println("f_call::::::::")
		class := []string{}
		if find := strings.Index(s,"::");find>-1{
			class = strings.Split(s,"::")
		}
		if find := strings.Index(s,"new ");find>-1{
			class = strings.Split(s,"new")
		}

		if len(class)>0{
			//fmt.Println("string>",s)
			//os.Exit(1 )
			op, _ := regexp.Compile(`->(.*)(\(.*\))` )
			op_php := op.FindAllSubmatch([]byte(class[1]),-1)
			//op := strings.Split(class[1],"->")
			go_class_do := ""
			//fmt.Println("ddddddddddddddddd",op_php)
			os.Exit(1)
			for _,val := range op_php {
				fmt.Println("ddddddddddddddddd",val)
				os.Exit(1)
				for _,v := range val{
					cval, _ := regexp.Compile(`(\(.*\))` )
					var_php := cval.FindSubmatch([]byte(v))
					fmt.Println("string :::",string(var_php[1]))
					if len(var_php)>0 {
						go_class_do += s_vstring( string(v),life.lang )
					}else{
						fmt.Println("error :::",val)
					}
				}
			}
			return class[0]+"."+go_class_do
		}
		return s
	}
	f_v := func(s string) string {
		fmt.Println("f_v::::::::")
		find := strings.Index(s,"=")
		//find1 := strings.Index(s,"=>")
		//find2 := strings.Index(s,"==")
		//find3 := strings.Index(s,"= >")
		if find>-1 { //&& find1==-1 && find2==-1 && find3==-1
			va := strings.Split(s,"=")
			return s_v(va[0])+f_call(va[1])
		}
		return s
	}
	f_return := func(s string) string {
		fmt.Println("f_return::::::::")
		if strings.Index(s,"return ")>-1{
			va := strings.Split(s,"return")
			if find := strings.Index(va[1],"(");find>-1{
				return "("+ s_vstring( va[1],life.lang ) +")"
			}else {
				return s_v(va[1])
			}
		}
		return s
	}
	dcode := ""
	phpDeal := func(str string,f deal_f) bool {
		if dcode = f(str);dcode!=str{
			return true
		}else {
			return false
		}
	}
	if  phpDeal(code,f_input) ||
		phpDeal(code,f_v) ||
		phpDeal(code,f_p) ||
		phpDeal(code,f_call)||
		phpDeal(code,f_return) {

		return dcode
	}else {
		return code
	}
}

func funcs(str string,go_func *[]string) {
	phpFuns := splitFunction(str)

	d_fun_name := func(code string) string {
		string_slice := strings.Split(code,"(")
		return string_slice[0]
	}
	d_fun_param := func(code string) string {
		code = strings.Replace(code,")","",-1)
		string_slice := strings.Split(code,"{")
		string_slice = strings.Split(string_slice[0],"(")
		string_slice = strings.Split(string_slice[1],",")
		var p []string
		for _,val := range string_slice {
			param := strings.Replace(val,"Request $request","context iris.Context",-1 )
			p = append(p, param)
		}
		return strings.Join(p,"")
	}
	is_block := func(code string) bool {
		if strings.Index(code,"{")>-1 && strings.Index(code,"}")>-1{
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
		params string
		returns string
		block []string
		class string
	}
	//phpFuns[1] = phpFuns[1][strings.Index(phpFuns[1],"{") : len(phpFuns[1])]

	for i:=1;i<len(phpFuns);i++ {
		fun_code := fun_c{}
		val := phpFuns[i]
		fmt.Println("function ::::",val)

		fun_code.fun_name = d_fun_name(val)
		fun_code.params = d_fun_param(val)
		CodeBlock := splitBlock(val[strings.Index(val,"{")+1:strings.LastIndex(val,"}")])
		fmt.Println("CODE BLOCK::::",len(CodeBlock),CodeBlock)

		CodeRow := []string{}
		for _,val := range CodeBlock {

			if is_block(val) {
				block_before := val[0:strings.Index(val,"{")]
				val = val[strings.Index(val,"{"):]
				fun_code.block = append( fun_code.block,d_fun_block(block_before) )
			}

			CodeRow = splitRow(val)
			decode:=""
			for _, val := range CodeRow {

				fmt.Println("php code----------", val)
				decode = d_fun_block(val)
				fun_code.block = append( fun_code.block,decode+"\n" )
				fmt.Println("go code ----------", decode)
			}
		}
		fmt.Println("go fun ::::",fun_code)

	}

}

func (do *lang_life_struck) codeDeal(str string) string {
	if str==""{
		return str
	}
	do.code = str
	do.lang = *init_go( str )
	return core(do)
}

func (code *class_struck) structDeal(str string) string {
	if str==""{
		return str
	}
	str = strings.Replace(str, "<?php", "",-1 )
	str = strings.Replace(str, "?>", "",-1 )

	packageName(str, &code.package_name )
	imports(str, &code.imports )
	funcs(str, &code.funcs )
	fmt.Println( code )
	os.Exit(1)

	return str
}

func splitRow(str string) []string {
	string_slice := strings.Split(str,";")
	return string_slice
}

func splitRoute(str string) []string  {
	string_slice := strings.Split(str,"\\")
	return string_slice
}

func splitFunction(str string) []string  {

	string_slice := strings.Split(str,"function")
	return string_slice
}
func splitBlock(str string) []string  {

	p1, _ := regexp.Compile("{" )
	all_ix1 := p1.FindAllIndex([]byte(str), -1)
	p2, _ := regexp.Compile("}" )
	all_ix2 := p2.FindAllIndex([]byte(str), -1)
	if len(all_ix1)!=len(all_ix2){
		fmt.Println("errort",all_ix1,all_ix2,str)
	}
	if len(all_ix1)==0&&len(all_ix2)==0{
		return []string{str}
	}
	restr := []string{}
	strl := 0
	strr := 0
	for i:=1;i<=len(all_ix1);i++{
		// 查找最近的 { }
		strr = strl + strings.Index(str,"}")+1
		if strr==all_ix2[i-1][1]  {
			// {} {}
			pos_block := strings.Index(str,"{")
			if pos_block>-1{
				pos_block = strings.LastIndex(str[0:pos_block],";")
			}
			restr = append(restr, str[0:pos_block+1] )
			restr = append(restr, str[pos_block+1:all_ix2[i-1][1]-strl] )
			str = str[all_ix2[i-1][1]-strl:]
			strl = all_ix2[i-1][1]
			//os.Exit(1)
		}else{
			// { { } }
			fmt.Println("dddddddddddd",strr,all_ix2[i-1][0],str)
			os.Exit(1)
		}
		//for j:=0;j<len(all_ix2);j++{
		//
		//	//strl = strings.Index(str,"{")
		//
		//	fmt.Println( all_ix1[j],all_ix2[j] )
		//
		//
		//		str2 = str[strr:]
		//	}else{
		//
		//
		//	}
		//}
	}
	restr = append(restr, str )
	//for _,val := range restr{
	//	fmt.Println("row------------->",val+"\n")
	//}
	//fmt.Println(str,restr)
	//os.Exit(1)
	return  restr
}
func load( file string ) (string) {
	fmt.Println("loading file...",file)
	if ok,_ := helper.PathExists(file);!ok{
		return ""
	}
	conent := helper.IoutilRead(file)
	return conent
}

func main()  {

	flag.Parse()
	file := flag.Arg(0)
	//str_slices := splitRow(load(file))
	CodeBlock := &class_struck{ "",[]string{},[]string{} }
	CodeBlock.structDeal(load(file))
	//for _,val := range str_slices {
	//	val = structDeal(val)
	//	fmt.Println(val)
	//	//fmt.Println(splitRoute(val)[len(splitRoute(val))-1])
	//	os.Exit(1 )
	//}
	os.Exit(1 )
}
