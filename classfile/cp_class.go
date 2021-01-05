package classfile

type ConstantClassInfo struct {
	cp        *ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name(reader *ClassReader) string {
	return self.cp.getUtf8(self.nameIndex)
}
