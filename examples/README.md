# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or are testable even if some parts are not relevant for the documentation.

* `provider/provider.tf` example file for the provider index page
* `resources/<TYPE>/resource.tf` example file for the named source page
* `data-sources/<TYPE>/data-source.tf` example file for the named data source page
- `functions/<TYPE>/function.tf` example file for the named function page

Replace `<TYPE>` with the name of the resource, data source, or function. For example: `examples/functions/rfc3339_parse/function.tf`.
