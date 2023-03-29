describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/closet')
  })

  it('clicks on "outift builder" button on navbar and navigates to builder page', () => {
    cy.get('#outfit-builder').click('center')
    cy.url().should('include', '/builder')
    cy.url().should('eq', 'http://localhost:4200/builder')
  })
})