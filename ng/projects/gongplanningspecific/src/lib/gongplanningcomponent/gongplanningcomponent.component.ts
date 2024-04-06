import { Component, Input } from '@angular/core';

import * as gongplanning from 'gongplanning'
import * as gongtable from 'gongtable'


@Component({
  selector: 'lib-gongplanningcomponent',
  templateUrl: './gongplanningcomponent.component.html',
  styleUrls: ['./gongplanningcomponent.component.css']
})
export class GongplanningcomponentComponent {

  StacksNames = gongplanning.StackName
  FormNames = gongplanning.FormNames

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  @Input() GONG__StackPath: string = ""

}
