package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Product represents an item in the store
type Product struct {
	id    int
	name  string
	price float64
}

// CartItem represents a product in the shopping cart
type CartItem struct {
	product  Product
	quantity int
}

// Model holds the state of the application
type Model struct {
	products   []Product
	cart       []CartItem
	selected   int
	view       string // "products" or "cart"
	list       list.Model
	width      int
	height     int
	totalPrice float64
}

// ================================
// PRODUCT LIST ITEM IMPLEMENTATION
// ================================

type productItem struct {
	product Product
}

func (i productItem) FilterValue() string {
	return i.product.name
}

func (i productItem) Title() string {
	return fmt.Sprintf("%s - $%.2f", i.product.name, i.product.price)
}

func (i productItem) Description() string {
	return fmt.Sprintf("Product ID: %d", i.product.id)
}

// ================================
// INITIALIZATION
// ================================

func initialModel() Model {
	// Sample products
	products := []Product{
		{1, "Laptop", 999.99},
		{2, "Wireless Mouse", 29.99},
		{3, "USB-C Cable", 15.99},
		{4, "Monitor Stand", 49.99},
		{5, "Mechanical Keyboard", 129.99},
		{6, "Webcam", 79.99},
	}

	// Convert products to list items
	items := make([]list.Item, len(products))
	for i, p := range products {
		items[i] = productItem{p}
	}

	// Create list
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Shopping Products"

	return Model{
		products: products,
		cart:     []CartItem{},
		selected: 0,
		view:     "products",
		list:     l,
	}
}

// ================================
// BUBBLETEA INTERFACE METHODS
// ================================

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// Quit application
		case "ctrl+c", "q":
			return m, tea.Quit

		// Switch between views
		case "tab":
			if m.view == "products" {
				m.view = "cart"
			} else {
				m.view = "products"
			}
			return m, nil

		// Add to cart (only in products view)
		case "enter":
			if m.view == "products" {
				selected := m.list.SelectedItem()
				if selected != nil {
					product := selected.(productItem).product
					m = addToCart(m, product)
				}
			}
			return m, nil

		// Remove from cart (only in cart view)
		case "delete", "backspace":
			if m.view == "cart" && len(m.cart) > 0 {
				m.cart = append(m.cart[:m.selected], m.cart[m.selected+1:]...)
				if m.selected > 0 && m.selected >= len(m.cart) {
					m.selected--
				}
				m.calculateTotal()
			}
			return m, nil

		// Navigate cart
		case "up":
			if m.view == "cart" && m.selected > 0 {
				m.selected--
			}
			return m, nil

		case "down":
			if m.view == "cart" && m.selected < len(m.cart)-1 {
				m.selected++
			}
			return m, nil
		}

		// Let list handle other key inputs in products view
		if m.view == "products" {
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	var content string

	if m.view == "products" {
		content = m.productsView()
	} else {
		content = m.cartView()
	}

	return content
}

// ================================
// VIEW FUNCTIONS
// ================================

func (m Model) productsView() string {
	title := titleStyle.Render("🛍️  SHOPPING PLATFORM - PRODUCTS")
	help := helpStyle.Render("↑/↓: Navigate | Enter: Add to Cart | Tab: View Cart | Q: Quit")

	return fmt.Sprintf("%s\n%s\n\n%s\n\n%s",
		title,
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("─────────────────────────────────"),
		m.list.View(),
		help,
	)
}

func (m Model) cartView() string {
	title := titleStyle.Render("🛒 SHOPPING CART")
	help := helpStyle.Render("↑/↓: Navigate | Delete: Remove Item | Tab: Back to Products | Q: Quit")

	if len(m.cart) == 0 {
		emptyMsg := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Margin(2).
			Render("Your cart is empty. Add items from products!")

		return fmt.Sprintf("%s\n%s\n\n%s\n\n%s",
			title,
			lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("─────────────────────────────────"),
			emptyMsg,
			help,
		)
	}

	// Build cart items list
	cartContent := ""
	for i, item := range m.cart {
		selected := ""
		if i == m.selected {
			selected = " ← "
		}
		cartContent += fmt.Sprintf("  %s %s (x%d) - $%.2f\n",
			selected,
			item.product.name,
			item.quantity,
			item.product.price*float64(item.quantity),
		)
	}

	totalLine := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("46")).
		Render(fmt.Sprintf("TOTAL: $%.2f", m.totalPrice))

	return fmt.Sprintf("%s\n%s\n\n%s\n%s\n\n%s",
		title,
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("─────────────────────────────────"),
		cartContent,
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("─────────────────────────────────"),
		fmt.Sprintf("%s\n\n%s", totalLine, help),
	)
}

// ================================
// HELPER FUNCTIONS
// ================================

func addToCart(m Model, product Product) Model {
	// Check if product already in cart
	for i, item := range m.cart {
		if item.product.id == product.id {
			m.cart[i].quantity++
			m.calculateTotal()
			return m
		}
	}

	// Add new item to cart
	m.cart = append(m.cart, CartItem{
		product:  product,
		quantity: 1,
	})

	m.calculateTotal()
	return m
}

func (m *Model) calculateTotal() {
	m.totalPrice = 0
	for _, item := range m.cart {
		m.totalPrice += item.product.price * float64(item.quantity)
	}
}

// ================================
// STYLING
// ================================

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("213")).
		MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Italic(true)
)

// ================================
// MAIN
// ================================

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
