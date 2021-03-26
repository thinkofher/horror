package horror

import (
	"errors"
	"net/http"
)

// AdapterBuilder holds arguments for creating new horror.Adapter.
type AdapterBuilder struct {
	// BeforeError is the function that will be called by Adapter.WithError
	// method before using ServerHTTP method of horror.Error interface.
	BeforeError func(error, http.ResponseWriter, *http.Request)

	// AfterError is the function that will be called by Adapter.WithError
	// method after using ServerHTTP method of horror.Error interface.
	AfterError      func(error, http.ResponseWriter, *http.Request)
	InternalHandler func(error, http.ResponseWriter, *http.Request)
	WrapInternal    bool
}

// Adapter creates a custom way for adapting the horror.Handler to the
// standard http.Handler implementation with it's WithError method.
//
// You should always use New factory function with AdapterBuilder structure
// for creating your adapters.
type Adapter struct {
	beforeErr       func(error, http.ResponseWriter, *http.Request)
	afterErr        func(error, http.ResponseWriter, *http.Request)
	internalHandler func(error, http.ResponseWriter, *http.Request)
	wrapInternal    bool
}

// WithError wraps horror.Handler, adapts it and returns http.Handler that
// can be used with other APIs that relies on go standard library.
func (a Adapter) WithError(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h.ServeHTTP(w, r); err != nil {
			var horrorError Error
			if errors.As(err, &horrorError) {
				a.beforeErr(err, w, r)
				horrorError.ServeHTTP(w, r)
				a.afterErr(err, w, r)
				return
			} else {
				if a.wrapInternal {
					a.beforeErr(err, w, r)
				}
				a.internalHandler(err, w, r)
				if a.wrapInternal {
					a.afterErr(err, w, r)
				}
				return
			}
		}
	})
}

// doNothing is function that literally does nothing.
func doNothing(error, http.ResponseWriter, *http.Request) {}

// NewAdapter returns pointer to Adapter that is safe to use and will
// not panic during runtime because of nil pointers.
//
// Using NewAdapter is only preferable way to create horror.Adapter.
func NewAdapter(b *AdapterBuilder) (a *Adapter) {
	// allocate default adapter
	a = &Adapter{
		beforeErr:       doNothing,
		afterErr:        doNothing,
		internalHandler: defaultInternalServerErrorHandler,
		wrapInternal:    b.WrapInternal,
	}

	// check for nil arguments
	if b.AfterError != nil {
		a.afterErr = b.AfterError
	}
	if b.BeforeError != nil {
		a.beforeErr = b.BeforeError
	}
	if b.InternalHandler != nil {
		a.internalHandler = b.InternalHandler
	}

	// return adapter that is safe to use and will
	// not panic because of nil function pointers
	return
}
