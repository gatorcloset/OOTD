const username = 'michelle'
const password = 'michelle'

describe('Closet Item Page', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/login')

    // Log in first
    cy.get('#user').type(username)
    cy.get('#pass').type(password)
    cy.get('#login').click('center')
  })
  
  it('on the jeans page, only clothing under the jeans category is displayed', () => {
    cy.contains('mat-card', 'Bottoms').click()

    cy.get('mat-card').each(($card) => {
      // Check that each card is of category 'Bottoms'
      cy.wrap($card).click()
      cy.get('#categories').contains('Bottom')

      // Close dialog
      cy.get('body').click('left', {multiple: true})
      
    });
  })
})