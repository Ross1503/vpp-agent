// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package interfaces

import "reflect"

var Types = map[string]reflect.Type{
	"CreateLoopback":                   reflect.TypeOf((*CreateLoopback)(nil)).Elem(),
	"CreateLoopbackInstance":           reflect.TypeOf((*CreateLoopbackInstance)(nil)).Elem(),
	"CreateLoopbackInstanceReply":      reflect.TypeOf((*CreateLoopbackInstanceReply)(nil)).Elem(),
	"CreateLoopbackReply":              reflect.TypeOf((*CreateLoopbackReply)(nil)).Elem(),
	"CreateSubif":                      reflect.TypeOf((*CreateSubif)(nil)).Elem(),
	"CreateSubifReply":                 reflect.TypeOf((*CreateSubifReply)(nil)).Elem(),
	"CreateVlanSubif":                  reflect.TypeOf((*CreateVlanSubif)(nil)).Elem(),
	"CreateVlanSubifReply":             reflect.TypeOf((*CreateVlanSubifReply)(nil)).Elem(),
	"DeleteLoopback":                   reflect.TypeOf((*DeleteLoopback)(nil)).Elem(),
	"DeleteLoopbackReply":              reflect.TypeOf((*DeleteLoopbackReply)(nil)).Elem(),
	"DeleteSubif":                      reflect.TypeOf((*DeleteSubif)(nil)).Elem(),
	"DeleteSubifReply":                 reflect.TypeOf((*DeleteSubifReply)(nil)).Elem(),
	"InterfaceNameRenumber":            reflect.TypeOf((*InterfaceNameRenumber)(nil)).Elem(),
	"InterfaceNameRenumberReply":       reflect.TypeOf((*InterfaceNameRenumberReply)(nil)).Elem(),
	"SwInterfaceAddDelAddress":         reflect.TypeOf((*SwInterfaceAddDelAddress)(nil)).Elem(),
	"SwInterfaceAddDelAddressReply":    reflect.TypeOf((*SwInterfaceAddDelAddressReply)(nil)).Elem(),
	"SwInterfaceClearStats":            reflect.TypeOf((*SwInterfaceClearStats)(nil)).Elem(),
	"SwInterfaceClearStatsReply":       reflect.TypeOf((*SwInterfaceClearStatsReply)(nil)).Elem(),
	"SwInterfaceDetails":               reflect.TypeOf((*SwInterfaceDetails)(nil)).Elem(),
	"SwInterfaceDump":                  reflect.TypeOf((*SwInterfaceDump)(nil)).Elem(),
	"SwInterfaceEvent":                 reflect.TypeOf((*SwInterfaceEvent)(nil)).Elem(),
	"SwInterfaceGetTable":              reflect.TypeOf((*SwInterfaceGetTable)(nil)).Elem(),
	"SwInterfaceGetTableReply":         reflect.TypeOf((*SwInterfaceGetTableReply)(nil)).Elem(),
	"SwInterfaceSetFlags":              reflect.TypeOf((*SwInterfaceSetFlags)(nil)).Elem(),
	"SwInterfaceSetFlagsReply":         reflect.TypeOf((*SwInterfaceSetFlagsReply)(nil)).Elem(),
	"SwInterfaceSetMacAddress":         reflect.TypeOf((*SwInterfaceSetMacAddress)(nil)).Elem(),
	"SwInterfaceSetMacAddressReply":    reflect.TypeOf((*SwInterfaceSetMacAddressReply)(nil)).Elem(),
	"SwInterfaceSetMtu":                reflect.TypeOf((*SwInterfaceSetMtu)(nil)).Elem(),
	"SwInterfaceSetMtuReply":           reflect.TypeOf((*SwInterfaceSetMtuReply)(nil)).Elem(),
	"SwInterfaceSetRxMode":             reflect.TypeOf((*SwInterfaceSetRxMode)(nil)).Elem(),
	"SwInterfaceSetRxModeReply":        reflect.TypeOf((*SwInterfaceSetRxModeReply)(nil)).Elem(),
	"SwInterfaceSetTable":              reflect.TypeOf((*SwInterfaceSetTable)(nil)).Elem(),
	"SwInterfaceSetTableReply":         reflect.TypeOf((*SwInterfaceSetTableReply)(nil)).Elem(),
	"SwInterfaceSetUnnumbered":         reflect.TypeOf((*SwInterfaceSetUnnumbered)(nil)).Elem(),
	"SwInterfaceSetUnnumberedReply":    reflect.TypeOf((*SwInterfaceSetUnnumberedReply)(nil)).Elem(),
	"SwInterfaceTagAddDel":             reflect.TypeOf((*SwInterfaceTagAddDel)(nil)).Elem(),
	"SwInterfaceTagAddDelReply":        reflect.TypeOf((*SwInterfaceTagAddDelReply)(nil)).Elem(),
	"VlibCounter":                      reflect.TypeOf((*VlibCounter)(nil)).Elem(),
	"VnetCombinedCounter":              reflect.TypeOf((*VnetCombinedCounter)(nil)).Elem(),
	"VnetInterfaceCombinedCounters":    reflect.TypeOf((*VnetInterfaceCombinedCounters)(nil)).Elem(),
	"VnetInterfaceSimpleCounters":      reflect.TypeOf((*VnetInterfaceSimpleCounters)(nil)).Elem(),
	"VnetPerInterfaceCombinedCounters": reflect.TypeOf((*VnetPerInterfaceCombinedCounters)(nil)).Elem(),
	"VnetPerInterfaceSimpleCounters":   reflect.TypeOf((*VnetPerInterfaceSimpleCounters)(nil)).Elem(),
	"VnetSimpleCounter":                reflect.TypeOf((*VnetSimpleCounter)(nil)).Elem(),
	"WantInterfaceEvents":              reflect.TypeOf((*WantInterfaceEvents)(nil)).Elem(),
	"WantInterfaceEventsReply":         reflect.TypeOf((*WantInterfaceEventsReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewCreateLoopback":                   reflect.ValueOf(NewCreateLoopback),
	"NewCreateLoopbackInstance":           reflect.ValueOf(NewCreateLoopbackInstance),
	"NewCreateLoopbackInstanceReply":      reflect.ValueOf(NewCreateLoopbackInstanceReply),
	"NewCreateLoopbackReply":              reflect.ValueOf(NewCreateLoopbackReply),
	"NewCreateSubif":                      reflect.ValueOf(NewCreateSubif),
	"NewCreateSubifReply":                 reflect.ValueOf(NewCreateSubifReply),
	"NewCreateVlanSubif":                  reflect.ValueOf(NewCreateVlanSubif),
	"NewCreateVlanSubifReply":             reflect.ValueOf(NewCreateVlanSubifReply),
	"NewDeleteLoopback":                   reflect.ValueOf(NewDeleteLoopback),
	"NewDeleteLoopbackReply":              reflect.ValueOf(NewDeleteLoopbackReply),
	"NewDeleteSubif":                      reflect.ValueOf(NewDeleteSubif),
	"NewDeleteSubifReply":                 reflect.ValueOf(NewDeleteSubifReply),
	"NewInterfaceNameRenumber":            reflect.ValueOf(NewInterfaceNameRenumber),
	"NewInterfaceNameRenumberReply":       reflect.ValueOf(NewInterfaceNameRenumberReply),
	"NewSwInterfaceAddDelAddress":         reflect.ValueOf(NewSwInterfaceAddDelAddress),
	"NewSwInterfaceAddDelAddressReply":    reflect.ValueOf(NewSwInterfaceAddDelAddressReply),
	"NewSwInterfaceClearStats":            reflect.ValueOf(NewSwInterfaceClearStats),
	"NewSwInterfaceClearStatsReply":       reflect.ValueOf(NewSwInterfaceClearStatsReply),
	"NewSwInterfaceDetails":               reflect.ValueOf(NewSwInterfaceDetails),
	"NewSwInterfaceDump":                  reflect.ValueOf(NewSwInterfaceDump),
	"NewSwInterfaceEvent":                 reflect.ValueOf(NewSwInterfaceEvent),
	"NewSwInterfaceGetTable":              reflect.ValueOf(NewSwInterfaceGetTable),
	"NewSwInterfaceGetTableReply":         reflect.ValueOf(NewSwInterfaceGetTableReply),
	"NewSwInterfaceSetFlags":              reflect.ValueOf(NewSwInterfaceSetFlags),
	"NewSwInterfaceSetFlagsReply":         reflect.ValueOf(NewSwInterfaceSetFlagsReply),
	"NewSwInterfaceSetMacAddress":         reflect.ValueOf(NewSwInterfaceSetMacAddress),
	"NewSwInterfaceSetMacAddressReply":    reflect.ValueOf(NewSwInterfaceSetMacAddressReply),
	"NewSwInterfaceSetMtu":                reflect.ValueOf(NewSwInterfaceSetMtu),
	"NewSwInterfaceSetMtuReply":           reflect.ValueOf(NewSwInterfaceSetMtuReply),
	"NewSwInterfaceSetRxMode":             reflect.ValueOf(NewSwInterfaceSetRxMode),
	"NewSwInterfaceSetRxModeReply":        reflect.ValueOf(NewSwInterfaceSetRxModeReply),
	"NewSwInterfaceSetTable":              reflect.ValueOf(NewSwInterfaceSetTable),
	"NewSwInterfaceSetTableReply":         reflect.ValueOf(NewSwInterfaceSetTableReply),
	"NewSwInterfaceSetUnnumbered":         reflect.ValueOf(NewSwInterfaceSetUnnumbered),
	"NewSwInterfaceSetUnnumberedReply":    reflect.ValueOf(NewSwInterfaceSetUnnumberedReply),
	"NewSwInterfaceTagAddDel":             reflect.ValueOf(NewSwInterfaceTagAddDel),
	"NewSwInterfaceTagAddDelReply":        reflect.ValueOf(NewSwInterfaceTagAddDelReply),
	"NewVnetInterfaceCombinedCounters":    reflect.ValueOf(NewVnetInterfaceCombinedCounters),
	"NewVnetInterfaceSimpleCounters":      reflect.ValueOf(NewVnetInterfaceSimpleCounters),
	"NewVnetPerInterfaceCombinedCounters": reflect.ValueOf(NewVnetPerInterfaceCombinedCounters),
	"NewVnetPerInterfaceSimpleCounters":   reflect.ValueOf(NewVnetPerInterfaceSimpleCounters),
	"NewWantInterfaceEvents":              reflect.ValueOf(NewWantInterfaceEvents),
	"NewWantInterfaceEventsReply":         reflect.ValueOf(NewWantInterfaceEventsReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{
	"VlAPIVersion": reflect.ValueOf(VlAPIVersion),
}
