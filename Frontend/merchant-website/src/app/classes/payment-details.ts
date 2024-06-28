export class PaymentDetails {
  constructor(
    public firstName: string,
    public lastName: string,
    public cardNumber: string,
    public expiryDate: string,
    public amount: number,
    public currencyCode: string,
    public cvv: string
  ) {}
}
