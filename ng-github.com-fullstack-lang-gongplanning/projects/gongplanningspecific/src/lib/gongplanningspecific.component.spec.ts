import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongplanningspecificComponent } from './gongplanningspecific.component';

describe('GongplanningspecificComponent', () => {
  let component: GongplanningspecificComponent;
  let fixture: ComponentFixture<GongplanningspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongplanningspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongplanningspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
