// sample_spec.js created with Cypress
//
// Start writing your Cypress tests below

describe('Navigation Menu Tests', () => {
    it('Open Home Page', () => {
        cy.visit('http://localhost:4200/');

        cy.contains('Home').click();
        cy.url().should('include', '/');
    })

    it('Open HowItWorks', () =>{
        cy.contains('How It Works').click();
        cy.url().should('include', '/how-it-works');
    })

    it('Open Login', () =>{
        cy.contains('Sign In').click();
        cy.url().should('include', '/login');
    })

    it('Open Pricing', () =>{
        cy.contains('Sign In').click();
        cy.url().should('include', '/pricing');
    })
  })

describe('Sign Up User', () => {
    it('Navigate to Sign Up', () => {
        cy.contains('Sign In').click();
        cy.url().should('include', '/login');
    })

    it('Populate Form', () => {
        cy.get('input[name="fname"]').type('Siva').should('have.value', "Siva");
        cy.get('input[name="lname"]').type('praneeth').should('have.value', "praneeth");
        cy.get('input[name="email"]').type('Siva.praneeth@gmail.com').should('have.value', "Siva.praneeth@gmail.com");
        cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");

        cy.contains('Submit').click();
        cy.url().should('include', '/popup-message');

        cy.get('h1').should('contain', 'Succesfully Signed-Up');
        cy.contains('Continue').click();
    })
})
describe('Sign In User', () => {
    it('Navigate to Log In', () => {
        cy.contains('Log In').click();
        cy.url().should('include', '/login')
    })

    it('Populate Form', () => {
        cy.get('input[name="email"]').type("Siva.praneeth@gmail.com").should('have.value', 'Siva.praneeth@gmail.com');
        cy.get('input[name="password"]').type('Qwerty123').should('have.value', "Qwerty123");
    })

    it('Log In', () => {
        cy.get('button[name="loginbutton"]').click();
    })
})