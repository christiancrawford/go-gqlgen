package main

import (
	// The github.com/99designs/gqlgen package contains the core gqlgen code generation logic.
	_ "github.com/99designs/gqlgen"

	// 	The github.com/99designs/gqlgen/graphql/introspection package enables GraphQL introspection
	//	queries, which allow querying the schema for metadata.

	_ "github.com/99designs/gqlgen/graphql/introspection"
)

// By importing these packages, the gqlgen code generator will have access to the necessary code
// to generate a GraphQL server based on the schema. The imports are prefixed with _ to indicate
// they are only needed for side effects and their APIs are not directly used in this file.

// So in summary, this imports gqlgen packages required for schema-based code generation.
