/// <reference types="cypress" />

context('Register', () => {
    beforeEach(() => {
      cy.visit('http://localhost:3000/register')
    })
  
    it('The link to go to login page works', () => {
        // On Register page

        // tests the Login link in Already have an account?
        cy.contains('Login').click()
        cy.url().should('include', '/login')

        // go back, on the Register page
        cy.go('back')
    })

    // it('Require user to fill out all the fields in the Register form', () => {
       
    // })
})