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

package uuid ;import (_gg "crypto/rand";_a "encoding/hex";_e "io";);var _ab =_gg .Reader ;func _eb (_ea []byte ,_fb UUID ){_a .Encode (_ea ,_fb [:4]);_ea [8]='-';_a .Encode (_ea [9:13],_fb [4:6]);_ea [13]='-';_a .Encode (_ea [14:18],_fb [6:8]);_ea [18]='-';
_a .Encode (_ea [19:23],_fb [8:10]);_ea [23]='-';_a .Encode (_ea [24:],_fb [10:]);};func (_d UUID )String ()string {var _cd [36]byte ;_eb (_cd [:],_d );return string (_cd [:])};func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_ce :=_e .ReadFull (_ab ,uuid [:]);
if _ce !=nil {return _cdc ,_ce ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};var _cdc UUID ;var Nil =_cdc ;type UUID [16]byte ;func MustUUID ()UUID {uuid ,_cb :=NewUUID ();if _cb !=nil {panic (_cb );};return uuid ;};
