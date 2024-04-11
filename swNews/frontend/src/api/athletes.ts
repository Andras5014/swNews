import { get, type RequestConfig } from '@/utils/request' 

export interface Athlete {
    country: string,
    firstName: string,
    lastName: string,
    gender: string,
    dob: string,
    countryCode: string,
}


export const getAthletes = (config?: RequestConfig) => get<Athlete[]>('/athletes', config)
