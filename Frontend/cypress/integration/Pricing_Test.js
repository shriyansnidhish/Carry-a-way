// sample_spec.js created with Cypress
//
// Start writing your Cypress tests below

describe('Navigation Menu Tests', () => {
    it('Open Pricing Page', () => {
        cy.visit('http://localhost:4200/pricing');

        cy.contains('Pricing').click();
        cy.url().should('include', '/pricing');
    })

    it('Open Bag calculator', () =>{
        cy.contains('Bag calculator').click();
        cy.url().should('include', '/bag-calculator');
    })

    it('Open Login', () =>{
        cy.contains('Sign In').click();
        cy.url().should('include', '/login');
    })

    it('Open Help', () =>{
        cy.contains('Help').click();
        cy.url().should('include', '/help');
    })
  })

describe('Pricing Page', () => {
    it('Navigate to Pricing', () => {
        cy.contains('Pricing').click();
        cy.url().should('include', '/pricing');
    })

    it('Populate Form', () => {
        cy.get('input[name="from"]').type('Gainesville, FL, US').should('have.value', "Gainesville, FL, US");
        cy.get('input[name="to"]').type('New York, NY, US').should('have.value', "New York, NY, US");
        cy.get('input[name="date"]').type('02/06/2022').should('have.value', "02/06/2022");
        // cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");

        cy.contains('Submit').click();
        cy.url().should('include', '/popup-message');

        cy.get('h1').should('contain', 'Succesfully Signed-Up');
        cy.contains('Continue').click();
    })
})
//  describe('Pricing', () => {
//      it('Navigate to Pricing', () => {
//          cy.contains('Pricing').click();
//          cy.url().should('include', '/pricing')
//      })

//      it('Populate Form', () => {
//          cy.get('input[name="email"]').type("Siva.praneeth@gmail.com").should('have.value', 'Siva.praneeth@gmail.com');
//          cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");
//      })

//      it('Log In', () => {
//          cy.get('button[name="loginbutton"]').click();
//      })
//   })

