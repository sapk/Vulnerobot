//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/2009/01/xml.xsd
package go_Xml

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type TxsdLang xsdt.String

//	Since TxsdLang is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdLang) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdLang is just a simple String type, this merely returns the current string value.
func (me TxsdLang) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdLang's alias type xsdt.String.
func (me TxsdLang) ToXsdtString() xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Lang struct {
	Lang TxsdLang `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

type TxsdSpace xsdt.NCName

//	This convenience method just performs a simple type conversion to TxsdSpace's alias type xsdt.NCName.
func (me TxsdSpace) ToXsdtNCName() xsdt.NCName { return xsdt.NCName(me) }

//	Returns true if the value of this enumerated TxsdSpace is "default".
func (me TxsdSpace) IsDefault() bool { return me.String() == "default" }

//	Returns true if the value of this enumerated TxsdSpace is "preserve".
func (me TxsdSpace) IsPreserve() bool { return me.String() == "preserve" }

//	Since TxsdSpace is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdSpace) Set(s string) { (*xsdt.NCName)(me).Set(s) }

//	Since TxsdSpace is just a simple String type, this merely returns the current string value.
func (me TxsdSpace) String() string { return xsdt.NCName(me).String() }

type XsdGoPkgHasAttr_Space struct {
	Space TxsdSpace `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
}

type XsdGoPkgHasAttr_Base struct {
	Base xsdt.AnyURI `xml:"http://www.w3.org/XML/1998/namespace base,attr"`
}

type XsdGoPkgHasAttr_Id struct {
	Id xsdt.Id `xml:"http://www.w3.org/XML/1998/namespace id,attr"`
}

type XsdGoPkgHasAtts_SpecialAttrs struct {
	XsdGoPkgHasAttr_Space

	XsdGoPkgHasAttr_Id

	XsdGoPkgHasAttr_Base

	XsdGoPkgHasAttr_Lang
}

type TxsdLangTxsdLangUnion xsdt.String

//	TxsdLang is an XSD union-type of several types. This is a simple type conversion to TxsdLangTxsdLangUnion, but keep in mind the actual value may or may not be a valid TxsdLangTxsdLangUnion value.
func (me TxsdLang) ToTxsdLangTxsdLangUnion() TxsdLangTxsdLangUnion { return TxsdLangTxsdLangUnion(me) }

//	Since TxsdLangTxsdLangUnion is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdLangTxsdLangUnion) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdLangTxsdLangUnion is just a simple String type, this merely returns the current string value.
func (me TxsdLangTxsdLangUnion) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdLangTxsdLangUnion's alias type xsdt.String.
func (me TxsdLangTxsdLangUnion) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdLangTxsdLangUnion is "".
func (me TxsdLangTxsdLangUnion) Is() bool { return me.String() == "" }

//	TxsdLang is an XSD union-type of several types. This is a simple type conversion to XsdtLanguage, but keep in mind the actual value may or may not be a valid XsdtLanguage value.
func (me TxsdLang) ToXsdtLanguage() xsdt.Language { return xsdt.Language(me) }

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasCdata func(*XsdGoPkgHasCdata, bool) error
}
