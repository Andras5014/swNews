import axios from "axios"
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
export interface Response<T = any> {
  code: number
  data: T
  msg: string
  error: string
}

export interface RequestOptions {
  showLoadingBar?: boolean
  errorCallBack?: (error: Response) => void
  showErrorMessage?: boolean
}

// 拓展自定义请求配置
interface ExpandAxiosRequestConfig<D = any> extends AxiosRequestConfig<D> {
  interceptorHooks?: InterceptorHooks
  requestOptions?: RequestOptions
}
export interface RequestConfig<D = any> extends AxiosRequestConfig<D> {
  requestOptions?: RequestOptions

}
// 拓展axios请求配置
interface ExpandInternalAxiosRequestConfig<D = any> extends InternalAxiosRequestConfig<D> {
  intercetorHooks?: InterceptorHooks
  requestOptions?: RequestOptions
}

// 拓展axios响应配置
interface ExpandAxiosResponse<T = any, D = any> extends AxiosResponse<T, D> {
  config: ExpandInternalAxiosRequestConfig<D>
}
export interface InterceptorHooks {
  requestInterceptor?: (config: ExpandInternalAxiosRequestConfig) => ExpandInternalAxiosRequestConfig
  requestInterceptorCatch?: (config: ExpandInternalAxiosRequestConfig) => any
  responseInterceptor?: (response: ExpandAxiosResponse) => AxiosResponse | Promise<AxiosResponse>
  responseInterceptorCatch?: (response: ExpandAxiosResponse) => any
}


const transform: InterceptorHooks = {
  requestInterceptor: (config) => {
    if (config.requestOptions?.showLoadingBar) {
      // TODO show loading bar
    }
    return config
  },
  requestInterceptorCatch: (config) => {
    if (config.requestOptions?.showErrorMessage) {
      // TODO show error message
    }
    return Promise.reject(config)
  },

  responseInterceptor: (response) => {
    if (response.config.requestOptions?.showLoadingBar) {
      // TODO hide loading bar
    }
    if (response.data.code !== 0) {
      // 请求出现业务错误
      if (response.config.requestOptions?.showErrorMessage) {
        // TODO show error message
      }

      if (response.config.requestOptions?.errorCallBack) {
        response.config.requestOptions.errorCallBack(response.data)
      }
      return Promise.reject(response.data)
    }
    // 请求成功, 直接返回自定义返回类型的data
    return response.data.data
  },

  responseInterceptorCatch: (err) => {
    if (err.config.requestOptions?.showLoadingBar) {
      // TODO hide loading bar
    }
    if (err.config.requestOptions?.showErrorMessage) {
      // TODO show error message
    }
    return Promise.reject(err.config)
  }
}

class Request {
  private _instance: AxiosInstance
  private _defaultConfig: ExpandAxiosRequestConfig = {
    baseURL: '/api',
    timeout: 5000,
    requestOptions: {
      showLoadingBar: true,
      showErrorMessage: true
    }
  }

  private _interceptorHooks?: InterceptorHooks

  constructor(config: ExpandAxiosRequestConfig) {
    this._instance = axios.create(Object.assign(this._defaultConfig, config))
    this._interceptorHooks = config.interceptorHooks
    this.setupInterceptors()
  }

  private setupInterceptors() {
    this._instance.interceptors.request.use(this._interceptorHooks?.requestInterceptor, this._interceptorHooks?.requestInterceptorCatch)
    this._instance.interceptors.response.use(this._interceptorHooks?.responseInterceptor, this._interceptorHooks?.responseInterceptorCatch)
  }

  public request(config: RequestConfig): Promise<AxiosResponse> {
    return this._instance.request(config)
  }

  public get<T = any>(url:string, config?:RequestConfig): Promise<T> {
    return this._instance.get(url, config)
  }

  public post<T = any>(url:string, data?:any, config?:RequestConfig): Promise<T> {
    return this._instance.post(url, data, config)
  }

  public delete<T = any>(url:string, config?:RequestConfig): Promise<T> {
    return this._instance.delete(url, config)
  }

  public put<T = any>(url:string, data?:any, config?:RequestConfig): Promise<T> {
    return this._instance.put(url, data, config)
  }
}

const instance = new Request({
  interceptorHooks: transform
})
 
export const request: (config: RequestConfig) => Promise<AxiosResponse> = (config) => instance.request(config)
export const get: <T = any>(url:string, config?:RequestConfig) => Promise<T> = (url, config) => instance.get(url, config)
export const post: <T = any>(url:string, data?:any, config?:RequestConfig) => Promise<T> = (url, data, config) => instance.post(url, data, config)
export const del: <T = any>(url:string, config?:RequestConfig) => Promise<T> = (url, config) => instance.delete(url, config)
export const put: <T = any>(url:string, data?:any, config?:RequestConfig) => Promise<T> = (url, data, config) => instance.put(url, data, config)

