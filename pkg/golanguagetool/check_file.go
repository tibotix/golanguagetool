package golanguagetool

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"unsafe"

	"github.com/go-openapi/strfmt"
	text_processing "github.com/tibotix/golanguagetool/internal/text"
	"github.com/tibotix/golanguagetool/pkg/api/operations"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Match struct {
	*operations.PostCheckOKBodyMatchesItems0
	LineNumber int
}

type CheckResults struct {
	Language *operations.PostCheckOKBodyLanguage
	Matches  []*Match
	Software *operations.PostCheckOKBodySoftware
}

// type CheckResults operations.PostCheckOKBody

type CheckLevel int

const (
	CheckLevelDefault = iota
	CheckLevelPicky
)

var checkLevelValues = map[CheckLevel]string{
	CheckLevelDefault: "default",
	CheckLevelPicky:   "picky",
}

func (level CheckLevel) String() string {
	return checkLevelValues[level]
}

type CheckOptions struct {
	Language           string
	Dicts              []string
	MotherTongue       string
	PreferredVariants  []string
	EnabledRules       []int
	DisabledRules      []int
	EnabledCategories  []int
	DisabledCategories []int
	EnabledOnly        bool
	Level              CheckLevel
}

var defaultCheckOptions = CheckOptions{
	Language:           "auto",
	Dicts:              []string{},
	MotherTongue:       "",
	PreferredVariants:  []string{"en-US", "de-DE"},
	EnabledRules:       []int{},
	DisabledRules:      []int{},
	EnabledCategories:  []int{},
	DisabledCategories: []int{},
	EnabledOnly:        false,
	Level:              CheckLevelDefault,
}

func i2s_array(i []int) []string {
	s := make([]string, len(i))
	for _, v := range i {
		s = append(s, strconv.Itoa(v))
	}
	return s
}

func ByteArrayToString(buf []byte) string {
	return unsafe.String(unsafe.SliceData(buf), len(buf))
}

type FileType int

const (
	FileTypePlain = iota
	FileTypeMarkdown
)

type textTransformerDesc struct {
	fileExtensions  []string
	textTransformer text_processing.TextTransformer
}

var fileTypeMappings = map[FileType]textTransformerDesc{
	FileTypePlain: {
		fileExtensions:  []string{"plain", "txt"},
		textTransformer: text_processing.TransformTextPlain,
	},
	FileTypeMarkdown: {
		fileExtensions:  []string{"markdown", "md"},
		textTransformer: text_processing.TransformTextMarkdown,
	},
}

var SupportedFileTypes = func() []string {
	l := make([]string, 10)
	for _, v := range fileTypeMappings {
		l = append(l, v.fileExtensions...)
	}
	return l
}()

func GetFileTypeFromString(s string) FileType {
	for fileType, v := range fileTypeMappings {
		if slices.Contains(v.fileExtensions, s) {
			return fileType
		}
	}
	return FileTypePlain
}

func DetermineFileType(file string) FileType {
	parts := strings.Split(file, ".")
	return GetFileTypeFromString(parts[len(parts)-1])
}

func (f FileType) String() string {
	for fileType, v := range fileTypeMappings {
		if f == fileType {
			return v.fileExtensions[0]
		}
	}
	return "unknown"
}

type Text struct {
	Contents []byte
	FileType FileType
}

func (client *Client) transformToUTF8(data []byte) ([]byte, error) {
	transformer := unicode.BOMOverride(unicode.UTF8.NewDecoder())
	data, _, err := transform.Bytes(transformer, data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (client *Client) CheckText(t Text, options *CheckOptions) (*CheckResults, error) {
	contents, err := client.transformToUTF8(t.Contents)
	if err != nil {
		return nil, err
	}
	// If file is empty, early return with no results
	if len(contents) == 0 {
		return &CheckResults{}, nil
	}

	if options == nil {
		options = &defaultCheckOptions
	}

	desc, ok := fileTypeMappings[t.FileType]
	if !ok {
		return nil, errors.New("unknown file type")
	}
	result, err := desc.textTransformer(contents)
	if err != nil {
		return nil, err
	}
	data_str := ByteArrayToString(result.Data)

	dicts := strings.Join(options.Dicts, ",")
	preferredVariants := strings.Join(options.PreferredVariants, ",")
	enabledRules := strings.Join(i2s_array(options.EnabledRules), ",")
	disabledRules := strings.Join(i2s_array(options.DisabledRules), ",")
	enabledCategories := strings.Join(i2s_array(options.EnabledCategories), ",")
	disabledCategories := strings.Join(i2s_array(options.DisabledCategories), ",")
	level := options.Level.String()

	params := operations.PostCheckParams{
		APIKey:             (*strfmt.Password)(&client.ApiKey),
		Username:           &client.Username,
		Text:               &data_str,
		Language:           options.Language,
		Dicts:              &dicts,
		MotherTongue:       &options.MotherTongue,
		PreferredVariants:  &preferredVariants,
		EnabledRules:       &enabledRules,
		DisabledRules:      &disabledRules,
		EnabledCategories:  &enabledCategories,
		DisabledCategories: &disabledCategories,
		EnabledOnly:        &options.EnabledOnly,
		Level:              &level,
	}

	resp, err := client.ApiClient.Operations.PostCheck(&params)
	if err != nil {
		return nil, err
	}

	matches := make([]*Match, len(resp.Payload.Matches))
	for i, m := range resp.Payload.Matches {
		matches[i] = &Match{
			PostCheckOKBodyMatchesItems0: m,
			LineNumber:                   result.LineBeginnings.LookupLine(int(*m.Offset)),
		}
	}

	checkResult := &CheckResults{
		Language: resp.Payload.Language,
		Matches:  matches,
		Software: resp.Payload.Software,
	}
	return checkResult, nil
}
