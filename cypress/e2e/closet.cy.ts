describe('My Closet Page', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/closet')
  })

  it('visits OOTD "your closet" page', () => {
    cy.visit('http://localhost:4200/closet')
  })

  it('clicks on bottoms cateogry and navigates to closet/bottoms', () => {
    cy.get('mat-grid-tile').first().click('center')
    cy.url().should('include', '/closet/bottoms')
    cy.url().should('eq', 'http://localhost:4200/closet/bottoms')
  })
})