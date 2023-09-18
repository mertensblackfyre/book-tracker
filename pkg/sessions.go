package pkg

import (

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
)

// Create a global variable for the session manager.

var SessionManager *scs.Session

func InitSessions() {
	SessionManager = scs.NewSession()
	SessionManager.Store = memstore.New()
}
