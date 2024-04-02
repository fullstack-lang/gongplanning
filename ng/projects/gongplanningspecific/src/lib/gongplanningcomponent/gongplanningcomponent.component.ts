import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-gongplanningcomponent',
  templateUrl: './gongplanningcomponent.component.html',
  styleUrls: ['./gongplanningcomponent.component.css']
})
export class GongplanningcomponentComponent {

  @Input() GONG__StackPath: string = ""

}
