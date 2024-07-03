package golanguagetool

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"strings"
	"unsafe"

	"github.com/go-openapi/strfmt"
	text_processing "github.com/tibotix/golanguagetool/internal/text"
	"github.com/tibotix/golanguagetool/pkg/api/operations"
)

// type Match struct {
// }

//	type CheckResults struct {
//		checkedLanguageCode  string
//		detectedLanguageCode string
//		matches              []Match
//	}
type CheckResults operations.PostCheckOKBody

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

func (client *Client) OpenFile(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// transformer := unicode.BOMOverride(unicode.UTF8.NewDecoder())
	// data, _, err = transform.Bytes(transformer, data)
	// if err != nil {
	// 	return nil, err
	// }
	return data, nil
	// unicode.UTF16(unicode.BigEndian, unicode.UseBOM)

}

func (client *Client) CheckText(t Text, options *CheckOptions) (*CheckResults, error) {
	// If file is empty, early return with no results
	if len(t.Contents) == 0 {
		return &CheckResults{}, nil
	}

	if options == nil {
		options = &defaultCheckOptions
	}

	// fmt.Println(string(t.Contents[:]))
	desc, ok := fileTypeMappings[t.FileType]
	if !ok {
		return nil, errors.New("unknown file type")
	}
	result, err := desc.textTransformer(t.Contents)
	if err != nil {
		return nil, err
	}
	// fmt.Println()
	// fmt.Println(string(result.data[:]))
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

	// fmt.Printf("%+v\n", resp)
	return (*CheckResults)(resp.Payload), nil
}
