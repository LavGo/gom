package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool *ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}
func (self *ClassFile) Magic() uint32 {
	return self.magic
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass) //
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) SuperClassName() string {
	return self.constantPool.getClassName(self.superClass)
}
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for index := range interfaceNames {
		interfaceNames[index] = self.constantPool.getClassName(self.interfaces[index])
	}
	return interfaceNames
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) readConstantPool(reader *ClassReader) {
	self.constantPool = &ConstantPool{cf: self}
	self.constantPool.read(reader)
}
