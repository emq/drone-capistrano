package main

import "testing"

func TestBundlerNotQuiet(t *testing.T) {
	vargs := Params{}
	vargs.Debug = true
	args := bundlerArgs(vargs)

	if containsString(args, "--quiet") {
		t.Error("Not expecting --quiet argument to be one of bundler arguments", args)
	}
}

func TestBundlerQuiet(t *testing.T) {
	vargs := Params{}
	args := bundlerArgs(vargs)

	if !containsString(args, "--quiet") {
		t.Error("Expecting --quiet argument to be one of bundler arguments", args)
	}
}

func TestBundlerPath(t *testing.T) {
	vargs := Params{}
	vargs.BundlePath = "foo/bar"
	args := bundlerArgs(vargs)

	if !containsString(args, "--path") {
		t.Error("Missing --path argument in bundler arguments", args)
	}
	if !containsString(args, "foo/bar") {
		t.Error("Missing --path argument value in bundler arguments", args)
	}
}

func TestBundlerArgs(t *testing.T) {
	vargs := Params{}
	vargs.BundleArgs = "--without development"
	args := bundlerArgs(vargs)

	if !containsString(args, "--without development") {
		t.Error("Expecting --without development argument to be one of bundler arguments", args)
	}
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
