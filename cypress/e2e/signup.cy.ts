describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/signup')
  })

  it('displays error message for required fields that are empty', () => {
    cy.get('input[id=firstname]').click()
    cy.get('input[id=lastname]').click()
    cy.get('input[id=username]').click()
    cy.get('input[id=password]').click()
    cy.get('#card').click()
    
    cy.get('#firstError').should('be.visible').contains("You must enter a first name")
    cy.get('#lastError').should('be.visible').contains("You must enter a last name")
    cy.get('#userError').should('be.visible').contains("You must enter a username")
    cy.get('#passError').should('be.visible').contains("You must enter a password")
  })


  it('displays error message for invalid name and username format', () => {
    cy.get('input[id=firstname]').type('this is my first name')
    cy.get('input[id=username]').type('invalid username')
    cy.get('#card').click()
    cy.get('#firstError').should('be.visible').contains("Please enter a valid first name")
    cy.get('#userError').should('be.visible').contains("Not a valid username")
  })
  

  it ('fill out form and enter', function() {
    cy.get('input[id=firstname]').type('new')
    cy.get('input[id=lastname]').type('user')
    cy.get('input[id=username]').type('newuser')
    cy.get('input[id=password]').type('password')
    cy.get('.login-signup-page-button').first().click('center')
  })
})