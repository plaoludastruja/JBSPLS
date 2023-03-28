import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth.service';

@Component({
  selector: 'app-homepage-menu',
  templateUrl: './homepage-menu.component.html',
  styleUrls: ['./homepage-menu.component.css']
})
export class HomepageMenuComponent implements OnInit {

  userRole: String='';
  constructor(public authService: AuthService) { }

  ngOnInit(): void {
    this.userRole=String(this.authService.decodeToken()?.role);
    console.log(this.userRole);
  }

}
