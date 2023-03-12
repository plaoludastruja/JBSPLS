import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateFlightComponent } from './create-flight/create-flight.component';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
//import { DateTimePickerModule, MaskedDateTimeService } from '@syncfusion/ej2-angular-calendars';
import { DateTimePickerModule } from '@smart-webcomponents-angular/datetimepicker';


@NgModule({
  declarations: [
    CreateFlightComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    DateTimePickerModule,
    MatFormFieldModule,
    MatInputModule
  ]
})
export class AdminModuleModule { }
