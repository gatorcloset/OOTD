describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1372, 768)
    cy.visit('http://localhost:4200/login')
  })

  it('visits login page', () => {
    cy.visit('http://localhost:4200/login')
  })

  it('displays error message for required field that are empty', () => {
    cy.get('#user').click()
    cy.get('#pass').click()
    cy.get('#card').click()
    cy.get('#userError').should('be.visible').contains("You must enter a username")
    cy.get('#passError').should('be.visible').contains("You must enter a password")
  })

  it('displays error message for invalid username format', () => {
    cy.get('#user').type('first last')
    cy.get('#pass').type('1234')
    cy.get('#userError').should('be.visible').contains("Not a valid username")
  })

  
  it('displays error message for incorrect username or password', () => {
    cy.get('#user').type('test')
    cy.get('#pass').type('test')
    cy.get('#login').click('center')
    cy.get('#invalidLogin').should('be.visible').contains("Sorry, the username or password you entered is incorrect. Please try again.")
  })

  it('removes error message for correct username or password and send POST request', () => {
    cy.get('#user').type('blah')
    cy.get('#pass').type('blah')
    cy.get('#login').click('center')
  })

  it('clicks on sign up and navigates to sign up page', () => {
    cy.get('button').eq(1).click('center')
    cy.url().should('include', '/signup')
    cy.url().should('eq', 'http://localhost:4200/signup')
  })

})