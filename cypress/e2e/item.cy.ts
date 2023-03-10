import { CATEGORIES } from "src/app/mock-data/mock-categories"

describe('Closet Item Page', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/closet')
  })
  
  it('on the jeans page, only clothing under the jeans category is displayed', () => {
    cy.visit('http://localhost:4200/closet/jeans')
    cy.get('mat-grid-tile')
    cy.wrap({ category: CATEGORIES[0].name }).its('category').should('eq', 'Jeans')
  })

  // delete this comment
})