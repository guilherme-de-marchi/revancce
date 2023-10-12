import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BatchComponent } from './batch.component';

describe('BatchComponent', () => {
  let component: BatchComponent;
  let fixture: ComponentFixture<BatchComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [BatchComponent]
    });
    fixture = TestBed.createComponent(BatchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
