export interface Flight {
    id: string,
    start: string,
    startPlace: string,
    end: string,
    endPlace: string,
    maxNumberOfPlaces: number,
    remaining: number,
    pricePerPlace: number
}