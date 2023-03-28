
import { ClosetComponent } from "src/app/closet/closet.component"
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatCardModule} from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatDividerModule} from '@angular/material/divider';
import {MatGridListModule} from '@angular/material/grid-list';

describe('closet-component.cy.ts', () => {
  beforeEach(() => {
    cy.viewport(1366, 768)
  })

  it('test', () => {
    cy.mount(ClosetComponent, {
      imports: [
        MatToolbarModule,
        MatCardModule,
        MatIconModule,
        MatButtonModule,
        MatDividerModule,
        MatGridListModule
      ],
    })

  })
})