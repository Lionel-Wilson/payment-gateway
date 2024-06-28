import { Component } from '@angular/core';
import { PaymentDetails } from '../classes/payment-details';
import { PaymentGatewayService } from '../services/payment-gateway.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  paymentDetailsModel = new PaymentDetails('', '', '', '', 0, '', '');
  paymentResponse: string | null = null; // Variable to hold the payment response
  paymentErrors: string[] = [];

  constructor(private _paymentGatewayService: PaymentGatewayService) {}

  onSubmit(): void {
    this._paymentGatewayService
      .processPayment(this.paymentDetailsModel)
      .subscribe(
        (data) => {
          this.paymentResponse =
            "Payment successful! Here's the payment id: " + data.id;
          this.paymentErrors = []; // Clear any previous errors on success
        },
        (errorResponse) => {
          if (errorResponse.status == 402) {
            this.paymentErrors = [
              'Payment failed! ' + errorResponse.error.responseSummary,
            ];
            this.paymentResponse = ''; // Clear success message on error
            return;
          }

          this.paymentErrors = errorResponse.error.errors;

          this.paymentResponse = ''; // Clear success message on error
        }
      );
  }
}
