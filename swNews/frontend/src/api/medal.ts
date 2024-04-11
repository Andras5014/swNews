import { get, type RequestConfig } from '@/utils/request' 

export interface Medal {
    country: string,
    countryCode: string,
    Rank: number,
    gold: number,
    silver: number,
    bronze: number,
    totalCount: number,
}

export const getMedal = (config?: RequestConfig) => get<Medal[]>('/medals', config)