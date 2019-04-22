package codegen

func (o *Object) UniqueFields() map[string]*Field {
	m := map[string]*Field{}

	for _, f := range o.Fields {
		if f.IsResolver {
			m[f.GoFieldName] = f
		}
	}

	return m
}
