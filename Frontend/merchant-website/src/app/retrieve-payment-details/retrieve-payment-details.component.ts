import { Component } from '@angular/core';
import { PaymentGatewayService } from '../services/payment-gateway.service';
import { PaymentDetails } from '../interfaces/payment-details';

@Component({
  selector: 'app-retrieve-payment-details',
  templateUrl: './retrieve-payment-details.component.html',
  styleUrl: './retrieve-payment-details.component.css',
})
export class RetrievePaymentDetailsComponent {
  constructor(private _paymentGatewayService: PaymentGatewayService) {}

  paymentDetailsResponse: PaymentDetails | undefined; // Variable to hold the payment response
  id: string | null = null; // Variable to hold the payment response
  errorString: string = '';

  onSubmit(): void {
    this._paymentGatewayService.retrievePaymentDetails(this.id).subscribe(
      (data) => {
        this.paymentDetailsResponse = data;
        this.errorString = '';
      },
      (errorResponse) => {
        this.errorString = errorResponse.error;
        this.paymentDetailsResponse = undefined;
      }
    );
  }
}
