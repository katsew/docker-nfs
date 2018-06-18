package exports

import "testing"

func TestParseEmptyLine(t *testing.T) {
	c, err := Parse("")
	if err != nil {
		t.Fatal(err)
	}
	if !c.isEmpty() {
		t.Fatalf("expect empty, but actual %+v", c)
	}
}

func TestParseComment(t *testing.T) {
	c, err := Parse("# This is a comment")
	if err != nil {
		t.Fatal(err)
	}
	if !c.isComment() {
		t.Fatalf("expect comment, but actual %+v", c)
	}
}
