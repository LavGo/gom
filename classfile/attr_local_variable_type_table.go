package classfile

type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}
type LocalVariableTypeTableEntry struct {
	startPc uint16
	length uint16
	nameIndex uint16
	signatureIndex uint16
	index uint16
}
func (self *LocalVariableTypeTableAttribute)readInfo(reader *ClassReader){
	len:=reader.readUint16()
	self.localVariableTypeTable=make([]*LocalVariableTypeTableEntry,len)
	for index:= range self.localVariableTypeTable{
		self.localVariableTypeTable[index]=&LocalVariableTypeTableEntry{
			startPc:reader.readUint16(),
			length:reader.readUint16(),
			nameIndex:reader.readUint16(),
			signatureIndex:reader.readUint16(),
			index:reader.readUint16(),
		}
	}
}