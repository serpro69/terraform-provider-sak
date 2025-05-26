package provider_test

import (
	"fmt"
	"math/big"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestYamlDecode(t *testing.T) {
	t.Parallel()

	outputName := "test"

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testYamlDecodeConfig("testdata/decode_single.yaml"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(outputName, knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_plain":              knownvalue.StringExact("This is a plain string"),
							"string_single_quoted":      knownvalue.StringExact("This string contains special characters like : or #"),
							"string_double_quoted":      knownvalue.StringExact("This string can have \n newlines and \t tabs."),
							"string_block_literal":      knownvalue.StringExact("This is a literal block scalar.\nNewlines are preserved.\nIndentation matters here.\n"),
							"string_block_folded":       knownvalue.StringExact("This is a folded block scalar. Newlines are typically folded into spaces, but an empty line creates a paragraph break.\nLike this.\n"),
							"null_value":                knownvalue.Null(),
							"empty_value":               knownvalue.Null(),
							"boolean_true":              knownvalue.Bool(true),
							"boolean_false":             knownvalue.Bool(false),
							"integer_decimal":           knownvalue.NumberExact(big.NewFloat(12345)),
							"integer_octal":             knownvalue.NumberExact(big.NewFloat(63)),  // YAML 0o77
							"integer_hexadecimal":       knownvalue.NumberExact(big.NewFloat(255)), // YAML 0xFF
							"float_fixed":               knownvalue.NumberExact(big.NewFloat(2.5)),
							"float_exponential":         knownvalue.NumberExact(new(big.Float).SetFloat64(6.022e+8)),
							"timestamp_iso8601":         knownvalue.StringExact("2023-10-26 14:30:00 +0000 UTC"),
							"timestamp_space_separated": knownvalue.StringExact("2023-10-26 14:30:00 -05:00"),
							"timestamp_canonical":       knownvalue.StringExact("2001-12-15 02:59:43.1 +0000 UTC"),
							"date_simple":               knownvalue.StringExact("2023-10-26 00:00:00 +0000 UTC"),
							"binary_data":               knownvalue.StringExact("hello, world"),
							"map": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"submap": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"another_string": knownvalue.StringExact("another string"),
								}),
								"string": knownvalue.StringExact("a string"),
								"int":    knownvalue.NumberExact(big.NewFloat(42)),
							}),
							"array": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("item1"),
								knownvalue.StringExact("item2"),
								knownvalue.StringExact("item3"),
							}),
							"array_with_map": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"key":   knownvalue.StringExact("key1"),
									"value": knownvalue.NumberExact(big.NewFloat(1)),
								}),
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"key":   knownvalue.StringExact("key2"),
									"value": knownvalue.NumberExact(big.NewFloat(2)),
								}),
							}),
						}),
					})),
				},
			},
			{
				Config: testYamlDecodeConfig("testdata/decode_multi.yaml"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(outputName, knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_plain":              knownvalue.StringExact("This is a plain string"),
							"string_single_quoted":      knownvalue.StringExact("This string contains special characters like : or #"),
							"string_double_quoted":      knownvalue.StringExact("This string can have \n newlines and \t tabs."),
							"string_block_literal":      knownvalue.StringExact("This is a literal block scalar.\nNewlines are preserved.\nIndentation matters here.\n"),
							"string_block_folded":       knownvalue.StringExact("This is a folded block scalar. Newlines are typically folded into spaces, but an empty line creates a paragraph break.\nLike this.\n"),
							"null_value":                knownvalue.Null(),
							"empty_value":               knownvalue.Null(),
							"boolean_true":              knownvalue.Bool(true),
							"boolean_false":             knownvalue.Bool(false),
							"integer_decimal":           knownvalue.NumberExact(big.NewFloat(12345)),
							"integer_octal":             knownvalue.NumberExact(big.NewFloat(63)),  // YAML 0o77
							"integer_hexadecimal":       knownvalue.NumberExact(big.NewFloat(255)), // YAML 0xFF
							"float_fixed":               knownvalue.NumberExact(big.NewFloat(2.5)),
							"float_exponential":         knownvalue.NumberExact(new(big.Float).SetFloat64(6.022e+8)),
							"timestamp_iso8601":         knownvalue.StringExact("2023-10-26 14:30:00 +0000 UTC"),
							"timestamp_space_separated": knownvalue.StringExact("2023-10-26 14:30:00 -05:00"),
							"timestamp_canonical":       knownvalue.StringExact("2001-12-15 02:59:43.1 +0000 UTC"),
							"date_simple":               knownvalue.StringExact("2023-10-26 00:00:00 +0000 UTC"),
							"binary_data":               knownvalue.StringExact("hello, world"),
							"map": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"submap": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"another_string": knownvalue.StringExact("another string"),
								}),
								"string": knownvalue.StringExact("a string"),
								"int":    knownvalue.NumberExact(big.NewFloat(42)),
							}),
							"array": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.StringExact("item1"),
								knownvalue.StringExact("item2"),
								knownvalue.StringExact("item3"),
							}),
							"array_with_map": knownvalue.ListExact([]knownvalue.Check{
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"key":   knownvalue.StringExact("key1"),
									"value": knownvalue.NumberExact(big.NewFloat(1)),
								}),
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"key":   knownvalue.StringExact("key2"),
									"value": knownvalue.NumberExact(big.NewFloat(2)),
								}),
							}),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"string_plain":         knownvalue.StringExact("This is a plain string"),
							"string_single_quoted": knownvalue.StringExact("This string contains special characters like : or #"),
							"string_double_quoted": knownvalue.StringExact("This string can have \n newlines and \t tabs."),
							"string_block_literal": knownvalue.StringExact("This is a literal block scalar.\nNewlines are preserved.\nIndentation matters here.\n"),
							"map": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"submap": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"test": knownvalue.StringExact("foo---bar"),
								}),
							}),
							"another_map": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"bool":                knownvalue.Bool(false),
								"embedded_pseudo_doc": knownvalue.StringExact("---\ntest: document"),
							}),
						}),
					})),
				},
			},
			{
				Config:      testYamlDecodeConfig("testdata/decode_invalid.yaml"),
				ExpectError: regexp.MustCompile(`Invalid\s+YAML\s+document`),
			},
		},
	})
}

func testYamlDecodeConfig(filename string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf(`
output "test" {
  value = provider::sak::yaml_decode(file(%q))
}`, path.Join(cwd, filename))
}
