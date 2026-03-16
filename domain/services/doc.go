package services

// Package services contiene los servicios propios del dominio, donde se implementa la lógica de negocio de la aplicación.
// Aquí se definen las operaciones que pueden realizarse sobre los modelos del dominio,
// interactuando con los puertos para acceder a los datos necesarios.
// Los servicios son responsables de coordinar las acciones entre los modelos y los adaptadores de datos,
// asegurando que las reglas de negocio se apliquen correctamente.

/*
Ejemplo de servicio
type ExampleService struct {
	repo ports.ExamplePort
	anotherDependency AnotherDependencyType
}

func (s *ExampleService) NewService(repo ports.ExamplePort, another....) *ExampleService {
	return &ExampleService{
		repo: repo,
		anotherDependency: anotherDependency,
	}
}

func (s *ExampleService) SomeBusinessLogicMethod(param1 Type1, param2 Type2) (ReturnType, error) {
	// Implementación de la lógica de negocio
	// Utilizando s.repo para acceder a los datos necesarios
	// y aplicando las reglas del dominio.
}
*/
