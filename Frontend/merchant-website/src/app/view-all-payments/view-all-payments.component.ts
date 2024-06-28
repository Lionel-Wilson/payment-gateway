import { Component, OnInit } from '@angular/core';
import { PaymentGatewayService } from '../services/payment-gateway.service';
import { PaymentDetailsDTO } from '../interfaces/payment-details';

@Component({
  selector: 'app-view-all-payments',
  templateUrl: './view-all-payments.component.html',
  styleUrl: './view-all-payments.component.css',
})
export class ViewAllPaymentsComponent implements OnInit {
  payments: PaymentDetailsDTO[] = [];
  errorMessage: string = '';
  constructor(private _paymentGatewayService: PaymentGatewayService) {}

  ngOnInit(): void {
    this.getAllPayments();
  }

  getAllPayments(): void {
    this._paymentGatewayService.retrievePayments().subscribe(
      (data: PaymentDetailsDTO[]) => {
        this.payments = data;
      },
      (errorResponse) => {
        this.errorMessage = errorResponse.error;
      }
    );
  }
}
