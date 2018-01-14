package classfile

type ConstantNameAndType struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndType)readInfo(reader *ClassReader)  {
	self.nameIndex=reader.readUint16()
	self.descriptorIndex=reader.readUint16()
}