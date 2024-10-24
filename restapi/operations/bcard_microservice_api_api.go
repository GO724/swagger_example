// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"swagger_example/restapi/operations/admins"
)

// NewBcardMicroserviceAPIAPI creates a new BcardMicroserviceAPI instance
func NewBcardMicroserviceAPIAPI(spec *loads.Document) *BcardMicroserviceAPIAPI {
	return &BcardMicroserviceAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AdminsGetBanksHandler: admins.GetBanksHandlerFunc(func(params admins.GetBanksParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetBanks has not yet been implemented")
		}),
		AdminsGetCardsHandler: admins.GetCardsHandlerFunc(func(params admins.GetCardsParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetCards has not yet been implemented")
		}),
		AdminsGetPersonsHandler: admins.GetPersonsHandlerFunc(func(params admins.GetPersonsParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetPersons has not yet been implemented")
		}),
		AdminsAddCardHandler: admins.AddCardHandlerFunc(func(params admins.AddCardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.AddCard has not yet been implemented")
		}),
		AdminsDelBankHandler: admins.DelBankHandlerFunc(func(params admins.DelBankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelBank has not yet been implemented")
		}),
		AdminsDelPersonHandler: admins.DelPersonHandlerFunc(func(params admins.DelPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelPerson has not yet been implemented")
		}),
		AdminsDelСardHandler: admins.DelСardHandlerFunc(func(params admins.DelСardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.DelСard has not yet been implemented")
		}),
		AdminsGetBankHandler: admins.GetBankHandlerFunc(func(params admins.GetBankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetBank has not yet been implemented")
		}),
		AdminsGetCardHandler: admins.GetCardHandlerFunc(func(params admins.GetCardParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetCard has not yet been implemented")
		}),
		AdminsGetPersonHandler: admins.GetPersonHandlerFunc(func(params admins.GetPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.GetPerson has not yet been implemented")
		}),
		AdminsIsValidHandler: admins.IsValidHandlerFunc(func(params admins.IsValidParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.IsValid has not yet been implemented")
		}),
		AdminsNewBankHandler: admins.NewBankHandlerFunc(func(params admins.NewBankParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.NewBank has not yet been implemented")
		}),
		AdminsNewPersonHandler: admins.NewPersonHandlerFunc(func(params admins.NewPersonParams) middleware.Responder {
			return middleware.NotImplemented("operation admins.NewPerson has not yet been implemented")
		}),
	}
}

/*
BcardMicroserviceAPIAPI Sample OpenAPI2.0 spec for demo microservice bcard:
Check valid paycard (exist & not expiried)
CRUD Card, Person, Bank & check valid
*/
type BcardMicroserviceAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AdminsGetBanksHandler sets the operation handler for the get banks operation
	AdminsGetBanksHandler admins.GetBanksHandler
	// AdminsGetCardsHandler sets the operation handler for the get cards operation
	AdminsGetCardsHandler admins.GetCardsHandler
	// AdminsGetPersonsHandler sets the operation handler for the get persons operation
	AdminsGetPersonsHandler admins.GetPersonsHandler
	// AdminsAddCardHandler sets the operation handler for the add card operation
	AdminsAddCardHandler admins.AddCardHandler
	// AdminsDelBankHandler sets the operation handler for the del bank operation
	AdminsDelBankHandler admins.DelBankHandler
	// AdminsDelPersonHandler sets the operation handler for the del person operation
	AdminsDelPersonHandler admins.DelPersonHandler
	// AdminsDelСardHandler sets the operation handler for the del сard operation
	AdminsDelСardHandler admins.DelСardHandler
	// AdminsGetBankHandler sets the operation handler for the get bank operation
	AdminsGetBankHandler admins.GetBankHandler
	// AdminsGetCardHandler sets the operation handler for the get card operation
	AdminsGetCardHandler admins.GetCardHandler
	// AdminsGetPersonHandler sets the operation handler for the get person operation
	AdminsGetPersonHandler admins.GetPersonHandler
	// AdminsIsValidHandler sets the operation handler for the is valid operation
	AdminsIsValidHandler admins.IsValidHandler
	// AdminsNewBankHandler sets the operation handler for the new bank operation
	AdminsNewBankHandler admins.NewBankHandler
	// AdminsNewPersonHandler sets the operation handler for the new person operation
	AdminsNewPersonHandler admins.NewPersonHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BcardMicroserviceAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BcardMicroserviceAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BcardMicroserviceAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BcardMicroserviceAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BcardMicroserviceAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BcardMicroserviceAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BcardMicroserviceAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BcardMicroserviceAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BcardMicroserviceAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BcardMicroserviceAPIAPI
func (o *BcardMicroserviceAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AdminsGetBanksHandler == nil {
		unregistered = append(unregistered, "admins.GetBanksHandler")
	}
	if o.AdminsGetCardsHandler == nil {
		unregistered = append(unregistered, "admins.GetCardsHandler")
	}
	if o.AdminsGetPersonsHandler == nil {
		unregistered = append(unregistered, "admins.GetPersonsHandler")
	}
	if o.AdminsAddCardHandler == nil {
		unregistered = append(unregistered, "admins.AddCardHandler")
	}
	if o.AdminsDelBankHandler == nil {
		unregistered = append(unregistered, "admins.DelBankHandler")
	}
	if o.AdminsDelPersonHandler == nil {
		unregistered = append(unregistered, "admins.DelPersonHandler")
	}
	if o.AdminsDelСardHandler == nil {
		unregistered = append(unregistered, "admins.DelСardHandler")
	}
	if o.AdminsGetBankHandler == nil {
		unregistered = append(unregistered, "admins.GetBankHandler")
	}
	if o.AdminsGetCardHandler == nil {
		unregistered = append(unregistered, "admins.GetCardHandler")
	}
	if o.AdminsGetPersonHandler == nil {
		unregistered = append(unregistered, "admins.GetPersonHandler")
	}
	if o.AdminsIsValidHandler == nil {
		unregistered = append(unregistered, "admins.IsValidHandler")
	}
	if o.AdminsNewBankHandler == nil {
		unregistered = append(unregistered, "admins.NewBankHandler")
	}
	if o.AdminsNewPersonHandler == nil {
		unregistered = append(unregistered, "admins.NewPersonHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BcardMicroserviceAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BcardMicroserviceAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BcardMicroserviceAPIAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BcardMicroserviceAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BcardMicroserviceAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BcardMicroserviceAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the bcard microservice API API
func (o *BcardMicroserviceAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BcardMicroserviceAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/banks"] = admins.NewGetBanks(o.context, o.AdminsGetBanksHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cards"] = admins.NewGetCards(o.context, o.AdminsGetCardsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/persons"] = admins.NewGetPersons(o.context, o.AdminsGetPersonsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cards"] = admins.NewAddCard(o.context, o.AdminsAddCardHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/banks/{inn}"] = admins.NewDelBank(o.context, o.AdminsDelBankHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/persons/{inn}"] = admins.NewDelPerson(o.context, o.AdminsDelPersonHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/cards/{inn}"] = admins.NewDelСard(o.context, o.AdminsDelСardHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/banks/{inn}"] = admins.NewGetBank(o.context, o.AdminsGetBankHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cards/{inn}"] = admins.NewGetCard(o.context, o.AdminsGetCardHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/persons/{inn}"] = admins.NewGetPerson(o.context, o.AdminsGetPersonHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/card/validator/{id}"] = admins.NewIsValid(o.context, o.AdminsIsValidHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/banks"] = admins.NewNewBank(o.context, o.AdminsNewBankHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/persons"] = admins.NewNewPerson(o.context, o.AdminsNewPersonHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BcardMicroserviceAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BcardMicroserviceAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BcardMicroserviceAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BcardMicroserviceAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BcardMicroserviceAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}