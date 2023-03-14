import { Component, OnInit } from '@angular/core';
import { NgbDateStruct, NgbCalendar, NgbDatepickerModule } from '@ng-bootstrap/ng-bootstrap';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IFlight } from 'src/app/shared/material/Flight';

@Component({
  selector: 'app-create-flight',
  templateUrl: './create-flight.component.html',
  styleUrls: ['./create-flight.component.css']
})
export class CreateFlightComponent implements OnInit {

  public dateValue: Date = new Date("06/06/2022 15:30");
  public form: FormGroup;
  public errorMessage: String = ""

  constructor(
    fb: FormBuilder) {
      this.form = fb.group({
        startDateTime: ['', [Validators.required]],
        startPlace: ['', [Validators.required, Validators.pattern("^[A-Z].*$")]],
        endDateTime: ['', [Validators.required]],
        endPlace: ['', [Validators.required, Validators.pattern("^[A-Z].*$")]],
        maxNumberOfPlaces: ['', [Validators.required]],
        pricePerPlace: ['', [Validators.required]]
      });
     }
     
  get startDateTime(){
    return this.form.get('startDateTime');
  }
  get startPlace(){
    return this.form.get('startPlace');
  }
  get endDateTime(){
    return this.form.get('endDateTime');
  }
  get endPlace(){
    return this.form.get('endPlace');
  }
  get maxNumberOfPlaces(){
    return this.form.get('maxNumberOfPlaces');
  }
  get pricePerPlace(){
    return this.form.get('pricePerPlace');
  }

  ngOnInit(): void {
  }

  register(){
    let flight: IFlight = {
      start: this.startDateTime?.value,
      startPlace: String(this.startPlace?.value),
      end: this.endDateTime?.value,
      endPlace: String(this.endPlace?.value),
      maxNumberOfPlaces: this.maxNumberOfPlaces?.value,
      pricePerPlace: this.pricePerPlace?.value
    }
    console.log(flight)
    if(this.form.valid && this.compareDates(this.startDateTime, this.endDateTime)){
      this.errorMessage=""
    }else{
      if(this.errorMessage=="") this.errorMessage="All fields must be entered"
    }
  }

  compareDates(start: any, end: any): boolean{
    const startDate = new Date(start.value).valueOf();
    const endDate = new Date(end.value).valueOf();
    const now = new Date().valueOf();
    if(startDate <= now){
      this.errorMessage="Start date must be in future";
      return false;
    }
    if(endDate <= now){
      this.errorMessage="End date must be in future";
      return false;
    }
    if(startDate - endDate > 0){
      this.errorMessage="End date must be after start date"
      return false;
    }
    return true;
  }

}
