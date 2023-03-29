import { Component, OnInit } from '@angular/core';
import { LoginDTO } from 'src/app/shared/model/DTO/loginDTO';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginDto: LoginDTO = new LoginDTO;
  errorMessage: string = ''
  constructor(private userService:UserService) { }

  ngOnInit(): void {
  }

  onLogin():void{
    if(!this.CheckFields()){
      this.errorMessage ="All fields must be entered"
    }else{
      this.userService.login(this.loginDto)
    }
    
  }

  CheckFields(): boolean {
    if ( this.loginDto.email == '' || this.loginDto.password == '' ){
        return false
    }
    return true
}
}
