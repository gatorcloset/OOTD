describe('Edit and delete clothing item', () => {
  const username = 'aneesha'
  const password = 'aneesha'

  beforeEach(() => {
    cy.viewport(1372, 768)
    cy.visit('http://localhost:4200/login')

    // Log in first
    cy.get('#user').type(username)
    cy.get('#pass').type(password)
    cy.get('#login').click('center')

    
  })

  it("Edit and delete an outfit", () => {
    // Navigate to outfit
    cy.get('mat-grid-tile').contains('Outfits').click()
    cy.get('mat-grid-tile').first().click()

    // Edit name 
    cy.get('#name').clear().type("Edit outfit name")

    // Edit outfit
    cy.get('#accessories-next-btn')
      .click()
      .wait(300)
    cy.get('#tops-next-btn')
      .click()
      .wait(300)
    cy.get('#bottoms-next-btn')
      .click()
      .wait(300)
    cy.get('#shoes-next-btn')
      .click()
      .wait(300)

    // Save outfit
    cy.get('#save-outfit').click('center')

    // Navigate back to outfits and view edited outfit
    cy.get('mat-grid-tile').contains('Outfits').click()
  })




})