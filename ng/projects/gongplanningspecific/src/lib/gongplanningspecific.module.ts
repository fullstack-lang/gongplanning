import { NgModule } from '@angular/core';
import { GongplanningcomponentComponent } from './gongplanningcomponent/gongplanningcomponent.component';

import { GongsvgModule } from 'gongsvg'
import { GongsvgspecificModule } from 'gongsvgspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    GongplanningcomponentComponent,
  ],
  imports: [
    AngularSplitModule,

    GongsvgModule,
    GongsvgspecificModule,

    GongtreeModule,
    GongtreespecificModule,
  ],
  exports: [
    GongplanningcomponentComponent
  ]
})
export class GongplanningspecificModule { }
