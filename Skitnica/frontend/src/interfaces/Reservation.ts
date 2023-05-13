export default class Reservation{
    id: string = "";
    accomodationId: string = "";
    username: string = "";
    startDate: string = "";
    endDate: string = "";
    guestNumber: number = 0;
    status: string = "";
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.accomodationId = obj.accomodationId;
            this.username = obj.username;
            this.startDate = obj.startDate;
            this.endDate = obj.endDate;
            this.guestNumber = obj.guestNumber;
            this.status = obj.status;
        }
    }
}
