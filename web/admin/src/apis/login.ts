import { $http } from '.'

export const loginApi = async (data: { username: string, password: string }) => {
    const res = await $http({
        method: 'post',
        url: '/login',
        data
    })
    return res
}