package ast // Mini

import (
	"errors"
	"fmt"

	"github.com/wk8/go-ordered-map/v2"
)

var StructTable = make(map[string]*orderedmap.OrderedMap[string, Type])
var SymbolTable = make(map[string]Type)

func TypeCheck(root *Root) error {
	return generateStructTable(root.Types)
}

func generateStructTable(types []*TypeDeclaration) (err error) {
	for _, v := range types {
		id := v.Id

		// Check for re-defined
		if _, ok := StructTable[id]; ok {
			e := fmt.Errorf("%s: error: re-definition of 'struct %s'", v.Position, id)
			err = errors.Join(err, e)
			continue
		}

		// Check validity of fields
		om, e := populateFields(id, v.Fields)
		err = errors.Join(err, e)

		if e != nil {
			continue
		}

		// Add to struct table
		fmt.Printf("Adding to table: struct %v\n", id)
		StructTable[id] = om
	}

	return
}

func populateFields(structId string,
	fields []*Declaration) (om *orderedmap.OrderedMap[string, Type], err error) {

	// Create fields map
	om = orderedmap.New[string, Type]()

	// Loop through all fields
	for _, f := range fields {
		name := f.Name

		// Check that struct field names are valid
		if v, ok := f.Type.(StructType); ok {
			id := v.Id

			if _, present := StructTable[id]; !present && id != structId {
				e := fmt.Errorf("%s: error: struct '%s' not defined (yet?)", f.Position, id)
				err = errors.Join(err, e)
				continue
			}
		}

		// Add to the field map (and check for re-definition)
		if _, present := om.Set(name, f.Type); present {
			e := fmt.Errorf("%s: error: re-definition of field '%s'", f.Position, name)
			err = errors.Join(err, e)
		}
	}

	return
}
