package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "ecommerce_user"
	dbPassword = "root123"
	dbName     = "ecommerce"
)

var db *sql.DB

func main() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer una conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	router := gin.Default()

	// Configurar CORS middleware
	router.Use(cors.Default())

	// Rutas para insertar datos
	router.POST("/categorias", insertarCategoria)
	router.POST("/usuarios", insertarUsuario)
	router.POST("/productos", insertarProducto)

	// Rutas para actualizar datos
	router.PUT("/categorias/:id", actualizarCategoria)
	router.PUT("/usuarios/:id", actualizarUsuario)
	router.PUT("/productos/:id", actualizarProducto)

	// Rutas para eliminar datos
	router.DELETE("/categorias/:id", eliminarCategoria)
	router.DELETE("/usuarios/:id", eliminarUsuario)
	router.DELETE("/productos/:id", eliminarProducto)

	// Rutas para obtener datos
	router.GET("/categorias", obtenerCategorias)
	router.GET("/usuarios", obtenerUsuarios)
	router.GET("/productos", obtenerProductos)

	// Ejecutar el servidor
	router.Run(":8080")
}

type Categoria struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type Usuario struct {
	ID       int    `json:"id"`
	Tipo     string `json:"tipo"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	CategoriaID int     `json:"categoria_id"`
}

func insertarCategoria(c *gin.Context) {
	var nuevaCategoria Categoria
	if err := c.ShouldBindJSON(&nuevaCategoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO categorias (nombre) VALUES (?)", nuevaCategoria.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoría creada exitosamente"})
}

func insertarUsuario(c *gin.Context) {
	var nuevoUsuario Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO usuarios (tipo, nombre, email, password) VALUES (?, ?, ?, ?)", nuevoUsuario.Tipo, nuevoUsuario.Nombre, nuevoUsuario.Email, nuevoUsuario.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente"})
}

// Ejemplo de código en Go para verificar que la categoría existe
func insertarProducto(c *gin.Context) {
	var producto Producto
	if err := c.BindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el cuerpo de la solicitud"})
		return
	}

	// Verificar si la categoría existe
	var idCategoria int
	err := db.QueryRow("SELECT id FROM categorias WHERE id = ?", producto.CategoriaID).Scan(&idCategoria)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Categoría no existe"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al verificar la categoría: %v", err)})
		}
		return
	}

	// Insertar el producto si la categoría es válida
	query := `INSERT INTO productos (nombre, precio, stock, categoria_id) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, producto.Nombre, producto.Precio, producto.Stock, producto.CategoriaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al insertar el producto: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto insertado exitosamente"})
}

func actualizarCategoria(c *gin.Context) {
	id := c.Param("id")
	var categoria Categoria

	// Parsear JSON del cuerpo de la solicitud
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ejecutar consulta de actualización
	_, err := db.Exec("UPDATE categorias SET nombre = ? WHERE id = ?", categoria.Nombre, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoría actualizada correctamente"})
}

func actualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	var usuario Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE usuarios SET tipo = ?, nombre = ?, email = ?, password = ? WHERE id = ?", usuario.Tipo, usuario.Nombre, usuario.Email, usuario.Password, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}

func actualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE productos SET nombre = ?, precio = ?, stock = ?, categoria_id = ? WHERE id = ?", producto.Nombre, producto.Precio, producto.Stock, producto.CategoriaID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}

func obtenerCategorias(c *gin.Context) {
	rows, err := db.Query("SELECT id, nombre FROM categorias")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var categoria Categoria
		if err := rows.Scan(&categoria.ID, &categoria.Nombre); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categorias = append(categorias, categoria)
	}

	c.JSON(http.StatusOK, categorias)
}

func obtenerUsuarios(c *gin.Context) {
	rows, err := db.Query("SELECT id, tipo, nombre, email, password FROM usuarios")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Tipo, &usuario.Nombre, &usuario.Email, &usuario.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	c.JSON(http.StatusOK, usuarios)
}

func obtenerProductos(c *gin.Context) {
	rows, err := db.Query("SELECT id, nombre, precio, stock, categoria_id FROM productos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var producto Producto
		if err := rows.Scan(&producto.ID, &producto.Nombre, &producto.Precio, &producto.Stock, &producto.CategoriaID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		productos = append(productos, producto)
	}

	c.JSON(http.StatusOK, productos)
}

func eliminarCategoria(c *gin.Context) {
	id := c.Param("id")

	// Ejecutar consulta de eliminación
	_, err := db.Exec("DELETE FROM categorias WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada correctamente"})
}

func eliminarUsuario(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}

func eliminarProducto(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM productos WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
