import axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { ElLoading } from 'element-plus'
import "element-plus/theme-chalk/el-loading.css";

export const httpInstance = axios.create({})

export type BackResponse = {
    data: any
    code: number
    message: string
    success: boolean
}

// 设置请求根路径
httpInstance.defaults.baseURL = import.meta.env.VITE_BASE_URL

// 响应拦截器
export const $http = async (config: AxiosRequestConfig) => {
    const loadingInstance = ElLoading.service()
    try {
        const axiosResponse = await httpInstance<BackResponse>(config)
        const backResponse = axiosResponse.data
        if (!backResponse?.success) {
            let errTitle: string = 'ERROR'
            if (backResponse?.code === 401) {
                errTitle = '未授权'
                /* ElMessage.error('未授权，请重新登录') */
                // 重定向到login页面
            } else if (backResponse?.code === 403) {
                errTitle = '无权限'
            } else if (backResponse?.code === 404) {
                errTitle = '未发现'
            } else if (backResponse?.code === 500) {
                errTitle = '服务器错误'
            } else {
                errTitle = '未知错误'
            }
            const error = new Error(errTitle)
            throw error
        }
        return backResponse
    } catch (error) {
        if (error instanceof AxiosError) {
            error.message = '网络错误'
            console.log(error)
        }
        throw error
    } finally {
        loadingInstance.close()
    }
}
