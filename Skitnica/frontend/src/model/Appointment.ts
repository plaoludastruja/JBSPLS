export default interface Appointment{
    id : string;
    accomodationId: string;
    start: string;
    end: string;
    priceType: string;
    price: number;
    status: string;
}