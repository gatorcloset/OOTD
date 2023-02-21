import { CATEGORIES } from "src/app/mock-data/mock-categories"

describe('My Closet Page', () => {
  context('desktop display size', () => {
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

    it('on the jeans page, only clothing under the jeans category is displayed', () => {
      cy.visit('http://localhost:4200/closet/jeans')
      cy.get('mat-grid-tile')
      cy.wrap({ category: CATEGORIES[0].name }).its('category').should('eq', 'Jeans')
    })
  })
})