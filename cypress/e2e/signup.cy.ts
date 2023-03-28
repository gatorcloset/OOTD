describe('template spec', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
    cy.visit('http://localhost:4200/signup')

    cy.request('POST', 'http://localhost:9000/users', { 
      firstname: 'Michelle',
      lastname: 'Taing',
      username: 'michelle.taing', 
      password: 'hellothere123'
    })
      .its('body')
      .as('currentUser')
  })

  it ('fill out form and enter', function() {
    const { firstname, lastname, username, password} = this['currentUser'];
    cy.get('input[id=firstname]').type(firstname)
    cy.get('input[id=lastname]').type(lastname)
    cy.get('input[id=email]').type('mleetaing@gmail.com')
    cy.get('input[id=username]').type(username)
    cy.get('input[id=password]').type(password)
    cy.get('.login-button').first().click('center')

  })
})