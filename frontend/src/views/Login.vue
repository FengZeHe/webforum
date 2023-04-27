<template>
    <div class="main">
        <div class="container">
            <h2 class="form-title">登陆</h2>
            <div class="form-group">
                <label for="name">用户名</label>
                <input type="text" class="form-control" v-model="username" name="name" id="name" placeholder="请输入用户名">
            </div>
            <div class="form-group">
                <label for="pass">密码</label>
                <input type="text" class="form-control" v-model="password" name="pass" id="pass" placeholder="请输入密码">
            </div>
            <div class="form-btn">
                <button type="button" class="btn btn-info" @click="submit">提交</button>
            </div>
        </div>
    </div>
</template>
<script>
export default{
    name:"Login",
    data(){
        return{
            username:"",
            password:"",
            submitted:false
        }
    },
    methods:{
        submit(){
            this.$axios({
                methods:'post',
                url:'login',
                data: JSON.stringify({
                    username: this.username,
                    password: this.password
                })
            }).then((res)=>{
                console.log(res.data)
                if(res.code == 1000){
                    localStorage.setItem("loginResult",JSON.stringify(res.data));
                    this.$store.commit("login",res.data);
                    this.$router.push({path:this.redirect || '/'})
                }else{
                    console.log(res.msg)
                }
            }).catch((error)=>{
                console.log(error)
            })
        }
    }
}

</script>