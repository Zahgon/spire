package agentpathtemplate

import (
	"fmt"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
)

var funcList = []string{
	"abbrev",
	"abbrevboth",
	"trunc",
	"trim",
	"upper",
	"lower",
	"title",
	"untitle",
	"substr",
	"repeat",
	"trimAll",
	"trimSuffix",
	"trimPrefix",
	"nospace",
	"initials",
	"swapcase",
	"snakecase",
	"camelcase",
	"kebabcase",
	"wrap",
	"wrapWith",
	"contains",
	"hasPrefix",
	"hasSuffix",
	"quote",
	"squote",
	"cat",
	"indent",
	"nindent",
	"replace",
	"plural",
	"sha1sum",
	"sha256sum",
	"adler32sum",
	"toString",
	"seq",
	"splitList",
	"toStrings",
	"join",
	"sortAlpha",
	"default",
	"empty",
	"coalesce",
	"all",
	"any",
	"compact",
	"mustCompact",
	"ternary",
	"base",
	"dir",
	"clean",
	"ext",
	"isAbs",
	"b64enc",
	"b64dec",
	"b32enc",
	"b32dec",
	"tuple",
	"list",
	"dict",
	"get",
	"set",
	"unset",
	"hasKey",
	"pluck",
	"keys",
	"pick",
	"omit",
	"merge",
	"mergeOverwrite",
	"mustMerge",
	"mustMergeOverwrite",
	"values",
	"append",
	"push",
	"mustAppend",
	"mustPush",
	"prepend",
	"mustPrepend",
	"first",
	"mustFirst",
	"rest",
	"mustRest",
	"last",
	"mustLast",
	"initial",
	"mustInitial",
	"reverse",
	"mustReverse",
	"uniq",
	"mustUniq",
	"without",
	"mustWithout",
	"has",
	"mustHas",
	"slice",
	"mustSlice",
	"concat",
	"dig",
	"chunk",
	"mustChunk",
	"uuidv4",
	"fail",
	"regexMatch",
	"mustRegexMatch",
	"regexFindAll",
	"mustRegexFindAll",
	"regexFind",
	"mustRegexFind",
	"regexReplaceAll",
	"mustRegexReplaceAll",
	"regexReplaceAllLiteral",
	"mustRegexReplaceAllLiteral",
	"regexSplit",
	"mustRegexSplit",
	"regexQuoteMeta",
	"urlParse",
	"urlJoin",
}

var ourMap = make(template.FuncMap)

func init() {
	sprigMap := sprig.TxtFuncMap()
	for _, f := range funcList {
		if fn, ok := sprigMap[f]; ok {
			ourMap[f] = fn
		} else {
			panic(fmt.Errorf("missing sprig function %q", f))
		}
	}
}

// Parse parses an agent path template. It changes the behavior for missing
// keys to return an error instead of the default behavior, which renders a
// value that requires percent-encoding to include in a URI, which is against
// the SPIFFE specification.
func Parse(text string) (*Template, error) { _ = "STUB: not implemented"; return nil, nil }

// MustParse parses an agent path template. It changes the behavior for missing
// keys to return an error instead of the default behavior, which renders a
// value that requires percent-encoding to include in a URI, which is against
// the SPIFFE specification. If parsing fails, the function panics.
func MustParse(text string) *Template { _ = "STUB: not implemented"; return nil }

type Template struct {
	tmpl *template.Template
}

func (t *Template) Execute(args any) (string, error) { _ = "STUB: not implemented"; return "", nil }
