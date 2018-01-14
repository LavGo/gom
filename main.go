package main

import (
	"fmt"
	"github.com/LavGo/gom/classpath"
	"github.com/LavGo/gom/cmd"
	"strings"
	"github.com/LavGo/gom/classfile"
)

func main() {
	cmm := cmd.ParseCmd()
	if cmm.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmm.HelpFlag || cmm.Class == "" {
		cmd.PrintUsage()
	} else {
		startGom(cmm)
	}
}
func startGom(cm *cmd.Cmd) {
	cp := classpath.Parse(cm.XjreOption, cm.CpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cp, cm.Class, cm.Args)
	className := strings.Replace(cm.Class, ".", "/", -1)
	cf:=loadClass(className,cp)
	printClassInfo(cf)
}
func loadClass(className string,cp *classpath.Classpath)*classfile.ClassFile{
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", className)
		return nil
	}
	cf,err:=classfile.Parse(classData)
	if err !=nil{
		panic(err)
	}

	//fmt.Printf("class data:%v\n", classData)
	return cf
}

func printClassInfo(cf *classfile.ClassFile){
	fmt.Printf("Magic: 0x%x\n",cf.Magic())
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool().Infos()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlags())
	fmt.Printf("this class: %v\n",cf.ClassName())
	fmt.Printf("super class: %v\n",cf.SuperClassName())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n",len(cf.Fields()))
	for _,f:= range cf.Fields(){

		fmt.Printf("    AccessFlags: 0x%x\n",f.AcccessFlags())
		fmt.Printf("    Name: %s\n",f.Name())
		fmt.Printf("    Descriptor: %s\n",f.Descriptor())
		fmt.Printf("    Attributes:\n")
		for _,a:=range f.AttributeInfos(){
			fmt.Printf("        %v\n",a)
		}

	}
	fmt.Printf("methods count: %v\n",len(cf.Methods()))
	for _,m:= range cf.Methods(){
		fmt.Printf("    AccessFlags: 0x%x\n",m.AcccessFlags())
		fmt.Printf("    Name: %s\n",m.Name())
		fmt.Printf("    Descriptor: %s\n",m.Descriptor())
		fmt.Printf("    Attributes:\n")
		for _,a:=range m.AttributeInfos(){
			fmt.Printf("        %v\n",a)
		}
	}
}