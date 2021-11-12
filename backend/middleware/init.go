package middleware

import "github.com/justinas/alice"

var Mdw alice.Chain

func init() {
	Mdw = alice.New(Auth)
}
