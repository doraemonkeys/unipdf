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

package endian ;import (_c "encoding/binary";_b "unsafe";);func IsBig ()bool {return _bc };func init (){const _f =int (_b .Sizeof (0));_fc :=1;_bca :=(*[_f ]byte )(_b .Pointer (&_fc ));if _bca [0]==0{_bc =true ;ByteOrder =_c .BigEndian ;}else {ByteOrder =_c .LittleEndian ;
};};var (ByteOrder _c .ByteOrder ;_bc bool ;);func IsLittle ()bool {return !_bc };