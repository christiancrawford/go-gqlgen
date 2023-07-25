package graph

import (
	"context"

	"github.com/christiancrawford/go-gqlgen/graph/common"
	"github.com/christiancrawford/go-gqlgen/graph/customTypes"
	"github.com/christiancrawford/go-gqlgen/graph/generated"
)

// This defines the resolver function for the CreateTodo mutation.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes.Todo, error) {

	// It gets the database context and creates a new Todo object with the provided text and default false done value.
	context := common.GetContext(ctx)

	// It saves the new Todo to the database using the Create method.
	todo := &customTypes.Todo{
		Text: text,
		Done: false,
	}

	// If there is an error, it returns the error.
	err := context.Database.Create(&todo).Error
	if err != nil {
		return nil, err
	}

	// Otherwise, it returns the newly created Todo object.
	return todo, nil
}

// This code defines the resolver function for the UpdateTodo mutation.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input customTypes.TodoInput) (*customTypes.Todo, error) {

	// It gets the database context and creates a Todo object from the input data.
	context := common.GetContext(ctx)

	// It saves the Todo to the database using the Save method, which performs an update if the record already exists.
	todo := &customTypes.Todo{
		ID:   input.ID,
		Text: input.Text,
		Done: input.Done,
	}

	// If there is an error saving, it returns the error.
	err := context.Database.Save(&todo).Error
	if err != nil {
		return nil, err
	}

	// Otherwise, it returns the updated Todo object.
	return todo, nil
}

// This defines the resolver function for the DeleteTodo mutation.
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (*customTypes.Todo, error) {

	// It gets the database context and finds the Todo to delete by ID using the Where clause.
	context := common.GetContext(ctx)

	// It deletes the Todo from the database using the Delete method.
	var todo *customTypes.Todo
	err := context.Database.Where("id = ?", todoID).Delete(&todo).Error

	// If there is an error deleting, it returns the error.
	if err != nil {
		return nil, err
	}

	// Otherwise, it returns the deleted Todo object.
	return todo, nil
}

// This defines the resolver function for the GetTodos query to retrieve all todos:
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes.Todo, error) {

	// It gets the database context
	context := common.GetContext(ctx)

	// It declares a todos slice to hold the results
	var todos []*customTypes.Todo

	// It queries the database using Find() to get all Todo records
	err := context.Database.Find(&todos).Error

	// If there is an error, it returns the error
	if err != nil {
		return nil, err
	}

	// Otherwise, it returns the todos slice
	return todos, nil
}

// This defines the resolver function for the GetTodo query to retrieve a single todo by ID:
func (r *queryResolver) GetTodo(ctx context.Context, todoID int) (*customTypes.Todo, error) {

	// It gets the database context
	context := common.GetContext(ctx)

	// Declares a todo pointer to hold the result
	var todo *customTypes.Todo

	// Queries the database by ID using Where and Find
	err := context.Database.Where("id = ?", todoID).Find(&todo).Error

	// If there is an error, returns the error
	if err != nil {
		return nil, err
	}

	// Otherwise returns the found todo
	return todo, nil
}

// This code is defining the Mutation resolver method on the Resolver struct.
func (r *Resolver) Mutation() generated.MutationResolver {

	// It returns a pointer to a mutationResolver struct, passing the Resolver struct pointer as a field.
	return &mutationResolver{r}

}

// This code defines the Query resolver method on the Resolver struct.
func (r *Resolver) Query() generated.QueryResolver {

	// It returns a pointer to a queryResolver struct, passing the Resolver struct pointer as a field.
	return &queryResolver{r}

}

// This mutationResolver struct embeds the original Resolver, so it has access to all the fields and methods on Resolver.
type mutationResolver struct{ *Resolver }

// The queryResolver struct embeds the original Resolver, so it has access to all the fields and methods on Resolver.
type queryResolver struct{ *Resolver }
