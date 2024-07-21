import { Component, Input } from '@angular/core';

import * as gongplanning from '../../../../gongplanning/src/public-api'
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api'

// import { GongsvgModule } from 'gongsvg'
// import { GongsvgspecificModule } from 'gongsvgspecific'


import { AngularSplitModule } from 'angular-split';
import { GongsvgModule } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvg/src/public-api';
import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming';

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';


@Component({
  selector: 'lib-gongplanningcomponent',
  standalone: true,
  templateUrl: './gongplanningcomponent.component.html',
  styleUrls: ['./gongplanningcomponent.component.css'],
  imports: [
    AngularSplitModule,

    TreeComponent,
    GongsvgDiagrammingComponent,
    MaterialFormComponent,


    // GongsvgModule,
    // GongsvgDiagrammingComponent,

    // GongtableModule,
    // GongtablespecificModule,

    // GongtreeModule,
    // GongtreespecificModule,
  ]
})
export class GongplanningcomponentComponent {

  StacksNames = gongplanning.StackName
  FormNames = gongplanning.FormNames

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  @Input() GONG__StackPath: string = ""

}
