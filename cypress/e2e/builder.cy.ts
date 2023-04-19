describe('Name, Build, and Save an outfit', () => {
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

  it('clicks on "outift builder" button on navbar and navigates to builder page', () => {
    cy.wait(300)
    cy.get('#outfit-builder').click('center')
    cy.url().should('include', '/builder')
    cy.url().should('eq', 'http://localhost:4200/builder')
    cy.wait(500)
  })

  it('names, creates, and saves an outfit', () => {
    cy.get('#outfit-builder').click('center')
    cy.get('#name').type('Party Outfit')
    cy.get('#accessories-next-btn')
      .click()
      .click()
      .click()
      .wait(300)
    cy.get('#tops-next-btn')
      .click()
      .click()
      .click()
      .wait(300)
    cy.get('#bottoms-next-btn')
      .click()
      .click()
      .wait(300)
    cy.get('#shoes-next-btn')
      .click()
      .click()
      .wait(500)
    cy.get('.save-btn').click()
  })

})