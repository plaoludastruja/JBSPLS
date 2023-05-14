export default interface ReservationDto{
    id: string;
    accomodationId: string;
    username: string;
    startDate: string;
    endDate: string;
    guestNumber: number;
    cancellationNum: number;
}