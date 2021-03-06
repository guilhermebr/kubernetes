/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1alpha1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg4_resource "k8s.io/client-go/pkg/api/resource"
	pkg1_unversioned "k8s.io/client-go/pkg/api/unversioned"
	pkg2_v1 "k8s.io/client-go/pkg/api/v1"
	pkg3_types "k8s.io/client-go/pkg/types"
	pkg5_intstr "k8s.io/client-go/pkg/util/intstr"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg4_resource.Quantity
		var v1 pkg1_unversioned.TypeMeta
		var v2 pkg2_v1.ObjectMeta
		var v3 pkg3_types.UID
		var v4 pkg5_intstr.IntOrString
		var v5 time.Time
		_, _, _, _, _, _ = v0, v1, v2, v3, v4, v5
	}
}

func (x *PetSet) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[3] = true
			yyq2[4] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} else {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.ObjectMeta
					yy11.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					yy13 := &x.Spec
					yy13.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("spec"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy14 := &x.Spec
					yy14.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yy16 := &x.Status
					yy16.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy17 := &x.Status
					yy17.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PetSet) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym18 := z.DecBinary()
	_ = yym18
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct19 := r.ContainerType()
		if yyct19 == codecSelferValueTypeMap1234 {
			yyl19 := r.ReadMapStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl19, d)
			}
		} else if yyct19 == codecSelferValueTypeArray1234 {
			yyl19 := r.ReadArrayStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl19, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PetSet) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys20Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys20Slc
	var yyhl20 bool = l >= 0
	for yyj20 := 0; ; yyj20++ {
		if yyhl20 {
			if yyj20 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys20Slc = r.DecodeBytes(yys20Slc, true, true)
		yys20 := string(yys20Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys20 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv23 := &x.ObjectMeta
				yyv23.CodecDecodeSelf(d)
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = PetSetSpec{}
			} else {
				yyv24 := &x.Spec
				yyv24.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = PetSetStatus{}
			} else {
				yyv25 := &x.Status
				yyv25.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys20)
		} // end switch yys20
	} // end for yyj20
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PetSet) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj26 int
	var yyb26 bool
	var yyhl26 bool = l >= 0
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv29 := &x.ObjectMeta
		yyv29.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = PetSetSpec{}
	} else {
		yyv30 := &x.Spec
		yyv30.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = PetSetStatus{}
	} else {
		yyv31 := &x.Status
		yyv31.CodecDecodeSelf(d)
	}
	for {
		yyj26++
		if yyhl26 {
			yyb26 = yyj26 > l
		} else {
			yyb26 = r.CheckBreak()
		}
		if yyb26 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj26-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PetSetSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym32 := z.EncBinary()
		_ = yym32
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep33 := !z.EncBinary()
			yy2arr33 := z.EncBasicHandle().StructToArray
			var yyq33 [5]bool
			_, _, _ = yysep33, yyq33, yy2arr33
			const yyr33 bool = false
			yyq33[0] = x.Replicas != nil
			yyq33[1] = x.Selector != nil
			yyq33[3] = len(x.VolumeClaimTemplates) != 0
			var yynn33 int
			if yyr33 || yy2arr33 {
				r.EncodeArrayStart(5)
			} else {
				yynn33 = 2
				for _, b := range yyq33 {
					if b {
						yynn33++
					}
				}
				r.EncodeMapStart(yynn33)
				yynn33 = 0
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[0] {
					if x.Replicas == nil {
						r.EncodeNil()
					} else {
						yy35 := *x.Replicas
						yym36 := z.EncBinary()
						_ = yym36
						if false {
						} else {
							r.EncodeInt(int64(yy35))
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq33[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("replicas"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Replicas == nil {
						r.EncodeNil()
					} else {
						yy37 := *x.Replicas
						yym38 := z.EncBinary()
						_ = yym38
						if false {
						} else {
							r.EncodeInt(int64(yy37))
						}
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[1] {
					if x.Selector == nil {
						r.EncodeNil()
					} else {
						yym40 := z.EncBinary()
						_ = yym40
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Selector) {
						} else {
							z.EncFallback(x.Selector)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq33[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("selector"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Selector == nil {
						r.EncodeNil()
					} else {
						yym41 := z.EncBinary()
						_ = yym41
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Selector) {
						} else {
							z.EncFallback(x.Selector)
						}
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy43 := &x.Template
				yy43.CodecEncodeSelf(e)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("template"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy44 := &x.Template
				yy44.CodecEncodeSelf(e)
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[3] {
					if x.VolumeClaimTemplates == nil {
						r.EncodeNil()
					} else {
						yym46 := z.EncBinary()
						_ = yym46
						if false {
						} else {
							h.encSlicev1_PersistentVolumeClaim(([]pkg2_v1.PersistentVolumeClaim)(x.VolumeClaimTemplates), e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq33[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("volumeClaimTemplates"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.VolumeClaimTemplates == nil {
						r.EncodeNil()
					} else {
						yym47 := z.EncBinary()
						_ = yym47
						if false {
						} else {
							h.encSlicev1_PersistentVolumeClaim(([]pkg2_v1.PersistentVolumeClaim)(x.VolumeClaimTemplates), e)
						}
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym49 := z.EncBinary()
				_ = yym49
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.ServiceName))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("serviceName"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym50 := z.EncBinary()
				_ = yym50
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.ServiceName))
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PetSetSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym51 := z.DecBinary()
	_ = yym51
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct52 := r.ContainerType()
		if yyct52 == codecSelferValueTypeMap1234 {
			yyl52 := r.ReadMapStart()
			if yyl52 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl52, d)
			}
		} else if yyct52 == codecSelferValueTypeArray1234 {
			yyl52 := r.ReadArrayStart()
			if yyl52 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl52, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PetSetSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys53Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys53Slc
	var yyhl53 bool = l >= 0
	for yyj53 := 0; ; yyj53++ {
		if yyhl53 {
			if yyj53 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys53Slc = r.DecodeBytes(yys53Slc, true, true)
		yys53 := string(yys53Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys53 {
		case "replicas":
			if r.TryDecodeAsNil() {
				if x.Replicas != nil {
					x.Replicas = nil
				}
			} else {
				if x.Replicas == nil {
					x.Replicas = new(int32)
				}
				yym55 := z.DecBinary()
				_ = yym55
				if false {
				} else {
					*((*int32)(x.Replicas)) = int32(r.DecodeInt(32))
				}
			}
		case "selector":
			if r.TryDecodeAsNil() {
				if x.Selector != nil {
					x.Selector = nil
				}
			} else {
				if x.Selector == nil {
					x.Selector = new(pkg1_unversioned.LabelSelector)
				}
				yym57 := z.DecBinary()
				_ = yym57
				if false {
				} else if z.HasExtensions() && z.DecExt(x.Selector) {
				} else {
					z.DecFallback(x.Selector, false)
				}
			}
		case "template":
			if r.TryDecodeAsNil() {
				x.Template = pkg2_v1.PodTemplateSpec{}
			} else {
				yyv58 := &x.Template
				yyv58.CodecDecodeSelf(d)
			}
		case "volumeClaimTemplates":
			if r.TryDecodeAsNil() {
				x.VolumeClaimTemplates = nil
			} else {
				yyv59 := &x.VolumeClaimTemplates
				yym60 := z.DecBinary()
				_ = yym60
				if false {
				} else {
					h.decSlicev1_PersistentVolumeClaim((*[]pkg2_v1.PersistentVolumeClaim)(yyv59), d)
				}
			}
		case "serviceName":
			if r.TryDecodeAsNil() {
				x.ServiceName = ""
			} else {
				x.ServiceName = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys53)
		} // end switch yys53
	} // end for yyj53
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PetSetSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj62 int
	var yyb62 bool
	var yyhl62 bool = l >= 0
	yyj62++
	if yyhl62 {
		yyb62 = yyj62 > l
	} else {
		yyb62 = r.CheckBreak()
	}
	if yyb62 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.Replicas != nil {
			x.Replicas = nil
		}
	} else {
		if x.Replicas == nil {
			x.Replicas = new(int32)
		}
		yym64 := z.DecBinary()
		_ = yym64
		if false {
		} else {
			*((*int32)(x.Replicas)) = int32(r.DecodeInt(32))
		}
	}
	yyj62++
	if yyhl62 {
		yyb62 = yyj62 > l
	} else {
		yyb62 = r.CheckBreak()
	}
	if yyb62 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.Selector != nil {
			x.Selector = nil
		}
	} else {
		if x.Selector == nil {
			x.Selector = new(pkg1_unversioned.LabelSelector)
		}
		yym66 := z.DecBinary()
		_ = yym66
		if false {
		} else if z.HasExtensions() && z.DecExt(x.Selector) {
		} else {
			z.DecFallback(x.Selector, false)
		}
	}
	yyj62++
	if yyhl62 {
		yyb62 = yyj62 > l
	} else {
		yyb62 = r.CheckBreak()
	}
	if yyb62 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Template = pkg2_v1.PodTemplateSpec{}
	} else {
		yyv67 := &x.Template
		yyv67.CodecDecodeSelf(d)
	}
	yyj62++
	if yyhl62 {
		yyb62 = yyj62 > l
	} else {
		yyb62 = r.CheckBreak()
	}
	if yyb62 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.VolumeClaimTemplates = nil
	} else {
		yyv68 := &x.VolumeClaimTemplates
		yym69 := z.DecBinary()
		_ = yym69
		if false {
		} else {
			h.decSlicev1_PersistentVolumeClaim((*[]pkg2_v1.PersistentVolumeClaim)(yyv68), d)
		}
	}
	yyj62++
	if yyhl62 {
		yyb62 = yyj62 > l
	} else {
		yyb62 = r.CheckBreak()
	}
	if yyb62 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ServiceName = ""
	} else {
		x.ServiceName = string(r.DecodeString())
	}
	for {
		yyj62++
		if yyhl62 {
			yyb62 = yyj62 > l
		} else {
			yyb62 = r.CheckBreak()
		}
		if yyb62 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj62-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PetSetStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym71 := z.EncBinary()
		_ = yym71
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep72 := !z.EncBinary()
			yy2arr72 := z.EncBasicHandle().StructToArray
			var yyq72 [2]bool
			_, _, _ = yysep72, yyq72, yy2arr72
			const yyr72 bool = false
			yyq72[0] = x.ObservedGeneration != nil
			var yynn72 int
			if yyr72 || yy2arr72 {
				r.EncodeArrayStart(2)
			} else {
				yynn72 = 1
				for _, b := range yyq72 {
					if b {
						yynn72++
					}
				}
				r.EncodeMapStart(yynn72)
				yynn72 = 0
			}
			if yyr72 || yy2arr72 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq72[0] {
					if x.ObservedGeneration == nil {
						r.EncodeNil()
					} else {
						yy74 := *x.ObservedGeneration
						yym75 := z.EncBinary()
						_ = yym75
						if false {
						} else {
							r.EncodeInt(int64(yy74))
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq72[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("observedGeneration"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.ObservedGeneration == nil {
						r.EncodeNil()
					} else {
						yy76 := *x.ObservedGeneration
						yym77 := z.EncBinary()
						_ = yym77
						if false {
						} else {
							r.EncodeInt(int64(yy76))
						}
					}
				}
			}
			if yyr72 || yy2arr72 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym79 := z.EncBinary()
				_ = yym79
				if false {
				} else {
					r.EncodeInt(int64(x.Replicas))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("replicas"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym80 := z.EncBinary()
				_ = yym80
				if false {
				} else {
					r.EncodeInt(int64(x.Replicas))
				}
			}
			if yyr72 || yy2arr72 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PetSetStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym81 := z.DecBinary()
	_ = yym81
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct82 := r.ContainerType()
		if yyct82 == codecSelferValueTypeMap1234 {
			yyl82 := r.ReadMapStart()
			if yyl82 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl82, d)
			}
		} else if yyct82 == codecSelferValueTypeArray1234 {
			yyl82 := r.ReadArrayStart()
			if yyl82 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl82, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PetSetStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys83Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys83Slc
	var yyhl83 bool = l >= 0
	for yyj83 := 0; ; yyj83++ {
		if yyhl83 {
			if yyj83 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys83Slc = r.DecodeBytes(yys83Slc, true, true)
		yys83 := string(yys83Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys83 {
		case "observedGeneration":
			if r.TryDecodeAsNil() {
				if x.ObservedGeneration != nil {
					x.ObservedGeneration = nil
				}
			} else {
				if x.ObservedGeneration == nil {
					x.ObservedGeneration = new(int64)
				}
				yym85 := z.DecBinary()
				_ = yym85
				if false {
				} else {
					*((*int64)(x.ObservedGeneration)) = int64(r.DecodeInt(64))
				}
			}
		case "replicas":
			if r.TryDecodeAsNil() {
				x.Replicas = 0
			} else {
				x.Replicas = int32(r.DecodeInt(32))
			}
		default:
			z.DecStructFieldNotFound(-1, yys83)
		} // end switch yys83
	} // end for yyj83
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PetSetStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj87 int
	var yyb87 bool
	var yyhl87 bool = l >= 0
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.ObservedGeneration != nil {
			x.ObservedGeneration = nil
		}
	} else {
		if x.ObservedGeneration == nil {
			x.ObservedGeneration = new(int64)
		}
		yym89 := z.DecBinary()
		_ = yym89
		if false {
		} else {
			*((*int64)(x.ObservedGeneration)) = int64(r.DecodeInt(64))
		}
	}
	yyj87++
	if yyhl87 {
		yyb87 = yyj87 > l
	} else {
		yyb87 = r.CheckBreak()
	}
	if yyb87 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Replicas = 0
	} else {
		x.Replicas = int32(r.DecodeInt(32))
	}
	for {
		yyj87++
		if yyhl87 {
			yyb87 = yyj87 > l
		} else {
			yyb87 = r.CheckBreak()
		}
		if yyb87 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj87-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PetSetList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym91 := z.EncBinary()
		_ = yym91
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep92 := !z.EncBinary()
			yy2arr92 := z.EncBasicHandle().StructToArray
			var yyq92 [4]bool
			_, _, _ = yysep92, yyq92, yy2arr92
			const yyr92 bool = false
			yyq92[0] = x.Kind != ""
			yyq92[1] = x.APIVersion != ""
			yyq92[2] = true
			var yynn92 int
			if yyr92 || yy2arr92 {
				r.EncodeArrayStart(4)
			} else {
				yynn92 = 1
				for _, b := range yyq92 {
					if b {
						yynn92++
					}
				}
				r.EncodeMapStart(yynn92)
				yynn92 = 0
			}
			if yyr92 || yy2arr92 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq92[0] {
					yym94 := z.EncBinary()
					_ = yym94
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq92[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym95 := z.EncBinary()
					_ = yym95
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr92 || yy2arr92 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq92[1] {
					yym97 := z.EncBinary()
					_ = yym97
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq92[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym98 := z.EncBinary()
					_ = yym98
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr92 || yy2arr92 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq92[2] {
					yy100 := &x.ListMeta
					yym101 := z.EncBinary()
					_ = yym101
					if false {
					} else if z.HasExtensions() && z.EncExt(yy100) {
					} else {
						z.EncFallback(yy100)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq92[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy102 := &x.ListMeta
					yym103 := z.EncBinary()
					_ = yym103
					if false {
					} else if z.HasExtensions() && z.EncExt(yy102) {
					} else {
						z.EncFallback(yy102)
					}
				}
			}
			if yyr92 || yy2arr92 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym105 := z.EncBinary()
					_ = yym105
					if false {
					} else {
						h.encSlicePetSet(([]PetSet)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym106 := z.EncBinary()
					_ = yym106
					if false {
					} else {
						h.encSlicePetSet(([]PetSet)(x.Items), e)
					}
				}
			}
			if yyr92 || yy2arr92 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PetSetList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym107 := z.DecBinary()
	_ = yym107
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct108 := r.ContainerType()
		if yyct108 == codecSelferValueTypeMap1234 {
			yyl108 := r.ReadMapStart()
			if yyl108 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl108, d)
			}
		} else if yyct108 == codecSelferValueTypeArray1234 {
			yyl108 := r.ReadArrayStart()
			if yyl108 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl108, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PetSetList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys109Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys109Slc
	var yyhl109 bool = l >= 0
	for yyj109 := 0; ; yyj109++ {
		if yyhl109 {
			if yyj109 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys109Slc = r.DecodeBytes(yys109Slc, true, true)
		yys109 := string(yys109Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys109 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_unversioned.ListMeta{}
			} else {
				yyv112 := &x.ListMeta
				yym113 := z.DecBinary()
				_ = yym113
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv112) {
				} else {
					z.DecFallback(yyv112, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv114 := &x.Items
				yym115 := z.DecBinary()
				_ = yym115
				if false {
				} else {
					h.decSlicePetSet((*[]PetSet)(yyv114), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys109)
		} // end switch yys109
	} // end for yyj109
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PetSetList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj116 int
	var yyb116 bool
	var yyhl116 bool = l >= 0
	yyj116++
	if yyhl116 {
		yyb116 = yyj116 > l
	} else {
		yyb116 = r.CheckBreak()
	}
	if yyb116 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj116++
	if yyhl116 {
		yyb116 = yyj116 > l
	} else {
		yyb116 = r.CheckBreak()
	}
	if yyb116 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj116++
	if yyhl116 {
		yyb116 = yyj116 > l
	} else {
		yyb116 = r.CheckBreak()
	}
	if yyb116 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_unversioned.ListMeta{}
	} else {
		yyv119 := &x.ListMeta
		yym120 := z.DecBinary()
		_ = yym120
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv119) {
		} else {
			z.DecFallback(yyv119, false)
		}
	}
	yyj116++
	if yyhl116 {
		yyb116 = yyj116 > l
	} else {
		yyb116 = r.CheckBreak()
	}
	if yyb116 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv121 := &x.Items
		yym122 := z.DecBinary()
		_ = yym122
		if false {
		} else {
			h.decSlicePetSet((*[]PetSet)(yyv121), d)
		}
	}
	for {
		yyj116++
		if yyhl116 {
			yyb116 = yyj116 > l
		} else {
			yyb116 = r.CheckBreak()
		}
		if yyb116 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj116-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSlicev1_PersistentVolumeClaim(v []pkg2_v1.PersistentVolumeClaim, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv123 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy124 := &yyv123
		yy124.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSlicev1_PersistentVolumeClaim(v *[]pkg2_v1.PersistentVolumeClaim, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv125 := *v
	yyh125, yyl125 := z.DecSliceHelperStart()
	var yyc125 bool
	if yyl125 == 0 {
		if yyv125 == nil {
			yyv125 = []pkg2_v1.PersistentVolumeClaim{}
			yyc125 = true
		} else if len(yyv125) != 0 {
			yyv125 = yyv125[:0]
			yyc125 = true
		}
	} else if yyl125 > 0 {
		var yyrr125, yyrl125 int
		var yyrt125 bool
		if yyl125 > cap(yyv125) {

			yyrg125 := len(yyv125) > 0
			yyv2125 := yyv125
			yyrl125, yyrt125 = z.DecInferLen(yyl125, z.DecBasicHandle().MaxInitLen, 368)
			if yyrt125 {
				if yyrl125 <= cap(yyv125) {
					yyv125 = yyv125[:yyrl125]
				} else {
					yyv125 = make([]pkg2_v1.PersistentVolumeClaim, yyrl125)
				}
			} else {
				yyv125 = make([]pkg2_v1.PersistentVolumeClaim, yyrl125)
			}
			yyc125 = true
			yyrr125 = len(yyv125)
			if yyrg125 {
				copy(yyv125, yyv2125)
			}
		} else if yyl125 != len(yyv125) {
			yyv125 = yyv125[:yyl125]
			yyc125 = true
		}
		yyj125 := 0
		for ; yyj125 < yyrr125; yyj125++ {
			yyh125.ElemContainerState(yyj125)
			if r.TryDecodeAsNil() {
				yyv125[yyj125] = pkg2_v1.PersistentVolumeClaim{}
			} else {
				yyv126 := &yyv125[yyj125]
				yyv126.CodecDecodeSelf(d)
			}

		}
		if yyrt125 {
			for ; yyj125 < yyl125; yyj125++ {
				yyv125 = append(yyv125, pkg2_v1.PersistentVolumeClaim{})
				yyh125.ElemContainerState(yyj125)
				if r.TryDecodeAsNil() {
					yyv125[yyj125] = pkg2_v1.PersistentVolumeClaim{}
				} else {
					yyv127 := &yyv125[yyj125]
					yyv127.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj125 := 0
		for ; !r.CheckBreak(); yyj125++ {

			if yyj125 >= len(yyv125) {
				yyv125 = append(yyv125, pkg2_v1.PersistentVolumeClaim{}) // var yyz125 pkg2_v1.PersistentVolumeClaim
				yyc125 = true
			}
			yyh125.ElemContainerState(yyj125)
			if yyj125 < len(yyv125) {
				if r.TryDecodeAsNil() {
					yyv125[yyj125] = pkg2_v1.PersistentVolumeClaim{}
				} else {
					yyv128 := &yyv125[yyj125]
					yyv128.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj125 < len(yyv125) {
			yyv125 = yyv125[:yyj125]
			yyc125 = true
		} else if yyj125 == 0 && yyv125 == nil {
			yyv125 = []pkg2_v1.PersistentVolumeClaim{}
			yyc125 = true
		}
	}
	yyh125.End()
	if yyc125 {
		*v = yyv125
	}
}

func (x codecSelfer1234) encSlicePetSet(v []PetSet, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv129 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy130 := &yyv129
		yy130.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSlicePetSet(v *[]PetSet, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv131 := *v
	yyh131, yyl131 := z.DecSliceHelperStart()
	var yyc131 bool
	if yyl131 == 0 {
		if yyv131 == nil {
			yyv131 = []PetSet{}
			yyc131 = true
		} else if len(yyv131) != 0 {
			yyv131 = yyv131[:0]
			yyc131 = true
		}
	} else if yyl131 > 0 {
		var yyrr131, yyrl131 int
		var yyrt131 bool
		if yyl131 > cap(yyv131) {

			yyrg131 := len(yyv131) > 0
			yyv2131 := yyv131
			yyrl131, yyrt131 = z.DecInferLen(yyl131, z.DecBasicHandle().MaxInitLen, 800)
			if yyrt131 {
				if yyrl131 <= cap(yyv131) {
					yyv131 = yyv131[:yyrl131]
				} else {
					yyv131 = make([]PetSet, yyrl131)
				}
			} else {
				yyv131 = make([]PetSet, yyrl131)
			}
			yyc131 = true
			yyrr131 = len(yyv131)
			if yyrg131 {
				copy(yyv131, yyv2131)
			}
		} else if yyl131 != len(yyv131) {
			yyv131 = yyv131[:yyl131]
			yyc131 = true
		}
		yyj131 := 0
		for ; yyj131 < yyrr131; yyj131++ {
			yyh131.ElemContainerState(yyj131)
			if r.TryDecodeAsNil() {
				yyv131[yyj131] = PetSet{}
			} else {
				yyv132 := &yyv131[yyj131]
				yyv132.CodecDecodeSelf(d)
			}

		}
		if yyrt131 {
			for ; yyj131 < yyl131; yyj131++ {
				yyv131 = append(yyv131, PetSet{})
				yyh131.ElemContainerState(yyj131)
				if r.TryDecodeAsNil() {
					yyv131[yyj131] = PetSet{}
				} else {
					yyv133 := &yyv131[yyj131]
					yyv133.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj131 := 0
		for ; !r.CheckBreak(); yyj131++ {

			if yyj131 >= len(yyv131) {
				yyv131 = append(yyv131, PetSet{}) // var yyz131 PetSet
				yyc131 = true
			}
			yyh131.ElemContainerState(yyj131)
			if yyj131 < len(yyv131) {
				if r.TryDecodeAsNil() {
					yyv131[yyj131] = PetSet{}
				} else {
					yyv134 := &yyv131[yyj131]
					yyv134.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj131 < len(yyv131) {
			yyv131 = yyv131[:yyj131]
			yyc131 = true
		} else if yyj131 == 0 && yyv131 == nil {
			yyv131 = []PetSet{}
			yyc131 = true
		}
	}
	yyh131.End()
	if yyc131 {
		*v = yyv131
	}
}
