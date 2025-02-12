import axios from '../../plugins/axios'

export const getCart = () =>
{
    return axios.get('/cart')
}

export const addToCart = data =>
{
    return axios.post('/cart', data)
}