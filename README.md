Here’s the updated README, replacing **Gomour** with **CSSGo** and adjusting imports, examples, and naming conventions accordingly:

---

# CSSGo

**CSSGo** is a Go-native, type-safe CSS library inspired by [Elm CSS](https://package.elm-lang.org/packages/rtfeldman/elm-css/latest/). Just as Elm CSS integrates seamlessly with Elm’s HTML package, CSSGo integrates with [Gomponents](https://github.com/maragudk/gomponents), enabling developers to create modular, reusable CSS directly in Go.

CSSGo brings the power of type-safe CSS into the Go ecosystem, allowing you to write styles alongside your components while ensuring correctness and maintainability.

---

## **Why CSSGo?**

Writing CSS in Go comes with challenges:
- Inline styles often lack structure, leading to duplication and hard-to-maintain code.
- String-based CSS is prone to typos and errors (e.g., `colro: blie;`).
- CSS files disconnected from components can create context-switching and modularity issues.

Inspired by **Elm CSS**, CSSGo solves these problems by:
- **Ensuring Type Safety**: Only valid CSS properties and values can be used, reducing runtime errors.
- **Encouraging Modularity**: Keep styles close to your components, promoting reuse and reducing duplication.
- **Integrating Seamlessly**: Just as Elm CSS works with Elm's HTML package, CSSGo works effortlessly with Gomponents.

---

## **Features**

- **Type-Safe Properties**: Define styles with confidence—invalid values are caught at compile time.
- **Reusable Values**: Use predefined constants for colors, border styles, and more.
- **Inline and Scoped Styles**: Create both inline styles and reusable `<style>` rules.
- **Gomponents Integration**: Designed to work natively with Gomponents for a seamless developer experience.

---

## **Installation**

Install CSSGo via `go get`:

```bash
go get github.com/avearmin/cssgo
```

---

## **Examples**

### **Imports**
Before diving into examples, let’s look at the required imports:

```go
import (
	g "maragu.dev/gomponents" // Gomponents
	ghtml "maragu.dev/gomponents/html" // HTML elements from Gomponents 
	c "github.com/avearmin/cssgo" // CSSGo for CSS styling
	chtml "github.com/avearmin/cssgo/html" // Custom HTML helpers for CSSGo 	
)
```

---

### **1. Inline Styles with Gomponents**

#### **With CSSGo**
```go
func CardComponent(title, content string) g.Node {
	return ghtml.Div(
		chtml.Style(
			c.Width(c.PX(300)),
			c.Height(c.PX(200)),
			c.BackgroundColor(c.White),
			c.Border(c.PX(2), c.Solid, c.Hex(0xcccccc)),
			c.Padding(c.PX(20), nil, nil, nil),
		),
		ghtml.H2(
			g.Text(title),
			chtml.Style(c.TextColor(c.Blue)),
		),
		ghtml.P(
			g.Text(content),
			c.Style(c.TextColor(c.Gray)),
		),
	)
}
```

#### **Without CSSGo**
```go
func CardComponent(title, content string) g.Node {
	return ghtml.Div(
		g.Attr("style", `width: 300px; height: 200px; background-color: white; border: 2px solid #cccccc; padding: 20px;`),
		ghtml.H2(
			g.Text(title),
			g.Attr("style", `color: blue;`),
		),
		ghtml.P(
			g.Text(content),
			g.Attr("style", `color: gray;`),
		),
	)
}
```

**Why CSSGo?**
- Reduce human error with type-safe properties and values.
- Modularize reusable constants like `cssgo.White` and `cssgo.PX(20)`.

---

### **2. Scoped Styles with `<style>` Rules**

Define reusable CSS rules colocated with your components.

```go
func StyledButtonComponent() g.Node {
	return csshtml.StyleEl(
		// Define button styles
		c.El("button").Or(c.Class("foo")).Props(
			c.BackgroundColor(c.Blue),
			c.TextColor(c.White),
			c.Padding(c.PX(10), c.PX(20), nil, nil),
		),
	)
}
```

Generated CSS:
```css
button, .foo {
    background-color: blue;
    color: white;
    padding: 10px 20px;
}
```

**Why CSSGo?**
- Simplifies the use of comma-separated selectors with `Or`.
- Keeps CSS colocated with the components they affect.

---

### **3. Grouping Properties**

Reuse grouped properties for consistency across components.

```go
func CardComponent(title, content string) g.Node {
	cardStyle := c.GroupProps(
		c.Width(c.PX(300)),
		c.Height(cssgo.PX(200)),
		c.BackgroundColor(c.White),
		c.Border(c.PX(2), c.Solid, c.Hex(0xcccccc)),
		c.Padding(c.PX(20), nil, nil, nil),
	)

	return ghtml.Div(
		chtml.Style(cardStyle),
		ghtml.H2(
			g.Text(title),
		),
		ghtml.P(
			g.Text(content),
		),
	)
}
```

---

## **Roadmap**

1. Add support for more CSS properties (e.g., `box-shadow`, `grid`).
2. Integrate advanced selectors like `nth-child` and `::after`.
3. Provide utilities for theming and design systems.

---

## **Contributing**

In the future, I'd like to open this up to contributions, but as of now, I am still ironing out quirks.