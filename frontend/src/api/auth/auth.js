import axios from '../../plugins/axios'

export const login =  data => 
{
    return axios.post('/login', data)
}

export const signup =  data =>
{
    return  axios.post('/signup', data)
}