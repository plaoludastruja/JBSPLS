export default class AccomodationRating{
    id: string = "";
    email: string = "";
    accomodationId: string = "";
    rating: number = 0;
    date: string = ""
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.accomodationId = obj.accomodationId;
            this.email = obj.email;
            this.rating = obj.rating;
            this.date = obj.date;
        }
    }
}