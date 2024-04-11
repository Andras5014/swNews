import { get, post, type RequestConfig } from '@/utils/request' 

export interface Comment {
    content: string,
    time: number,
    name: string,
}

export interface CommentList {
    pageNo: number,
    list: Comment[],
    totalPage: number,
    pageSize: number,
}


export const getCommentList = (page:number, config?: RequestConfig) => get<CommentList>('/comments/' + page, config)

export const addComment = (data: Comment, config?: RequestConfig) => post('/comments', data, config)