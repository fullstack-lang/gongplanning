import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongplanningcomponentComponent } from './gongplanningcomponent.component';

describe('GongplanningcomponentComponent', () => {
  let component: GongplanningcomponentComponent;
  let fixture: ComponentFixture<GongplanningcomponentComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongplanningcomponentComponent]
    });
    fixture = TestBed.createComponent(GongplanningcomponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
