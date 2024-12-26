package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/microsoft/go-mssqldb" // Driver para SQL Server
)

// Estructura para representar un libro
type Libro struct {
	LibroID   int     `json:"libro_id"`
	Titulo    string  `json:"titulo"`
	Autor     string  `json:"autor"`
	Categoria string  `json:"categoria"`
	Año       int     `json:"año"`
	Prestado  bool    `json:"prestado"`
	Precio    float64 `json:"precio"`
}

// Estructura para representar un usuario
type Usuario struct {
	UsuarioID int    `json:"usuario_id"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Telefono  string `json:"telefono"`
}

var db *sql.DB // Variable global para la conexión a la base de datos

// Conexión a la base de datos
func conectarBaseDeDatos() error {
	var err error
	connStr := "sqlserver://localhost:1433?database=BibliotecaDB&trusted_connection=true"
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error al verificar la conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
	return nil
}

// Handler de bienvenida
func bienvenida(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenido al sistema de biblioteca virtual"))
}

// Obtener todos los libros
func obtenerLibros(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT LibroID, Titulo, Autor, Categoria, Año, Prestado, Precio FROM Libros")
	if err != nil {
		http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
		log.Println("Error al obtener libros:", err)
		return
	}
	defer rows.Close()

	var libros []Libro
	for rows.Next() {
		var libro Libro
		err := rows.Scan(&libro.LibroID, &libro.Titulo, &libro.Autor, &libro.Categoria, &libro.Año, &libro.Prestado, &libro.Precio)
		if err != nil {
			http.Error(w, "Error al leer los datos de los libros", http.StatusInternalServerError)
			log.Println("Error al leer los datos de los libros:", err)
			return
		}
		libros = append(libros, libro)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

// Crear un nuevo libro
func crearLibro(w http.ResponseWriter, r *http.Request) {
	var nuevoLibro Libro
	err := json.NewDecoder(r.Body).Decode(&nuevoLibro)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Asegúrate de que el libro tenga todos los datos requeridos
	if nuevoLibro.Titulo == "" || nuevoLibro.Autor == "" || nuevoLibro.Categoria == "" || nuevoLibro.Año == 0 || nuevoLibro.Precio == 0 {
		http.Error(w, "Faltan campos requeridos", http.StatusBadRequest)
		return
	}

	// Si 'Prestado' no se ha especificado, lo asignamos a 'FALSE' (o '0')
	if nuevoLibro.Prestado == "" {
		nuevoLibro.Prestado = "FALSE"
	}

	// Consulta SQL para insertar el libro en la tabla 'Libros'
	query := `INSERT INTO Libros (Titulo, Autor, Categoria, Año, Prestado, precio) 
              VALUES (@Titulo, @Autor, @Categoria, @Año, @Prestado, @precio)`

	// Ejecutar la consulta con los parámetros adecuados
	_, err = db.Exec(query,
		sql.Named("Titulo", nuevoLibro.Titulo),
		sql.Named("Autor", nuevoLibro.Autor),
		sql.Named("Categoria", nuevoLibro.Categoria),
		sql.Named("Año", nuevoLibro.Año),
		sql.Named("Prestado", nuevoLibro.Prestado),
		sql.Named("Precio", nuevoLibro.Precio),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear el libro: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Libro creado con éxito"))
}

// Actualizar un libro
func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var libro Libro
	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		http.Error(w, "Error al decodificar los datos", http.StatusBadRequest)
		return
	}

	// Consulta SQL corregida
	query := `UPDATE libros SET titulo = @titulo, autor = @autor, categoria = @categoria, año = @año, prestado = @prestado, precio = @precio WHERE LibroID = @id`

	// Ejecutar la consulta con los parámetros adecuados
	_, err = db.Exec(query, sql.Named("id", id), sql.Named("titulo", libro.Titulo), sql.Named("autor", libro.Autor),
		sql.Named("categoria", libro.Categoria), sql.Named("año", libro.Año), sql.Named("prestado", libro.Prestado), sql.Named("precio", libro.Precio))

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el libro: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Libro actualizado exitosamente")
}

// Eliminar un libro
func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Consulta SQL corregida
	query := `DELETE FROM libros WHERE LibroID = @id`

	// Ejecutar la consulta con el parámetro adecuado
	_, err := db.Exec(query, sql.Named("id", id))

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el libro: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Libro eliminado exitosamente")
}

// Agregar un nuevo usuario
func agregarUsuario(w http.ResponseWriter, r *http.Request) {
	var nuevoUsuario Usuario
	err := json.NewDecoder(r.Body).Decode(&nuevoUsuario)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Usando la sintaxis correcta para SQL Server
	query := `INSERT INTO Usuarios (Nombre, Email, Telefono) VALUES (@Nombre, @Email, @Telefono)`
	_, err = db.Exec(query, sql.Named("Nombre", nuevoUsuario.Nombre), sql.Named("Email", nuevoUsuario.Email), sql.Named("Telefono", nuevoUsuario.Telefono))
	if err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		log.Println("Error al crear el usuario:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuario creado con éxito"))
}

// Eliminar un usuario
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.URL.Query().Get("id")
	if usuarioID == "" {
		http.Error(w, "Falta el parámetro 'id'", http.StatusBadRequest)
		return
	}

	// Verificar qué valor tiene el parámetro 'id' y eliminar caracteres no deseados
	usuarioID = strings.TrimSpace(usuarioID)
	log.Printf("Valor recibido para 'id': %s", usuarioID)

	// Asegurarse de que el 'id' sea un número entero
	id, err := strconv.Atoi(usuarioID)
	if err != nil {
		http.Error(w, "El ID del usuario debe ser un número entero", http.StatusBadRequest)
		log.Printf("Error al convertir el ID del usuario: %v", err)
		return
	}

	// Usando la sintaxis correcta para SQL Server
	query := `DELETE FROM Usuarios WHERE UsuarioID = @UsuarioID`
	_, err = db.Exec(query, sql.Named("UsuarioID", id))
	if err != nil {
		http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
		log.Println("Error al eliminar el usuario:", err)
		return
	}

	w.Write([]byte("Usuario eliminado con éxito"))
}

// Configurar y ejecutar el servidor
func main() {
	// Conectar a la base de datos
	err := conectarBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}

	// Configurar rutas
	http.HandleFunc("/", bienvenida)
	http.HandleFunc("/libros", obtenerLibros)
	http.HandleFunc("/libro/crear", crearLibro)
	http.HandleFunc("/libro/actualizar", actualizarLibro)
	http.HandleFunc("/libro/eliminar", eliminarLibro)
	http.HandleFunc("/usuario/agregar", agregarUsuario)
	http.HandleFunc("/usuario/eliminar", eliminarUsuario)

	// Iniciar el servidor
	fmt.Println("Servidor escuchando en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
