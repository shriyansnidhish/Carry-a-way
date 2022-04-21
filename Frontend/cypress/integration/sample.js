/// <reference types="Cypress" />

describe('Carry-A-Way', () => {
    it('should display the app name on the home page', () => {
      cy.visit('/'); // go to the home page
  
      // get the rocket element and verify that the app name is in it
      cy.get('.a')
        .should('contain.text', 'Pricing');
    });
  });