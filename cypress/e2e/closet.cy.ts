describe('My Closet Page', () => {
  const username = 'michelle'
  const password = 'michelle'
  
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/login')

    // Log in first
    cy.get('#user').type(username)
    cy.get('#pass').type(password)
    cy.get('#login').click('center')
  })

  it('clicks on bottoms cateogry and navigates to closet/bottoms', () => {
    cy.get('mat-grid-tile').first().click('center')
    cy.url().should('include', '/closet/bottoms')
    cy.url().should('eq', 'http://localhost:4200/closet/bottoms')
  })
})