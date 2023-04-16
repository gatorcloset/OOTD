describe('Edit and delete clothing item', () => {
  const username = 'michelle'
  const password = 'michelle'

  beforeEach(() => {
    cy.viewport(1372, 768)
    cy.visit('http://localhost:4200/login')

    // Log in first
    cy.get('#user').type(username)
    cy.get('#pass').type(password)
    cy.get('#login').click('center')

    
  })

  it('Edit clothing item image, name, and category', () => {
    // Select clothing item
    cy.get('mat-grid-tile').first().click()
    cy.get('mat-grid-tile').last().click()

    // Change image
    cy.get('input[type="file"]').selectFile('cypress/fixtures/mock-shirt.jpg')

    // Change name
    cy.get('#name').clear().type('Button up')

    // Change category
    cy.get('#categories').click();
    cy.contains('Top').click();
    
    // cy.get('#save-button').click()
  })

  it('Check edited item', () => {
    // Locate item just edited
    cy.contains('mat-card', 'Tops').click()
    cy.contains('mat-card', 'Button up').click()

    cy.get('#categories').contains('Top')
  })




})