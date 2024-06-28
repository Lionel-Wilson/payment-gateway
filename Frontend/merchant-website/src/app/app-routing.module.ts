import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { RetrievePaymentDetailsComponent } from './retrieve-payment-details/retrieve-payment-details.component';
import { ViewAllPaymentsComponent } from './view-all-payments/view-all-payments.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'home', component: HomeComponent },
  {
    path: 'retrieve-payment-details',
    component: RetrievePaymentDetailsComponent,
  },
  {
    path: 'view-all-payments',
    component: ViewAllPaymentsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
