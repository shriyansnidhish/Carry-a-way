import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { HowItWorksComponent } from './how-it-works/how-it-works.component';
import { LoginComponent } from './login/login.component';
import { PricingComponent } from './pricing/pricing.component';

var routes: Routes = [
  // {path:'',component: HomePageComponent},
  {path:'',component: HomePageComponent},
  {path:'how-it-works',component: HowItWorksComponent},
  {path:'login',component: LoginComponent},      
  {path:'pricing',component: PricingComponent}
//   {path:'dashboard',
//   component: HomePageComponent,
//   children:[
//     {path:'',component: HomePageComponent},
//     {path:'how-it-works',component: HowItWorksComponent},
//     {path:'login',component: LoginComponent},      
//     {path:'pricing',component: HomePageComponent}
//   ]
// }
]; 


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
