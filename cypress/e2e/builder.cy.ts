describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
  })

  it('clicks on "outift builder" button on navbar and navigates to builder page', () => {
    cy.visit('http://localhost:4200/closet')
    cy.wait(300)
    cy.get('#outfit-builder').click('center')
    cy.url().should('include', '/builder')
    cy.url().should('eq', 'http://localhost:4200/builder')
    cy.wait(500)
  })

  it('clicks through the tops, bottoms, and shoes with the next button', () => {
    cy.visit('http://localhost:4200/builder')
    for(let n = 0; n < 5; n ++){
      cy.get('#tops-next-btn')
        .click()
        .wait(300)
      cy.get('#bottoms-next-btn')
        .click()
        .wait(300)
        cy.get('#bottoms-next-btn')
        .click()
        .wait(300)
      cy.get('#shoes-next-btn')
        .click()
        .wait(300)
    }
  })

  it('clicks through the tops, bottoms, and shoes with the previous button', () => {
    cy.visit('http://localhost:4200/builder')
    for(let n = 0; n < 5; n ++){
      cy.get('#tops-prev-btn')
        .click()
        .wait(300)
        cy.get('#tops-prev-btn')
        .click()
        .wait(300)
      cy.get('#bottoms-prev-btn')
        .click()
        .wait(300)
      cy.get('#shoes-prev-btn')
        .click()
        .wait(300)
    }
  })
})