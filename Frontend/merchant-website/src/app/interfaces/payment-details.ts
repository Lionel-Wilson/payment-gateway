export interface PaymentDetailsDTO {
  id: string;
  firstName: string;
  lastName: string;
  cardNumber: string;
  expiryDate: string;
  amount: number;
  currencyCode: string;
  status: string;
  statusCode: number;
}
