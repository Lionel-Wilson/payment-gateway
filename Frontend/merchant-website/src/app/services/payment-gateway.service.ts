import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { PaymentDetails } from '../classes/payment-details';
import { PaymentDetailsDTO } from '../interfaces/payment-details';

@Injectable({
  providedIn: 'root',
})
export class PaymentGatewayService {
  _url = 'http://localhost:8080/api/v1/payments';

  constructor(private _http: HttpClient) {}

  processPayment(paymentDetails: PaymentDetails) {
    return this._http.post<any>(this._url, paymentDetails);
  }
  retrievePaymentDetails(id: string | null) {
    return this._http.get<any>(this._url + '/' + id);
  }
  retrievePayments() {
    return this._http.get<PaymentDetailsDTO[]>(this._url);
  }
}
