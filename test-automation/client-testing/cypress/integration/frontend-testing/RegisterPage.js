/// <reference types="cypress" />

context('Register', () => {
    beforeEach(() => {
      cy.visit('http://localhost:3000/register')
    })
  
    it('The link to go to Login page works', () => {
      // tests the Login link in Already have an account?
      cy.contains('Login').click()
      cy.url().should('include', '/login')
      // go back, on the Register page
      cy.go('back')
    })

    it('Require user to fill out all the fields in the Register form', () => {
      // all fields are empty
      cy.get('[type="submit"]').click() 
      cy.get('[name=first_name]').invoke('prop', 'validationMessage')
      .should('include', 'Please fill out')

      // only the First Name field is filled out
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[type="submit"]').click() 
      cy.get('[name=last_name]').invoke('prop', 'validationMessage')
      .should('include', 'Please fill out')

      // at least 1 field is empty
      cy.get('[name=last_name]').type('Lam') 
      cy.get('[name=email]').type('nhien@email.com') 
      cy.get('[type="submit"]').click() 
      cy.get('[name=password]').invoke('prop', 'validationMessage')
      .should('include', 'Please fill out')
    })

    it('Check valid email address- missing @', () => {
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[name=last_name]').type('Lam') 

      // missing '@'
      cy.get('[name=email]').type('nhienlamemail')
      cy.get('[type="submit"]').click() 
      cy.get('[name=email]').invoke('prop', 'validationMessage')
      .should('include', "missing an '@'")
    })

    it('Check valid email address- missing a part following @', () => {
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[name=last_name]').type('Lam') 

      // missing a part following ‘@’
      cy.get('[name=email]').type('nhienlam@')
      cy.get('[type="submit"]').click() 
      cy.get('[name=email]').invoke('prop', 'validationMessage')
      .should('include', "a part following '@'")
    })

    it('Check valid email address- missing a part followed by @', () => {
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[name=last_name]').type('Lam') 

      // missing a part followed by ‘@’
      cy.get('[name=email]').type('@email')
      cy.get('[type="submit"]').click() 
      cy.get('[name=email]').invoke('prop', 'validationMessage')
      .should('include', "a part followed by '@'")
    })

    it('Register feature works - successfully create an account', () => {
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[name=last_name]').type('Lam') 
      cy.get('[name=email]').type('newemail@sjsu.edu')
      cy.get('[name=password]').type('newpassword')
      cy.get('[type="submit"]').click() 
      // should redirect to login
      cy.url().should('include', '/login')
    })


    it('No redirect because an account with this email address already existed', () => {
      cy.get('[name=first_name]').type('Nhien') 
      cy.get('[name=last_name]').type('Lam') 
      cy.get('[name=email]').type('nhienlam@sjsu.edu')
      cy.get('[name=password]').type('newpassword')
      cy.get('[type="submit"]').click() 
      
      // should stay at Register
      cy.url().should('include', '/register')
      cy.contains('Create your account').should('exist')
    })
})