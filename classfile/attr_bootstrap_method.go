package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}

func (self *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	bootstrapMethodsLength := reader.readUint32()
	self.bootstrapMethods = make([]*BootstrapMethod, bootstrapMethodsLength)
	for i := range self.bootstrapMethods {
		self.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}
