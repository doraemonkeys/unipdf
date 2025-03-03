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

package textshaping ;import (_c "github.com/unidoc/garabic";_ca "golang.org/x/text/unicode/bidi";_d "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_e :=_ca .Paragraph {};_e .SetString (text );_cg ,_g :=_e .Order ();if _g !=nil {return "",_g ;};for _eg :=0;_eg < _cg .NumRuns ();_eg ++{_f :=_cg .Run (_eg );_da :=_f .String ();if _f .Direction ()==_ca .RightToLeft {var (_fg =_c .Shape (_da );
_egb =[]rune (_fg );_aa =make ([]rune ,len (_egb )););_fc :=0;for _ee :=len (_egb )-1;_ee >=0;_ee --{_aa [_fc ]=_egb [_ee ];_fc ++;};_da =string (_aa );text =_d .Replace (text ,_d .TrimSpace (_f .String ()),_da ,1);};};return text ,nil ;};