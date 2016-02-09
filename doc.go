/*
Pacakge pqerror provides constants for the Postgres Error Codes, to be used with the Golang Postgres driver https://godoc.org/github.com/lib/pq

Example Usage

	pqError := err.(*pq.Error)
	
	switch pqError.Code {
	case pqerror.CodeSyntaxError:
		//@TODO
	case pqerror.CodeIntegrityConstraintViolationUniqueViolation:
		//@TODO
	default:
		//@TODO
	}
*/
package pqerror
