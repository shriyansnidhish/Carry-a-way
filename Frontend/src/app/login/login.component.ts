import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EROFS } from 'constants';
// import { any } from 'cypress/types/bluebird';
import { error } from 'protractor';
import { LoginService } from './login.service';
import { loginUser } from './signIn.model';
import { newUser } from './signUp.model';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username: string = "";
  userpassword: string = "";
  logInUser:string = "";
  logInPassword:string = "";
  formType: string = "signUp";
  mySelection = "signUp";
  registerUser: newUser[] = [];
  loginUser: loginUser[] = [];
  loggedInUser: newUser[] = [];

  constructor(public router: Router, public loginService: LoginService) { }

  ngOnInit() {
  }

  LoginUser() {
    
    console.log("login user function entered");
    var thePromise = this.loginService.loginUsers();
    thePromise.then(
      (response) => {
        this.loginUser = response;
        this.loginService.allUsersLogin = response;
        console.log(JSON.stringify(response));
      },
      (error) => {
        console.log(JSON.stringify(error));
      });
    // this.loginService.registerUsers();
    console.log("login user function exited");

    // redirect to pricing page
    this.router.navigate(['pricing']);
    
    console.log("check the logged in user");
    var loggedInUserPromise = this.loginService.loggedInUser();
    loggedInUserPromise.then(
      (response) => {
        this.loggedInUser = response;
        this.loginService.allUsersLoggedIn = response;
        console.log(JSON.stringify(response));
      },
      (error) => {
        console.log(JSON.stringify(error));
      });
  }

  RegisterUser() {
    
    console.log("register user function entered");
    var thePromise = this.loginService.registerUsers();
    thePromise.then(
      (response) => {
        this.registerUser = response;
        this.loginService.allUsersRegister = response;
        console.log(JSON.stringify(response));
      },
      (error) => {
        console.log(JSON.stringify(error));
      });
    // this.loginService.registerUsers();
    console.log("register user function exited");

    // redirect to pricing page
    this.router.navigate(['pricing']);
  }

  onItemChange($event) {
    this.username ="";
  this.userpassword= "";
  this.logInUser = "";
  this.logInPassword = "";
    console.log($event.target.value);
    this.formType = $event.target.value;
  }

}
