import { CATEGORIES } from "src/app/mock-data/mock-categories"

describe('My Closet Page', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/closet')
  })

  it('visits OOTD "your closet" page', () => {
    cy.visit('http://localhost:4200/closet')
  })

  it('clicks on jeans cateogry and navigates to closet/jeans', () => {
    cy.get('mat-grid-tile').first().click('center')
    cy.url().should('include', '/closet/jeans')
    cy.url().should('eq', 'http://localhost:4200/closet/jeans')
  })
})