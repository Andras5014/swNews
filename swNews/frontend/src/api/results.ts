import { get } from '@/utils/request'

export interface Result {
    id: string,
    disciplineName: string,
    
}
export interface ResultDetail {
    name: string,
    utcTime: string,
    athleteGroups: AthleteGroup[],
}


export interface AthleteGroup {
    countryName: string,
    rank: number,
    points: string,
    pointBehind: string,
    athletes: Athlete[],
}

export interface Athlete {
    firstName: string,
    lastName: string,
    age: number,
}

export const getResults = () => get<Result[]>('/results')
export const getResultDetail = (id: string) => get<ResultDetail[]>(`/results/${id}`)