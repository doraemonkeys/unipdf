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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_b "encoding/json";_f "github.com/doraemonkeys/unipdf/v3/common";_eg "github.com/doraemonkeys/unipdf/v3/core";_df "github.com/doraemonkeys/unipdf/v3/model";_d "io";_e "os";);type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;
ImageValue *_df .Image `json:"-"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// FieldValues implements model.FieldValueProvider interface.
func (_fd *FieldData )FieldValues ()(map[string ]_eg .PdfObject ,error ){_bd :=make (map[string ]_eg .PdfObject );for _ ,_bdc :=range _fd ._fc {if len (_bdc .Value )> 0{_bd [_bdc .Name ]=_eg .MakeString (_bdc .Value );};};return _bd ,nil ;};

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_eag ,_bcd :=_e .Open (filePath );if _bcd !=nil {return nil ,_bcd ;};defer _eag .Close ();return LoadFromPDF (_eag );};

// SetImage assign model.Image to a specific field identified by fieldName.
func (_dgf *FieldData )SetImage (fieldName string ,img *_df .Image ,opt []string )error {_dfbg :=fieldValue {Name :fieldName ,ImageValue :img ,Options :opt };_dgf ._fc =append (_dgf ._fc ,_dfbg );return nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_fc []fieldValue };

// JSON returns the field data as a string in JSON format.
func (_abf FieldData )JSON ()(string ,error ){_dg ,_dfb :=_b .MarshalIndent (_abf ._fc ,"","\u0020\u0020\u0020\u0020");return string (_dg ),_dfb ;};

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_ec *FieldData )SetImageFromFile (fieldName string ,imagePath string ,opt []string )error {_eee ,_daa :=_e .Open (imagePath );if _daa !=nil {return _daa ;};defer _eee .Close ();_fgd ,_daa :=_df .ImageHandling .Read (_eee );if _daa !=nil {_f .Log .Error ("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_daa );
return _daa ;};return _ec .SetImage (fieldName ,_fgd ,opt );};

// FieldImageValues implements model.FieldImageProvider interface.
func (_da *FieldData )FieldImageValues ()(map[string ]*_df .Image ,error ){_dd :=make (map[string ]*_df .Image );for _ ,_dgd :=range _da ._fc {if _dgd .ImageValue !=nil {_dd [_dgd .Name ]=_dgd .ImageValue ;};};return _dd ,nil ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_fce ,_ba :=_e .Open (filePath );if _ba !=nil {return nil ,_ba ;};defer _fce .Close ();return LoadFromJSON (_fce );};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _d .Reader )(*FieldData ,error ){var _bb FieldData ;_c :=_b .NewDecoder (r ).Decode (&_bb ._fc );if _c !=nil {return nil ,_c ;};return &_bb ,nil ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _d .ReadSeeker )(*FieldData ,error ){_a ,_ae :=_df .NewPdfReader (rs );if _ae !=nil {return nil ,_ae ;};if _a .AcroForm ==nil {return nil ,nil ;};var _bf []fieldValue ;_ega :=_a .AcroForm .AllFields ();for _ ,_fe :=range _ega {var _bc []string ;
_ab :=make (map[string ]struct{});_db ,_egd :=_fe .FullName ();if _egd !=nil {return nil ,_egd ;};if _cb ,_aeb :=_fe .V .(*_eg .PdfObjectString );_aeb {_bf =append (_bf ,fieldValue {Name :_db ,Value :_cb .Decoded ()});continue ;};var _de string ;for _ ,_gd :=range _fe .Annotations {_bab ,_ag :=_eg .GetName (_gd .AS );
if _ag {_de =_bab .String ();};_cc ,_ff :=_eg .GetDict (_gd .AP );if !_ff {continue ;};_aa ,_ :=_eg .GetDict (_cc .Get ("\u004e"));for _ ,_ea :=range _aa .Keys (){_aed :=_ea .String ();if _ ,_fg :=_ab [_aed ];!_fg {_bc =append (_bc ,_aed );_ab [_aed ]=struct{}{};
};};_ad ,_ :=_eg .GetDict (_cc .Get ("\u0044"));for _ ,_fa :=range _ad .Keys (){_adb :=_fa .String ();if _ ,_cff :=_ab [_adb ];!_cff {_bc =append (_bc ,_adb );_ab [_adb ]=struct{}{};};};};_gb :=fieldValue {Name :_db ,Value :_de ,Options :_bc };_bf =append (_bf ,_gb );
};_aad :=FieldData {_fc :_bf };return &_aad ,nil ;};