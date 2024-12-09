package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "ecommerce_user"
	dbPassword = "root123"
	dbName     = "ecommerce"
)

func main() {
	// Conectarse a la base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer una conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	// Crear tablas dinámicamente
	crearTablas(db)
}

func crearTablas(db *sql.DB) {
	crearTablaCategorias := `
    CREATE TABLE IF NOT EXISTS categorias (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL
    );`

	crearTablaUsuarios := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id INT AUTO_INCREMENT PRIMARY KEY,
        tipo ENUM('cliente', 'proveedor', 'administrador') NOT NULL,
        nombre VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL
    );`

	crearTablaProductos := `
    CREATE TABLE IF NOT EXISTS productos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        precio DECIMAL(10, 2) NOT NULL,
        stock INT NOT NULL,
        categoria_id INT,
        FOREIGN KEY (categoria_id) REFERENCES categorias(id)
    );`

	crearTablaPedidos := `
    CREATE TABLE IF NOT EXISTS pedidos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        usuario_id INT,
        fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        total DECIMAL(10, 2) NOT NULL,
        estado ENUM('pendiente', 'completado', 'cancelado') NOT NULL,
        FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
    );`

	crearTablaPagos := `
    CREATE TABLE IF NOT EXISTS pagos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pedido_id INT,
        fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        monto DECIMAL(10, 2) NOT NULL,
        metodo ENUM('tarjeta', 'paypal', 'transferencia') NOT NULL,
        FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
    );`

	crearTablaEnvios := `
    CREATE TABLE IF NOT EXISTS envios (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pedido_id INT,
        direccion VARCHAR(255) NOT NULL,
        fecha_envio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        fecha_entrega TIMESTAMP,
        FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
    );`

	crearTablaCarritoCompras := `
    CREATE TABLE IF NOT EXISTS carrito_de_compras (
        id INT AUTO_INCREMENT PRIMARY KEY,
        usuario_id INT,
        producto_id INT,
        cantidad INT NOT NULL,
        FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
        FOREIGN KEY (producto_id) REFERENCES productos(id)
    );`

	crearTablaInventario := `
    CREATE TABLE IF NOT EXISTS inventario (
        id INT AUTO_INCREMENT PRIMARY KEY,
        producto_id INT,
        cantidad INT NOT NULL,
        fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (producto_id) REFERENCES productos(id)
    );`

	tablas := []string{crearTablaCategorias, crearTablaUsuarios, crearTablaProductos, crearTablaPedidos, crearTablaPagos, crearTablaEnvios, crearTablaCarritoCompras, crearTablaInventario}

	for _, tabla := range tablas {
		_, err := db.Exec(tabla)
		if err != nil {
			log.Fatalf("No se pudo crear la tabla: %v", err)
		}
	}
	fmt.Println("Tablas creadas exitosamente")
}
