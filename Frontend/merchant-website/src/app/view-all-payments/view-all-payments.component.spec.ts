import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewAllPaymentsComponent } from './view-all-payments.component';

describe('ViewAllPaymentsComponent', () => {
  let component: ViewAllPaymentsComponent;
  let fixture: ComponentFixture<ViewAllPaymentsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ViewAllPaymentsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewAllPaymentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
