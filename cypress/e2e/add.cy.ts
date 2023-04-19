describe('Add new clothing item', function () {
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

  it('Add item', () => {
    cy.get('#add-item > .mdc-button > .mdc-button__label').click()

    // Add item
    cy.get('#upload').within(() => {
      cy.get('input[type="file"]').selectFile('cypress/fixtures/mock-shorts.jpg')
    });
    cy.get('#name').type("Denim shorts")
    cy.get('#categories').click()
    cy.contains('Bottom').click()
    cy.get('#add-button').click('center')

  });


})