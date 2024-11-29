Here’s an example that explicitly mimics some common **TailwindCSS classes** as Go inline style utility functions. This demonstrates how you can use **CSSGo** to replicate Tailwind's utility-first design system, directly in Go, with type safety and seamless integration.

---

## **Why CSSGo Can Replace TailwindCSS**

TailwindCSS's utility-first approach enables rapid styling with classes like `p-4` or `rounded-md`. With **CSSGo**, you can define similar utilities as Go functions, making them:
- **Type-Safe**: No chance of misspelled class names or invalid values.
- **Customizable**: Functions can accept parameters for dynamic styling.
- **Colocated**: Styles live alongside components, improving readability and maintainability.

Here’s how **CSSGo** replicates some common TailwindCSS utilities.

---

### **Tailwind-Like Utilities**

```go
import (
	g "maragu.dev/gomponents"         // Gomponents
	ghtml "maragu.dev/gomponents/html" // Gomponents HTML helpers
	c "github.com/avearmin/cssgo"     // CSSGo for CSS styling
	chtml "github.com/avearmin/cssgo/html" // CSSGo HTML helpers
)

// Padding utilities
func P4() c.Property {
	return c.Padding(c.PX(16), nil, nil, nil) // Tailwind `p-4` = 16px padding
}

func Px2() c.Property {
	return c.Padding(nil, c.PX(8), nil, c.PX(8)) // Tailwind `px-2` = 8px horizontal padding
}

// Rounded utilities
func RoundedMd() c.Property {
	return c.BorderRadius(c.PX(6)) // Tailwind `rounded-md` = 6px radius
}

// Background color utility
func BgBlue500() c.Property {
	return c.BackgroundColor(c.Blue)
}

// Text color utility
func TextWhite() c.Property {
	return c.TextColor(c.White)
}

// Width/Height utility
func Size52() c.Property {
    return c.GroupProps(
        c.Width(REM(13)),
        c.Height(REM(13)),
    )
}

func ButtonComponent(label string) g.Node {
	return ghtml.Button(
		chtml.Style(
			BgBlue500(),  // Tailwind `bg-blue-500`
			TextWhite(),  // Tailwind `text-white`
			P4(),         // Tailwind `p-4`
			RoundedMd(),  // Tailwind `rounded-md`
		),
		g.Text(label),
	)
}

func CardComponent(title, content string) g.Node {
	return ghtml.Div(
		chtml.Style(
			Size52(),     // Tailwind `size-52`
			Px2(),        // Tailwind `px-2`
			RoundedMd(),  // Tailwind `rounded-md`
			c.BackgroundColor(c.White),
		),
		ghtml.H2(
			g.Text(title),
			chtml.Style(c.TextColor(c.Blue)),
		),
		ghtml.P(
			g.Text(content),
			chtml.Style(c.TextColor(c.Gray)),
		),
		ButtonComponent("Learn More"),
	)
}
```

---

### **Output**

Generated HTML for the button:
```html
<button style="background-color: blue; color: white; padding: 16px; border-radius: 6px;">
    Learn More
</button>
```

Generated HTML for the card:
```html
<div style="width: 13rem; height: 13rem; padding: 8px; border-radius: 6px; background-color: white;">
    <h2 style="color: blue;">Title</h2>
    <p style="color: gray;">Content</p>
    <button style="background-color: blue; color: white; padding: 16px; border-radius: 6px;">
        Learn More
    </button>
</div>
```

---

### **Advantages of CSSGo**
- **Dynamic and Reusable Utilities**: Utilities like `P4()` and `BgBlue500()` mimic Tailwind classes and can be customized dynamically in Go.
- **Colocated Styles**: Styles are colocated with components, improving code cohesion.
- **Type Safety**: Go ensures that invalid values or typos are caught at compile time.