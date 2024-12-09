package productos

// Definimos la estructura Producto
type Producto struct {
	id     int
	nombre string
	precio float64
	stock  int
}

// Método público para crear un nuevo Producto
func NuevoProducto(id int, nombre string, precio float64, stock int) *Producto {
	return &Producto{id: id, nombre: nombre, precio: precio, stock: stock}
}

// Métodos públicos para obtener los detalles del producto
func (p *Producto) ObtenerID() int {
	return p.id
}

func (p *Producto) ObtenerNombre() string {
	return p.nombre
}

func (p *Producto) ObtenerPrecio() float64 {
	return p.precio
}

func (p *Producto) ObtenerStock() int {
	return p.stock
}

// Método público para actualizar el stock
func (p *Producto) ActualizarStock(nuevoStock int) {
	p.stock = nuevoStock
}
