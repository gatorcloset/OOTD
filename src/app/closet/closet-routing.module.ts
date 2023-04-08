import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClosetComponent } from './closet.component';

import { ItemsComponent } from '../items/items.component';

/*
const closetRoutes: Routes = [
  { path: 'closet', component: ClosetComponent },
  { path: 'closet/:name', component: ItemsComponent }
];
*/

@NgModule({
  // imports: [RouterModule.forChild(closetRoutes)],
  exports: [RouterModule]
})

export class ClosetRoutingModule { }
