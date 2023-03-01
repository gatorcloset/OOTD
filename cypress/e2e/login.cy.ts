describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1372, 768)
    cy.visit('http://localhost:4200/login')
  })

  it('visits login page', () => {
    cy.visit('http://localhost:4200/login')
  })

  it('clicks on sign up and navigates to sign up page', () => {
    cy.get('button').eq(3).click('center')
    cy.url().should('include', '/signup')
    cy.url().should('eq', 'http://localhost:4200/signup')
  })

})