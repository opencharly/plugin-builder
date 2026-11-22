package builderkind

import (
	"os"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

// #BuilderInput reproduces core #Builder: there is no detect_config (a manager declaration lives
// on its candy's distro tag section) and no authored package output (what a builder produces is
// the builder plugin's own declaration, DESIGN.cue #BuilderAlt; opencharly/sdk#608). The plugin
// validates every authored builder body against this def, so a drift from core would accept a key
// core dropped.
func TestBuilderInputMatchesCoreBuilder(t *testing.T) {
	src, err := os.ReadFile("schema/builder.cue")
	if err != nil {
		t.Fatal(err)
	}
	ctx := cuecontext.New()
	schema := ctx.CompileBytes(src)
	if schema.Err() != nil {
		t.Fatalf("schema does not compile: %v", schema.Err())
	}
	def := schema.LookupPath(cue.ParsePath("#BuilderInput"))
	check := func(doc string) error {
		v := def.Unify(ctx.CompileString(doc))
		return v.Validate(cue.Concrete(false))
	}
	if err := check(`{detect_file: ["package.json"]}`); err != nil {
		t.Errorf("#BuilderInput rejects a plain builder node: %v", err)
	}
	for _, authored := range []string{
		`{copy_artifact: [{src: "/home/u", dst: "/home/u"}]}`,
		`{copy_binary: {src: "/usr/local/bin/pixi", dst: "/usr/local/bin/pixi"}}`,
	} {
		if err := check(authored); err == nil {
			t.Errorf("#BuilderInput accepts %s: the payload is the builder plugin's reply, never an authored field", authored)
		}
	}
	if err := check(`{detect_config: "aur"}`); err == nil {
		t.Error("#BuilderInput still accepts detect_config")
	}
	if err := check(`{package: {format: "pac", dir: "/tmp/aur-pkgs/x", glob: "*"}}`); err == nil {
		t.Error("#BuilderInput accepts an authored package output; it is the builder plugin's declaration")
	}
}
