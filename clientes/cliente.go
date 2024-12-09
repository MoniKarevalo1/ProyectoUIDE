package clientes

// Definimos la estructura Cliente
type Cliente struct {
	ID     int
	Nombre string
	Email  string
}

// Método público para crear un nuevo Cliente
func NuevoCliente(id int, nombre string, email string) *Cliente {
	return &Cliente{ID: id, Nombre: nombre, Email: email}
}

// Métodos públicos para obtener los detalles del cliente
func (c *Cliente) ObtenerID() int {
	return c.ID
}

func (c *Cliente) ObtenerNombre() string {
	return c.Nombre
}

func (c *Cliente) ObtenerEmail() string {
	return c.Email
}
