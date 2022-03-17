import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { loginUser } from "./signIn.model";
import { newUser } from "./signUp.model";

@Injectable()
export class LoginService {
    allUsersRegister: newUser[] = [];
    allUsersLogin: loginUser[] = [];
    allUsersLoggedIn: newUser[] = [];
    constructor(public httpSerObj: HttpClient) {

    }

// Register Users POST method
    registerUsers() {
        console.log("register user  in service entered");
        return this.httpSerObj.post<newUser[]>('http://localhost:8000/api/register'
            , {
                 fname : 'Abhilash',
                 lname : 'abhilash',
                 email : 'doesnotexistatall@gmail.com',
                 password : 'Abhi@1234'
            }
        ).toPromise();
    }

// Login Users POST method
    loginUsers() {
        console.log("login user  in service entered");
        return this.httpSerObj.post<loginUser[]>('http://localhost:8000/api/login'
        ,{
            email : 'doesnotexistatall@gmail.com',
            password : 'Abhi@1234'
        }
        ).toPromise();
    }

// Logged in user GET method 
    loggedInUser(){
        console.log("logged in user in service entered");
        return this.httpSerObj.get<newUser[]>('http://localhost:8000/api/user').toPromise();
    }
}