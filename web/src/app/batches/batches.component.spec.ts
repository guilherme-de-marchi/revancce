import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BatchesComponent } from './batches.component';

describe('BatchComponent', () => {
  let component: BatchesComponent;
  let fixture: ComponentFixture<BatchesComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [BatchesComponent]
    });
    fixture = TestBed.createComponent(BatchesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
