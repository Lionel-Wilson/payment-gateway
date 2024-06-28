import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RetrievePaymentDetailsComponent } from './retrieve-payment-details.component';

describe('RetrievePaymentDetailsComponent', () => {
  let component: RetrievePaymentDetailsComponent;
  let fixture: ComponentFixture<RetrievePaymentDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [RetrievePaymentDetailsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RetrievePaymentDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
