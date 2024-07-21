import { TestBed } from '@angular/core/testing';

import { GongplanningspecificService } from './gongplanningspecific.service';

describe('GongplanningspecificService', () => {
  let service: GongplanningspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongplanningspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
