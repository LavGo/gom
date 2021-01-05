package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	stratPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			stratPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
