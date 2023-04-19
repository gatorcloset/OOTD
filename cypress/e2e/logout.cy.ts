describe('Log out', () => {
  const username = 'aneesha'
  const password = 'aneesha'

  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/login')
    // Log in first
    cy.get('#user').type(username)
    cy.get('#pass').type(password)
    cy.get('#login').click('center')
  })

  it('clicks on user page and logs out', () => {
    cy.wait(300)
    cy.get('.user-icon').click('center')
    cy.url().should('include', '/user')
    cy.wait(700)
    cy.get('.logout').click('center')
    cy.url().should('include', '/login')
  })

})