<template>
    <div class="main">
        <div class="container">
            <h2 class="form-title">注册</h2>
            <div class="form-group">
                <label for="name"><span style="color:red;">* </span>用户名</label>
                <el-input type="text" require name="name" id="name" placeholder="请输入用户名" v-model="username"></el-input>
            </div>
            <div class="form-group">
                <label for="email"><span style="color:red;">* </span>邮箱</label>
                <el-input type="text" require name="email" id="email" placeholder="请输入邮箱" v-model="email"></el-input>
            </div>
            <div class="form-group">
                <label for="pass"><span style="color:red;">* </span>密码</label>
                <el-input type="password" require name="pass" id="pass" placeholder="请输入密码"
                    v-model="password"></el-input>
            </div>
            <div class="form-group">
                <label for="re_pass"><span style="color:red;">* </span>确认密码</label>
                <el-input type="password" require name="re_pass" id="re_pass" placeholder="请确认密码"
                    v-model="confirm_password"></el-input>
            </div>
            <div class="form-group">
                <label for="gender"><span style="color:red;">* </span>性别</label>
                <div id="gender">
                    <el-radio v-model="gender" :label="1">男</el-radio>
                    <el-radio v-model="gender" :label="2">女</el-radio>
                </div>
            </div>
            <div class="form-btn">
                <button type="button" class="btn btn-info" @click="submit">提交</button>
            </div>
        </div>
    </div>
</template>

<script>
export default ({
    name: "SignUp",
    data() {
        return {
            username: '',
            password: '',
            email: '',
            gender: 1,
            confirm_password: '',
            submitted: false
        };
    },
    methods: {
        submit() {
            this.$axios({
                method: 'POST',
                url: '/signup',
                data: {
                    username: this.username,
                    eamil: this.email,
                    gender: this.gender,
                    password: this.password,
                    confirm_password: this.confirm_password
                }
            }).then((res) => {
                console.log(res.data);
                if (res.conde == 1000) {
                    console.log('signup success')
                    this.$router.push({ name: "Home" });
                } else {
                    console.log(res.msg);
                }
            }).catch((error) => {
                console.log(error)
            })
        }

    }
})

</script>

