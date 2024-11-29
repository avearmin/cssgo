package cssgo

import (
	"io"
	"strings"
)

// SelectorNode represents a CSS selector node.
// Any type implementing this interface can be rendered as a valid CSS selector.
type SelectorNode interface {
	Node
	selectorNode()
}

// Selector is a concrete type that represents a CSS selector as a function.
// It implements the SelectorNode interface, allowing it to be rendered as CSS.
type Selector func(io.Writer) error

// RenderCSS writes the CSS representation of the selector to the provided io.Writer.
// Example: For a selector like `Class("foo")`, this method writes `.foo`.
//
// Parameters:
// - w (io.Writer): The destination where the CSS representation is written.
//
// Returns:
// - error: If an error occurs during the write operation, it is returned; otherwise, nil.
func (s Selector) RenderCSS(w io.Writer) error {
	return s(w)
}

func (s Selector) String() string {
	var b strings.Builder
	_ = s.RenderCSS(&b)
	return b.String()
}

func (s Selector) selectorNode() {}

// Or combines two selectors with a comma (`,`), creating a group selector.
// Example: `.class1, .class2`.
//
// Parameters:
// - other (Selector): The other selector to combine with this one.
//
// Returns:
// - Selector: A new Selector representing the combined group.
func (s Selector) Or(other Selector) Selector {
	return Selector(func(w io.Writer) error {
		if err := s.RenderCSS(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte(", ")); err != nil {
			return err
		}

		err := other.RenderCSS(w)

		return err
	})
}

// Props creates a CSS rule block for the selector with the given properties.
// Example: `.class { color: red; }`.
//
// Parameters:
// - properties (...PropertyNode): One or more CSS properties to include in the rule.
//
// Returns:
// - RuleNodeFunc: A function that renders the full CSS rule.
func (s Selector) Props(properties ...PropertyNode) RuleNodeFunc {
	return RuleNodeFunc(func(w io.Writer) error {
		if err := s.RenderCSS(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte("{")); err != nil {
			return err
		}

		for _, prop := range properties {
			if err := prop.RenderCSS(w); err != nil {
				return err
			}
		}

		_, err := w.Write([]byte("}"))

		return err
	})
}

// selector is a helper function for creating selectors with a given prefix (e.g., `.` or `#`).
// Example: selector(".", "class") -> `.class`
//
// Parameters:
// - selectorLiteral (string): The prefix for the selector (e.g., `.` for class, `#` for ID).
// - name (string): The name of the selector (e.g., `button`, `header`).
//
// Returns:
// - Selector: A Selector instance representing the full CSS selector.
func selector(selectorLiteral, name string) Selector {
	return Selector(func(w io.Writer) error {
		if selectorLiteral != "" {
			if _, err := w.Write([]byte(selectorLiteral)); err != nil {
				return err
			}
		}

		_, err := w.Write([]byte(name))
		return err
	})
}

// Class creates a CSS class selector.
// Example: Class("classname") -> `.classname`
//
// Parameters:
// - name (string): The class name.
//
// Returns:
// - Selector: A Selector instance representing the class selector.
func Class(name string) Selector {
	return selector(".", name)
}

// ID creates a CSS ID selector.
// Example: ID("idname") -> `#idname`
//
// Parameters:
// - name (string): The ID name.
//
// Returns:
// - Selector: A Selector instance representing the ID selector.
func ID(name string) Selector {
	return selector("#", name)
}

// El creates a CSS element selector.
// Example: El("div") -> `div`
//
// Parameters:
// - name (string): The element name (e.g., `div`, `h1`).
//
// Returns:
// - Selector: A Selector instance representing the element selector.
func El(name string) Selector {
	return selector("", name)
}
