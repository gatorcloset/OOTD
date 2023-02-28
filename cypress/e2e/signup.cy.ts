describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/signup')
  })

  it ('creates go mod file', () => {
    // cy.exec('cd backend/go/src/github.com/user && go build && ./user')
    // cy.exec('go build')
    // cy.exec('./user')

  })

  it('passes', () => {
    cy.visit('http://localhost:4200/signup')
  })

  // delete this comment
})