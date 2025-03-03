//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package pdfaid ;import (_gd "fmt";_e "github.com/trimmer-io/go-xmp/xmp";_b "github.com/doraemonkeys/unipdf/v3/model/xmputil/pdfaextension";);

// Namespaces implements xmp.Model interface.
func (_gdb *Model )Namespaces ()_e .NamespaceList {return _e .NamespaceList {Namespace }};func init (){_e .Register (Namespace ,_e .XmpMetadata );_b .RegisterSchema (Namespace ,&Schema )};

// SyncModel implements xmp.Model interface.
func (_ec *Model )SyncModel (d *_e .Document )error {return nil };

// SyncToXMP implements xmp.Model interface.
func (_fc *Model )SyncToXMP (d *_e .Document )error {return nil };

// Can implements xmp.Model interface.
func (_ba *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };

// GetTag implements xmp.Model interface.
func (_bac *Model )GetTag (tag string )(string ,error ){_cfa ,_de :=_e .GetNativeField (_bac ,tag );if _de !=nil {return "",_gd .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_de );};return _cfa ,nil ;};var _ _e .Model =(*Model )(nil );


// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel (d *_e .Document )(*Model ,error ){_d ,_f :=d .MakeModel (Namespace );if _f !=nil {return nil ,_f ;};return _d .(*Model ),nil ;};

// SetTag implements xmp.Model interface.
func (_df *Model )SetTag (tag ,value string )error {if _eac :=_e .SetNativeField (_df ,tag ,value );_eac !=nil {return _gd .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_eac );};return nil ;};

// SyncFromXMP implements xmp.Model interface.
func (_cf *Model )SyncFromXMP (d *_e .Document )error {return nil };var Namespace =_e .NewNamespace ("\u0070\u0064\u0066\u0061\u0069\u0064","\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f",NewModel );
var Schema =_b .Schema {NamespaceURI :Namespace .URI ,Prefix :Namespace .Name ,Schema :"\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :[]_b .Property {{Category :_b .PropertyCategoryInternal ,Description :"\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"\u0070\u0061\u0072\u0074",ValueType :_b .ValueTypeNameInteger },{Category :_b .PropertyCategoryInternal ,Description :"A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064",Name :"\u0061\u006d\u0064",ValueType :_b .ValueTypeNameText },{Category :_b .PropertyCategoryInternal ,Description :"C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065",ValueType :_b .ValueTypeNameText }},ValueType :nil };


// NewModel creates a new model.
func NewModel (name string )_e .Model {return &Model {}};

// CanTag implements xmp.Model interface.
func (_a *Model )CanTag (tag string )bool {_ ,_ag :=_e .GetNativeField (_a ,tag );return _ag ==nil };

// Model is the XMP model for the PdfA metadata.
type Model struct{Part int `xmp:"pdfaid:part"`;Conformance string `xmp:"pdfaid:conformance"`;};