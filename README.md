# Gomour

**Gomour** is a Go-native, type-safe CSS library inspired by [Elm CSS](https://package.elm-lang.org/packages/rtfeldman/elm-css/latest/). Just as Elm CSS integrates seamlessly with Elm’s HTML package, Gomour integrates with [Gomponents](https://github.com/maragudk/gomponents), enabling developers to create modular, reusable CSS directly in Go.

Gomour brings the power of type-safe CSS into the Go ecosystem, allowing you to write styles alongside your components while ensuring correctness and maintainability.

---

## **Why Gomour?**

Writing CSS in Go comes with challenges:
- Inline styles often lack structure, leading to duplication and hard-to-maintain code.
- String-based CSS is prone to typos and errors (e.g., `colro: blie;`).
- CSS files disconnected from components can create context-switching and modularity issues.

Inspired by **Elm CSS**, Gomour solves these problems by:
- **Ensuring Type Safety**: Only valid CSS properties and values can be used, reducing runtime errors.
- **Encouraging Modularity**: Keep styles close to your components, promoting reuse and reducing duplication.
- **Integrating Seamlessly**: Just as Elm CSS works with Elm's HTML package, Gomour works effortlessly with Gomponents.

---

## **Features**

- **Type-Safe Properties**: Define styles with confidence—invalid values are caught at compile time.
- **Reusable Values**: Use predefined constants for colors, border styles, and more.
- **Inline and Scoped Styles**: Create both inline styles and reusable `<style>` rules.
- **Gomponents Integration**: Designed to work natively with Gomponents for a seamless developer experience.

---

## **Installation**

Install Gomour via `go get`:

```bash
go get github.com/avearmin/gomour
```

---

## **Examples**

### **Imports**
Before diving into examples, let’s look at the required imports:

```go
import (
	g "maragu.dev/gomponents" // Gomponents
	ghtml "maragu.dev/gomponents/html" // HTML elements from gomponents
    gcss "github.com/avearmin/gomour"// Gomour for CSS styling
	gmhtml "github.com/avearmin/gomour/html" // Custom HTML helpers for Gomour 	
)
```

---

### **1. Inline Styles with Gomponents**

#### **With Gomour**
```go
func CardComponent(title, content string) g.Node {
	return ghtml.Div(
		gmhtml.Style(
			gcss.Width(gcss.PX(300)),
			gcss.Height(gcss.PX(200)),
			gcss.BackgroundColor(gcss.White),
			gcss.Border(gcss.PX(2), gcss.Solid, gcss.Hex(0xcccccc)),
			gcss.Padding(gcss.PX(20), nil, nil, nil),
		),
		ghtml.H2(
			ghtml.Text(title),
			gmhtml.Style(gcss.TextColor(gcss.Blue)),
		),
		ghtml.P(
			ghtml.Text(content),
			gmhtml.Style(gcss.TextColor(gcss.Gray)),
		),
	)
}
```

#### **Without Gomour**
```go
func CardComponent(title, content string) g.Node {
	return ghtml.Div(
		ghtml.Attr("style", `width: 300px; height: 200px; background-color: white; border: 2px solid #cccccc; padding: 20px;`),
		ghtml.H2(
			ghtml.Text(title),
			ghtml.Attr("style", `color: blue;`),
		),
		ghtml.P(
			ghtml.Text(content),
			ghtml.Attr("style", `color: gray;`),
		),
	)
}
```

**Why Gomour?**
- Reduce human error with type-safe properties and values.
- Modularize reusable constants like `gomour.White` and `gomour.PX(20)`.

---

### **2. Scoped Styles with `<style>` Rules**

Define reusable CSS rules colocated with your components.

```go
func StyledButtonComponent(label string) g.Node {
	return gmhtml.StyleEl(
			// Define button styles
			gcss.El("button").Or(gcss.Class("foo")).Props(
				gcss.BackgroundColor(gcss.Blue),
				gcss.TextColor(gcss.White),
				gcss.Padding(gcss.PX(10), gcss.PX(20), nil, nil),
			),
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

**Why Gomour?**
- Simplifies the use of comma-separated selectors with `Or`.
- Keeps CSS colocated with the components they affect.

---

### **3. Grouping Properties**

Reuse grouped properties for consistency across components.

```go
func CardComponent(title, content string) g.Node {
	cardStyle := gcss.GroupProps(
		gcss.Width(gomour.PX(300)),
		gcss.Height(gomour.PX(200)),
		gcss.BackgroundColor(gomour.White),
		gcss.Border(gomour.PX(2), gomour.Solid, gomour.Hex(0xcccccc)),
		gcss.Padding(gomour.PX(20), nil, nil, nil),
	)

	return g.Div(
		gmhtml.Style(cardStyle),
		ghtml.H2(
			ghtml.Text(title),
		),
		ghtml.P(
			ghtml.Text(content),
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

In the future I'd like to open this up to contributions, but as of now I am still ironing out quirks.