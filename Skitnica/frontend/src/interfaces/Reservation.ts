export default interface Reservation{
    id: string;
    accomodationId: string;
    username: string;
    startDate: string;
    endDate: string;
    guestNumber: number;
    status: string;
}