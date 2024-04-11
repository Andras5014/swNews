import { get } from '@/utils/request'

export interface Schedule {
    disciplineName: string,
    id: number,
    date: string,
    time: string,
    utcDateTime: string,
    name: string,
    phaseName: string,
    resultStatus: string,
}

export const getSchedule = () => get<Schedule[]>('/schedule')