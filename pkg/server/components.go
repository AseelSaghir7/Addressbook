package server

import (
	"github.com/addressBook/pkg/components/address_book"
)

// register all components(APIs) to server
func (s *Server) registerComponents() {
	s.registerAddressBookComponent()
}

// register address_book component
func (s *Server) registerAddressBookComponent() {

	// new address book component
	comp := address_book.New(s.db)

	s.router.POST("/ab/v1/address/new", comp.CreateAddress())
	s.router.GET("/ab/v1/address/:addressID", comp.GetAddress())
	s.router.DELETE("/ab/v1/address/:addressID", comp.RemoveAddress())
	s.router.GET("/ab/v1/addressBook", comp.GetAddressBook())
	s.router.GET("/ab/v1/addressBook/search", comp.GetSearchedAddressBook())
}
