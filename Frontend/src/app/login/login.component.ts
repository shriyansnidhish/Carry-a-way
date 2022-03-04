import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username:string="";
  userpassword:string="";
  
  constructor(public router : Router) { }

  ngOnInit() {
  }

  AuthenticateUser(){
    if(this.username === "admin" && this.userpassword === "admin"){
        // redirect to dashboard !
        this.router.navigate(['']);
    }
  }
}
