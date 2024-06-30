import { Component } from '@angular/core';
import { PaymentDetails } from '../classes/payment-details';
import { PaymentGatewayService } from '../services/payment-gateway.service';
import { ErrorResponse } from '../interfaces/error-response';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  paymentDetailsModel = new PaymentDetails('', '', '', '', 0, '', '');
  paymentResponse: string | null = null; // Variable to hold the payment response
  error: ErrorResponse = {
    statusCode: 0,
    message: '',
    errors: [],
  };

  constructor(private _paymentGatewayService: PaymentGatewayService) {}

  onSubmit(): void {
    this._paymentGatewayService
      .processPayment(this.paymentDetailsModel)
      .subscribe(
        (data) => {
          this.paymentResponse =
            "Payment successful! Here's the payment id: " + data.id;

          this.error.errors = []; // Clear any previous errors on success
        },
        (errorResponse) => {
          if (errorResponse.status == 402) {
            this.error.errors = [
              'Payment failed! ' + errorResponse.error.responseSummary,
            ];
            this.paymentResponse = ''; // Clear success message on error
            this.error.message = '';
            return;
          }

          this.error.errors = errorResponse.error.errors;
          this.error.message = errorResponse.error.message;

          this.paymentResponse = ''; // Clear success message on error*/
        }
      );
  }
}
