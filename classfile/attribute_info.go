package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp *ConstantPool) []AttributeInfo {
	attrLength := reader.readUint16()
	attributes := make([]AttributeInfo, attrLength)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}
func readAttribute(reader *ClassReader, cp *ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attributeInfo := newAttributeInfo(attrName, cp)
	if attributeInfo == nil {
		attributeInfo = &UnparsedAttribute{
			name:   attrName,
			length: attrLen,
		}
	}
	attributeInfo.readInfo(reader)
	return attributeInfo
}
func newAttributeInfo(attrName string, cp *ConstantPool) AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	// case "MethodParameters":
	// case "RuntimeInvisibleAnnotations":
	// case "RuntimeInvisibleParameterAnnotations":
	// case "RuntimeInvisibleTypeAnnotations":
	// case "RuntimeVisibleAnnotations":
	// case "RuntimeVisibleParameterAnnotations":
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	// case "SourceDebugExtension":
	// case "StackMapTable":
	case "Synthetic":
		return _attrSynthetic
	default:
		return nil // undefined attr
	}
}
