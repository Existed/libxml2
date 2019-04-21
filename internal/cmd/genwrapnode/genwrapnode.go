package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	var buf bytes.Buffer

	buf.WriteString("package dom")
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genwrapnode/genwrapnode.go. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	buf.WriteString("\n\"fmt\"\n")
	for _, lib := range []string{"github.com/Existed/libxml2/clib", "github.com/Existed/libxml2/types"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(lib))
	}
	buf.WriteString("\n)")

	nodeTypes := []string{
		`Namespace`,
		`Attribute`,
		`CDataSection`,
		`Comment`,
		`Element`,
		`Text`,
		`Pi`,
	}

	for _, typ := range nodeTypes {
		fmt.Fprintf(&buf, "\n\nfunc wrap%sNode(ptr uintptr) *%s {", typ, typ)
		fmt.Fprintf(&buf, "\nvar n %s", typ)
		buf.WriteString("\nn.ptr = ptr")
		buf.WriteString("\nreturn &n")
		buf.WriteString("\n}")
	}

	buf.WriteString("\n\n// WrapNode is a function created with the sole purpose of allowing")
	buf.WriteString("\n// go-libxml2 consumers that can generate a C.xmlNode pointer to")
	buf.WriteString("\n// create libxml2.Node types, e.g. go-xmlsec.")
	buf.WriteString("\nfunc WrapNode(n uintptr) (types.Node, error) {")
	buf.WriteString("\nswitch typ := clib.XMLGetNodeTypeRaw(n); typ {")

	for _, typ := range nodeTypes {
		// XXX hmm, this never existed. don't have time to debug right now.
		// possibly an omission bug?
		if typ == "Namespace" {
			continue
		}
		fmt.Fprintf(&buf, "\ncase clib.%sNode:", typ)
		fmt.Fprintf(&buf, "\nreturn wrap%sNode(n), nil", typ)
	}

	buf.WriteString("\ndefault:")
	buf.WriteString("\nreturn nil, fmt.Errorf(\"unknown node: %%d\", typ)")
	buf.WriteString("\n}")
	buf.WriteString("\n}")

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return err
	}

	var out io.Writer = os.Stdout
	args := os.Args
	if len(args) > 2 && args[1] == "--" {
		args = append(append([]string(nil), args[1:]...), args[2:]...)
	}

	if len(args) > 1 {
		f, err := os.OpenFile(args[1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrapf(err, `failed to open %s`, args[1])
		}
		defer f.Close()
		out = f
	}

	out.Write(src)
	return nil
}
