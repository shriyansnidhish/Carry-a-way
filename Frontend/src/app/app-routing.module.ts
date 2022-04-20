import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BagCalculatorComponent } from './bag-calculator/bag-calculator.component';
import { BookingOptionsComponent } from './booking-options/booking-options.component';
import { HelpComponent } from './help/help.component';
import { HomePageComponent } from './home-page/home-page.component';
import { HowItWorksComponent } from './how-it-works/how-it-works.component';
import { LoginComponent } from './login/login.component';
import { PricingComponent } from './pricing/pricing.component';

var routes: Routes = [
  // {path:'',component: HomePageComponent},
  {path:'',component: HomePageComponent},
  {path:'how-it-works',component: HowItWorksComponent},
  {path:'login',component: LoginComponent},      
  {path:'pricing',component: PricingComponent},
  {path:'bag-calculator',component: BagCalculatorComponent},
  {path:'booking-options',component:BookingOptionsComponent},
  {path:'help',component:HelpComponent}
]; 


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
