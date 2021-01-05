package classfile

type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}
type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (self *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfInnerClasses := reader.readUint16()
	self.classes = make([]*InnerClassInfo, numberOfInnerClasses)
	for index := range self.classes {
		self.classes[index] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
