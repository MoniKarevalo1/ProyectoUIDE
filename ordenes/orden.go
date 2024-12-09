package ordenes

import (
	"ProyectoUIDE/clientes"
	"ProyectoUIDE/productos"
	"fmt"
)

// Definimos la estructura Orden
type Orden struct {
	ID        int
	Cliente   *clientes.Cliente // Utiliza la estructura Cliente del paquete clientes
	Productos map[int]*productos.Producto
}

// Método público para crear una nueva Orden
func NuevaOrden(id int, cliente *clientes.Cliente) *Orden {
	return &Orden{ID: id, Cliente: cliente, Productos: make(map[int]*productos.Producto)}
}

// Método público para agregar un producto a la orden
func (o *Orden) AgregarProducto(p *productos.Producto, cantidad int) {
	if p.ObtenerStock() >= cantidad {
		o.Productos[p.ObtenerID()] = p
		p.ActualizarStock(p.ObtenerStock() - cantidad)
	} else {
		fmt.Println("No hay suficiente stock para el producto:", p.ObtenerNombre())
	}
}

// Método público para mostrar los productos en la orden
func (o *Orden) MostrarProductos() {
	fmt.Printf("Orden ID: %d\nCliente: %s\n", o.ID, o.Cliente.ObtenerNombre())
	fmt.Println("Productos en la orden:")
	for _, producto := range o.Productos {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock restante: %d\n",
			producto.ObtenerID(), producto.ObtenerNombre(), producto.ObtenerPrecio(), producto.ObtenerStock())
	}
}
