// sample_spec.js created with Cypress
//
// Start writing your Cypress tests below

/// <reference types="Cypress" />

describe('Navigation Menu Tests', () => {
    it('Open Home Page', () => {
        cy.visit('http://localhost:4200/');
    })

    it('Open HowItWorks', () =>{
        cy.visit('http://localhost:4200/how-it-works');
    })

    it('Open Login', () =>{
        cy.visit('http://localhost:4200/login');
    })

    it('Open Pricing', () =>{
        cy.visit('http://localhost:4200/pricing');
    })
  })

  describe('My First Test', () => {
    it('Does not do much!', () => {
      expect(true).to.equal(true)
    })
  })

describe('Sign Up User', () => {
    it('Navigate to Sign Up', () => {
        cy.visit('http://localhost:4200/login');
    })

    // it('Populate Form', () => {
    //     cy.get('input[name="fname"]').type('Siva').should('have.value', "Siva");
    //     cy.get('input[name="lname"]').type('praneeth').should('have.value', "praneeth");
    //     cy.get('input[name="email"]').type('Siva.praneeth@gmail.com').should('have.value', "Siva.praneeth@gmail.com");
    //     cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");

    //     cy.contains('Submit').click();
    //     cy.url().should('include', '/popup-message');

    //     cy.get('h1').should('contain', 'Succesfully Signed-Up');
    //     cy.contains('Continue').click();
    // })
})

 describe('Sign In User', () => {
     it('Navigate to Log In', () => {
        cy.visit('http://localhost:4200/login');
     })

    //  it('Populate Form', () => {
    //      cy.get('input[name="email"]').type("Siva.praneeth@gmail.com").should('have.value', 'Siva.praneeth@gmail.com');
    //      cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");
    //  })

     it('Log In', () => {
         cy.visit('http://localhost:4200/pricing');
     })
  })

