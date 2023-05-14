export default interface SearchResult{
    Id : string;
    AccomodationId: string;
    Name: string;
    Location: string;
    Facilities: string;
    MinNumberOfGuests: number;
    MaxNumberOfGuests: number;
    TotalPrice: number;
    IsAutomaticApproved: string;
    Prices: number[];
    PriceType: string[];
    HostUsername: string;
    Image: string
}