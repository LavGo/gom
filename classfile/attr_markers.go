package classfile

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}

type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}
