package data

// Package donde se declaran los adaptadores de datos a los puertos del dominio.
// Este paquete contiene las implementaciones concretas de los puertos definidos en el dominio,
// permitiendo que los servicios del dominio interactúen con las fuentes de datos específicas.

/*
Ejemplo de adaptador de datos

type exampleAdapter struct {
	db: database.Connection // o cualquier otro tipo de conexión a la base de datos
}

//implemetamos el puerto definido en el dominio
func (a *exampleAdapter) NewExampleData(id string) ports.ExamplePort {
	return &exampleAdapter{
		db: db,
	}
}

func (a *exampleAdapter) GetByID(id string) (*models.ExampleModel, error) {
	// Implementación de la lógica para obtener un modelo por ID desde la base de datos
	// Utilizando a.db para acceder a la conexión de la base de datos
	// y devolviendo el modelo correspondiente o un error si ocurre.
}
*/
