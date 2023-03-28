import { TestBed } from '@angular/core/testing';

import { NewItemService } from './new-item.service';

describe('NewItemService', () => {
  let service: NewItemService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NewItemService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
