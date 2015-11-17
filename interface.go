package libxml2

/*
#cgo pkg-config: libxml-2.0
#include "libxml/tree.h"
#include "libxml/xpath.h"
#include <libxml/xpathInternals.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type XmlNodeType int

const (
	ElementNode XmlNodeType = iota + 1
	AttributeNode
	TextNode
	CDataSectionNode
	EntityRefNode
	EntityNode
	PiNode
	CommentNode
	DocumentNode
	DocumentTypeNode
	DocumentFragNode
	NotationNode
	HTMLDocumentNode
	DTDNode
	ElementDecl
	AttributeDecl
	EntityDecl
	NamespaceDecl
	XIncludeStart
	XIncludeEnd
	DocbDocumentNode
)

var (
	ErrNodeNotFound    = errors.New("node not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInvalidNodeName = errors.New("invalid node name")
)

// Node defines the basic DOM interface
type Node interface {
	// pointer() returns the underlying C pointer. Only we are allowed to
	// slice it, dice it, do whatever the heck with it.
	pointer() unsafe.Pointer

	AddChild(Node)
	AppendChild(Node) error
	ChildNodes() NodeList
	OwnerDocument() *Document
	FindNodes(string) (NodeList, error)
	FirstChild() Node
	HasChildNodes() bool
	IsSameNode(Node) bool
	LastChild() Node
	// Literal is almost the same as String(), except for things like Element
	// and Attribute nodes. String() will return the XML stringification of
	// these, but Literal() will return the "value" associated with them.
	Literal() string
	NextSibling() Node
	NodeName() string
	NodeType() XmlNodeType
	NodeValue() string
	ParetNode() Node
	PreviousSibling() Node
	SetNodeName(string)
	SetNodeValue(string)
	String() string
	TextContent() string
	ToString(int, bool) string
	ToStringC14N(bool) (string, error)
	Walk(func(Node) error)
}

type NodeList []Node

type xmlNode struct {
	ptr *C.xmlNode
}

type XmlNode struct {
	*xmlNode
}

type Attribute struct {
	*XmlNode
}

type CDataSection struct {
	*XmlNode
}

type Comment struct {
	*XmlNode
}

type Element struct {
	*XmlNode
}

type Document struct {
	ptr  *C.xmlDoc
	root *C.xmlNode
}

type Text struct {
	*XmlNode
}

type XPathObjectType int

const (
	XPathUndefined XPathObjectType = iota
	XPathNodeSet
	XPathBoolean
	XPathNumber
	XPathString
	XPathPoint
	XPathRange
	XPathLocationSet
	XPathUSers
	XPathXsltTree
)

type XPathObject struct {
	ptr *C.xmlXPathObject
	// This flag controls if the StringValue should use the *contents* (literal value)
	// of the nodeset instead of stringifying the node
	ForceLiteral bool
}

type XPathContext struct {
	ptr *C.xmlXPathContext
}

// XPathExpression is a compiled XPath.
type XPathExpression struct {
	ptr *C.xmlXPathCompExpr
	// This exists mainly for debugging purposes
	expr string
}

// ParseOption represents each of the parser option bit
type ParseOption int

// ParseOption represents the parser option bit set
type ParseOptions int

const (
	XmlParseRecover    ParseOption = 1 << iota /* recover on errors */
	XmlParseNoEnt                              /* substitute entities */
	XmlParseDTDLoad                            /* load the external subset */
	XmlParseDTDAttr                            /* default DTD attributes */
	XmlParseDTDValid                           /* validate with the DTD */
	XmlParseNoError                            /* suppress error reports */
	XmlParseNoWarning                          /* suppress warning reports */
	XmlParsePedantic                           /* pedantic error reporting */
	XmlParseNoBlanks                           /* remove blank nodes */
	XmlParseSAX1                               /* use the SAX1 interface internally */
	XmlParseXInclude                           /* Implement XInclude substitition  */
	XmlParseNoNet                              /* Forbid network access */
	XmlParseNoDict                             /* Do not reuse the context dictionnary */
	XmlParseNsclean                            /* remove redundant namespaces declarations */
	XmlParseNoCDATA                            /* merge CDATA as text nodes */
	XmlParseNoXIncNode                         /* do not generate XINCLUDE START/END nodes */
	XmlParseCompact                            /* compact small text nodes; no modification of the tree allowed afterwards (will possibly crash if you try to modify the tree) */
	XmlParseOld10                              /* parse using XML-1.0 before update 5 */
	XmlParseNoBaseFix                          /* do not fixup XINCLUDE xml:base uris */
	XmlParseHuge                               /* relax any hardcoded limit from the parser */
	XmlParseOldSAX                             /* parse using SAX2 interface before 2.7.0 */
	XmlParseIgnoreEnc                          /* ignore internal document encoding hint */
	XmlParseBigLines                           /* Store big lines numbers in text PSVI field */
	XmlParseMax
)

type Parser struct {
	Options ParseOptions
}