import axios from 'axios';
axios.defaults.baseURL = "/api/v1";

axios.interceptors.request.use((config)=>{
    let loginResult = JSON.parse(localStorage.getItem("loginResult"));
    if(loginResult){
        const token = loginResult.access_token  // 取出accessToken
        config.headers.Authorization = `Barer ${token}`;  //将accesstoken放在请求头里面
    }
    return config;
},(error)=>{
    return Promise.reject(error);
});

axios.interceptors.response.use(
    response =>{
        if(response.status === 200){
            return Promise.resolve(response.data);
        }else {
            return Promise.reject(response.data)
        }
    },
    (error)=>{
        console.log('error',error)
    }
)

export default axios;