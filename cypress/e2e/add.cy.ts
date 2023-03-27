describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1372, 768)
    cy.visit('http://localhost:4200/add')
  })

  it('visits add page', () => {
    cy.visit('http://localhost:4200/add')
  })

  it('fill out form', () => {
    cy.get('#upload').within(() => {
      cy.get('input[type="file"]').selectFile('cypress/fixtures/mock-shorts.jpg');
    });
    cy.get('#name').type("Denim shorts")
    cy.get('#categories').click();
    cy.contains('Bottom').click();
    cy.get('#add-button').click('center')
  });


})