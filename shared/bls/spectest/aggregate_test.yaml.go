// Code generated by yaml_to_go. DO NOT EDIT.
// source: aggregate.yaml

package spectest

type AggregateTest struct {
	Input  []string `json:"input"`
	Output string   `json:"output" ssz:"size=96"`
}