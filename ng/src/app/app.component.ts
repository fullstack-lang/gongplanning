import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongplanning from 'gongplanning'
import * as gongtable from 'gongtable'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongplanning Data/Model'
  gongplanning = "Gongplanning"
  view = this.gongplanning

  views: string[] = [this.gongplanning, this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongplanning"
  StackType = gongplanning.StackType

  StackNameForDiagrammer = gongplanning.StackName.StackNameDefault

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
