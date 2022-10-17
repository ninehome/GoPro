<template>
    <div class="head-content">

        <div class="l-contnet">

            <el-button @click="clickLeftIcon" icon="el-icon-menu" size="mini"></el-button>
            <span class="text">首页</span>

        </div>



        <div class="r-contnet">

            <el-dropdown>
                <span class="el-dropdown-link">
                <img class="userhead" src="../assets/images/icon_user.webp" />
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>个人中心</el-dropdown-item>
                    <el-dropdown-item>退出</el-dropdown-item>
            
                </el-dropdown-menu>
        </el-dropdown>

        </div>





    </div>
</template>





<script>
export default {
    data() {
        return {}
    },
    methods:{
        clickLeftIcon(){
            this.$store.commit('collapsMenu')  //因为vue 在main.js 中挂载了strore ，this可以获取到。 再调用strore方法修改值
        }
    }
}
</script>



<style lang="less" scoped>
.head-content {
    background-color: rgb(24, 28, 29);
    height: 60px;
    display: flex;
    justify-content: space-between; //X 轴两端显示
    align-items: center; //y轴垂直居中
    padding-left: 20px;
    border-right: none;
    border-left: none;

    .text {
        color: #fff;
        margin-left: 10px;
        font-size: 14px;
    }
    .dropMenu{
        margin-right: 20px;
    }

    .r-contnet{
        margin-right: 20px;
        .userhead{
           width: 40px;
           height: 40px;
           border-radius: 50%;
        }
    }
}




</style>
